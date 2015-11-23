# Build wink-mqtt
go get -v github.com/mannkind/wink-mqtt
GOOS=linux GOARCH=arm GOARM=5 go build -v -ldflags="-s -w" github.com/mannkind/wink-mqtt
upx -q wink-mqtt
