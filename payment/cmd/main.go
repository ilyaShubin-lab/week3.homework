package main

import (
	"log"
	"net"

	paymentv1API "boilerplates/payment/internal/api/payment/v1"
	paymentService "boilerplates/payment/internal/service/payment"
	paymentv1 "boilerplates/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = "50052"

func main() {
	// --- хранилище НЕТ

	// --- бизнес-логика
	svc := paymentService.NewService()
	// --- транспорт
	apiV1 := paymentv1API.NewAPI(svc)
	// --- сеть
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer()

	// вторая проверка соответствия контракту
	paymentv1.RegisterPaymentServiceServer(s, apiV1)

	// Чтобы grpcurl / Postman видели список методов без .proto
	reflection.Register(s)

	log.Printf("gRPC server listening on :%s", grpcPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("serve: %v", err)
	}
}
