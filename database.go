package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func initDatabase() *gorp.DbMap {
	// Get the database connection information.
	dbUser := os.GetEnv("dbuser")
	dbPass := os.GetEnv("dbpass")
	dbAddr := os.GetEnv("dbaddr")
	dbPort := os.GetEnv("dbport")
	// 
	db, err := sql.Open("postgres", fmt.Sprintf("user='%s' password='$s' dbname=mcs_stats host='%s' port='%s' sslmode=disable", dbUser, dBPass, dbAddr, dbPort))
	if err != nil {
		log.Fatal(err)
	}
	// Create a gorp mapping.
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	// Add tables to the mapping.
	dbmap.AddTableWithName(Ping{}, "pings").SetKeys(false, "ping_id")
	return dbmap
}
