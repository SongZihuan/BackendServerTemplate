$timestamp = [math]::Floor([decimal]((Get-Date).ToUniversalTime() | Get-Date -UFormat %s))
Set-Content -Path "build_date.txt" -Value $timestamp
