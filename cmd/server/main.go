package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "qfzack/go-web-starter/api/proto/server"
	"qfzack/go-web-starter/internal/server/config"
	"qfzack/go-web-starter/internal/server/container"
	"qfzack/go-web-starter/internal/server/router"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// service startup args
	configMode := flag.String("config-mode", "local", "config mode: local or remote")
	configPath := flag.String("config", "", "config file path for local, or URL for remote")
	httpPort := flag.Int("http-port", 8080, "HTTP server port")
	rpcPort := flag.Int("rpc-port", 50051, "gRPC server port")
	flag.Parse()

	// 1.load configurations
	cfg := loadConfig(*configMode, *configPath) //TODO: configure web server with cfg

	// 2.init HTTP server
	c := container.NewContainer()
	httpServer := setupHTTPServer(cfg, c, *httpPort)

	// 3.init gRPC server
	grpcServer := setupGRPCServer(cfg, c)

	// 4.servers startup
	if err := startServers(httpServer, grpcServer, *httpPort, *rpcPort); err != nil {
		log.Fatalf("Failed to start servers: %v", err)
	}
}

func loadConfig(configMode string, configPath string) *config.Config {
	var err error
	var cfg *config.Config

	switch configMode {
	case "local":
		cfg, err = loadLocalConfig(configPath)
	case "remote":
		if configPath == "" {
			log.Fatal("empty config path for remote mode")
		}
		cfg, err = fetchRemoteConfig(configPath)
	default:
		log.Fatalf("unsupported config mode: %s", configMode)
	}
	if err != nil {
		log.Fatalf("falied to load config: %v", err)
	}
	return cfg
}

// TODO: implement loadLocalConfig
func loadLocalConfig(configPath string) (*config.Config, error) {
	return &config.Config{}, nil
}

// TODO: implement fetchRemoteConfig
func fetchRemoteConfig(configPath string) (*config.Config, error) {
	return &config.Config{}, nil
}

func setupHTTPServer(cfg *config.Config, c *container.Container, httpPort int) *http.Server {
	if cfg.Enviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		router.RegisterUserRoutes(v1, c.UserHandler)
	}

	// health checkpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: r,
	}
}

func setupGRPCServer(cfg *config.Config, c *container.Container) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterMyRpcServer(s, c.RPCServer)

	if cfg.Enviroment != "production" {
		reflection.Register(s)
	}
	return s
}

func startServers(httpServer *http.Server, grpcServer *grpc.Server, httpPort int, rpcPort int) error {
	errChan := make(chan error, 2)

	// start http server
	go func() {
		log.Printf("Starting HTTP server on port: %d", httpPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("HTTP server failed: %w", err)
		}
	}()

	// start gRPC server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", rpcPort))
		if err != nil {
			errChan <- fmt.Errorf("failed to listen on gRPC port %d: %w", rpcPort, err)
			return
		}

		log.Printf("Starting gRPC server on port %d", rpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- fmt.Errorf("gRPC server failed: %w", err)
		}
	}()

	// wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		return err
	case sig := <-quit:
		log.Printf("Recieved signal: %v", sig)
		return gracefulShutDown(httpServer, grpcServer)
	}
}

func gracefulShutDown(httpServer *http.Server, grpcServer *grpc.Server) error {
	log.Println("Shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)
	errChan := make(chan error, 2)

	// shutdown http server
	go func() {
		defer wg.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			errChan <- fmt.Errorf("HTTP server shutdown failed: %w", err)
			return
		}
		log.Println("HTTP server stopped.")
		errChan <- nil
	}()

	// shutdown gRPC server
	go func() {
		defer wg.Done()
		done := make(chan struct{})
		go func() {
			grpcServer.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			log.Println("gRPC server stopped gracefully.")
		case <-ctx.Done():
			log.Println("gRPC server shutdown timeout, forcing stop...")
			grpcServer.Stop()
		}
		errChan <- nil
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	log.Println("All servers stopped.")
	return nil
}
