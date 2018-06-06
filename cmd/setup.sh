#!/bin/bash -e

# install golang
sudo apt-get update
sudo add-apt-repository ppa:gophers/archive -y
sudo apt-get update
sudo apt-get install golang-1.10-go -y
mkdir -p ~/go/src ~/go/bin
export GOPATH=~/go
export PATH=$PATH:/usr/lib/go-1.10/bin:$GOPATH/bin

# install Ponzu
cd $GOPATH/src
go get github.com/ponzu-cms/ponzu/...
# get git cli

# get tripezio project

# build ponzu

# start ponzu server
