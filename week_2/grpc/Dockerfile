# какой образ операционки(alpine) и с какими зависимлстями(golang:1.20.3) AS название_этапа
# все images берутся из dockerhub
FROM golang:1.20.3-alpine AS builder

# по этому пути скопируется все содержимое папки в которой находится dockerfile
COPY . /github.com/olezhek28/microservices_course/week_2/grpc/source/
# помечаю папку в которую все скопировал как рабочую директорию
WORKDIR /github.com/olezhek28/microservices_course/week_2/grpc/source/

# там запущу установку всех хависимостей
RUN go mod download
# и запущу билд
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

# второй этап, создаю образ только с операционкой
# гошка была нужна для билда приложения, а уже для работы бинарника она не понадобится
FROM alpine:latest

WORKDIR /root/
# копирую бинарник из предыдущего этапа в рабочую директорию этого образа
COPY --from=builder /github.com/olezhek28/microservices_course/week_2/grpc/source/bin/crud_server .

CMD ["./crud_server"]