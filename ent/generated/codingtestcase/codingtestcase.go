// Code generated by entc, DO NOT EDIT.

package codingtestcase

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the codingtestcase type in the database.
	Label = "coding_test_case"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// EdgeCodingProblem holds the string denoting the coding_problem edge name in mutations.
	EdgeCodingProblem = "coding_problem"
	// EdgeData holds the string denoting the data edge name in mutations.
	EdgeData = "data"
	// Table holds the table name of the codingtestcase in the database.
	Table = "coding_test_cases"
	// CodingProblemTable is the table that holds the coding_problem relation/edge.
	CodingProblemTable = "coding_test_cases"
	// CodingProblemInverseTable is the table name for the CodingProblem entity.
	// It exists in this package in order to avoid circular dependency with the "codingproblem" package.
	CodingProblemInverseTable = "coding_problems"
	// CodingProblemColumn is the table column denoting the coding_problem relation/edge.
	CodingProblemColumn = "coding_problem_test_cases"
	// DataTable is the table that holds the data relation/edge.
	DataTable = "coding_test_cases"
	// DataInverseTable is the table name for the CodingTestCaseData entity.
	// It exists in this package in order to avoid circular dependency with the "codingtestcasedata" package.
	DataInverseTable = "coding_test_case_data"
	// DataColumn is the table column denoting the data relation/edge.
	DataColumn = "coding_test_case_data_test_case"
)

// Columns holds all SQL columns for codingtestcase fields.
var Columns = []string{
	FieldID,
	FieldPoints,
	FieldPublic,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "coding_test_cases"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"coding_problem_test_cases",
	"coding_test_case_data_test_case",
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
	// DefaultPoints holds the default value on creation for the "points" field.
	DefaultPoints int
	// PointsValidator is a validator for the "points" field. It is called by the builders before save.
	PointsValidator func(int) error
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
)
