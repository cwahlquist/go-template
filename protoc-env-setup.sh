#!/bin/sh

GRPC_RELEASE_TAG=v1.12.x
GO_PROTO_RELEASE_TAG=v1.2.0

git clone -b ${GRPC_RELEASE_TAG} https://github.com/grpc/grpc /var/local/git/grpc && \
    cd /var/local/git/grpc && \
    git submodule update --init && \
    echo "--- installing protobuf ---" && \
    cd third_party/protobuf && \
    ./autogen.sh && ./configure --enable-shared && \
    make -j$(nproc) && make -j$(nproc) check && make install && make clean && ldconfig && \
    echo "--- installing grpc ---" && \
    cd /var/local/git/grpc && \
    make -j$(nproc) && make install && make clean && ldconfig && \
    git clone https://github.com/grpc/grpc-web /var/local/git/grpc/grpc-web && \
    cd /var/local/git/grpc/grpc-web && \
    echo "--- installing protoc-gen-go ---" && \
    make install-plugin && \
    go get github.com/golang/dep/cmd/dep && \
    go get -d -u github.com/golang/protobuf/protoc-gen-go && \
    git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GO_PROTO_RELEASE_TAG && \
    go install github.com/golang/protobuf/protoc-gen-go && \
    echo 'PATH=$PATH:$(go env GOPATH)/bin' > ~/.bashrc
