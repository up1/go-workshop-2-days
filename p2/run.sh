export GOPATH=$(pwd)
rm -rf ./bin
gofmt -w src/
go test -v -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
go install clinic
./bin/clinic