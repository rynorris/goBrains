#!/bin/bash

cd ${GOPATH%%:*}/src
mkdir -p github.com/banthar
cd github.com/banthar
git clone https://github.com/banthar/Go-SDL.git

