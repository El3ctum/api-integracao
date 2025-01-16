package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Service represents a service that interacts with a database
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type couchbaseService struct {
	db     *gocb.Cluster
	bucket *gocb.Bucket
}

const (
	ConnectionString = "DB_CONN_STR"
	UsernameKey      = "DB_USERNAME"
	PasswordKey      = "DB_PASSWORD"
	BucketName       = "DB_BUCKET_NAME"
	ScopeName        = "inventory"
)

func New() Service {
	// gocb.SetLogger(gocb.VerboseStdioLogger())

	connectionString := getEnvVar(ConnectionString)
	username := getEnvVar(UsernameKey)
	password := getEnvVar(PasswordKey)

	cluster, err := gocb.Connect("couchbase://"+connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Check if the specified scope exists
	if !checkScopeExists(cluster, BucketName, ScopeName) {
		fmt.Println("Inventory scope does not exist in the bucket. Ensure that you have the inventory scope in your travel-sample bucket.")
		os.Exit(1)
	}

	bucket := cluster.Bucket(BucketName)

	err = bucket.WaitUntilReady(time.Second*5, nil)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance := &couchbaseService{
		db:     cluster,
		bucket: bucket,
	}

	return dbInstance
}

func (s *couchbaseService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	// We'll ping the KV nodes in our cluster.
	pings, err := s.bucket.Ping(&gocb.PingOptions{
		ReportID:     "my-report",
		ServiceTypes: []gocb.ServiceType{gocb.ServiceTypeKeyValue},
		Context:      ctx,
	})
	if err != nil {
		panic(err)
	}

	for service, pingReports := range pings.Services {
		if service != gocb.ServiceTypeKeyValue {
			panic("we got a service type that we didn't ask for!")
		}

		for _, pingReport := range pingReports {
			if pingReport.State != gocb.PingStateOk {
				fmt.Printf(
					"Node %s at remote %s is not OK, error: %s, latency: %s\n",
					pingReport.ID, pingReport.Remote, pingReport.Error, pingReport.Latency.String(),
				)
			} else {
				fmt.Printf(
					"Node %s at remote %s is OK, latency: %s\n",
					pingReport.ID, pingReport.Remote, pingReport.Latency.String(),
				)
			}
		}
	}

	b, err := pings.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ping report JSON: %s\n", string(b))
	return map[string]string{
		"pingReport": string(b),
	}
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *couchbaseService) Close() error {
	log.Printf("Disconnected from database: %s", "BucketName")
	return s.db.Close(&gocb.ClusterCloseOptions{})
}

// GetScope returns a scope for the specified cluster, bucket, and scope name.
func GetScope(cluster *gocb.Cluster) *gocb.Scope {
	bucket := cluster.Bucket(BucketName)
	scope := bucket.Scope(ScopeName)
	return scope
}

// Helper function to retrieve an environment variable
func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" && (key == "DB_USERNAME" || key == "DB_PASSWORD" || key == "DB_CONN_STR") {
		fmt.Printf("Environment variable %s is empty.\n", key)
		os.Exit(1)
	}
	return value
}

// Function to check if a scope exists in a bucket
func checkScopeExists(cluster *gocb.Cluster, bucketName, scopeName string) bool {
	bucket := cluster.Bucket(bucketName)
	// Fetch all scopes in the bucket
	scopes, err := bucket.Collections().GetAllScopes(nil)
	if err != nil {
		fmt.Println("Error fetching scopes in the cluster. Ensure that the travel sample bucket exists in the cluster.")
		return false
	}

	// Check if the specified scope exists in the list of scopes
	for _, s := range scopes {
		if s.Name == scopeName {
			return true
		}
	}

	// Return false if the scope doesn't exist
	return false
}
