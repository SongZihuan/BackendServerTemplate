# Copyright 2025 BackendServerTemplate Authors. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

length=40
randomString=$(cat /dev/urandom | tr -dc 'a-z0-9' | fold -w $length | head -n 1)
echo "$randomString" > random_data.txt