#!/bin/sh
go build ./src/server

ng serve & ./server.exe
#./src/server/SWE-2023-Spring
#gin --port 4201 --path . --build ./src/server/ --i --all &