#!/bin/bash

# Copyright 2019 Adevinta

set -e

function delete_dir {
	for var in "$@"
	do
		echo "Deleting $var"
		if [ -d "$var" ]
		then
			rm -r "$var"/
		fi
	done
}

# Delete auto-generated content
delete_dir client swagger
