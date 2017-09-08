debug:
	export ARULA_DB_NAME="arula-test"
	cd ${GOPATH}/src && go build -o arulaservice.exe

db:
	sudo ./util/mongo.sh