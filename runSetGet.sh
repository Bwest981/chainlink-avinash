#!/bin/bash

git clone https://github.com/avinash-k16/chainlink.git
cd chainlink
go buil main.go
result=`go run .`
contract=$(echo $result | cut -d " " -f3)
val=$(echo $result | cut -d " " -f8)

JSON='{"Contract Address" : "%s", "Value set" : "%s"}\n'
printf "$JSON" "$contract" "$val" | tee output.json
