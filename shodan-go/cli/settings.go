package cli

import (
	"os"
	"path/filepath"
)

// SHODAN_CONFIG_DIR matches Python shodan.cli.settings
func ConfigDir() string {
	home, _ := os.UserHomeDir()
	legacy := filepath.Join(home, ".shodan")
	if st, err := os.Stat(legacy); err == nil && st.IsDir() {
		return legacy
	}
	return filepath.Join(home, ".config", "shodan")
}

// COLORIZE_FIELDS from Python
var ColorizeFields = map[string]string{
	"ip_str":    "green",
	"port":      "yellow",
	"data":      "white",
	"hostnames": "magenta",
	"org":       "cyan",
	"vulns":     "red",
}
