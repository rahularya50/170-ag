// Code generated by entc, DO NOT EDIT.

package codingsubmissionstaffdata

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the codingsubmissionstaffdata type in the database.
	Label = "coding_submission_staff_data"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExecutionID holds the string denoting the execution_id field in the database.
	FieldExecutionID = "execution_id"
	// FieldInput holds the string denoting the input field in the database.
	FieldInput = "input"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldStderr holds the string denoting the stderr field in the database.
	FieldStderr = "stderr"
	// FieldExitError holds the string denoting the exit_error field in the database.
	FieldExitError = "exit_error"
	// EdgeCodingSubmission holds the string denoting the coding_submission edge name in mutations.
	EdgeCodingSubmission = "coding_submission"
	// Table holds the table name of the codingsubmissionstaffdata in the database.
	Table = "coding_submission_staff_data"
	// CodingSubmissionTable is the table that holds the coding_submission relation/edge.
	CodingSubmissionTable = "coding_submissions"
	// CodingSubmissionInverseTable is the table name for the CodingSubmission entity.
	// It exists in this package in order to avoid circular dependency with the "codingsubmission" package.
	CodingSubmissionInverseTable = "coding_submissions"
	// CodingSubmissionColumn is the table column denoting the coding_submission relation/edge.
	CodingSubmissionColumn = "coding_submission_staff_data_coding_submission"
)

// Columns holds all SQL columns for codingsubmissionstaffdata fields.
var Columns = []string{
	FieldID,
	FieldExecutionID,
	FieldInput,
	FieldOutput,
	FieldStderr,
	FieldExitError,
}

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
	// OutputValidator is a validator for the "output" field. It is called by the builders before save.
	OutputValidator func(string) error
	// StderrValidator is a validator for the "stderr" field. It is called by the builders before save.
	StderrValidator func(string) error
	// ExitErrorValidator is a validator for the "exit_error" field. It is called by the builders before save.
	ExitErrorValidator func(string) error
)
