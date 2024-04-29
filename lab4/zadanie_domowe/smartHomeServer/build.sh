#!/usr/bin/env bash

echo "Build started"
mvn install || { echo "Build failed"; exit 1; }
server_id=$1
echo "Server id $server_id"
mvn exec:java -Dexec.mainClass=org.example.Main -Dexec.args="$server_id" || { echo "Failed to run server"; exit 1; }
