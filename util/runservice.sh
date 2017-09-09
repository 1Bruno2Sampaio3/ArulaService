GO=/usr/bin/go
GOPATH=/home/$USER/workspace/ArulaService/go/src

echo "Exporting environment vars"

export ARULA_DB_NAME="arula-test"
export ARULA_DB_PASSWORD="6666"
export ARULA_DB_USERNAME="admin"

echo "Executing build"

cd $GOPATH && ./src

exit