#!/bin/bash
go build -buildmode=c-shared -o GoChest.so CWrapper.go
python3 test.py