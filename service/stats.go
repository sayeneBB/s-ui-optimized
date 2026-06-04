package service

import (
	"time"

	"github.com/sayeneBB/s-ui/database"
	"github.com/sayeneBB/s-ui/database/model"

	"gorm.io/gorm"
)

type onlines struct {
	Inbound  []string `json:"inbound,omitempty"`
	User     []string `json:"user,omitempty"`
	Outbound []string `json:"outbound,omitempty"`
}

var onlineResources = &onlines{}

type StatsService struct {
}

func (s *StatsService) SaveStats(enableTraffic bool) error {
	if corePtr == nil || !corePtr.IsRunning() {
		return nil
	}
	box := corePtr.GetInstance()
	if box == nil {
		return nil
	}
	st := box.StatsTracker()
	if st == nil {
		return nil
	}
	stats := st.GetStats()

	// Reset onlines
	onlineResources.Inbound = nil
	onlineResources.Outbound = nil
	onlineResources.User = nil

	if len(*stats) == 0 {
		return nil
	}

	var err error
	db := database.GetDB()
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	for _, stat := range *stats {
		if stat.Resource == "user" {
			if stat.Direction {
				err = tx.Model(model.Client{}).Where("name = ?", stat.Tag).
					UpdateColumn("up", gorm.Expr("up + ?", stat.Traffic)).Error
			} else {
				err = tx.Model(model.Client{}).Where("name = ?", stat.Tag).
					UpdateColumn("down", gorm.Expr("down + ?", stat.Traffic)).Error
			}
			if err != nil {
				return err
			}
		}
		if stat.Direction {
			switch stat.Resource {
			case "inbound":
				onlineResources.Inbound = append(onlineResources.Inbound, stat.Tag)
			case "outbound":
				onlineResources.Outbound = append(onlineResources.Outbound, stat.Tag)
			case "user":
				onlineResources.User = append(onlineResources.User, stat.Tag)
			}
		}
	}

	if !enableTraffic {
		return nil
	}
	err = tx.Create(&stats).Error
	return err
}

func (s *StatsService) GetStats(resource string, tag string, limit int) ([]model.Stats, error) {
	currentTime := time.Now().Unix()
	timeDiff := currentTime - (int64(limit) * 3600)
	timeMin := timeDiff
	timeMax := currentTime

	db := database.GetDB()
	resources := []string{resource}
	if resource == "endpoint" {
		resources = []string{"inbound", "outbound"}
	}

	numBuckets := 30
	bucketSpan := (timeMax - timeMin) / int64(numBuckets)
	if bucketSpan <= 0 {
		bucketSpan = 1
	}

	// 1. Initialize all 60 output rows (30 buckets * 2 directions)
	downsampled := make([]model.Stats, 0, numBuckets*2)
	type key struct {
		bucket    int
		direction bool
	}
	idxMap := make(map[key]int)

	for i := 0; i < numBuckets; i++ {
		bucketStart := timeMin + int64(i)*bucketSpan
		for _, dir := range []bool{false, true} {
			idxMap[key{bucket: i, direction: dir}] = len(downsampled)
			downsampled = append(downsampled, model.Stats{
				DateTime:  bucketStart,
				Resource:  resource,
				Tag:       tag,
				Direction: dir,
				Traffic:   0,
			})
		}
	}

	// 2. Query aggregated results from SQLite
	type DBStats struct {
		BucketID   int  `gorm:"column:bucket_id"`
		Direction  bool `gorm:"column:direction"`
		AvgTraffic int64 `gorm:"column:avg_traffic"`
	}

	var dbResults []DBStats
	err := db.Model(model.Stats{}).
		Select("CAST((date_time - ?) / ? AS INTEGER) AS bucket_id, direction, CAST(AVG(traffic) AS INTEGER) AS avg_traffic", timeMin, bucketSpan).
		Where("resource IN ? AND tag = ? AND date_time >= ? AND date_time <= ?", resources, tag, timeMin, timeMax).
		Group("bucket_id, direction").
		Scan(&dbResults).Error

	if err != nil {
		return nil, err
	}

	// 3. Fill the aggregated results into the pre-allocated downsampled slice
	for _, r := range dbResults {
		bId := r.BucketID
		if bId < 0 {
			bId = 0
		}
		if bId >= numBuckets {
			bId = numBuckets - 1
		}
		k := key{bucket: bId, direction: r.Direction}
		if idx, ok := idxMap[k]; ok {
			downsampled[idx].Traffic = r.AvgTraffic
		}
	}

	return downsampled, nil
}

func (s *StatsService) GetOnlines() (onlines, error) {
	return *onlineResources, nil
}
func (s *StatsService) DelOldStats(days int) error {
	oldTime := time.Now().AddDate(0, 0, -(days)).Unix()
	db := database.GetDB()
	return db.Where("date_time < ?", oldTime).Delete(model.Stats{}).Error
}
