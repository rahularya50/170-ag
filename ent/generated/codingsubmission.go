// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CodingSubmission is the model entity for the CodingSubmission schema.
type CodingSubmission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Status holds the value of the "status" field.
	Status codingsubmission.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CodingSubmissionQuery when eager-loading is set.
	Edges                            CodingSubmissionEdges `json:"edges"`
	coding_submission_author         *int
	coding_submission_coding_problem *int
}

// CodingSubmissionEdges holds the relations/edges for other nodes in the graph.
type CodingSubmissionEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// CodingProblem holds the value of the coding_problem edge.
	CodingProblem *CodingProblem `json:"coding_problem,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CodingSubmissionEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// CodingProblemOrErr returns the CodingProblem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CodingSubmissionEdges) CodingProblemOrErr() (*CodingProblem, error) {
	if e.loadedTypes[1] {
		if e.CodingProblem == nil {
			// The edge coding_problem was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: codingproblem.Label}
		}
		return e.CodingProblem, nil
	}
	return nil, &NotLoadedError{edge: "coding_problem"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CodingSubmission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case codingsubmission.FieldID:
			values[i] = new(sql.NullInt64)
		case codingsubmission.FieldCode, codingsubmission.FieldStatus:
			values[i] = new(sql.NullString)
		case codingsubmission.ForeignKeys[0]: // coding_submission_author
			values[i] = new(sql.NullInt64)
		case codingsubmission.ForeignKeys[1]: // coding_submission_coding_problem
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CodingSubmission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CodingSubmission fields.
func (cs *CodingSubmission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case codingsubmission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cs.ID = int(value.Int64)
		case codingsubmission.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				cs.Code = value.String
			}
		case codingsubmission.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cs.Status = codingsubmission.Status(value.String)
			}
		case codingsubmission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coding_submission_author", value)
			} else if value.Valid {
				cs.coding_submission_author = new(int)
				*cs.coding_submission_author = int(value.Int64)
			}
		case codingsubmission.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coding_submission_coding_problem", value)
			} else if value.Valid {
				cs.coding_submission_coding_problem = new(int)
				*cs.coding_submission_coding_problem = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the CodingSubmission entity.
func (cs *CodingSubmission) QueryAuthor() *UserQuery {
	return (&CodingSubmissionClient{config: cs.config}).QueryAuthor(cs)
}

// QueryCodingProblem queries the "coding_problem" edge of the CodingSubmission entity.
func (cs *CodingSubmission) QueryCodingProblem() *CodingProblemQuery {
	return (&CodingSubmissionClient{config: cs.config}).QueryCodingProblem(cs)
}

// Update returns a builder for updating this CodingSubmission.
// Note that you need to call CodingSubmission.Unwrap() before calling this method if this CodingSubmission
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *CodingSubmission) Update() *CodingSubmissionUpdateOne {
	return (&CodingSubmissionClient{config: cs.config}).UpdateOne(cs)
}

// Unwrap unwraps the CodingSubmission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *CodingSubmission) Unwrap() *CodingSubmission {
	tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("generated: CodingSubmission is not a transactional entity")
	}
	cs.config.driver = tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *CodingSubmission) String() string {
	var builder strings.Builder
	builder.WriteString("CodingSubmission(")
	builder.WriteString(fmt.Sprintf("id=%v", cs.ID))
	builder.WriteString(", code=")
	builder.WriteString(cs.Code)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", cs.Status))
	builder.WriteByte(')')
	return builder.String()
}

// CodingSubmissions is a parsable slice of CodingSubmission.
type CodingSubmissions []*CodingSubmission

func (cs CodingSubmissions) config(cfg config) {
	for _i := range cs {
		cs[_i].config = cfg
	}
}
