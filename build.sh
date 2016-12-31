# Build wink-local
go get -v github.com/mannkind/wink-local
GOOS=linux GOARCH=arm GOARM=5 go build -v -ldflags="-s -w" github.com/mannkind/wink-local
upx -q wink-local
