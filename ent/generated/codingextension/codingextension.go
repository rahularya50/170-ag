// Code generated by entc, DO NOT EDIT.

package codingextension

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the codingextension type in the database.
	Label = "coding_extension"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// EdgeStudent holds the string denoting the student edge name in mutations.
	EdgeStudent = "student"
	// EdgeCodingProblem holds the string denoting the coding_problem edge name in mutations.
	EdgeCodingProblem = "coding_problem"
	// Table holds the table name of the codingextension in the database.
	Table = "coding_extensions"
	// StudentTable is the table that holds the student relation/edge.
	StudentTable = "coding_extensions"
	// StudentInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	StudentInverseTable = "users"
	// StudentColumn is the table column denoting the student relation/edge.
	StudentColumn = "coding_extension_student"
	// CodingProblemTable is the table that holds the coding_problem relation/edge.
	CodingProblemTable = "coding_extensions"
	// CodingProblemInverseTable is the table name for the CodingProblem entity.
	// It exists in this package in order to avoid circular dependency with the "codingproblem" package.
	CodingProblemInverseTable = "coding_problems"
	// CodingProblemColumn is the table column denoting the coding_problem relation/edge.
	CodingProblemColumn = "coding_extension_coding_problem"
)

// Columns holds all SQL columns for codingextension fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeadline,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "coding_extensions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"coding_extension_student",
	"coding_extension_coding_problem",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "170-ag/ent/generated/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
