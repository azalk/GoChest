#!/bin/bash
go build -buildmode=c-shared -o GoChest.so CWrapper.go
python3 setup.py build_ext install
cp /Users/lukas/Documents/Projects/GoChest/GoChest.so /Library/Frameworks/Python.framework/Versions/3.8/lib/python3.8/site-packages/GoChest-0.51-py3.8-macosx-10.9-x86_64.egg