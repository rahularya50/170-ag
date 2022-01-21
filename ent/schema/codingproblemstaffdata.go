package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CodingProblemStaffData holds the schema definition for the CodingProblemStaffData entity.
type CodingProblemStaffData struct {
	ent.Schema
}

// Fields of the CodingProblemStaffData.
func (CodingProblemStaffData) Fields() []ent.Field {
	return []ent.Field{
		field.Text("input"),
	}
}

// Edges of the CodingProblemStaffData.
func (CodingProblemStaffData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("coding_problem", CodingProblem.Type).Required().Unique().Annotations(entgql.Bind()),
	}
}
