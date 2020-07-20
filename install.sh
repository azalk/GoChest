#!/bin/bash
go build -buildmode=c-shared -o libGoChest.so GoChest.go Changepoint.go
python3 build.py
echo "Step 2"
python3 setup.py install --force
