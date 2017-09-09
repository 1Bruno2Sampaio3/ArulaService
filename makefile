debug:
	cd ${GOPATH}/src && go build
db:
	echo "Deleting logs"
	cd mongodb/logs && echo "" > mongo.log
	sudo ./util/mongo.sh