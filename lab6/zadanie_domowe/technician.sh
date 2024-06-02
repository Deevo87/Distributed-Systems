#!/usr/bin/zsh

if [ -n "$1" ]; then
  args="$1"
  java -jar target/zadanie_domowe-1.0-SNAPSHOT-technician-jar-with-dependencies.jar "$args"
else
  java -jar target/zadanie_domowe-1.0-SNAPSHOT-technician-jar-with-dependencies.jar
fi