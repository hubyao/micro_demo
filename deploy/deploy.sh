#!/bin/bash
###
 # @Author       : jianyao
 # @Date         : 2020-07-21 12:05:29
 # @LastEditTime : 2020-07-21 12:24:12
 # @Description  : file content
###


# 接收参数一,动作命令
p1=("docker","build","run","deploy") # docker=容器化.build=编译,run=运行,deploy=发布

parameterError=true

# shellcheck disable=SC2068
for i in ${array[@]}
do
   [ "$i" == "$var" ] && $parameterError=false
done

# shellcheck disable=SC1066
if  $parameterError
then
    echo "参数错误格式 make [docker|build|run|deploy]"
    exit
fi

APP_NAME=$1
echo $APP_NAME