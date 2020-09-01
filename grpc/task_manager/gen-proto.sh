#!/bin/bash
CURRENT_DIR=$GOPATH/src/golang_start/grpc/task_manager
for x in $(find ${CURRENT_DIR}/proto/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/proto -I /usr/local/include --go_out=plugins=grpc:${CURRENT_DIR} ${x}/*.proto
done