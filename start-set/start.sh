docker pull mysql

docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -d -p 3306:3306 mysql:latest

docker exec -it mysql-container bash

mysql -u root -p root

