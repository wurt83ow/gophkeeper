package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	flagRunAddr, flagDataBaseDSN, flagLogLevel,
	flagHTTPSCertFile, flagHTTPSKeyFile string
	flagEnableHTTPS bool
}

func NewOptions() *Options {
	return new(Options)
}

// parseFlags handles command line arguments
// and stores their values in the corresponding variables.
func (o *Options) ParseFlags() {
	regStringVar(&o.flagRunAddr, "a", ":8080", "address and port to run server")
	regStringVar(&o.flagDataBaseDSN, "d", "", "")
	regStringVar(&o.flagLogLevel, "l", "info", "log level")
	regStringVar(&o.flagHTTPSCertFile, "r", "", "path to https cert file")
	regStringVar(&o.flagHTTPSKeyFile, "k", "", "path to https key file")
	regBoolVar(&o.flagEnableHTTPS, "s", false, "enable https")

	// parse the arguments passed to the server into registered variables
	flag.Parse()

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		o.flagRunAddr = envRunAddr
	}

	if envDataBaseDSN := os.Getenv("DATABASE_URI"); envDataBaseDSN != "" {
		o.flagDataBaseDSN = envDataBaseDSN
	}

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		o.flagLogLevel = envLogLevel
	}

	if envHTTPSCertFile := os.Getenv("HTTPS_CERT_FILE"); envHTTPSCertFile != "" {
		o.flagHTTPSCertFile = envHTTPSCertFile
	}

	if envHTTPSKeyFile := os.Getenv("HTTPS_KEY_FILE"); envHTTPSKeyFile != "" {
		o.flagHTTPSKeyFile = envHTTPSKeyFile
	}

	if envEnableHTTPS := os.Getenv("ENABLE_HTTPS"); envEnableHTTPS != "" {
		// Assuming "ENABLE_HTTPS" should be a boolean value
		enableHTTPS, err := strconv.ParseBool(envEnableHTTPS)
		if err == nil {
			o.flagEnableHTTPS = enableHTTPS
		} else {
			// Handle the error (failed to parse as boolean)
			fmt.Println("Failed to parse ENABLE_HTTPS as a boolean value:", err)
		}
	}

}

func (o *Options) RunAddr() string {
	return getStringFlag("a")
}

func (o *Options) DataBaseDSN() string {
	return getStringFlag("d")
}

func (o *Options) LogLevel() string {
	return getStringFlag("l")
}

// HTTPSCertFile returns the path to HTTPS cert file.
func (o *Options) HTTPSCertFile() string {
	return getStringFlag("r")
}

// HTTPSCertFile returns the path to HTTPS key file.
func (o *Options) HTTPSKeyFile() string {
	return getStringFlag("k")
}

// EnableHTTPS returns whether HTTPS is enabled.
func (o *Options) EnableHTTPS() bool {
	return getBoolFlag("s")
}

func regStringVar(p *string, name string, value string, usage string) {
	if flag.Lookup(name) == nil {
		flag.StringVar(p, name, value, usage)
	}
}

// regBoolVar registers a bool flag with the specified name, default value, and usage string.
func regBoolVar(p *bool, name string, value bool, usage string) {
	if flag.Lookup(name) == nil {
		flag.BoolVar(p, name, value, usage)
	}
}
func getStringFlag(name string) string {
	return flag.Lookup(name).Value.(flag.Getter).Get().(string)
}

// getBoolFlag retrieves the bool value of the specified flag.
func getBoolFlag(name string) bool {
	return flag.Lookup(name).Value.(flag.Getter).Get().(bool)
}

// GetAsString reads an environment or returns a default value.
func GetAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
