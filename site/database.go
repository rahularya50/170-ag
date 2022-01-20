package site

import (
	ent "170-ag/ent/generated"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func GetEntClient() (*ent.Client, error) {
	env := os.Getenv("ENV")
	if env == "dev" {
		return ent.Open(
			"sqlite3",
			"file:app.db?&_fk=1",
		)
	} else {
		// This connects to the Cloud SQL Proxy on the same Pod, which has an encrypted connection to the actual db
		// This is why it is OK for SSL to be disabled!
		return ent.Open(
			"postgres",
			"host=localhost port=5432 user=autograder-web@formidable-gate-337712.iam dbname=autograder sslmode=disable",
		)
	}
}
