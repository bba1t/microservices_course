package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// пакет где лежит сгенерированный код
	desc "github.com/olezhek28/microservices_course/week_1/grpc/pkg/note_v1"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	// Этот вызов создает gRPC-клиента и открывает соединение с сервером по указанному адресу.
	// grpc.Dial — это метод, который устанавливает соединение с gRPC-сервером.
	// grpc.WithTransportCredentials(insecure.NewCredentials()) — указывает, что для подключения не используются безопасные соединения (без TLS/SSL).
	// переменная conn будет хранить объект соединения
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close() // закрывает соединение с gRPC-сервером

	// Создается клиент для вызова RPC методов сервиса NoteV1, используя ранее установленное соединение conn
	client := desc.NewNoteV1Client(conn)

	// Создает контекст с тайм-аутом.
	// context.Background() — это базовый контекст, который не содержит данных.
	// time.Second — тайм-аут в 1 секунду. Это ограничение по времени для вызова RPC, чтобы он не выполнялся бесконечно.
	// ctx — переменная, которая хранит контекст с тайм-аутом.
	// cancel — функция для отмены этого контекста, которая должна быть вызвана для освобождения ресурсов, когда работа завершена.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Вызов RPC метода Get на сервере через gRPC-клиента.
	// client.Get — это вызов метода Get, который реализован на стороне сервера (получение заметки по ID).
	// &desc.GetRequest{Id: noteID} — это аргумент запроса, который содержит структуру GetRequest с полем Id.
	// r — будет содержать ответ от сервера (информацию о заметке), если запрос выполнится успешно.
	r, err := client.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get note by id: %v", err)
	}

	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetNote()))
}
