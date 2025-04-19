# Copyright 2025 BackendServerTemplate Authors. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

if [ -d "./.git" ]; then
    last_commit="$(git rev-parse HEAD 2>/dev/null)"
    last_tag="$(git describe --tags --abbrev=0 2>/dev/null)"

    if [ -n "$last_tag" ]; then
      last_tag_commit="$(git rev-list -n 1 "$last_tag" 2>/dev/null)"

      echo "$last_commit" > commit_data.txt
      echo "$last_tag" > tag_data.txt
      echo "$last_tag_commit" > tag_commit_data.txt
    else
      echo "$last_commit" > commit_data.txt
      touch tag_data.txt
      touch tag_commit_data.txt
    fi
else
    touch commit_data.txt
    touch tag_data.txt
    touch tag_commit_data.txt
fi

touch "VERSION"
