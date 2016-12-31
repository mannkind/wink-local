go test -v ./...

go test -coverprofile=controller/.coverprofile github.com/mannkind/wink-local/controller
go test -coverprofile=handlers/.coverprofile github.com/mannkind/wink-local/handlers
gover . .coverprofile
go tool cover -html=.coverprofile 
find . -name ".coverprofile" -exec rm {} \;
