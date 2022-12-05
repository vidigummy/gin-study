CREATE DATABASE RPG;

CREATE user 'vidi'@'%' identified BY 'vidi';

GRANT ALL privileges ON *.* TO 'vidi'@'%';

FLUSH PRIVILEGES;