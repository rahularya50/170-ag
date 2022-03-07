//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	conf := gen.Config{
		Target:  "generated",
		Package: "170-ag/ent/generated",
	}
	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames("privacy", "schema/snapshot", "sql/upsert", "entql", "sql/modifier"),
	}

	if err := entc.Generate("./schema", &conf, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
