if (Test-Path -Path ".\.git" -PathType Container) {
    # 如果 .git 目录存在
    $last_commit = git rev-parse HEAD 2>$null
    $last_tag = git describe --tags --abbrev=0 2>$null

    if (-not [string]::IsNullOrEmpty($last_tag)) {
        # 如果有标签
        $last_tag_commit = git rev-list -n 1 $last_tag 2>$null

        Set-Content -Path "commit_data.txt" -Value $last_commit -Encoding UTF8
        Set-Content -Path "tag_data.txt" -Value $last_tag -Encoding UTF8
        Set-Content -Path "tag_commit_data.txt" -Value $last_tag_commit -Encoding UTF8
    } else {
        # 如果没有标签
        Set-Content -Path "commit_data.txt" -Value $last_commit -Encoding UTF8
        New-Item -Path "tag_data.txt" -ItemType File -Force
        New-Item -Path "tag_commit_data.txt" -ItemType File -Force
    }
} else {
    # 如果 .git 目录不存在
    New-Item -Path "commit_data.txt" -ItemType File -Force
    New-Item -Path "tag_data.txt" -ItemType File -Force
    New-Item -Path "tag_commit_data.txt" -ItemType File -Force
}

# 创建 VERSION 文件
New-Item -Path "VERSION" -ItemType File -Force
