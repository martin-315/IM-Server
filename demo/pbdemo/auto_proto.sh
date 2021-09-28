#!/usr/bin/env bash

source ./proto_dir.cfg

for ((i = 0; i < ${#all_proto[*]}; i++)); do
  proto=${all_proto[$i]}
  generate=${all_generate[$i]}
#  echo ./$general
  protoc --go_out=plugins=grpc:./$generate $proto
  s=`echo $proto | sed 's/ //g'`
  v=${s//proto/pb.go}
  protoc-go-inject-tag -input=./$v
  echo "protoc --go_out=plugins=grpc:." $proto
done
echo "proto file generate success..."
