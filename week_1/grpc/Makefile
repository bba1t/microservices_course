LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	# protoc-gen-go генерирует код из прото, устанавливаю его не в корень бинарников, а в этот пакет
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	# protoc-gen-go-grpc генерирует grpc сервер и клиент, тоже устанавливаю в пакет
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/note_v1	# папка куда все сгенерируется

	# 1. путь где лежит прото файл
	# 2. хочу сгенерировать го файл в папку пкг		3. подтверждение папки
	# 4. показываю, где установлен protoc-gen-go
	# 5. хочу сгенерировать grpc клиент/сервер в папку пкг	  6. подтверждение папки
	# 7. показываю, где установлен protoc-gen-go-grpc
	# 8. указываю путь до прото файла
	protoc --proto_path api/note_v1 \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/note_v1/note.proto