package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"dora-dev-test/consumer"
	"dora-dev-test/data"
	"dora-dev-test/generator"
	"dora-dev-test/postgres"
	"dora-dev-test/service"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	port = 8090
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	tickCh := make(chan data.Tick)
	wg.Go(func() {
		generator.GenerateTick(ctx, tickCh)
	})

	ds, err := postgres.NewDataStore("postgresql://postgres:mysecretpassword@localhost:5432/postgres")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to data store: %v", err))
	}
	con := consumer.NewConsumer(ds)
	wg.Go(func() {
		con.Start(ctx, tickCh) // saving ticks and candles
	})

	svc := service.NewService()

	mux := http.NewServeMux()
	// /ticks/{asset_id}?start={start}&end={end}&limit={limit}&offset={offset}
	mux.HandleFunc("GET /ticks/{asset_id}", svc.GetTicks)
	// /candles/{asset_id}?start={start}&end={end}&granularity={granularity}&limit={limit}
	mux.HandleFunc("GET /candles/{asset_id}", svc.GetCandles)
	// /health
	mux.HandleFunc("GET /health", svc.HealthCheck)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func() {
		log.Printf("Starting server on port %d", port)
		log.Fatal(server.ListenAndServe())
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	cancel()
	wg.Wait()

	log.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v, forcing shutdown", err)
	}

	log.Println("Graceful shutdown complete.")
}
