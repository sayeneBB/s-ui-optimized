# Set encoding to UTF-8
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

Write-Host "=== Starting Git Sync Process ==="
$repoPath = "C:\Users\stefa\.gemini\antigravity\scratch\s-ui"
Set-Location $repoPath

# 1. Config identity
Write-Host "Configuring user identity..."
git config user.name "sayeneBB"
git config user.email "sayeneBB@users.noreply.github.com"

# 2. Stage changes
Write-Host "Staging changes..."
git add -A

# 3. Commit changes if there are any
Write-Host "Checking status..."
$status = git status --porcelain
if ($status) {
    Write-Host "Changes detected. Committing..."
    git commit -m "Update s-ui repository changes"
} else {
    Write-Host "No changes to commit."
}

# 4. Push changes
Write-Host "Pushing to origin main..."
git push origin main

Write-Host "=== Git Sync Completed ==="
