## Structure

```bash
./demo
├── README.md
├── client  # client service
│   └── main.go
├── rpc  # gRPC module
│   ├── rpc.pb.go
│   ├── rpc.proto  # defined proto file
│   └── rpc_grpc.pb.go
└── server  # server service
    └── main.go
```

## Prerequisites

Install [protocol buffer compiler](https://protobuf.dev/installation/):

```bash
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+
```

Inatsll go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Workflow

1. Define your functions in `rpc/rpc.proto` file
2. Generate Go code from the proto file with protoc:

```bash
protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    ./api/proto/server/server.proto
```

3. Create server service to implement the gRPC funcs in `server/main.go`
4. Create client service to call the gRPC funcs
