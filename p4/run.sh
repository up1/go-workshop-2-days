set GOPATH=%cd%
rm -rf ./bin
gofmt -w src/
go test -v -cover -coverprofile=coverage.out ./...
go tool cover -html=coverdage.out
go install clinic
./bin/clinic