#!/bin/bash

set -e

export GOOS=linux
export GOARCH=amd64

ScriptPath=$(cd `dirname $0` && pwd)
ProjectPath=$ScriptPath

rm client
rm server
go build -o "$ProjectPath/client" --race "$ProjectPath/client.go"
go build -o "$ProjectPath/server" "$ProjectPath/server.go"

