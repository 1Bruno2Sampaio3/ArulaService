#!/bin/bash
clear
echo "Executando script."

echo "Excluindo porta ativa"
kill $(lsof -t -i:7778)

echo "Executando conexão"
mongod --directoryperdb --dbpath $HOME/workspace/ArulaService/mongodb --logpath $HOME/workspace/ArulaService/mongodb/logs/mongo.log --logappend --rest --port 7778