#!/bin/sh

cd $HOME/gopath/src
mkdir -p github.com/banthar
cd github.com/banthar
git clone https://github.com/banthar/Go-SDL.git
cd Go-SDL && make install
