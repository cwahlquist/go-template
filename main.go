package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "api/go"

	h "go-template/handler"
	m "go-template/model"
	s "go-template/service"

	"google.golang.org/grpc"

	// needed for postman proxy
	_ "github.com/jnewmano/grpc-json-proxy/codec"
)

var (
	port = flag.Int("port", 31400, "The server grpc port")
)

func main() {
	// get env vars
	flag.Parse()

	// start listening tcp:host:port
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	// create db interface, in this example we are using a map instead of a actual database session
	session := make(map[string]m.Todo)

	// inject dependencies

	// initialize service layer
	srv := s.NewService()

	hnd := h.NewHandler(srv)

	// create grpc server and apply middleware
	grpcServer := grpc.NewServer()

	// register missions PB with grpcServer
	pb.RegisterGoTemplateServiceServer(grpcServer, hnd)

	// create http server
	http.HandleFunc("/health", healthHandler)
	go func() {
		if err := http.ListenAndServe(":80", nil); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	log.Printf("Service started on 0.0.0.0:%d", *port)

	// start gRPC server
	err = grpcServer.Serve(listen)
	if err != nil {
		panic("gRpc Server failed to start")
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}
