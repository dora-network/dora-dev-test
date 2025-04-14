package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"

	api "dora-dev-test/api/v1"
	"dora-dev-test/consumer"
	"dora-dev-test/data"
	"dora-dev-test/generator"
	"dora-dev-test/publisher"
	"dora-dev-test/redisds"
	"dora-dev-test/service"

	"github.com/redis/go-redis/v9"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
)

const (
	port = 8090
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))

	tickCh := make(chan data.Tick)
	slog.Debug("starting tick generator")
	go generator.GenerateTick(context.Background(), tickCh)
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
	)
	if err != nil {
		panic(err)
	}

	slog.Debug("building redis client")
	var redisClient *redis.Client
	{
		redisClient = redis.NewClient(&redis.Options{Addr: ":6379"})
	}
	slog.Debug("redis client built")

	con := consumer.NewConsumer(client, redisds.NewDataStore(redisClient))
	con.Start(context.Background())
	pub := publisher.NewTickPublisher(client, kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil))
	pub.Start(context.Background(), tickCh, "incoming_prices")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	api.RegisterDoraDevTestServiceServer(grpcServer, service.NewService())
	log.Fatal(grpcServer.Serve(lis))
}
