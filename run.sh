#!/bin/bash

function g() {
    go mod init $0
    go run .
}

for d in $(ls | grep ex)
do
  echo "---- Run $d --------------------"
  cd $d
  go mod init $(basename $PWD)
  go run .
  cd ..
done
