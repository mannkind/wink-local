# Build wink-local
REPO="github.com/mannkind/wink-local"
go get -v $REPO
cd $GOPATH/src/$REPO

GOOS=linux GOARCH=arm GOARM=5 go build -v -ldflags="-s -w" .
upx -q wink-local
cd web
npm install
npm run dist
cd -
