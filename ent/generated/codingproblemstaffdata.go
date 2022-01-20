// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingproblemstaffdata"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CodingProblemStaffData is the model entity for the CodingProblemStaffData schema.
type CodingProblemStaffData struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Input holds the value of the "input" field.
	Input string `json:"input,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CodingProblemStaffDataQuery when eager-loading is set.
	Edges CodingProblemStaffDataEdges `json:"edges"`
}

// CodingProblemStaffDataEdges holds the relations/edges for other nodes in the graph.
type CodingProblemStaffDataEdges struct {
	// CodingProblem holds the value of the coding_problem edge.
	CodingProblem *CodingProblem `json:"coding_problem,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CodingProblemOrErr returns the CodingProblem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CodingProblemStaffDataEdges) CodingProblemOrErr() (*CodingProblem, error) {
	if e.loadedTypes[0] {
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
func (*CodingProblemStaffData) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case codingproblemstaffdata.FieldID:
			values[i] = new(sql.NullInt64)
		case codingproblemstaffdata.FieldInput:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CodingProblemStaffData", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CodingProblemStaffData fields.
func (cpsd *CodingProblemStaffData) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case codingproblemstaffdata.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cpsd.ID = int(value.Int64)
		case codingproblemstaffdata.FieldInput:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field input", values[i])
			} else if value.Valid {
				cpsd.Input = value.String
			}
		}
	}
	return nil
}

// QueryCodingProblem queries the "coding_problem" edge of the CodingProblemStaffData entity.
func (cpsd *CodingProblemStaffData) QueryCodingProblem() *CodingProblemQuery {
	return (&CodingProblemStaffDataClient{config: cpsd.config}).QueryCodingProblem(cpsd)
}

// Update returns a builder for updating this CodingProblemStaffData.
// Note that you need to call CodingProblemStaffData.Unwrap() before calling this method if this CodingProblemStaffData
// was returned from a transaction, and the transaction was committed or rolled back.
func (cpsd *CodingProblemStaffData) Update() *CodingProblemStaffDataUpdateOne {
	return (&CodingProblemStaffDataClient{config: cpsd.config}).UpdateOne(cpsd)
}

// Unwrap unwraps the CodingProblemStaffData entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cpsd *CodingProblemStaffData) Unwrap() *CodingProblemStaffData {
	tx, ok := cpsd.config.driver.(*txDriver)
	if !ok {
		panic("generated: CodingProblemStaffData is not a transactional entity")
	}
	cpsd.config.driver = tx.drv
	return cpsd
}

// String implements the fmt.Stringer.
func (cpsd *CodingProblemStaffData) String() string {
	var builder strings.Builder
	builder.WriteString("CodingProblemStaffData(")
	builder.WriteString(fmt.Sprintf("id=%v", cpsd.ID))
	builder.WriteString(", input=")
	builder.WriteString(cpsd.Input)
	builder.WriteByte(')')
	return builder.String()
}

// CodingProblemStaffDataSlice is a parsable slice of CodingProblemStaffData.
type CodingProblemStaffDataSlice []*CodingProblemStaffData

func (cpsd CodingProblemStaffDataSlice) config(cfg config) {
	for _i := range cpsd {
		cpsd[_i].config = cfg
	}
}