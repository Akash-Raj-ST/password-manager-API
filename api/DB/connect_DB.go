package DB

import (
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func ConnectDB() *gocql.Session{
	log.Println("Starting connection...",os.Getenv("DB_PRIVATE_DOMAIN"));
	cluster := gocql.NewCluster(os.Getenv("DB_PRIVATE_DOMAIN")) 
	
	cluster.ProtoVersion = 4
    cluster.DisableInitialHostLookup = true
    cluster.Consistency = gocql.Quorum
    cluster.CQLVersion = "3.4.5"
    cluster.IgnorePeerAddr = true
    cluster.DefaultIdempotence = true
    cluster.Timeout = time.Second * 30
    cluster.ConnectTimeout = time.Second * 30
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Session started successfully.")

	return session;
}