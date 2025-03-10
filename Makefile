BINARY_PATH=${EXEC_DIRECTORY}/bin/api
EXEC_DIRECTORY=cmd/api
DOCKER_IMAGE_NAME=sos-api

default: run

.PHONY: build
build:
	go build -o ${BINARY_PATH} ${EXEC_DIRECTORY}/main.go

.PHONY: run
run: build
	./${BINARY_PATH}

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test ./...


########## DOCKER
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} -f build/Dockerfile .

.PHONY: docker-compose-up
docker-compose-up:
	docker compose -f ./build/docker-compose.yaml up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker compose -f ./build/docker-compose.yaml down
##########

########## OPEN API
.PHONY: swag-init
swag-init:
	swag init -g ${EXEC_DIRECTORY}/main.go -o ./api/docs
##########