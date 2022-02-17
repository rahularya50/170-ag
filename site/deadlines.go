package site

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingextension"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/user"
	"context"
	"encoding/csv"
	"errors"
	"strings"
	"time"
)

const layout = "15:04:05 02/01/06"

func GetSubmissionDeadlineForStudent(ctx context.Context, problem *ent.CodingProblem, student *ent.User) (*time.Time, error) {
	extension, err := problem.QueryExtensions().Where(codingextension.HasStudentWith(user.IDEQ(student.ID))).Only(ctx)
	_, not_found := err.(*ent.NotFoundError)
	if err != nil && !not_found {
		return nil, err
	}
	if extension != nil {
		return &extension.Deadline, nil
	}
	return &problem.Deadline, nil
}

func SetProblemExtensions(ctx context.Context, client *ent.Client, problem *ent.CodingProblem, roster string) error {
	reader := csv.NewReader(strings.NewReader(roster))
	header, err := reader.Read()
	if err != nil {
		return err
	}
	email_index := -1
	deadline_index := -1
	for index, heading := range header {
		switch strings.ToLower(heading) {
		case "email":
			email_index = index
		case "deadline":
			deadline_index = index
		}
	}
	if email_index == -1 || deadline_index == -1 {
		return errors.New("missing email or deadline column")
	}

	tx, err := client.Tx(ctx)
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.CodingExtension.Delete().
		Where(codingextension.HasCodingProblemWith(codingproblem.ID(problem.ID))).
		Exec(ctx)
	if err != nil {
		return err
	}

	builders := []*ent.CodingExtensionCreate{}

	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, row := range rows {
		email := row[email_index]
		deadline, err := time.Parse(layout, row[deadline_index])
		if err != nil {
			return err
		}
		student, err := tx.User.Query().Where(user.Email(email)).Only(ctx)
		_, not_found := err.(*ent.NotFoundError)
		if not_found {
			// create the user
			student, err = tx.User.Create().SetEmail(email).Save(ctx)
			if err != nil {
				return err
			}
		}
		builders = append(
			builders,
			tx.CodingExtension.Create().
				SetCodingProblem(problem).
				SetStudent(student).
				SetDeadline(deadline),
		)
	}
	err = tx.CodingExtension.CreateBulk(builders...).Exec(ctx)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func ExportProblemExtensions(ctx context.Context, problem *ent.CodingProblem) (string, error) {
	out := &strings.Builder{}
	writer := csv.NewWriter(out)
	writer.Write([]string{"Email", "Deadline"})
	extensions, err := problem.QueryExtensions().WithStudent().All(ctx)
	if err != nil {
		return "", err
	}
	for _, extension := range extensions {
		writer.Write([]string{extension.Edges.Student.Email, extension.Deadline.Format(layout)})
	}
	writer.Flush()
	if err = writer.Error(); err != nil {
		return "", err
	}
	return out.String(), nil
}
