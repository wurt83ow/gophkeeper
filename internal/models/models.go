package models

// Key is an alias for string and represents a key used in various contexts.
type Key string

// Response describes the server's response.
type Response struct {
	Result string `json:"result"`
}
