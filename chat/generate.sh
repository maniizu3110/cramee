#!/bin/bash
export PATH=$PATH:/Users/user/go/bin/

protoc chat/chat.proto --go_out=plugins=grpc:.