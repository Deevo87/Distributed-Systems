#!/usr/bin/env bash

JAVA_PATH="../smartHomeServer/src/main/java/org/example/Ice"

echo "Ice for Java in ${JAVA_PATH}"
mkdir -p $JAVA_PATH
slice2java server.ice --output-dir $JAVA_PATH
