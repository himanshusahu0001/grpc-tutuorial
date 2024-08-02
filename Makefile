gen: 
	go generate ./proto/...

run:
	go run main.go 

clean:
	rm proto/service_grpc.pb.go  || true
	rm proto/service.pb.go || true
	rm -rf proto/swagger/* || true
	rm proto/service.pb.gw.go || true
	rm proto/service.gw.client.go || true