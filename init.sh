#!/bin/sh
# Get the name of the MySQL container

set -e

MYSQL_CONTAINER=$(docker ps -a | grep mysql | awk '{print $1}')
MYSQL_PASS=secret
# Create the user

echo $MYSQL_CONTAINER

echo $MYSQL_PASS && echo "Creating user"
docker exec -it $MYSQL_CONTAINER mysql -u root -p$MYSQL_PASS -e "CREATE USER 'dev'@'%' IDENTIFIED BY 'dev1234';"
# Grant the user privileges
docker exec -it $MYSQL_CONTAINER mysql -u root -p$MYSQL_PASS -e "GRANT ALL PRIVILEGES ON *.* TO 'dev'@'%' WITH GRANT OPTION;"
docker exec -it $MYSQL_CONTAINER mysql -u root -p$MYSQL_PASS -e "FLUSH PRIVILEGES;"
echo "User Created"