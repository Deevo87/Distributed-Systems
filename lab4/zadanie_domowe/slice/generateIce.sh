#!/usr/bin/env bash

JAVA_PATH="../smartHomeServer/src/main/java"
PYTHON_PATH="../Client/"

echo "Ice for Java in ${JAVA_PATH}"
mkdir -p $JAVA_PATH
slice2java generate.ice --output-dir $JAVA_PATH

echo "Ice for Python in ${PYTHON_PATH}"
mkdir -p $PYTHON_PATH
slice2py generate.ice --output-dir $PYTHON_PATH