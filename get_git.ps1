# Copyright 2025 BackendServerTemplate Authors. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

if (Test-Path -Path ".\.git" -PathType Container) {
    $last_commit = git rev-parse HEAD 2>$null
    $last_tag = git describe --tags --abbrev=0 2>$null

    if (-not [string]::IsNullOrEmpty($last_tag)) {
        $last_tag_commit = git rev-list -n 1 $last_tag 2>$null

        Set-Content -Path "commit_data.txt" -Value $last_commit -Encoding UTF8
        Set-Content -Path "tag_data.txt" -Value $last_tag -Encoding UTF8
        Set-Content -Path "tag_commit_data.txt" -Value $last_tag_commit -Encoding UTF8
    } else {
        Set-Content -Path "commit_data.txt" -Value $last_commit -Encoding UTF8
        New-Item -Path "tag_data.txt" -ItemType File -Force
        New-Item -Path "tag_commit_data.txt" -ItemType File -Force
    }
} else {
    New-Item -Path "commit_data.txt" -ItemType File -Force
    New-Item -Path "tag_data.txt" -ItemType File -Force
    New-Item -Path "tag_commit_data.txt" -ItemType File -Force
}

# 创建 VERSION 文件
New-Item -Path "VERSION" -ItemType File -Force
