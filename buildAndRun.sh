#!/bin/bash

# This is useful if you want to quickly build and run this to do some testing of any changes you are making
docker build -t dummy-external-adapter .  
docker run --rm -p 6060:6060 dummy-external-adapter