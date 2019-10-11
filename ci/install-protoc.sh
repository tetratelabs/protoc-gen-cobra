#!/bin/bash

set -ex

curl -sSL "https://github.com/protocolbuffers/protobuf/releases/download/v3.10.0/protoc-3.10.0-linux-x86_64.zip" -o /tmp/protoc.zip
unzip /tmp/protoc.zip -d /tmp/protoc
sudo mv /tmp/protoc/bin/protoc /usr/local/bin/protoc
