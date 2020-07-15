#!/usr/bin/bash
dir=$(ls -l ./ |awk '/^d/ {print $NF}')
for i in ${dir}
do
echo -n "{$i}  ->"


#protoc --proto_path=. --micro_out=. --go_out=.  ./greeter/greeter.proto
#protoc  --micro_out=. --go_out=. $i/*.proto
protoc --proto_path=. --micro_out=. --go_out=.  ./$i/*.proto
done

