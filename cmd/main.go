package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/nullexp/finman-auth-service/internal/adapter/driven"
	grpcDriver "github.com/nullexp/finman-auth-service/internal/adapter/driver/grpc"
	authv1 "github.com/nullexp/finman-auth-service/internal/adapter/driver/grpc/proto/auth/v1"
	driver "github.com/nullexp/finman-auth-service/internal/adapter/driver/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	userServiceAddr := os.Getenv("USER_SERVICE_ADDR")
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

	log.Println("User service address: ", userServiceAddr)
	conn, err := establishGRPCConnection(userServiceAddr, 10)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create UserService client

	userService := driven.NewUserService(conn)
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

// establishGRPCConnection establishes a gRPC connection with retry mechanism
func establishGRPCConnection(serverAddr string, retryAttempts int) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	for i := 0; i < retryAttempts; i++ {
		conn, err = grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // insecure for test purpose
		if err == nil {
			log.Println("connected")
			return conn, nil
		}
		log.Printf("Failed to connect (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second) // Retry after 2 seconds
	}
	return nil, err
}
