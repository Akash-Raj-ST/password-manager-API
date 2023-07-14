package DB

import (
	"fmt"
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

	keyspaceQuery := fmt.Sprintf(`
		CREATE KEYSPACE IF NOT EXISTS %s
		WITH REPLICATION = {
			'class' : 'SimpleStrategy',
			'replication_factor' : 1
		};`,os.Getenv("DB_KEYSPACE"))

	err = session.Query(keyspaceQuery).Exec()

	if err != nil {
		log.Fatal("Error creating keyspace:", err)
	}

	log.Println("Keyspace created successfully")

	// Use the newly created keyspace
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")
	log.Println("keyspace ",os.Getenv("DB_KEYSPACE"))

	session, err = cluster.CreateSession()

	if err != nil {
		log.Println("Error connecting to mykeyspace:", err)
	}
	
	
	createUserTable(session);
	
	log.Println("Session started successfully.")

	return session;
}


func createUserTable(session *gocql.Session){
	var tableQuery string;
	var err error;

	// Define the user-defined type (UDT)
	err = session.Query(fmt.Sprintf(`
		CREATE TYPE IF NOT EXISTS %s.itemUDT (
			key TEXT,
			value TEXT,
			is_secret BOOLEAN
		)`,os.Getenv("DB_KEYSPACE"))).Exec();

	if err != nil {
		log.Fatal("Error creating user-defined type:", err)
	}

	log.Println("User-defined type created successfully")

	err = session.Query(fmt.Sprintf(`
		CREATE TYPE IF NOT EXISTS %s.dataUDT (
			username TEXT,
			password TEXT,
			items List<FROZEN<itemUDT>>
		)`,os.Getenv("DB_KEYSPACE"))).Exec();

	if err != nil {
		log.Fatal("Error creating user-defined type:", err)
	}

	log.Println("User-defined type created successfully")
	
	// Create a user table
	tableQuery = `
		CREATE TABLE IF NOT EXISTS user(
			userID UUID,
			username TEXT PRIMARY KEY,
			password TEXT,
			data FROZEN<dataUDT>,
		);`

	err = session.Query(tableQuery).Exec()

	if err != nil {
		log.Fatal("Error creating user table:", err)
	}

	log.Println("User table created successfully")
}

