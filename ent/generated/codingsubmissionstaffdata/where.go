// Code generated by entc, DO NOT EDIT.

package codingsubmissionstaffdata

import (
	"170-ag/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ExecutionID applies equality check predicate on the "execution_id" field. It's identical to ExecutionIDEQ.
func ExecutionID(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutionID), v))
	})
}

// Input applies equality check predicate on the "input" field. It's identical to InputEQ.
func Input(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// Output applies equality check predicate on the "output" field. It's identical to OutputEQ.
func Output(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutput), v))
	})
}

// Stderr applies equality check predicate on the "stderr" field. It's identical to StderrEQ.
func Stderr(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStderr), v))
	})
}

// ExitError applies equality check predicate on the "exit_error" field. It's identical to ExitErrorEQ.
func ExitError(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExitError), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ExecutionIDEQ applies the EQ predicate on the "execution_id" field.
func ExecutionIDEQ(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExecutionID), v))
	})
}

// ExecutionIDNEQ applies the NEQ predicate on the "execution_id" field.
func ExecutionIDNEQ(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExecutionID), v))
	})
}

// ExecutionIDIn applies the In predicate on the "execution_id" field.
func ExecutionIDIn(vs ...int64) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExecutionID), v...))
	})
}

// ExecutionIDNotIn applies the NotIn predicate on the "execution_id" field.
func ExecutionIDNotIn(vs ...int64) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExecutionID), v...))
	})
}

// ExecutionIDGT applies the GT predicate on the "execution_id" field.
func ExecutionIDGT(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExecutionID), v))
	})
}

// ExecutionIDGTE applies the GTE predicate on the "execution_id" field.
func ExecutionIDGTE(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExecutionID), v))
	})
}

// ExecutionIDLT applies the LT predicate on the "execution_id" field.
func ExecutionIDLT(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExecutionID), v))
	})
}

// ExecutionIDLTE applies the LTE predicate on the "execution_id" field.
func ExecutionIDLTE(v int64) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExecutionID), v))
	})
}

// ExecutionIDIsNil applies the IsNil predicate on the "execution_id" field.
func ExecutionIDIsNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExecutionID)))
	})
}

// ExecutionIDNotNil applies the NotNil predicate on the "execution_id" field.
func ExecutionIDNotNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExecutionID)))
	})
}

// InputEQ applies the EQ predicate on the "input" field.
func InputEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// InputNEQ applies the NEQ predicate on the "input" field.
func InputNEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInput), v))
	})
}

// InputIn applies the In predicate on the "input" field.
func InputIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInput), v...))
	})
}

// InputNotIn applies the NotIn predicate on the "input" field.
func InputNotIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInput), v...))
	})
}

// InputGT applies the GT predicate on the "input" field.
func InputGT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInput), v))
	})
}

// InputGTE applies the GTE predicate on the "input" field.
func InputGTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInput), v))
	})
}

// InputLT applies the LT predicate on the "input" field.
func InputLT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInput), v))
	})
}

// InputLTE applies the LTE predicate on the "input" field.
func InputLTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInput), v))
	})
}

// InputContains applies the Contains predicate on the "input" field.
func InputContains(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInput), v))
	})
}

// InputHasPrefix applies the HasPrefix predicate on the "input" field.
func InputHasPrefix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInput), v))
	})
}

// InputHasSuffix applies the HasSuffix predicate on the "input" field.
func InputHasSuffix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInput), v))
	})
}

// InputEqualFold applies the EqualFold predicate on the "input" field.
func InputEqualFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInput), v))
	})
}

// InputContainsFold applies the ContainsFold predicate on the "input" field.
func InputContainsFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInput), v))
	})
}

// OutputEQ applies the EQ predicate on the "output" field.
func OutputEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutput), v))
	})
}

// OutputNEQ applies the NEQ predicate on the "output" field.
func OutputNEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutput), v))
	})
}

// OutputIn applies the In predicate on the "output" field.
func OutputIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutput), v...))
	})
}

// OutputNotIn applies the NotIn predicate on the "output" field.
func OutputNotIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutput), v...))
	})
}

// OutputGT applies the GT predicate on the "output" field.
func OutputGT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutput), v))
	})
}

// OutputGTE applies the GTE predicate on the "output" field.
func OutputGTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutput), v))
	})
}

// OutputLT applies the LT predicate on the "output" field.
func OutputLT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutput), v))
	})
}

// OutputLTE applies the LTE predicate on the "output" field.
func OutputLTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutput), v))
	})
}

// OutputContains applies the Contains predicate on the "output" field.
func OutputContains(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutput), v))
	})
}

// OutputHasPrefix applies the HasPrefix predicate on the "output" field.
func OutputHasPrefix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutput), v))
	})
}

// OutputHasSuffix applies the HasSuffix predicate on the "output" field.
func OutputHasSuffix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutput), v))
	})
}

// OutputIsNil applies the IsNil predicate on the "output" field.
func OutputIsNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOutput)))
	})
}

