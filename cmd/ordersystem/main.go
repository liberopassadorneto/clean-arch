package main

import (
	"fmt"
	"github.com/liberopassadorneto/clean-arch/db_connection"
	"github.com/liberopassadorneto/clean-arch/migrations"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/liberopassadorneto/clean-arch/configs"
	"github.com/liberopassadorneto/clean-arch/internal/event/handler"
	"github.com/liberopassadorneto/clean-arch/internal/infra/graph"
	"github.com/liberopassadorneto/clean-arch/internal/infra/grpc/pb"
	"github.com/liberopassadorneto/clean-arch/internal/infra/grpc/service"
	"github.com/liberopassadorneto/clean-arch/internal/infra/web/webserver"
	"github.com/liberopassadorneto/clean-arch/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := db_connection.ConnectDB(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Run migrations
	migrations.CreateOrdersTable(db)
	migrations.InsertSampleData(db)

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrdersUseCase := NewListOrdersUseCase(db)

	webServer := webserver.NewWebServer(config.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webServer.AddHandler("/order", webOrderHandler.Create)
	webServer.AddHandler("/order", webOrderHandler.List)
	fmt.Println("Starting web server on port", config.WebServerPort)
	go webServer.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *listOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", config.GraphQLServerPort)
	http.ListenAndServe(":"+config.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
