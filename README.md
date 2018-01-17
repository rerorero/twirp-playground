```
protoc --twirp_out=. --go_out=. ./rpc/haberdasher/service.proto
go run main.go
```
```
curl localhost:8080/twirp/twirp.example.haberdasher.Haberdasher/MakeHat -H 'Content-Type: application/json' -d '{"inches":1}'
```
