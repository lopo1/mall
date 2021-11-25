; auth
protoc   --go_out=plugins=grpc,paths=source_relative:. ./auth/proto/*.proto
protoc   --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./auth/proto/auth.yaml:. ./auth/proto/*.proto

; goods
protoc   --go_out=plugins=grpc,paths=source_relative:. ./goods/proto/*.proto
protoc   --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./goods/proto/goods.yaml:. ./goods/proto/*.proto

; inventory
protoc   --go_out=plugins=grpc,paths=source_relative:. ./inventory/proto/*.proto
protoc   --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./inventory/proto/inventory.yaml:. ./inventory/proto/*.proto

; order
protoc   --go_out=plugins=grpc,paths=source_relative:. ./order/proto/*.proto
protoc   --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./order/proto/order.yaml:. ./userop/proto/order.proto

; userop
protoc   --go_out=plugins=grpc,paths=source_relative:. ./userop/proto/*.proto
protoc   --grpc-gateway_out=paths=source_relative,grpc_api_configuration=./userop/proto/userfav.yaml:. ./userop/proto/userfav.proto

pause