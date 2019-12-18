generate:\
  protoc -I=./src --go_out=plugins=grpc\:./out ./src/hello.proto

# protoc -I=./src --go_out=plugins=grpc:./out ./src/hello.proto