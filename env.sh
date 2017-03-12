#!/bin/bash

# Get Current Real Path
basepath=`pwd`

# Setting GOPATH
export GOPATH="$basepath/3rd/golang/"

# Get GoLang Package
while read line; do
    echo " Goang Package" $line
    go get $line
done < $basepath/3rd/golang/README
