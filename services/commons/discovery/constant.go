package discovery

import "os"

var (
	consulEndpoint string
)

func init() {
	consulEndpoint = os.Getenv("CONSUL_ENDPOINT")
	if consulEndpoint == "" {
		consulEndpoint = "consul:8500"
	}
}
