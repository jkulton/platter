package options

import (
	"flag"
	"log"
	"os"
	"strings"
)

type Options struct {
	Port      string
	Directory string
	Address   string
	Auth      string
}

// ParseOptions ...
func ParseOptions() Options {
	const (
		DefaultPort      = "8000"
		DefaultDirectory = "."
	)

	port := flag.String("p", DefaultPort, "port used for serving files")
	directory := flag.String("d", DefaultDirectory, "directory of files to serve")
	address := flag.String("a", "0.0.0.0", "address to serve on")
	auth := flag.String("auth", "", "(optional) user:pass formatted authorization (for basic auth)")
	flag.Parse()

	// Validate format of -auth argument, exit if provided but not valid
	if *auth != "" {
		split := strings.Split(*auth, ":")
		user := split[0]
		pass := split[len(split)-1]

		if user == "" || pass == "" || len(split) != 2 {
			log.Fatal("auth must be a string formatted as user:pass")
			os.Exit(1)
		}
	}

	return Options{
		Port:      *port,
		Directory: *directory,
		Address:   *address,
		Auth:      *auth,
	}
}
