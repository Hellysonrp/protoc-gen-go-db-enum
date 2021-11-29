# protoc-gen-go-db-enum

This protobuf compiler plugin generates the `Valuer` and `Scanner` interfaces for enums defined in the `proto` files.

It is highly based of the code of [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go/tree/b92717ecb630d4a4824b372bf98c729d87311a4d/cmd/protoc-gen-go).
I just got the code, stripped out any unnecessary parts, and built this plugin.

# Usage

Install it using the `go install` command:
> go install github.com/Hellysonrp/protoc-gen-go-db-enum

Some usage examples:
> protoc --plugin protoc-gen-go-db-enum --go-db-enum_out=output example.proto

> protoc --plugin protoc-gen-go-db-enum --go-db-enum_out=paths=source_relative:output example.proto

If you have problems with `protoc` not finding the plugin in the `PATH`, I recommend passing the absolute path to the plugin:
> protoc --plugin ${HOME}/go/bin/protoc-gen-go-db-enum --go-db-enum_out=output example.proto

> protoc --plugin ${HOME}/go/bin/protoc-gen-go-db-enum --go-db-enum_out=paths=source_relative:output example.proto
