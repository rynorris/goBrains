#!/bin/sh

cd $HOME/gopath/src
mkdir -p github.com/banthar/Go-SDL
cd github.com/banthar/Go-SDL
git clone git@github.com:banthar/Go-SDL.git
make
