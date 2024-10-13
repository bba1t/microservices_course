package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	// Пакет где лежит сгенерированный код
	desc "github.com/olezhek28/microservices_course/week_1/grpc/pkg/note_v1"
)

const grpcPort = 50051

// Нужно методами этой структуры соответствовать интерфейсу сгенерированному из proto файла в note_grpc.pb.go(89 строка)
type server struct {
	desc.UnimplementedNoteV1Server
}

// Ручка Get ... первый параметр всегда контекст, второй тот самый proto message
// В сгенерированном файле node_grpc.pb.go уже есть все методы, реализованные как заглушки, но при имплементации, они перекрываются

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	// печатаю входящий proto message
	log.Printf("Note id: %d", req.GetId())

	// возвращаю нод, сгенерировав поля
	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Content:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	// net.Listen("tcp", ...) — открывает сетевое соединение, которое будет "слушать" входящие запросы по протоколу TCP на указанном порту.
	// fmt.Sprintf(":%d", grpcPort) — форматирует строку для указания порта в формате ":50051".
	// lis — объект типа net.Listener, который будет слушать входящие соединения.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()  // экземпляр gRPC сервера, который будет обрабатывать входящие от клиента gRPC-запросы
	reflection.Register(s) // включаю возможность серверу рассказать какие методы у него есть (строка 223)

	// Когда gRPC-клиент отправляет запрос на сервис NoteV1, сервер s будет перенаправлять этот запрос к методам, реализованным в структуре server{}.
	// Это нужно для того, чтобы gRPC-сервер знал, какую логику использовать для обработки запросов клиентов.
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	// блокирующий метод, закидываю лис, в котором порт
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
