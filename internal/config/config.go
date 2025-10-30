package config

import (
	"os"
)

// LoadConfig loads minimal configuration from environment variables.
// Expand this to use a dotenv library or a proper config struct as needed.
func LoadConfig() {
	// Example: read DB path or port from env
	_ = os.Getenv("DATABASE_URL")
	_ = os.Getenv("PORT")
}
