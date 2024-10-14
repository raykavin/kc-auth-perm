package web

import (
	"fmt"
	"keycloak-example/confs"
	"keycloak-example/controllers"
	"keycloak-example/middlewares"
	"log"
	"net/http"
	"strings"
	"time"
)

// Setup initializes the OIDC provider and OAuth2 configuration, then starts the web server.
func Setup() error {
	config := confs.Default.Application
	webConfig := config.Web

	log.Println("setting up server routes...")
	setupRoutes()

	logMsg := fmt.Sprintf("server is now running on port :%d", webConfig.Listen)

	// Start the server with or without SSL based on the config
	server := createServer(webConfig.Listen)
	if webConfig.Ssl {
		logMsg += " using SSL"
		log.Println(logMsg)
		return startSSLServer(server, webConfig.Crt, webConfig.Key)
	}

	log.Println(logMsg)
	return server.ListenAndServe()
}

// createServer sets up the HTTP server with timeouts.
func createServer(port uint) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func startSSLServer(server *http.Server, certificate, key string) error {
	if strings.TrimSpace(certificate) == "" || strings.TrimSpace(key) == "" {
		return fmt.Errorf("cannot start server: invalid certificate or key path")
	}
	return server.ListenAndServeTLS(certificate, key)
}

// setupRoutes registers the HTTP handlers for authorization and callback.
func setupRoutes() {
	personController := controllers.NewPersonController()
	http.HandleFunc("GET /api/people", middlewares.IsAuthorized("khambalia-admin", personController.List))
}
