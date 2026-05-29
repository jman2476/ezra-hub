#! /bin/bash
go build -C ./app/server -o ezra_server
docker build ./app/server -t ezra_server:latest