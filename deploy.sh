#!/bin/bash

#git reset --hard HEAD && git fetch && git checkout $1 && git pull origin $1
#sudo ./isubata/db/init.sh
(
cd webapp/go
make
)

## replaces log data 
LOGPATH=/var/log
NOW=`date +'%H-%M-%S'`
sudo cp $LOGPATH/nginx/access.log $LOGPATH/nginx/access-$NOW.log
sudo sh -c 'echo "" > /var/log/nginx/access.log'
#
sudo cp $LOGPATH/mariadb/slow.log $LOGPATH/mariadb/slow-$NOW.log
sudo sh -c 'echo "" > /var/log/mariadb/slow.log'
#
#sudo cp conf/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf

echo 'systemctl are restarting...'
sudo systemctl restart mariadb.service
sudo systemctl restart torb.go.service 
sudo systemctl restart nginx.service
echo 'Finished to restart!!'

(
cd bench
bin/bench -remotes=127.0.0.1 -output result.json
)
jq . < bench/result.json
sudo /usr/local/bin/alp -f /var/log/nginx/access.log -r --sum | head -n 30
#sudo mysqldumpslow -s t /var/log/mariadb/slow.log | head -n 5
#cd isubata/bench && ./bin/bench -remotes=127.0.0.1 -output result.json
#jq . < result.json
