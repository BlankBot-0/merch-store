package main

import (
	"Merch/internal/grpc/server/merch_store"
	"Merch/internal/postgres/mock"
	pb "Merch/pkg/api/v1"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	//cfg := config.MustLoad()
	//ctx := context.Background()
	grpcServer := grpc.NewServer()

	//conn, err := postgres.Connect(ctx, cfg.Dsn)
	//if err != nil {
	//	log.Fatalf("db connection failed: %s", err)
	//}

	service := merch_store.NewService(merch_store.Deps{
		Repository: &mock.MerchPlatformMock{},
	})

	gatewayMux := runtime.NewServeMux()

	pb.RegisterMerchStoreServer(grpcServer, service)

	reflection.Register(grpcServer)

	port := 7000
	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen tcp: %s", err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Printf("running grpc server on port %d\n", port)
		if err := grpcServer.Serve(tcpListener); err != nil {
			log.Printf("failed to serve grpc server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()

		conn, err := grpc.DialContext(
			context.Background(),
			tcpListener.Addr().String(),
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed dial server: %s", err)
		}

		if err := pb.RegisterMerchStoreHandler(context.Background(), gatewayMux, conn); err != nil {
			log.Fatalf("failed to register gateway: %s", err)
		}

		mux := http.NewServeMux()
		mux.Handle("/", gatewayMux)
		httpPort := 7001
		gwServer := &http.Server{
			Addr:    fmt.Sprintf(":%d", httpPort),
			Handler: mux,
		}

		log.Printf("running http server on port %d\n", httpPort)

		if err := gwServer.ListenAndServe(); err != nil {
			log.Printf("failed to serve http server: %s", err)
		}
	}()

	wg.Wait()
	fmt.Println("server is finished")
}
