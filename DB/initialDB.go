package DB

import (
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func Connect() *gocql.Session{
	log.Println("Starting connection...");
	cluster := gocql.NewCluster(os.Getenv("DB_PRIVATE_DOMAIN")) 
	
	cluster.ProtoVersion = 4
    cluster.DisableInitialHostLookup = true
    cluster.Consistency = gocql.Quorum
    cluster.CQLVersion = "3.4.5"
    cluster.IgnorePeerAddr = true
    cluster.DefaultIdempotence = true
    cluster.Timeout = time.Second * 30
    cluster.ConnectTimeout = time.Second * 30


	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	keyspaceQuery := `
		CREATE KEYSPACE IF NOT EXISTS password_api
		WITH REPLICATION = {
			'class' : 'SimpleStrategy',
			'replication_factor' : 1
		};`

	err = session.Query(keyspaceQuery).Exec()

	if err != nil {
		log.Fatal("Error creating keyspace:", err)
	}

	log.Println("Keyspace created successfully")

	// Use the newly created keyspace
	cluster.Keyspace = "password_api"
	session, err = cluster.CreateSession()

	if err != nil {
		log.Println("Error connecting to mykeyspace:", err)
	}
	defer session.Close()

	// Create a user table
	tableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name TEXT,
			email TEXT
		);`
	err = session.Query(tableQuery).Exec()
	if err != nil {
		log.Println("Error creating user table:", err)
	}

	log.Println("User table created successfully")

	log.Println("Session started executed successfully.")

	return session;
}
