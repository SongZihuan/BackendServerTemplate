# Copyright 2025 BackendServerTemplate Authors. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

$characters = 'abcdefghijklmnopqrstuvwxyz0123456789'
$length = 40

$randomString = -join (1..$length | ForEach-Object { $characters[(Get-Random -Maximum $characters.Length)] })

Set-Content -Path "random_data.txt" -Value $randomString -Encoding UTF8
