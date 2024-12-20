#!/bin/bash

rm -rf lottery_server
git pull
go build -tags release
nohup ./lottery_server > lottery_server.log 2>&1 &