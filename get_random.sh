length=40
randomString=$(cat /dev/urandom | tr -dc 'a-z0-9' | fold -w $length | head -n 1)
echo "$randomString" > random_data.txt