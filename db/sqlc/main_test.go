package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5433/simple_bank?sslmode=disable"
)

// build db connection
var testStore *Store

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(connPool)
	db := testStore.db
	if db == nil {
		log.Fatalf("error..")
	}
	fmt.Println("good! test pass")
	os.Exit(m.Run())
}
