#!/usr/bin/bash
dir=$(ls -l ./ |awk '/^d/ {print $NF}')
for i in ${dir}
do
echo -n "{$i}  ->"

# protoc --go_out=. user.proto

protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. $i/*.proto
done

