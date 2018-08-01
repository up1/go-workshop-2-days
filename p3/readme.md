# Learning Go 2 days with somkiat.cc
owner of p3 (Apipol Sukgler)

# Before start
```
go get gopkg.in/mgo.v2
```
# Run Project
```
export GOPATH=$(pwd) // current directory p3/
go run src/main.go
```
# Try REST API
## Echo
```
server := router.SetupRouteEcho()
server.Start(":3000")
```
## Gin
```
server := router.SetupRouteGin()
server.Run(":3000")
```
## Standard Library
```
mux := router.SetupRouteStandardLibrary()
http.ListenAndServe(":3000", mux)
```

#GO STEP
```
go get ....
gofmt -w src
go clean -testcache
go test -v -cover -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
go install patient // go build -o ./bin/patient patient
./bin/patient
```