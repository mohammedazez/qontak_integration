arg = $(filter-out $@,$(MAKECMDGOALS))
this_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

proto:
	@protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. \app/grpcHandler/fiber.proto

proto_s:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/grpcHandler/fiber.proto

run:
	@go run main.go $(call arg)

build:
	@go build

run-nohup: build
	@nohup ./qontak_integration $(call arg) &

docker-compose:
	@docker-compose up -d

docker-compose-down:
	@docker-compose down

docker:
	@docker build . -t qontak_integration
	@docker image prune --filter label=stage-qontak_integration=builder -f

docker-run:
	docker run -v $(this_dir)logs:/app/logs -d --restart always --env-file .env --hostname qontak_integration --name qontak_integration -p 9098:8888 -e TZ=Asia/Jakarta qontak_integration

docker-stop:
	@docker stop qontak_integration

clear-container:
	@docker rm -f qontak_integration

docker-stop-rm: docker-stop clear-container

clear-image:
	@docker rmi -f qontak_integration

clear-docker: clear-container clear-image

docker-exec:
	@docker exec -it qontak_integration sh

volume:
	@docker volume create $(call arg)

run-redis:
	@docker run --detach --restart always --name redis_fiber --hostname redis.fiber redis redis-server

run-redis-volume:
	@docker run --detach --restart always -v $(call arg):/data --name redis_fiber --hostname redis.fiber redis redis-server

redis-cli:
	@docker run -it  --link redis_fiber:redis --rm redis redis-cli -h redis -p 6379

stop-redis:
	@docker stop redis_fiber
	@docker rm redis_fiber