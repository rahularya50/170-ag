package site

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/codingsubmission"
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"time"
)

func NumEnqueuedSubmissions(c context.Context, client *ent.Client) (int, error) {
	return client.CodingSubmission.
		Query().
		Where(codingsubmission.StatusEQ(codingsubmission.StatusQueued)).
		Count(c)
}

func DequeueSubmission(c context.Context, client *ent.Client) (*ent.CodingSubmission, *ent.CodingSubmissionStaffData, error) {
	tx, err := client.Tx(c)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	submission, err := tx.CodingSubmission.
		Query().
		Where(codingsubmission.StatusEQ(codingsubmission.StatusQueued)).
		First(c)
	if err != nil {
		return nil, nil, err
	}
	nonce, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return nil, nil, err
	}
	staff_data, err := submission.StaffData(c)
	if err != nil {
		return nil, nil, err
	}
	staff_data, err = staff_data.Update().SetExecutionID(nonce.Int64()).Save(c)
	if err != nil {
		return nil, nil, err
	}
	submission, err = submission.Update().SetStatus(codingsubmission.StatusRunning).Save(c)
	if err != nil {
		return nil, nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, nil, err
	}
	return submission.Unwrap(), staff_data.Unwrap(), nil
}

func CleanupStalled(c context.Context, client *ent.Client, timeout time.Duration) error {
	return client.CodingSubmission.Update().
		Where(
			codingsubmission.StatusEQ(codingsubmission.StatusRunning),
			codingsubmission.UpdateTimeLT(time.Now().Add(-timeout))).
		SetStatus(codingsubmission.StatusInternalError).
		Exec(c)
}
