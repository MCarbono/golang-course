curl -H 'Content-Type: application/json' -d '{ "id": "1", "price": 100, "tax": 1}' -X POST http://localhost:8000/order

genereate proto
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto