DB_NAME=arula-test
DB_USERNAME=adm
DB_PW=6666
DB_DROP=db.dropDatabase();
DB_CREATE=db.createCollection('users');db.createCollection('companies');db.createCollection('jobs');

debug:
	echo Buildando Service
	cd ${GOPATH}/src && go build
db:
	sudo ./util/mongo.sh
db_reset:
	echo "Deleting logs"
	cd mongodb/logs && echo "" > mongo.log
	sudo ./util/mongo.sh
	mongo --port 7778 $(DB_NAME) --eval "$(DB_DROP)$(DB_CREATE)"
	exit

