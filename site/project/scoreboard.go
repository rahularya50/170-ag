package project

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/projectscore"
	"170-ag/privacyrules"
	"context"
	"crypto/hmac"
	"encoding/json"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type Scoreboard struct {
	Entries []ScoreboardEntry
}

type ScoreboardEntry struct {
	TeamName  string
	TeamScore float64
}

type scoreboardHandler struct {
	client *ent.Client
}

func HandlerCheckingAuthorizationToken(h http.Handler, accessToken string) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("X-Authorization-Token")
		if clientToken == "" {
			rw.WriteHeader(http.StatusUnauthorized)
		} else if !hmac.Equal([]byte(accessToken), []byte(clientToken)) {
			rw.WriteHeader(http.StatusForbidden)
		} else {
			r = r.WithContext(privacyrules.NewContextWithAccessToken(
				r.Context(), privacyrules.CloudflareCacheAccessToken,
			))
			h.ServeHTTP(rw, r)
		}
	})

}

func ScoreboardHandler(client *ent.Client) *scoreboardHandler {
	return &scoreboardHandler{client: client}
}

func (s *scoreboardHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if os.Getenv("ENV") == "dev" {
		r = r.WithContext(privacyrules.NewContextWithAccessToken(
			r.Context(), privacyrules.CloudflareCacheAccessToken,
		))
	}
	values := r.URL.Query()
	raw_case_id := values.Get("caseID")
	var case_id *int32

	if raw_case_id != "" {
		parsed_case_id, err := strconv.Atoi(raw_case_id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		parsed_case_id_i32 := int32(parsed_case_id)
		case_id = &parsed_case_id_i32
	}

	raw_case_type := values.Get("caseType")
	var case_type *projectscore.Type
	if raw_case_type != "" {
		parsed_case_type := projectscore.Type(raw_case_type)
		if err := projectscore.TypeValidator(parsed_case_type); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		case_type = &parsed_case_type
	}

	scoreboard, err := ExportScoreboard(r.Context(), s.client, case_id, case_type)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	out, err := json.Marshal(scoreboard)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}

func ExportScoreboard(ctx context.Context, client *ent.Client, case_id *int32, case_type *projectscore.Type) (*Scoreboard, error) {
	query := client.ProjectScore.Query()
	if case_id != nil {
		query = query.Where(projectscore.CaseID(*case_id))
	}
	if case_type != nil {
		query = query.Where(projectscore.TypeEQ(*case_type))
	}
	scores, err := query.WithTeam().All(ctx)
	if err != nil {
		return nil, err
	}
	scoreboard := &Scoreboard{Entries: []ScoreboardEntry{}}
	for _, score := range scores {
		scoreboard.Entries = append(scoreboard.Entries, ScoreboardEntry{
			TeamName:  score.Edges.Team.Name,
			TeamScore: score.Score,
		})
	}
	// break ties by team name
	sort.Slice(scoreboard.Entries,
		func(i, j int) bool {
			return scoreboard.Entries[i].TeamName < scoreboard.Entries[j].TeamName
		},
	)
	// then sort by score
	sort.SliceStable(scoreboard.Entries,
		func(i, j int) bool {
			return scoreboard.Entries[i].TeamScore < scoreboard.Entries[j].TeamScore
		},
	)
	return scoreboard, nil
}