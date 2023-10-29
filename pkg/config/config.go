package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

// AppConfig struct for webapp config
type AppConfig struct {
	Server struct {
		// Host is the local machine IP address to bind the HTTP server to
		Host	string	`yaml:"host"`
		// Port is the local machine TCP Port to bind the HTTP server to
		Port	string	`yaml:"posrt"`
		Timeout	struct {
			// Server is the general server timeout to use for graceful shutdown
			Server	time.Duration	`yaml:"timeout"`
			// Read is the amount of time to wait until an HTTP server read operation is cancelled
			Read	time.Duration	`yaml:"read"`
			// Write is the amount of time to wait until an HTTP server write operation is cancelled
			Write	time.Duration	`yaml:"write"`
			// Idle is the amount of time to wait until an idle HTTP session is closed
			Idle	time.Duration	`yaml:"idle"`
		}	`yaml:"timeout"`
	}	`yaml:"server"`
	// Database is the local machine Postgres credentials to use for database connection
	Database struct {
		// PSQL local machine user
		DbUser	string	`yaml:"dbuser"`
		// PSQL local machine password
		DbPwd	string	`yaml:"dbpwd"`
		// PSQL local machine IP address to bind the HTTP server
		DbHost	string	`yaml:"dbhost"`
		// PSQL local machine port to bind the HTTP server
		DbPort	int		`yaml:"dbport"`
		// PSQL local machine database name
		DbName	string	`yaml:"dbname"`
		SSLmode string	`yaml:"sslmode"`
		MaxOpenConnections	int	`yaml:"maxopenconns"`
		MaxIdleConnections	int	`yaml:"maxidleconns"`
	}	`yaml:"database"`
}

// LoadConfig return a new decoded AppConfig struct
func LoadConfig(configpath string) (cfg *AppConfig, err error) {
	viper.AddConfigPath(configpath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}

// Listen will kick-start the http server
func (cfg *AppConfig) Listen(){
	var lisChan = make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout.Server * time.Second)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr: ":" + cfg.Server.Port,
		ReadTimeout: cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: cfg.Server.Timeout.Write * time.Second,
		IdleTimeout: cfg.Server.Timeout.Idle * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(lisChan, os.Interrupt, syscall.SIGTSTP, syscall.SIGTERM, syscall.SIGHUP)

	// Alert the user that the server is starting
    log.Printf("[init] Server is starting on %s...🚀\n", server.Addr)

	// Run the server on a new goroutine
	// go func() {
	// 	if err := server.ListenAndServe(); err != nil {
	// 		if err == http.ErrServerClosed {
				
	// 		} else {
	// 			log.Fatalf("[ERROR] Server failed to start due to err: %v", err)
	// 		}
	// 	}
	// }()

	// Block on this channel listeninf for those previously defined syscalls assign
    // to variable so we can let the user know why the server is shutting down
    interruption := <-lisChan

    // If we get one of the pre-prescribed syscalls, gracefully terminate the server
    // while alerting the user
    log.Printf("[INFO] Server is shutting down due to %+v\n", interruption)
	
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("[INFO] Server was unable to gracefully shutdown due to err: %+v", err)
    }
}
