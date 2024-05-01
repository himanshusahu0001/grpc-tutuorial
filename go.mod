module github.com/user/test

go 1.20

require (
	github.com/akuity/grpc-gateway-client v0.0.0-20231116134900-80c401329778
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/alevinval/sse v1.0.1 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
)

replace github.com/user/test/proto/ => ./proto/
