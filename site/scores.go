package site

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingextension"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/user"
	"context"
	"encoding/csv"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

type score = struct {
	Email  string `sql:"email"`
	Name   string `sql:"name"`
	Points int    `sql:"max"`
}

func ExportStudentScores(ctx context.Context, problem *ent.CodingProblem) (string, error) {
	scores, err := queryStudentScores(ctx, problem)
	if err != nil {
		return "", err
	}

	out := &strings.Builder{}
	writer := csv.NewWriter(out)
	writer.Write([]string{"Email", "Name", "Points"})
	for _, score := range scores {
		writer.Write([]string{score.Email, score.Name, fmt.Sprint(score.Points)})
	}
	writer.Flush()
	if err = writer.Error(); err != nil {
		return "", err
	}
	return out.String(), nil
}

func queryStudentScores(ctx context.Context, problem *ent.CodingProblem) ([]score, error) {
	var scores []score
	err := problem.QuerySubmissions().Modify(func(s *sql.Selector) {
		problems := sql.Table(codingproblem.Table)
		extensions := sql.Table(codingextension.Table)
		users := sql.Table(user.Table)
		s.
			// join with problem
			Join(problems).
			On(s.C(codingsubmission.CodingProblemColumn), problems.C(codingproblem.FieldID)).

			// join with extensions, if available
			LeftJoin(extensions).
			On(s.C(codingsubmission.AuthorColumn), extensions.C(codingextension.StudentColumn)).

			// remove invalid (but nonnull) extensions
			Where(
				sql.Or(
					sql.EQ(extensions.C(codingextension.CodingProblemColumn), problem.ID),
					sql.IsNull(extensions.C(codingextension.CodingProblemColumn)),
				)).

			// filter to submissions made before deadline
			Where(
				sql.Or(
					sql.LT(s.C(codingsubmission.FieldCreateTime), sql.Raw(problems.C(codingproblem.FieldDeadline))),
					sql.LT(s.C(codingsubmission.FieldCreateTime), sql.Raw(extensions.C(codingextension.FieldDeadline))),
				)).

			// join with user
			Join(users).
			On(s.C(codingsubmission.AuthorColumn), users.C(user.FieldID)).

			// aggregate and select
			GroupBy(codingsubmission.AuthorColumn, user.FieldEmail, users.C(user.FieldName)).
			Select(user.FieldEmail, users.C(user.FieldName), sql.Max(codingsubmission.FieldPoints))
	}).Scan(ctx, &scores)
	return scores, err
}
