docker_up:
	docker-compose up -d
docker_down:
	docker-compose down
run_proto:
	protoc --go_out=. --go-grpc_out=. user/user.proto