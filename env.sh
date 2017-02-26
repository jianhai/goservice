#!/bin/bash

LOCALE=$(dirname $(readlink -f "$0"))

export GOPATH="$LOCALE/3rd/golang/"
