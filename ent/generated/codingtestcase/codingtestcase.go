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
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// FieldVisible holds the string denoting the visible field in the database.
	FieldVisible = "visible"
	// EdgeCodingProblem holds the string denoting the coding_problem edge name in mutations.
	EdgeCodingProblem = "coding_problem"
	// Table holds the table name of the codingtestcase in the database.
	Table = "coding_test_cases"
	// CodingProblemTable is the table that holds the coding_problem relation/edge. The primary key declared below.
	CodingProblemTable = "coding_problem_test_cases"
	// CodingProblemInverseTable is the table name for the CodingProblem entity.
	// It exists in this package in order to avoid circular dependency with the "codingproblem" package.
	CodingProblemInverseTable = "coding_problems"
)

// Columns holds all SQL columns for codingtestcase fields.
var Columns = []string{
	FieldID,
	FieldInput,
	FieldOutput,
	FieldPoints,
	FieldVisible,
}

var (
	// CodingProblemPrimaryKey and CodingProblemColumn2 are the table columns denoting the
	// primary key for the coding_problem relation (M2M).
	CodingProblemPrimaryKey = []string{"coding_problem_id", "coding_test_case_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// PointsValidator is a validator for the "points" field. It is called by the builders before save.
	PointsValidator func(int) error
)
