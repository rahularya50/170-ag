// Code generated by entc, DO NOT EDIT.

package generated

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"170-ag/ent/generated/predicate"
	"170-ag/ent/generated/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 7)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   codingdraft.Table,
			Columns: codingdraft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingdraft.FieldID,
			},
		},
		Type: "CodingDraft",
		Fields: map[string]*sqlgraph.FieldSpec{
			codingdraft.FieldCreateTime: {Type: field.TypeTime, Column: codingdraft.FieldCreateTime},
			codingdraft.FieldUpdateTime: {Type: field.TypeTime, Column: codingdraft.FieldUpdateTime},
			codingdraft.FieldCode:       {Type: field.TypeString, Column: codingdraft.FieldCode},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   codingproblem.Table,
			Columns: codingproblem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingproblem.FieldID,
			},
		},
		Type: "CodingProblem",
		Fields: map[string]*sqlgraph.FieldSpec{
			codingproblem.FieldCreateTime: {Type: field.TypeTime, Column: codingproblem.FieldCreateTime},
			codingproblem.FieldUpdateTime: {Type: field.TypeTime, Column: codingproblem.FieldUpdateTime},
			codingproblem.FieldName:       {Type: field.TypeString, Column: codingproblem.FieldName},
			codingproblem.FieldStatement:  {Type: field.TypeString, Column: codingproblem.FieldStatement},
			codingproblem.FieldReleased:   {Type: field.TypeBool, Column: codingproblem.FieldReleased},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   codingsubmission.Table,
			Columns: codingsubmission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingsubmission.FieldID,
			},
		},
		Type: "CodingSubmission",
		Fields: map[string]*sqlgraph.FieldSpec{
			codingsubmission.FieldCreateTime: {Type: field.TypeTime, Column: codingsubmission.FieldCreateTime},
			codingsubmission.FieldUpdateTime: {Type: field.TypeTime, Column: codingsubmission.FieldUpdateTime},
			codingsubmission.FieldCode:       {Type: field.TypeString, Column: codingsubmission.FieldCode},
			codingsubmission.FieldStatus:     {Type: field.TypeEnum, Column: codingsubmission.FieldStatus},
			codingsubmission.FieldPoints:     {Type: field.TypeInt, Column: codingsubmission.FieldPoints},
			codingsubmission.FieldResults:    {Type: field.TypeJSON, Column: codingsubmission.FieldResults},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   codingsubmissionstaffdata.Table,
			Columns: codingsubmissionstaffdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingsubmissionstaffdata.FieldID,
			},
		},
		Type: "CodingSubmissionStaffData",
		Fields: map[string]*sqlgraph.FieldSpec{
			codingsubmissionstaffdata.FieldCreateTime:  {Type: field.TypeTime, Column: codingsubmissionstaffdata.FieldCreateTime},
			codingsubmissionstaffdata.FieldUpdateTime:  {Type: field.TypeTime, Column: codingsubmissionstaffdata.FieldUpdateTime},
			codingsubmissionstaffdata.FieldExecutionID: {Type: field.TypeInt64, Column: codingsubmissionstaffdata.FieldExecutionID},
			codingsubmissionstaffdata.FieldInput:       {Type: field.TypeString, Column: codingsubmissionstaffdata.FieldInput},
			codingsubmissionstaffdata.FieldOutput:      {Type: field.TypeString, Column: codingsubmissionstaffdata.FieldOutput},
			codingsubmissionstaffdata.FieldStderr:      {Type: field.TypeString, Column: codingsubmissionstaffdata.FieldStderr},
			codingsubmissionstaffdata.FieldExitError:   {Type: field.TypeString, Column: codingsubmissionstaffdata.FieldExitError},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   codingtestcase.Table,
			Columns: codingtestcase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcase.FieldID,
			},
		},
		Type: "CodingTestCase",
		Fields: map[string]*sqlgraph.FieldSpec{
			codingtestcase.FieldCreateTime: {Type: field.TypeTime, Column: codingtestcase.FieldCreateTime},
			codingtestcase.FieldUpdateTime: {Type: field.TypeTime, Column: codingtestcase.FieldUpdateTime},
			codingtestcase.FieldPoints:     {Type: field.TypeInt, Column: codingtestcase.FieldPoints},
			codingtestcase.FieldPublic:     {Type: field.TypeBool, Column: codingtestcase.FieldPublic},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   codingtestcasedata.Table,
			Columns: codingtestcasedata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: codingtestcasedata.FieldID,
			},
		},
		Type: "CodingTestCaseData",
		Fields: map[string]*sqlgraph.FieldSpec{
			codingtestcasedata.FieldCreateTime: {Type: field.TypeTime, Column: codingtestcasedata.FieldCreateTime},
			codingtestcasedata.FieldUpdateTime: {Type: field.TypeTime, Column: codingtestcasedata.FieldUpdateTime},
			codingtestcasedata.FieldInput:      {Type: field.TypeString, Column: codingtestcasedata.FieldInput},
			codingtestcasedata.FieldOutput:     {Type: field.TypeString, Column: codingtestcasedata.FieldOutput},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateTime: {Type: field.TypeTime, Column: user.FieldCreateTime},
			user.FieldUpdateTime: {Type: field.TypeTime, Column: user.FieldUpdateTime},
			user.FieldEmail:      {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldName:       {Type: field.TypeString, Column: user.FieldName},
			user.FieldIsStaff:    {Type: field.TypeBool, Column: user.FieldIsStaff},
		},
	}
	graph.MustAddE(
		"author",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingdraft.AuthorTable,
			Columns: []string{codingdraft.AuthorColumn},
			Bidi:    false,
		},
		"CodingDraft",
		"User",
	)
	graph.MustAddE(
		"coding_problem",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingdraft.CodingProblemTable,
			Columns: []string{codingdraft.CodingProblemColumn},
			Bidi:    false,
		},
		"CodingDraft",
		"CodingProblem",
	)
	graph.MustAddE(
		"drafts",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.DraftsTable,
			Columns: []string{codingproblem.DraftsColumn},
			Bidi:    false,
		},
		"CodingProblem",
		"CodingDraft",
	)
	graph.MustAddE(
		"test_cases",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   codingproblem.TestCasesTable,
			Columns: []string{codingproblem.TestCasesColumn},
			Bidi:    false,
		},
		"CodingProblem",
		"CodingTestCase",
	)
	graph.MustAddE(
		"submissions",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   codingproblem.SubmissionsTable,
			Columns: []string{codingproblem.SubmissionsColumn},
			Bidi:    false,
		},
		"CodingProblem",
		"CodingSubmission",
	)
	graph.MustAddE(
		"author",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.AuthorTable,
			Columns: []string{codingsubmission.AuthorColumn},
			Bidi:    false,
		},
		"CodingSubmission",
		"User",
	)
	graph.MustAddE(
		"coding_problem",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   codingsubmission.CodingProblemTable,
			Columns: []string{codingsubmission.CodingProblemColumn},
			Bidi:    false,
		},
		"CodingSubmission",
		"CodingProblem",
	)
	graph.MustAddE(
		"staff_data",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingsubmission.StaffDataTable,
			Columns: []string{codingsubmission.StaffDataColumn},
			Bidi:    false,
		},
		"CodingSubmission",
		"CodingSubmissionStaffData",
	)
	graph.MustAddE(
		"coding_submission",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   codingsubmissionstaffdata.CodingSubmissionTable,
			Columns: []string{codingsubmissionstaffdata.CodingSubmissionColumn},
			Bidi:    false,
		},
		"CodingSubmissionStaffData",
		"CodingSubmission",
	)
	graph.MustAddE(
		"coding_problem",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   codingtestcase.CodingProblemTable,
			Columns: []string{codingtestcase.CodingProblemColumn},
			Bidi:    false,
		},
		"CodingTestCase",
		"CodingProblem",
	)
	graph.MustAddE(
		"data",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   codingtestcase.DataTable,
			Columns: []string{codingtestcase.DataColumn},
			Bidi:    false,
		},
		"CodingTestCase",
		"CodingTestCaseData",
	)
	graph.MustAddE(
		"test_case",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   codingtestcasedata.TestCaseTable,
			Columns: []string{codingtestcasedata.TestCaseColumn},
			Bidi:    false,
		},
		"CodingTestCaseData",
		"CodingTestCase",
	)
	graph.MustAddE(
		"drafts",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DraftsTable,
			Columns: []string{user.DraftsColumn},
			Bidi:    false,
		},
		"User",
		"CodingDraft",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cdq *CodingDraftQuery) addPredicate(pred func(s *sql.Selector)) {
	cdq.predicates = append(cdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CodingDraftQuery builder.
func (cdq *CodingDraftQuery) Filter() *CodingDraftFilter {
	return &CodingDraftFilter{cdq}
}

// addPredicate implements the predicateAdder interface.
func (m *CodingDraftMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CodingDraftMutation builder.
func (m *CodingDraftMutation) Filter() *CodingDraftFilter {
	return &CodingDraftFilter{m}
}

// CodingDraftFilter provides a generic filtering capability at runtime for CodingDraftQuery.
type CodingDraftFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CodingDraftFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CodingDraftFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(codingdraft.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CodingDraftFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(codingdraft.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CodingDraftFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(codingdraft.FieldUpdateTime))
}

// WhereCode applies the entql string predicate on the code field.
func (f *CodingDraftFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(codingdraft.FieldCode))
}

// WhereHasAuthor applies a predicate to check if query has an edge author.
func (f *CodingDraftFilter) WhereHasAuthor() {
	f.Where(entql.HasEdge("author"))
}

// WhereHasAuthorWith applies a predicate to check if query has an edge author with a given conditions (other predicates).
func (f *CodingDraftFilter) WhereHasAuthorWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("author", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasCodingProblem applies a predicate to check if query has an edge coding_problem.
func (f *CodingDraftFilter) WhereHasCodingProblem() {
	f.Where(entql.HasEdge("coding_problem"))
}

// WhereHasCodingProblemWith applies a predicate to check if query has an edge coding_problem with a given conditions (other predicates).
func (f *CodingDraftFilter) WhereHasCodingProblemWith(preds ...predicate.CodingProblem) {
	f.Where(entql.HasEdgeWith("coding_problem", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (cpq *CodingProblemQuery) addPredicate(pred func(s *sql.Selector)) {
	cpq.predicates = append(cpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CodingProblemQuery builder.
func (cpq *CodingProblemQuery) Filter() *CodingProblemFilter {
	return &CodingProblemFilter{cpq}
}

// addPredicate implements the predicateAdder interface.
func (m *CodingProblemMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CodingProblemMutation builder.
func (m *CodingProblemMutation) Filter() *CodingProblemFilter {
	return &CodingProblemFilter{m}
}

// CodingProblemFilter provides a generic filtering capability at runtime for CodingProblemQuery.
type CodingProblemFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CodingProblemFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CodingProblemFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(codingproblem.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CodingProblemFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(codingproblem.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CodingProblemFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(codingproblem.FieldUpdateTime))
}

// WhereName applies the entql string predicate on the name field.
func (f *CodingProblemFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(codingproblem.FieldName))
}

// WhereStatement applies the entql string predicate on the statement field.
func (f *CodingProblemFilter) WhereStatement(p entql.StringP) {
	f.Where(p.Field(codingproblem.FieldStatement))
}

// WhereReleased applies the entql bool predicate on the released field.
func (f *CodingProblemFilter) WhereReleased(p entql.BoolP) {
	f.Where(p.Field(codingproblem.FieldReleased))
}

// WhereHasDrafts applies a predicate to check if query has an edge drafts.
func (f *CodingProblemFilter) WhereHasDrafts() {
	f.Where(entql.HasEdge("drafts"))
}

// WhereHasDraftsWith applies a predicate to check if query has an edge drafts with a given conditions (other predicates).
func (f *CodingProblemFilter) WhereHasDraftsWith(preds ...predicate.CodingDraft) {
	f.Where(entql.HasEdgeWith("drafts", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasTestCases applies a predicate to check if query has an edge test_cases.
func (f *CodingProblemFilter) WhereHasTestCases() {
	f.Where(entql.HasEdge("test_cases"))
}

// WhereHasTestCasesWith applies a predicate to check if query has an edge test_cases with a given conditions (other predicates).
func (f *CodingProblemFilter) WhereHasTestCasesWith(preds ...predicate.CodingTestCase) {
	f.Where(entql.HasEdgeWith("test_cases", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasSubmissions applies a predicate to check if query has an edge submissions.
func (f *CodingProblemFilter) WhereHasSubmissions() {
	f.Where(entql.HasEdge("submissions"))
}

// WhereHasSubmissionsWith applies a predicate to check if query has an edge submissions with a given conditions (other predicates).
func (f *CodingProblemFilter) WhereHasSubmissionsWith(preds ...predicate.CodingSubmission) {
	f.Where(entql.HasEdgeWith("submissions", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (csq *CodingSubmissionQuery) addPredicate(pred func(s *sql.Selector)) {
	csq.predicates = append(csq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CodingSubmissionQuery builder.
func (csq *CodingSubmissionQuery) Filter() *CodingSubmissionFilter {
	return &CodingSubmissionFilter{csq}
}

// addPredicate implements the predicateAdder interface.
func (m *CodingSubmissionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CodingSubmissionMutation builder.
func (m *CodingSubmissionMutation) Filter() *CodingSubmissionFilter {
	return &CodingSubmissionFilter{m}
}

// CodingSubmissionFilter provides a generic filtering capability at runtime for CodingSubmissionQuery.
type CodingSubmissionFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CodingSubmissionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CodingSubmissionFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(codingsubmission.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CodingSubmissionFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(codingsubmission.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CodingSubmissionFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(codingsubmission.FieldUpdateTime))
}

// WhereCode applies the entql string predicate on the code field.
func (f *CodingSubmissionFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(codingsubmission.FieldCode))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *CodingSubmissionFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(codingsubmission.FieldStatus))
}

// WherePoints applies the entql int predicate on the points field.
func (f *CodingSubmissionFilter) WherePoints(p entql.IntP) {
	f.Where(p.Field(codingsubmission.FieldPoints))
}

// WhereResults applies the entql json.RawMessage predicate on the results field.
func (f *CodingSubmissionFilter) WhereResults(p entql.BytesP) {
	f.Where(p.Field(codingsubmission.FieldResults))
}

// WhereHasAuthor applies a predicate to check if query has an edge author.
func (f *CodingSubmissionFilter) WhereHasAuthor() {
	f.Where(entql.HasEdge("author"))
}

// WhereHasAuthorWith applies a predicate to check if query has an edge author with a given conditions (other predicates).
func (f *CodingSubmissionFilter) WhereHasAuthorWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("author", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasCodingProblem applies a predicate to check if query has an edge coding_problem.
func (f *CodingSubmissionFilter) WhereHasCodingProblem() {
	f.Where(entql.HasEdge("coding_problem"))
}

// WhereHasCodingProblemWith applies a predicate to check if query has an edge coding_problem with a given conditions (other predicates).
func (f *CodingSubmissionFilter) WhereHasCodingProblemWith(preds ...predicate.CodingProblem) {
	f.Where(entql.HasEdgeWith("coding_problem", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasStaffData applies a predicate to check if query has an edge staff_data.
func (f *CodingSubmissionFilter) WhereHasStaffData() {
	f.Where(entql.HasEdge("staff_data"))
}

// WhereHasStaffDataWith applies a predicate to check if query has an edge staff_data with a given conditions (other predicates).
func (f *CodingSubmissionFilter) WhereHasStaffDataWith(preds ...predicate.CodingSubmissionStaffData) {
	f.Where(entql.HasEdgeWith("staff_data", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (cssdq *CodingSubmissionStaffDataQuery) addPredicate(pred func(s *sql.Selector)) {
	cssdq.predicates = append(cssdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CodingSubmissionStaffDataQuery builder.
func (cssdq *CodingSubmissionStaffDataQuery) Filter() *CodingSubmissionStaffDataFilter {
	return &CodingSubmissionStaffDataFilter{cssdq}
}

// addPredicate implements the predicateAdder interface.
func (m *CodingSubmissionStaffDataMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CodingSubmissionStaffDataMutation builder.
func (m *CodingSubmissionStaffDataMutation) Filter() *CodingSubmissionStaffDataFilter {
	return &CodingSubmissionStaffDataFilter{m}
}

// CodingSubmissionStaffDataFilter provides a generic filtering capability at runtime for CodingSubmissionStaffDataQuery.
type CodingSubmissionStaffDataFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CodingSubmissionStaffDataFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CodingSubmissionStaffDataFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CodingSubmissionStaffDataFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CodingSubmissionStaffDataFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldUpdateTime))
}

// WhereExecutionID applies the entql int64 predicate on the execution_id field.
func (f *CodingSubmissionStaffDataFilter) WhereExecutionID(p entql.Int64P) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldExecutionID))
}

// WhereInput applies the entql string predicate on the input field.
func (f *CodingSubmissionStaffDataFilter) WhereInput(p entql.StringP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldInput))
}

// WhereOutput applies the entql string predicate on the output field.
func (f *CodingSubmissionStaffDataFilter) WhereOutput(p entql.StringP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldOutput))
}

// WhereStderr applies the entql string predicate on the stderr field.
func (f *CodingSubmissionStaffDataFilter) WhereStderr(p entql.StringP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldStderr))
}

// WhereExitError applies the entql string predicate on the exit_error field.
func (f *CodingSubmissionStaffDataFilter) WhereExitError(p entql.StringP) {
	f.Where(p.Field(codingsubmissionstaffdata.FieldExitError))
}

// WhereHasCodingSubmission applies a predicate to check if query has an edge coding_submission.
func (f *CodingSubmissionStaffDataFilter) WhereHasCodingSubmission() {
	f.Where(entql.HasEdge("coding_submission"))
}

// WhereHasCodingSubmissionWith applies a predicate to check if query has an edge coding_submission with a given conditions (other predicates).
func (f *CodingSubmissionStaffDataFilter) WhereHasCodingSubmissionWith(preds ...predicate.CodingSubmission) {
	f.Where(entql.HasEdgeWith("coding_submission", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (ctcq *CodingTestCaseQuery) addPredicate(pred func(s *sql.Selector)) {
	ctcq.predicates = append(ctcq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CodingTestCaseQuery builder.
func (ctcq *CodingTestCaseQuery) Filter() *CodingTestCaseFilter {
	return &CodingTestCaseFilter{ctcq}
}

// addPredicate implements the predicateAdder interface.
func (m *CodingTestCaseMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CodingTestCaseMutation builder.
func (m *CodingTestCaseMutation) Filter() *CodingTestCaseFilter {
	return &CodingTestCaseFilter{m}
}

// CodingTestCaseFilter provides a generic filtering capability at runtime for CodingTestCaseQuery.
type CodingTestCaseFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CodingTestCaseFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CodingTestCaseFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(codingtestcase.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CodingTestCaseFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(codingtestcase.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CodingTestCaseFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(codingtestcase.FieldUpdateTime))
}

// WherePoints applies the entql int predicate on the points field.
func (f *CodingTestCaseFilter) WherePoints(p entql.IntP) {
	f.Where(p.Field(codingtestcase.FieldPoints))
}

// WherePublic applies the entql bool predicate on the public field.
func (f *CodingTestCaseFilter) WherePublic(p entql.BoolP) {
	f.Where(p.Field(codingtestcase.FieldPublic))
}

// WhereHasCodingProblem applies a predicate to check if query has an edge coding_problem.
func (f *CodingTestCaseFilter) WhereHasCodingProblem() {
	f.Where(entql.HasEdge("coding_problem"))
}

// WhereHasCodingProblemWith applies a predicate to check if query has an edge coding_problem with a given conditions (other predicates).
func (f *CodingTestCaseFilter) WhereHasCodingProblemWith(preds ...predicate.CodingProblem) {
	f.Where(entql.HasEdgeWith("coding_problem", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasData applies a predicate to check if query has an edge data.
func (f *CodingTestCaseFilter) WhereHasData() {
	f.Where(entql.HasEdge("data"))
}

// WhereHasDataWith applies a predicate to check if query has an edge data with a given conditions (other predicates).
func (f *CodingTestCaseFilter) WhereHasDataWith(preds ...predicate.CodingTestCaseData) {
	f.Where(entql.HasEdgeWith("data", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (ctcdq *CodingTestCaseDataQuery) addPredicate(pred func(s *sql.Selector)) {
	ctcdq.predicates = append(ctcdq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CodingTestCaseDataQuery builder.
func (ctcdq *CodingTestCaseDataQuery) Filter() *CodingTestCaseDataFilter {
	return &CodingTestCaseDataFilter{ctcdq}
}

// addPredicate implements the predicateAdder interface.
func (m *CodingTestCaseDataMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CodingTestCaseDataMutation builder.
func (m *CodingTestCaseDataMutation) Filter() *CodingTestCaseDataFilter {
	return &CodingTestCaseDataFilter{m}
}

// CodingTestCaseDataFilter provides a generic filtering capability at runtime for CodingTestCaseDataQuery.
type CodingTestCaseDataFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CodingTestCaseDataFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *CodingTestCaseDataFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(codingtestcasedata.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *CodingTestCaseDataFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(codingtestcasedata.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *CodingTestCaseDataFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(codingtestcasedata.FieldUpdateTime))
}

// WhereInput applies the entql string predicate on the input field.
func (f *CodingTestCaseDataFilter) WhereInput(p entql.StringP) {
	f.Where(p.Field(codingtestcasedata.FieldInput))
}

// WhereOutput applies the entql string predicate on the output field.
func (f *CodingTestCaseDataFilter) WhereOutput(p entql.StringP) {
	f.Where(p.Field(codingtestcasedata.FieldOutput))
}

// WhereHasTestCase applies a predicate to check if query has an edge test_case.
func (f *CodingTestCaseDataFilter) WhereHasTestCase() {
	f.Where(entql.HasEdge("test_case"))
}

// WhereHasTestCaseWith applies a predicate to check if query has an edge test_case with a given conditions (other predicates).
func (f *CodingTestCaseDataFilter) WhereHasTestCaseWith(preds ...predicate.CodingTestCase) {
	f.Where(entql.HasEdgeWith("test_case", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *UserFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereIsStaff applies the entql bool predicate on the is_staff field.
func (f *UserFilter) WhereIsStaff(p entql.BoolP) {
	f.Where(p.Field(user.FieldIsStaff))
}

// WhereHasDrafts applies a predicate to check if query has an edge drafts.
func (f *UserFilter) WhereHasDrafts() {
	f.Where(entql.HasEdge("drafts"))
}

// WhereHasDraftsWith applies a predicate to check if query has an edge drafts with a given conditions (other predicates).
func (f *UserFilter) WhereHasDraftsWith(preds ...predicate.CodingDraft) {
	f.Where(entql.HasEdgeWith("drafts", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
