#!/bin/bash

$LOCAL=$(git rev-parse HEAD)
$REMOTE=$(git rev-parse @{u})


git fetch
if [$LOCAL != $REMOTE]
then
	echo "Out of date..."
	echo "Pulling, building, and restarting"
	git pull
	# Assumes Makefile shipped or included
	SRVR=$(make)
	# could specifcy more paths here, but fine for now
	export GIN_MODE=release
	# Go should build it to match the basename of working directory	
	./$(basename "`pwd`") &

else
	echo "Up to date with remote"
fi
