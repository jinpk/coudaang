package main

import (
	"context"
	"io/fs"
	"log"
	"net"
	"net/http"

	"github.com/coudaang/coudaang/docs"
	authv1 "github.com/coudaang/coudaang/packages/api/auth/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := do(); err != nil {
		log.Fatal(err.Error())
	}
}

func do() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServer := grpc.NewServer()

	authAPIServer := authv1.NewAuthAPIServer()
	authv1.RegisterAuthAPIServer(grpcServer, authAPIServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	go func() {
		log.Fatal(grpcServer.Serve(listener))
	}()

	grpcConn, err := grpc.DialContext(
		ctx,
		listener.Addr().String(),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	gwMux := runtime.NewServeMux()
	if err != nil {
		return err
	}

	if err := authv1.RegisterAuthAPIHandler(ctx, gwMux, grpcConn); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwMux)

	swagger, err := fs.Sub(docs.SwaggerUI, "swagger-ui")
	if err != nil {
		return err
	}

	mux.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.FS(swagger))))

	return http.ListenAndServe(":8081", mux)
}
