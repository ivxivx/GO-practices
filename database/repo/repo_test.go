package repo

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/ivxivx/go-practices/database"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.uber.org/goleak"
)

var dbURL string

func TestMain(m *testing.M) {
	ctx := context.Background()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	username := "pgtestuser"
	password := "pgtestpassword"

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16.3",
		Env: []string{
			"POSTGRES_USER=" + username,
			"POSTGRES_PASSWORD=" + password,
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	resource.Expire(120)

	dbURL = fmt.Sprintf("postgres://%s:%s@%s", username, password, net.JoinHostPort("localhost", resource.GetPort("5432/tcp")))

	//+ net.JoinHostPort(resource.GetBoundIP("26257/tcp"), resource.GetPort("26257/tcp")

	if err := pool.Retry(func() error {
		pgi, err := pginit.New(dbURL)
		if err != nil {
			return err
		}

		cPool, errCP := pgi.ConnPool(ctx)
		if errCP == nil {
			cPool.Close()
		}

		return errCP
	}); err != nil {
		log.Fatalf("Could not connect to database container: %s", err)
	}

	leak := flag.Bool("leak", false, "use leak detector")
	flag.Parse()

	if *leak {
		goleak.VerifyTestMain(m,
			goleak.IgnoreTopFunction("github.com/jackc/pgx/v5/pgxpool.(*Pool).backgroundHealthCheck"),
			goleak.IgnoreTopFunction("github.com/jackc/pgx/v5/pgxpool.(*Pool).triggerHealthCheck"),
			goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
			goleak.IgnoreTopFunction("net/http.(*persistConn).roundTrip"),
			goleak.IgnoreTopFunction("net/http.(*persistConn).writeLoop"),
			goleak.IgnoreTopFunction("internal/poll.runtime_pollWait"),
		)

		return
	}

	exitCode := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(exitCode)
}

func createDatabase(dbNamePrefix string) (*Repository, error) {
	ctx := context.Background()

	randomString := gonanoid.MustGenerate("abcdefghijklmnopqrstuvwxyz", 10)

	dbName := fmt.Sprintf("%s_%s", dbNamePrefix, randomString)

	// database name must be lowercase
	dbName = strings.ToLower(dbName)

	newDBURL := fmt.Sprintf("%s/%s", dbURL, dbName+"?sslmode=disable")

	// create database
	repo, err := NewRepository(ctx, dbURL)
	if err != nil {
		log.Printf("could not connect to db: %v", err)

		return nil, err
	}

	if _, err := repo.pool.Exec(ctx, "CREATE DATABASE "+dbName); err != nil {
		log.Printf("could not create test database %s: %v", dbName, err)

		return nil, err
	}

	if _, _, err := database.Migrate(newDBURL); err != nil {
		log.Printf("could not migrate: %v", err)

		return nil, err
	}

	repo2, err := NewRepository(ctx, newDBURL)
	if err != nil {
		log.Printf("could not connect to db: %v", err)

		return nil, err
	}

	return repo2, nil
}
