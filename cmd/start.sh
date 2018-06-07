#!/bin/bash -e

sudo -s

# set this in ~/.profile
export GOPATH=~/go
export PATH=$PATH:/usr/lib/go-1.10/bin:$GOPATH/bin

# install Ponzu
cd $GOPATH/src
cd github.com/lsilvapvt/tripezio

# build ponzu
ponzu build

# start ponzu server
ponzu run --port 8080 --https --bind tripezio.com &
