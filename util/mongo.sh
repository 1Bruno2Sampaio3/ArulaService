#!/bin/bash
echo "Removing proccess on the port 7778"
kill $(lsof -t -i:7778)

sleep 1

echo "Building bd"
mongod --directoryperdb --dbpath $HOME/workspace/ArulaService/mongodb --logpath $HOME/workspace/ArulaService/mongodb/logs/mongo.log --logappend --rest --port 7778 &

sleep 2

exit