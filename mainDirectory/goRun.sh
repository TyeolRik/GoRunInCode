#!/bin/bash

# This Script is for running "go run [main file]" for another folder
nowPath=$(pwd)

# Input Arguments needed (Path. not Go file!)
cd $1
go mod init && go mod tidy && go mod vendor
go run . > "${nowPath}/output.txt"
