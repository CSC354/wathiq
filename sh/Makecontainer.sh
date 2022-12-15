#!/usr/bin/env sh

cd .. # Since using make
result=${PWD##*/}

docker run --net debate-net --name $result -d $result
