package cassandra

import (
	"time"

	"github.com/gocql/gocql"
)

type Client struct {
	session *gocql.Session
}

func NewClient(hosts []string, keyspace string) (*Client, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &Client{session}, nil
}
