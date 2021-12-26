package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v4/pgxpool"
)

// pg connection pool (thread-safe)
var Pool *pgxpool.Pool

// snowflake node (thread-safe)
var snowflakeNode *snowflake.Node

// This function will initialize everything needed for
// the operation of the database
func Init() {
	// initialize psql pool
	pool, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DATABASE"),
	))
	Pool = pool

	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	log.Println("Connected to the database.")

	// initialize snowflake generator
	node_id, err := strconv.Atoi(os.Getenv("SNOWFLAKE_NODE_ID"))
	if err != nil {
		log.Fatalf("Error parsing SNOWFLAKE_NODE_ID: %s\n", err)
		return
	}

	snowflake.Epoch = 1609459200000 // Jan 1 2021
	node, err := snowflake.NewNode(int64(node_id))
	if err != nil {
		log.Fatalf("Error initializing snowflake: %s\n", err)
		return
	}

	snowflakeNode = node
	log.Printf("Initialized snowflake node %d\n", node_id)
}

// This function will return a Snowflake ID
// to be used as primary key for tables that require it
func GenerateId() int64 {
	return snowflakeNode.Generate().Int64()
}
