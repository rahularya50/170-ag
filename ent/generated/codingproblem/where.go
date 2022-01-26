// Code generated by entc, DO NOT EDIT.

package codingproblem

import (
	"170-ag/ent/generated/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func IDGT(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Statement applies equality check predicate on the "statement" field. It's identical to StatementEQ.
func Statement(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatement), v))
	})
}

// Skeleton applies equality check predicate on the "skeleton" field. It's identical to SkeletonEQ.
func Skeleton(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkeleton), v))
	})
}

// Released applies equality check predicate on the "released" field. It's identical to ReleasedEQ.
func Released(v bool) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleased), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StatementEQ applies the EQ predicate on the "statement" field.
func StatementEQ(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatement), v))
	})
}

// StatementNEQ applies the NEQ predicate on the "statement" field.
func StatementNEQ(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatement), v))
	})
}

// StatementIn applies the In predicate on the "statement" field.
func StatementIn(vs ...string) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatement), v...))
	})
}

// StatementNotIn applies the NotIn predicate on the "statement" field.
func StatementNotIn(vs ...string) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatement), v...))
	})
}

// StatementGT applies the GT predicate on the "statement" field.
func StatementGT(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatement), v))
	})
}

// StatementGTE applies the GTE predicate on the "statement" field.
func StatementGTE(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatement), v))
	})
}

// StatementLT applies the LT predicate on the "statement" field.
func StatementLT(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatement), v))
	})
}

// StatementLTE applies the LTE predicate on the "statement" field.
func StatementLTE(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatement), v))
	})
}

// StatementContains applies the Contains predicate on the "statement" field.
func StatementContains(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatement), v))
	})
}

// StatementHasPrefix applies the HasPrefix predicate on the "statement" field.
func StatementHasPrefix(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatement), v))
	})
}

// StatementHasSuffix applies the HasSuffix predicate on the "statement" field.
func StatementHasSuffix(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatement), v))
	})
}

// StatementEqualFold applies the EqualFold predicate on the "statement" field.
func StatementEqualFold(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatement), v))
	})
}

// StatementContainsFold applies the ContainsFold predicate on the "statement" field.
func StatementContainsFold(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatement), v))
	})
}

// SkeletonEQ applies the EQ predicate on the "skeleton" field.
func SkeletonEQ(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkeleton), v))
	})
}

// SkeletonNEQ applies the NEQ predicate on the "skeleton" field.
func SkeletonNEQ(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkeleton), v))
	})
}

// SkeletonIn applies the In predicate on the "skeleton" field.
func SkeletonIn(vs ...string) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkeleton), v...))
	})
}

// SkeletonNotIn applies the NotIn predicate on the "skeleton" field.
func SkeletonNotIn(vs ...string) predicate.CodingProblem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CodingProblem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkeleton), v...))
	})
}

// SkeletonGT applies the GT predicate on the "skeleton" field.
func SkeletonGT(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkeleton), v))
	})
}

// SkeletonGTE applies the GTE predicate on the "skeleton" field.
func SkeletonGTE(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkeleton), v))
	})
}

// SkeletonLT applies the LT predicate on the "skeleton" field.
func SkeletonLT(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkeleton), v))
	})
}

// SkeletonLTE applies the LTE predicate on the "skeleton" field.
func SkeletonLTE(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkeleton), v))
	})
}

// SkeletonContains applies the Contains predicate on the "skeleton" field.
func SkeletonContains(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSkeleton), v))
	})
}

// SkeletonHasPrefix applies the HasPrefix predicate on the "skeleton" field.
func SkeletonHasPrefix(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSkeleton), v))
	})
}

// SkeletonHasSuffix applies the HasSuffix predicate on the "skeleton" field.
func SkeletonHasSuffix(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSkeleton), v))
	})
}

// SkeletonEqualFold applies the EqualFold predicate on the "skeleton" field.
func SkeletonEqualFold(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSkeleton), v))
	})
}

// SkeletonContainsFold applies the ContainsFold predicate on the "skeleton" field.
func SkeletonContainsFold(v string) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSkeleton), v))
	})
}

// ReleasedEQ applies the EQ predicate on the "released" field.
func ReleasedEQ(v bool) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleased), v))
	})
}

// ReleasedNEQ applies the NEQ predicate on the "released" field.
func ReleasedNEQ(v bool) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReleased), v))
	})
}

// HasDrafts applies the HasEdge predicate on the "drafts" edge.
func HasDrafts() predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DraftsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DraftsTable, DraftsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDraftsWith applies the HasEdge predicate on the "drafts" edge with a given conditions (other predicates).
func HasDraftsWith(preds ...predicate.CodingDraft) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DraftsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, DraftsTable, DraftsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTestCases applies the HasEdge predicate on the "test_cases" edge.
func HasTestCases() predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestCasesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestCasesTable, TestCasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestCasesWith applies the HasEdge predicate on the "test_cases" edge with a given conditions (other predicates).
func HasTestCasesWith(preds ...predicate.CodingTestCase) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestCasesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestCasesTable, TestCasesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubmissions applies the HasEdge predicate on the "submissions" edge.
func HasSubmissions() predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubmissionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SubmissionsTable, SubmissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubmissionsWith applies the HasEdge predicate on the "submissions" edge with a given conditions (other predicates).
func HasSubmissionsWith(preds ...predicate.CodingSubmission) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubmissionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SubmissionsTable, SubmissionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CodingProblem) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CodingProblem) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
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
func Not(p predicate.CodingProblem) predicate.CodingProblem {
	return predicate.CodingProblem(func(s *sql.Selector) {
		p(s.Not())
	})
}
