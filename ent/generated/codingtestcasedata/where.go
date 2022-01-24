// Code generated by entc, DO NOT EDIT.

package codingtestcasedata

import (
	"170-ag/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func IDGT(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Input applies equality check predicate on the "input" field. It's identical to InputEQ.
func Input(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// Output applies equality check predicate on the "output" field. It's identical to OutputEQ.
func Output(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutput), v))
	})
}

// InputEQ applies the EQ predicate on the "input" field.
func InputEQ(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInput), v))
	})
}

// InputNEQ applies the NEQ predicate on the "input" field.
func InputNEQ(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInput), v))
	})
}

// InputIn applies the In predicate on the "input" field.
func InputIn(vs ...string) predicate.CodingTestCaseData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func InputNotIn(vs ...string) predicate.CodingTestCaseData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func InputGT(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInput), v))
	})
}

// InputGTE applies the GTE predicate on the "input" field.
func InputGTE(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInput), v))
	})
}

// InputLT applies the LT predicate on the "input" field.
func InputLT(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInput), v))
	})
}

// InputLTE applies the LTE predicate on the "input" field.
func InputLTE(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInput), v))
	})
}

// InputContains applies the Contains predicate on the "input" field.
func InputContains(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInput), v))
	})
}

// InputHasPrefix applies the HasPrefix predicate on the "input" field.
func InputHasPrefix(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInput), v))
	})
}

// InputHasSuffix applies the HasSuffix predicate on the "input" field.
func InputHasSuffix(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInput), v))
	})
}

// InputEqualFold applies the EqualFold predicate on the "input" field.
func InputEqualFold(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInput), v))
	})
}

// InputContainsFold applies the ContainsFold predicate on the "input" field.
func InputContainsFold(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInput), v))
	})
}

// OutputEQ applies the EQ predicate on the "output" field.
func OutputEQ(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutput), v))
	})
}

// OutputNEQ applies the NEQ predicate on the "output" field.
func OutputNEQ(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutput), v))
	})
}

// OutputIn applies the In predicate on the "output" field.
func OutputIn(vs ...string) predicate.CodingTestCaseData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func OutputNotIn(vs ...string) predicate.CodingTestCaseData {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func OutputGT(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutput), v))
	})
}

// OutputGTE applies the GTE predicate on the "output" field.
func OutputGTE(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutput), v))
	})
}

// OutputLT applies the LT predicate on the "output" field.
func OutputLT(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutput), v))
	})
}

// OutputLTE applies the LTE predicate on the "output" field.
func OutputLTE(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutput), v))
	})
}

// OutputContains applies the Contains predicate on the "output" field.
func OutputContains(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutput), v))
	})
}

// OutputHasPrefix applies the HasPrefix predicate on the "output" field.
func OutputHasPrefix(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutput), v))
	})
}

// OutputHasSuffix applies the HasSuffix predicate on the "output" field.
func OutputHasSuffix(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutput), v))
	})
}

// OutputEqualFold applies the EqualFold predicate on the "output" field.
func OutputEqualFold(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutput), v))
	})
}

// OutputContainsFold applies the ContainsFold predicate on the "output" field.
func OutputContainsFold(v string) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutput), v))
	})
}

// HasTestCase applies the HasEdge predicate on the "test_case" edge.
func HasTestCase() predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestCaseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TestCaseTable, TestCaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestCaseWith applies the HasEdge predicate on the "test_case" edge with a given conditions (other predicates).
func HasTestCaseWith(preds ...predicate.CodingTestCase) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestCaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TestCaseTable, TestCaseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CodingTestCaseData) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CodingTestCaseData) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
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
func Not(p predicate.CodingTestCaseData) predicate.CodingTestCaseData {
	return predicate.CodingTestCaseData(func(s *sql.Selector) {
		p(s.Not())
	})
}