package main

import (
	"context"
	api "dora-dev-test/api/v1"
	"dora-dev-test/consumer"
	"dora-dev-test/data"
	"dora-dev-test/generator"
	"dora-dev-test/publisher"
	"dora-dev-test/redis"
	"dora-dev-test/service"
	"flag"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

const (
	port = 8090
)

var (
	seedBrokers = flag.String("brokers", "localhost:9092", "comma delimited list of seed brokers")
	topic       = flag.String("topic", "incoming_prices", "topic to consume from")
	group       = flag.String("group", "default", "group to consume within")
)

func main() {
	flag.Parse()

	tickCh := make(chan data.Tick)
	go generator.GenerateTick(context.Background(), tickCh)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(strings.Split(*seedBrokers, ",")...),
		kgo.ConsumerGroup(*group),
		kgo.ConsumeTopics(*topic),
	)

	if err != nil {
		panic(err)
	}

	ds := redis.NewDataStore()
	con := consumer.NewConsumer(client, ds)
	con.Start(context.Background())
	pub := publisher.NewTickPublisher(client, kgo.BasicLogger(os.Stderr, kgo.LogLevelInfo, nil))
	pub.Start(context.Background(), tickCh, "incoming_prices")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	api.RegisterDoraDevTestServiceServer(grpcServer, service.NewService(ds))
	log.Fatal(grpcServer.Serve(lis))
}
