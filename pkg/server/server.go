package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"rest-api/config"
	"rest-api/pkg/middlewares"
	"rest-api/pkg/repository/postgres"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, ur *postgres.UserRepository) *Server {
	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middlewares.RequestLoggerMiddleware(),
	)

	RegisterHTTPEndpoints(r, ur)

	// handler := Handler.NewHandler(r)

	var server *http.Server

	cert, err := tls.LoadX509KeyPair(cfg.CertPath, cfg.KeyPath)
	if err != nil {
		log.Fatal("Cannot load tls certificates: ", err)
	}

	server = &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	return &Server{
		httpServer: server,
	}
}

func StartServer(cfg *config.Config, s *Server) error {
	if cfg.Tls == `1` {
		log.Printf("Server running in HTTPS mode on port: %v", cfg.Port)
		err := s.httpServer.ListenAndServeTLS("", "")
		return err
	} else {
		log.Printf("Server running in HTTP mode on port: %v", cfg.Port)
		err := s.httpServer.ListenAndServe()
		return err
	}
}
