package project

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/projectscore"
	"170-ag/ent/generated/projectteam"
	"170-ag/privacyrules"
	"170-ag/proto/schemas"
	"170-ag/site"
	"context"
	"crypto/hmac"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthProjectRequest(client *ent.Client, gradescopeToken string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		metadata, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Metadata not found")
		}
		tokens, ok := metadata["authorization-token"]
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Token not found")
		}
		token := tokens[0]
		if !hmac.Equal([]byte(token), []byte(gradescopeToken)) {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
		}

		ctx = privacyrules.NewContextWithAccessToken(ctx, privacyrules.GradescopeProjectSubmissionAccessToken)
		ctx = site.ContextWithEntClient(ctx, client)
		return handler(ctx, req)
	}
}

type ProjectScoresServer struct {
	schemas.UnimplementedProjectScoresServer
	Client *ent.Client
}

func (s *ProjectScoresServer) RecordSubmission(ctx context.Context, submission *schemas.ProjectSubmission) (*schemas.SubmissionReply, error) {
	tx, err := s.Client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	team, err := tx.ProjectTeam.Create().
		SetTeamID(submission.TeamId).
		SetName(submission.TeamName).
		OnConflictColumns(projectteam.FieldName).
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return &schemas.SubmissionReply{Status: schemas.Status_DUPLICATE_TEAM_NAME}, nil
		}
		return nil, err
	}

	old_scores, err := tx.ProjectTeam.Query().
		Where(projectteam.ID(team)).
		QueryScores().
		All(ctx)
	if err != nil {
		return nil, err
	}

	score_lookup := make(map[caseKey]float64)
	for _, score := range old_scores {
		score_lookup[caseKey{CaseID: score.CaseID, CaseType: score.Type}] = score.Score
	}

	scores := make([]*ent.ProjectScoreCreate, len(submission.Score))
	to_invalidate := make([]caseKey, 0)
	for i, score := range submission.Score {
		caseType := projectscore.Type(strings.ToLower(score.CaseType.String()))
		if err := projectscore.TypeValidator(caseType); err != nil {
			return nil, err
		}
		key := caseKey{CaseID: score.CaseId, CaseType: caseType}
		best_score, ok := score_lookup[key]
		if !ok || best_score > score.Score {
			best_score = score.Score
			to_invalidate = append(to_invalidate, key)
		}
		scores[i] = tx.ProjectScore.Create().
			SetTeamID(team).
			SetCaseID(score.CaseId).
			SetType(caseType).
			SetScore(best_score)
	}

	if len(scores) > 0 {
		_, err = tx.ProjectScore.
			Delete().
			Where(projectscore.HasTeamWith(projectteam.ID(team))).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		err = tx.ProjectScore.CreateBulk(scores...).Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	err = invalidateCases(to_invalidate)
	if err != nil {
		log.Default().Println(err.Error())
		return &schemas.SubmissionReply{Status: schemas.Status_CACHE_INVALIDATION_FAILURE}, nil
	}

	return &schemas.SubmissionReply{Status: schemas.Status_SUCCESS}, nil
}
