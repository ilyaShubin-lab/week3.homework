package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	orderV1API "boilerplates/order/internal/api/order/v1"
	inventoryClient "boilerplates/order/internal/client/grpc/inventory/v1"
	paymentClient "boilerplates/order/internal/client/grpc/payment/v1"
	"boilerplates/order/internal/migrator"
	orderRepository "boilerplates/order/internal/repository/order"
	orderService "boilerplates/order/internal/service/order"
	orderV1 "boilerplates/shared/pkg/openapi/order/v1"
	inventoryv1 "boilerplates/shared/pkg/proto/inventory/v1"
	paymentv1 "boilerplates/shared/pkg/proto/payment/v1"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	dsn := getEnv("POSTGRES_DSN",
		"postgres://order-service-user:order-service-password@localhost:5432/order-service?sslmode=disable",
	)

	migrationsDir := getEnv("MIGRATIONS_DIR", "order/migrations")

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("open db for migrations: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("ping db: %v", err)
	}

	err = migrator.NewMigrator(sqlDB, migrationsDir).Up()
	if err != nil {
		log.Fatalf("migrations: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {

		log.Printf("close migrations db: %v", err)
	}

	log.Println("migrations applied")

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("create pgx pool: %v", err)
	}
	defer pool.Close()

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("ping pgx pool: %v", err)
	}

	inventoryConn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to inventory: %v", err)
	}
	defer func() {
		if closeErr := inventoryConn.Close(); closeErr != nil {
			log.Printf("failed to close inventory connection: %v", closeErr)
		}
	}()

	paymentConn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to payment: %v", err)
	}
	defer func() {
		if closeErr := paymentConn.Close(); closeErr != nil {
			log.Printf("failed to close payment connection: %v", closeErr)
		}
	}()

	invClient := inventoryClient.NewClient(
		inventoryv1.NewInventoryServiceClient(inventoryConn),
	)
	payClient := paymentClient.NewClient(
		paymentv1.NewPaymentServiceClient(paymentConn),
	)

	repo := orderRepository.NewRepository(pool)
	service := orderService.NewService(repo, invClient, payClient)
	api := orderV1API.NewAPI(service)

	orderServer, err := orderV1.NewServer(api)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: orderServer,
	}

	log.Println("HTTP server listening on :8080")

	if err = httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}
