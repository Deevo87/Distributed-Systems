#!/usr/bin/bash

# shellcheck disable=SC2145
mvn compile && mvn exec:java -Dexec.mainClass=org.solution.Main -Dexec.args="$@"

