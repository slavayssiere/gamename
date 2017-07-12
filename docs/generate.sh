#/bin/bash

export GOPATH=/Users/sebastienlavayssiere/go

cd $GOPATH/src/github.com/slavayssiere/gamename/player
~/go/bin/swagger -apiPackage github.com/slavayssiere/gamename/player
mv docs.go $GOPATH/src/github.com/slavayssiere/gamename/docs/
cd $GOPATH/src/github.com/slavayssiere/gamename/docs/ && go build && mv ./docs ../docs-ui/player


