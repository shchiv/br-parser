#!/usr/bin/env bash

docker build -f Dockerfile -t br-parser-base .

exit_code=$?

if [  "$exit_code" != "0" ]; then
	echo "Build failed"
else
	echo "Build successful"
fi

exit $exit_code