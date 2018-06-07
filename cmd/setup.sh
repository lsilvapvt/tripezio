#!/bin/bash -e

# install golang
sudo apt-get upgrade -y
sudo apt-get update -y
sudo add-apt-repository ppa:gophers/archive -y
sudo apt-get update -y
sudo apt-get install golang-1.10-go -y
mkdir -p ~/go/src ~/go/bin

sudo -s

# set this in ~/.profile
export GOPATH=~/go
export PATH=$PATH:/usr/lib/go-1.10/bin:$GOPATH/bin

# install Ponzu
cd $GOPATH/src
go get github.com/ponzu-cms/ponzu/...

# get tripezio project
go get github.com/lsilvapvt/tripezio
cd github.com/lsilvapvt/tripezio

# build ponzu
ponzu build

# start ponzu server
ponzu run --port 8080 --https --bind tripezio.com
