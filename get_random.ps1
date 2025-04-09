$characters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
$length = 40

$randomString = -join (1..$length | ForEach-Object { $characters[(Get-Random -Maximum $characters.Length)] })

Set-Content -Path "random_data.txt" -Value $randomString -Encoding UTF8
