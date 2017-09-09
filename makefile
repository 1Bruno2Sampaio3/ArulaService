debug:
	cd ${GOPATH}/src && go build

db:
	echo "Deleting logs"
	cd mongodb/logs && echo "" > mongo.log
	sudo ./util/mongo.sh
db_reset:
	echo "Deleting logs"
	cd mongodb/logs && echo "" > mongo.log
	sudo ./util/mongo.sh
	mongo --port 7778 arula-test --eval "db.dropDatabase();db.createCollection('users'); db.createCollection('companies'); db.createCollection('jobs');"
	exit

