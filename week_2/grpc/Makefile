LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/note_v1/note.proto

build:
	# на какой операционке билдить, какая архитектура,
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

copy-to-server:
	# сама утилита, что копировать, куда копировать, после двоеточия путь внутри сервера
	scp service_linux root@5.159.100.20:

docker-build-and-push:
	# Билдит образ(images) из описанного dockerfile, под линукс64, положит локально, -t.... это тэг
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/grpc/test-server:v0.0.1 .
	# Логинюсь на удаленном Container Registry, генерирую токен и заполняю команду
	docker login -u token -p CRgAAAAAZ3oEnicMZyG9vVbsVGuOpC4Y18SHjkDd cr.selcloud.ru/grpc
	# Пушу образ на удаленный Container Registry
	docker push cr.selcloud.ru/grpc/test-server:v0.0.1
