// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// CodingTestCase is the model entity for the CodingTestCase schema.
type CodingTestCase struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Points holds the value of the "points" field.
	Points int `json:"points,omitempty"`
	// Visibility holds the value of the "visibility" field.
	Visibility codingtestcase.Visibility `json:"visibility,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CodingTestCaseQuery when eager-loading is set.
	Edges                           CodingTestCaseEdges `json:"edges"`
	coding_problem_test_cases       *int
	coding_test_case_data_test_case *int
}

// CodingTestCaseEdges holds the relations/edges for other nodes in the graph.
type CodingTestCaseEdges struct {
	// CodingProblem holds the value of the coding_problem edge.
	CodingProblem *CodingProblem `json:"coding_problem,omitempty"`
	// Data holds the value of the data edge.
	Data *CodingTestCaseData `json:"data,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CodingProblemOrErr returns the CodingProblem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CodingTestCaseEdges) CodingProblemOrErr() (*CodingProblem, error) {
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

// DataOrErr returns the Data value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CodingTestCaseEdges) DataOrErr() (*CodingTestCaseData, error) {
	if e.loadedTypes[1] {
		if e.Data == nil {
			// The edge data was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: codingtestcasedata.Label}
		}
		return e.Data, nil
	}
	return nil, &NotLoadedError{edge: "data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CodingTestCase) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case codingtestcase.FieldID, codingtestcase.FieldPoints:
			values[i] = new(sql.NullInt64)
		case codingtestcase.FieldVisibility:
			values[i] = new(sql.NullString)
		case codingtestcase.FieldCreateTime, codingtestcase.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case codingtestcase.ForeignKeys[0]: // coding_problem_test_cases
			values[i] = new(sql.NullInt64)
		case codingtestcase.ForeignKeys[1]: // coding_test_case_data_test_case
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CodingTestCase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CodingTestCase fields.
func (ctc *CodingTestCase) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case codingtestcase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ctc.ID = int(value.Int64)
		case codingtestcase.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ctc.CreateTime = value.Time
			}
		case codingtestcase.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ctc.UpdateTime = value.Time
			}
		case codingtestcase.FieldPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field points", values[i])
			} else if value.Valid {
				ctc.Points = int(value.Int64)
			}
		case codingtestcase.FieldVisibility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				ctc.Visibility = codingtestcase.Visibility(value.String)
			}
		case codingtestcase.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coding_problem_test_cases", value)
			} else if value.Valid {
				ctc.coding_problem_test_cases = new(int)
				*ctc.coding_problem_test_cases = int(value.Int64)
			}
		case codingtestcase.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field coding_test_case_data_test_case", value)
			} else if value.Valid {
				ctc.coding_test_case_data_test_case = new(int)
				*ctc.coding_test_case_data_test_case = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCodingProblem queries the "coding_problem" edge of the CodingTestCase entity.
func (ctc *CodingTestCase) QueryCodingProblem() *CodingProblemQuery {
	return (&CodingTestCaseClient{config: ctc.config}).QueryCodingProblem(ctc)
}

// QueryData queries the "data" edge of the CodingTestCase entity.
func (ctc *CodingTestCase) QueryData() *CodingTestCaseDataQuery {
	return (&CodingTestCaseClient{config: ctc.config}).QueryData(ctc)
}

// Update returns a builder for updating this CodingTestCase.
// Note that you need to call CodingTestCase.Unwrap() before calling this method if this CodingTestCase
// was returned from a transaction, and the transaction was committed or rolled back.
func (ctc *CodingTestCase) Update() *CodingTestCaseUpdateOne {
	return (&CodingTestCaseClient{config: ctc.config}).UpdateOne(ctc)
}

// Unwrap unwraps the CodingTestCase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ctc *CodingTestCase) Unwrap() *CodingTestCase {
	tx, ok := ctc.config.driver.(*txDriver)
	if !ok {
		panic("generated: CodingTestCase is not a transactional entity")
	}
	ctc.config.driver = tx.drv
	return ctc
}

// String implements the fmt.Stringer.
func (ctc *CodingTestCase) String() string {
	var builder strings.Builder
	builder.WriteString("CodingTestCase(")
	builder.WriteString(fmt.Sprintf("id=%v", ctc.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ctc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ctc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", points=")
	builder.WriteString(fmt.Sprintf("%v", ctc.Points))
	builder.WriteString(", visibility=")
	builder.WriteString(fmt.Sprintf("%v", ctc.Visibility))
	builder.WriteByte(')')
	return builder.String()
}

// CodingTestCases is a parsable slice of CodingTestCase.
type CodingTestCases []*CodingTestCase

func (ctc CodingTestCases) config(cfg config) {
	for _i := range ctc {
		ctc[_i].config = cfg
	}
}
