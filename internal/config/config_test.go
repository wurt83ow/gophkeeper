package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions_ParseFlags(t *testing.T) {
	// Backup original command line arguments and restore them after the test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Prepare test command line arguments
	testArgs := []string{
		"app", "-a", ":8080", "-d", "testdb_env", "-l", "info",
		"-n", "test777", "-j", "test_key_env", "-r", "/path/to/cert_env.pem", "-k", "/path/to/key_env.pem", "-s",
	}
	os.Args = testArgs

	// Prepare test environment variables
	os.Setenv("RUN_ADDRESS", ":8080")
	os.Setenv("DATABASE_URI", "testdb_env")
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("FILE_STORAGE_PATH", "test777")
	os.Setenv("JWT_SIGNING_KEY", "test_key_env")
	os.Setenv("HTTPS_CERT_FILE", "/path/to/cert_env.pem")
	os.Setenv("HTTPS_KEY_FILE", "/path/to/key_env.pem")
	os.Setenv("ENABLE_HTTPS", "true")

	// Create an instance of Options
	options := NewOptions()

	// Parse command line arguments
	options.ParseFlags()

	// Check if the options are correctly set
	assert.Equal(t, ":8080", options.RunAddr())
	assert.Equal(t, "testdb_env", options.DataBaseDSN())
	assert.Equal(t, "info", options.LogLevel())
	assert.Equal(t, "test777", options.FileStoragePath())
	assert.Equal(t, "test_key_env", options.JWTSigningKey())
	assert.Equal(t, "/path/to/cert_env.pem", options.HTTPSCertFile())
	assert.Equal(t, "/path/to/key_env.pem", options.HTTPSKeyFile())
	assert.True(t, options.EnableHTTPS())

	// Reset the environment variables
	os.Unsetenv("RUN_ADDRESS")
	os.Unsetenv("DATABASE_URI")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("FILE_STORAGE_PATH")
	os.Unsetenv("JWT_SIGNING_KEY")
	os.Unsetenv("HTTPS_CERT_FILE")
	os.Unsetenv("HTTPS_KEY_FILE")
	os.Unsetenv("ENABLE_HTTPS")
}

func TestOptions_DefaultValues(t *testing.T) {
	// Backup original command line arguments and restore them after the test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// No command line arguments or environment variables set

	// Create an instance of Options
	options := NewOptions()

	// Parse command line arguments
	options.ParseFlags()

	// Check if the options are set to default values
	assert.Equal(t, ":8080", options.RunAddr())
	assert.Equal(t, "testdb_env", options.DataBaseDSN())
	assert.Equal(t, "info", options.LogLevel())
	assert.Equal(t, "test777", options.FileStoragePath())
	assert.Equal(t, "test_key_env", options.JWTSigningKey())
	assert.Equal(t, "/path/to/cert_env.pem", options.HTTPSCertFile())
	assert.Equal(t, "/path/to/key_env.pem", options.HTTPSKeyFile())
	assert.True(t, options.EnableHTTPS())
}
