#/bin/bash

export GOPATH=/Users/sebastienlavayssiere/go

cd $GOPATH/src/github.com/slavayssiere/gamename/player
~/go/bin/swagger -apiPackage github.com/slavayssiere/gamename/player -format swagger
mkdir -p $GOPATH/src/github.com/slavayssiere/gamename/player/public
mv index.json $GOPATH/src/github.com/slavayssiere/gamename/player/public/
# cd $GOPATH/src/github.com/slavayssiere/gamename/docs/ && go build && mv ./docs ../docs-ui/player


