function genProto {
    DOMAIN=$1
    SKIP_GATEWAY=$2
    PROTO_PATH=./${DOMAIN}/proto
    GO_OUT_PATH=./${DOMAIN}/proto
#    mkdir -p $GO_OUT_PATH

#    protoc   --go_out=plugins=grpc:$PROTO_PATH ${PROTO_PATH}/*.proto
    protoc   --go_out=plugins=grpc,paths=source_relative:. ./auth/proto/*.proto
    protoc   --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./auth/proto/auth.yaml:. ./auth/proto/*.proto


#    protoc -I=$PROTO_PATH --go_out=plugins=grpc,paths=source_relative:$GO_OUT_PATH user.proto
}

genProto auth
#genProto rental
#genProto blob 1
#genProto car