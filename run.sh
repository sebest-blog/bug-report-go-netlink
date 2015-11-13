#!/bin/bash

docker run --privileged -v `pwd`:/go -ti golang ./bug.sh
