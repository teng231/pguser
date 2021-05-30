protoc -I/usr/local/include -I. -I$GOPATH/src \
    --gogofaster_out=plugins=grpc:. user.proto

protoc-go-inject-tag -input=./user.pb.go