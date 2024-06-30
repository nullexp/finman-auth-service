package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/nullexp/finman-auth-service/internal/adapter/driven"
	grpcDriver "github.com/nullexp/finman-auth-service/internal/adapter/driver/grpc"
	authv1 "github.com/nullexp/finman-auth-service/internal/adapter/driver/grpc/proto/auth/v1"
	driver "github.com/nullexp/finman-auth-service/internal/adapter/driver/service"
	"github.com/nullexp/finman-auth-service/internal/port/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	log.Println("Starting the server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpireMinute := os.Getenv("JWT_EXPIRE_MINUTE")
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")
	duration, err := strconv.Atoi(jwtExpireMinute)
	if err != nil {
		log.Fatal("duration should be a valid number")
	}

	addr := fmt.Sprintf("%s:%v", ip, port)
	// Create a TCP listener
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	tokenService := driven.NewTokenService(jwtSecret, time.Duration(int(time.Minute)*duration))
	userService := driven.NewMockUserService()
	userService.SetGetUserResponse(&model.GetUserResponse{
		IsAdmin: true,
		Id:      uuid.New().String(),
	}, nil)

	authService := driver.NewAuthService(userService, tokenService)
	service := grpcDriver.NewAuthService(authService)

	// Register the Greeter service
	authv1.RegisterAuthServiceServer(s, service)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// Log and start the server
	log.Printf("gRPC server listening on %s", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
