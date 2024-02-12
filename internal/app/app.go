package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	authz "github.com/wurt83ow/gophkeeper-server/internal/authorization"
	"github.com/wurt83ow/gophkeeper-server/internal/bdkeeper"
	"github.com/wurt83ow/gophkeeper-server/internal/config"
	"github.com/wurt83ow/gophkeeper-server/internal/controllers"
	"github.com/wurt83ow/gophkeeper-server/internal/logger"
	"github.com/wurt83ow/gophkeeper-server/internal/middleware"
	"github.com/wurt83ow/gophkeeper-server/internal/storage"
)

// Server represents the application server.
type Server struct {
	srv *http.Server
	ctx context.Context
}

// NewServer creates a new Server instance.
func NewServer(ctx context.Context) *Server {
	server := new(Server)
	server.ctx = ctx

	return server
}

// Serve starts the server.
func (server *Server) Serve() {
	// Create and initialize a new option instance
	option := config.NewOptions()
	option.ParseFlags()

	// Get a new logger
	nLogger, err := logger.NewLogger(option.LogLevel())
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize the keeper instance
	keeper, err := initializeKeeper(option.DataBaseDSN, nLogger)
	if err != nil {
		log.Fatalln(err)
	}
	defer keeper.Close()

	// Initialize the storage instance
	memoryStorage := initializeStorage(keeper, nLogger)

	authz := authz.NewJWTAuthz(option.JWTSigningKey(), nLogger)

	// Create a new controller to process incoming requests
	baseController := initializeBaseController(memoryStorage, option, nLogger, authz)

	// Create an instance of ChiServerOptions with your middleware
	options := controllers.ChiServerOptions{
		Middlewares: []controllers.MiddlewareFunc{
			authz.JWTAuthzMiddleware(memoryStorage, nLogger),
		},
	}

	// Create a handler with options
	genHandler := controllers.HandlerWithOptions(baseController, options)

	// Get a middleware for logging requests
	reqLog := middleware.NewReqLog(nLogger)

	// Create router and mount routes
	r := chi.NewRouter()
	r.Use(reqLog.RequestLogger)
	r.Mount("/", genHandler)

	// Configure and start the server
	startServer(server, r, option.RunAddr(), option.EnableHTTPS(),
		option.HTTPSCertFile(), option.HTTPSKeyFile())
}

func initializeKeeper(dataBaseDSN func() string, logger *logger.Logger) (*bdkeeper.BDKeeper, error) {
	return bdkeeper.NewBDKeeper(dataBaseDSN, logger, nil)
}

func initializeStorage(keeper storage.Keeper, logger *logger.Logger) *storage.MemoryStorage {
	if keeper == nil {
		return nil
	}

	return storage.NewMemoryStorage(keeper, logger)
}

func initializeBaseController(storage *storage.MemoryStorage, options *config.Options,
	logger *logger.Logger, authz *authz.JWTAuthz,
) *controllers.BaseController {
	return controllers.NewBaseController(storage, options, logger, authz)
}

func startServer(server *Server, router chi.Router, address string,
	enableHTTPS bool, HTTPSCertFile, HTTPSKeyFile string) {
	const (
		oneMegabyte = 1 << 20
		readTimeout = 3 * time.Second
	)

	server.srv = &http.Server{
		Addr:              address,
		Handler:           router,
		ReadHeaderTimeout: readTimeout,
		WriteTimeout:      readTimeout,
		IdleTimeout:       readTimeout,
		MaxHeaderBytes:    oneMegabyte, // 1 MB
	}

	log.Printf("Starting server at %s\n", address)

	// Start the HTTP/HTTPS server
	var err error
	if enableHTTPS {
		log.Printf("HTTPS enabled")
		err = server.srv.ListenAndServeTLS(HTTPSCertFile, HTTPSKeyFile)
	} else {
		log.Printf("HTTPS disabled")
		err = server.srv.ListenAndServe()
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}

}

// Shutdown gracefully shuts down the server.
func (server *Server) Shutdown() {
	log.Printf("server stopped")

	const shutdownTimeout = 5 * time.Second
	ctxShutDown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

	defer cancel()

	if err := server.srv.Shutdown(ctxShutDown); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server Shutdown Failed:%s", err)
		}
	}

	log.Println("server exited properly")
}
