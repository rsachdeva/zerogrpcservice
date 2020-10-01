# Plugin style is the right way with go_out currently

protoc -I=api/zero --go_out=plugins=grpc:third_party/zeropb zeroservice.proto

