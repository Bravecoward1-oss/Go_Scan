#!/bin/sh

go mod tidy

# shellcheck disable=SC2164
cd src/
go build -o Go_Scan -x
sudo chmod +x Go_Scan

echo Compile the complete!!!
