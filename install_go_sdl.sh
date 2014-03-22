#!/bin/sh

cd $HOME/gopath/src
mkdir -p github.com/banthar
cd github.com/banthar
yes yes | git clone git@github.com:banthar/Go-SDL.git
make
