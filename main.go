package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/jonasdacruz/lighthouses_aicontest/coms"
	"github.com/jonasdacruz/lighthouses_aicontest/engine"
	"google.golang.org/grpc"
)

const (
	timeoutForJoining = 1 * time.Second
)

func main() {
	listenAddress := flag.String("la", "", "game server listen address")
	flag.Parse()

	fmt.Println("Game server starting on", *listenAddress)

	if *listenAddress == "" {
		panic("addr is required")
	}

	lis, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(engine.UnaryLoggingInterceptor),
		grpc.StreamInterceptor(engine.StreamLoggingInterceptor),
	)

	ge := engine.NewGame()
	gs := engine.NewGameServer(ge)

	coms.RegisterGameServiceServer(grpcServer, gs)

	go func() {
		<-time.After(timeoutForJoining)
		grpcServer.Stop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println("players joining is closed, now the game starts!!!")

	ge.StartGame()

	fmt.Println("Game finished!!!")

}
