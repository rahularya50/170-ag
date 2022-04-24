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

	_, err = tx.ProjectTeam.Delete().Where(projectteam.TeamID(submission.TeamId)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	team, err := tx.ProjectTeam.Create().
		SetTeamID(submission.TeamId).
		SetName(submission.TeamName).
		Save(ctx)
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return &schemas.SubmissionReply{Status: schemas.Status_DUPLICATE_TEAM_NAME}, nil
		}
		return nil, err
	}

	scores := make([]*ent.ProjectScoreCreate, len(submission.Score))
	for i, score := range submission.Score {
		scores[i] = tx.ProjectScore.Create().
			SetTeam(team).
			SetCaseID(score.CaseId).
			SetType(projectscore.Type(strings.ToLower(score.CaseType.String()))).
			SetScore(score.Score)
		if err != nil {
			return nil, err
		}
	}

	_, err = tx.ProjectScore.CreateBulk(scores...).Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &schemas.SubmissionReply{Status: schemas.Status_SUCCESS}, nil
}