// OutputNotNil applies the NotNil predicate on the "output" field.
func OutputNotNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOutput)))
	})
}

// OutputEqualFold applies the EqualFold predicate on the "output" field.
func OutputEqualFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutput), v))
	})
}

// OutputContainsFold applies the ContainsFold predicate on the "output" field.
func OutputContainsFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutput), v))
	})
}

// StderrEQ applies the EQ predicate on the "stderr" field.
func StderrEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStderr), v))
	})
}

// StderrNEQ applies the NEQ predicate on the "stderr" field.
func StderrNEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStderr), v))
	})
}

// StderrIn applies the In predicate on the "stderr" field.
func StderrIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStderr), v...))
	})
}

// StderrNotIn applies the NotIn predicate on the "stderr" field.
func StderrNotIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStderr), v...))
	})
}

// StderrGT applies the GT predicate on the "stderr" field.
func StderrGT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStderr), v))
	})
}

// StderrGTE applies the GTE predicate on the "stderr" field.
func StderrGTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStderr), v))
	})
}

// StderrLT applies the LT predicate on the "stderr" field.
func StderrLT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStderr), v))
	})
}

// StderrLTE applies the LTE predicate on the "stderr" field.
func StderrLTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStderr), v))
	})
}

// StderrContains applies the Contains predicate on the "stderr" field.
func StderrContains(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStderr), v))
	})
}

// StderrHasPrefix applies the HasPrefix predicate on the "stderr" field.
func StderrHasPrefix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStderr), v))
	})
}

// StderrHasSuffix applies the HasSuffix predicate on the "stderr" field.
func StderrHasSuffix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStderr), v))
	})
}

// StderrIsNil applies the IsNil predicate on the "stderr" field.
func StderrIsNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStderr)))
	})
}

// StderrNotNil applies the NotNil predicate on the "stderr" field.
func StderrNotNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStderr)))
	})
}

// StderrEqualFold applies the EqualFold predicate on the "stderr" field.
func StderrEqualFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStderr), v))
	})
}

// StderrContainsFold applies the ContainsFold predicate on the "stderr" field.
func StderrContainsFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStderr), v))
	})
}

// ExitErrorEQ applies the EQ predicate on the "exit_error" field.
func ExitErrorEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExitError), v))
	})
}

// ExitErrorNEQ applies the NEQ predicate on the "exit_error" field.
func ExitErrorNEQ(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExitError), v))
	})
}

// ExitErrorIn applies the In predicate on the "exit_error" field.
func ExitErrorIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExitError), v...))
	})
}

// ExitErrorNotIn applies the NotIn predicate on the "exit_error" field.
func ExitErrorNotIn(vs ...string) predicate.CodingSubmissionStaffData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExitError), v...))
	})
}

// ExitErrorGT applies the GT predicate on the "exit_error" field.
func ExitErrorGT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExitError), v))
	})
}

// ExitErrorGTE applies the GTE predicate on the "exit_error" field.
func ExitErrorGTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExitError), v))
	})
}

// ExitErrorLT applies the LT predicate on the "exit_error" field.
func ExitErrorLT(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExitError), v))
	})
}

// ExitErrorLTE applies the LTE predicate on the "exit_error" field.
func ExitErrorLTE(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExitError), v))
	})
}

// ExitErrorContains applies the Contains predicate on the "exit_error" field.
func ExitErrorContains(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExitError), v))
	})
}

// ExitErrorHasPrefix applies the HasPrefix predicate on the "exit_error" field.
func ExitErrorHasPrefix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExitError), v))
	})
}

// ExitErrorHasSuffix applies the HasSuffix predicate on the "exit_error" field.
func ExitErrorHasSuffix(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExitError), v))
	})
}

// ExitErrorIsNil applies the IsNil predicate on the "exit_error" field.
func ExitErrorIsNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExitError)))
	})
}

// ExitErrorNotNil applies the NotNil predicate on the "exit_error" field.
func ExitErrorNotNil() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExitError)))
	})
}

// ExitErrorEqualFold applies the EqualFold predicate on the "exit_error" field.
func ExitErrorEqualFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExitError), v))
	})
}

// ExitErrorContainsFold applies the ContainsFold predicate on the "exit_error" field.
func ExitErrorContainsFold(v string) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExitError), v))
	})
}

// HasCodingSubmission applies the HasEdge predicate on the "coding_submission" edge.
func HasCodingSubmission() predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CodingSubmissionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CodingSubmissionTable, CodingSubmissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCodingSubmissionWith applies the HasEdge predicate on the "coding_submission" edge with a given conditions (other predicates).
func HasCodingSubmissionWith(preds ...predicate.CodingSubmission) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CodingSubmissionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CodingSubmissionTable, CodingSubmissionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CodingSubmissionStaffData) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CodingSubmissionStaffData) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CodingSubmissionStaffData) predicate.CodingSubmissionStaffData {
	return predicate.CodingSubmissionStaffData(func(s *sql.Selector) {
		p(s.Not())
	})
}
