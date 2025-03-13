package utils

import (
	"bank/configs"
	"fmt"
)

// ConnectionUrlBuilder builds connection URLs for services and databases.
func ConnectionUrlBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		// Fiber app URL
		url = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)

	case "postgresql":
		// PostgreSQL connection string
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.PostgreSQL.Host,
			cfg.PostgreSQL.Port,
			cfg.PostgreSQL.Username,
			cfg.PostgreSQL.Password,
			cfg.PostgreSQL.Database,
			cfg.PostgreSQL.SSLMode,
		)

	case "mysql":
		// MySQL connection string
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=%t",
			cfg.MySql.Username,
			cfg.MySql.Password,
			cfg.MySql.Host,
			cfg.MySql.Port,
			cfg.MySql.Database,
			cfg.MySql.ParseTime,
		)

	default:
		// Unknown service or database type
		return "", fmt.Errorf("error: connection URL builder doesn't support '%s'", stuff)
	}

	fmt.Println("URL", url)
	return url, nil
}
