#!/usr/bin/env bash

#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=123.56.66.22
port=33069
dbname=dev_$1
username=root
passwd=PXDN93VRKUm8TeE7

goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true