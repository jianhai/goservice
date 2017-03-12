#!/bin/bash

function isIgnore()
{
    line=$1
    if [[ ${line:0:1} == "#" ]]; then
        return 1
    else 
        return 0
    fi
}

CURRPATH=$(dirname $(readlink -f "$0"))

while read line; do
    isIgnore $line
    [[ $? != 1 ]] || continue

    echo "Running>   " $CURRPATH/$line
    $CURRPATH/$line
  
done < $CURRPATH/case
