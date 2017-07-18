#/bin/bash

export GOPATH=/Users/sebastienlavayssiere/go

cd $GOPATH/src/github.com/slavayssiere/gamename/player
swagger generate spec -o ./public/index.json

