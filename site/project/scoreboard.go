package project

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/projectscore"
	"170-ag/privacyrules"
	"context"
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
		// clientToken := r.Header.Get("X-Authorization-Token")
		// if clientToken == "" {
		// 	rw.WriteHeader(http.StatusUnauthorized)
		// } else if !hmac.Equal([]byte(accessToken), []byte(clientToken)) {
		// 	rw.WriteHeader(http.StatusForbidden)
		// } else
		{
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

type caseKey struct {
	caseID   int32
	caseType projectscore.Type
}

func ExportScoreboard(ctx context.Context, client *ent.Client, case_id *int32, case_type *projectscore.Type) (*Scoreboard, error) {
	query := client.ProjectScore.Query()
	scoreboard := &Scoreboard{Entries: []ScoreboardEntry{}}

	if case_type != nil {
		query = query.Where(projectscore.TypeEQ(*case_type))
	}
	if case_id == nil || case_type == nil {
		scores, err := query.WithTeam().All(ctx)
		if err != nil {
			return nil, err
		}
		scoresByCase := make(map[caseKey][]*ent.ProjectScore)
		for _, score := range scores {
			key := caseKey{caseID: score.CaseID, caseType: score.Type}
			scoresByCase[key] = append(scoresByCase[key], score)
		}
		for _, caseScores := range scoresByCase {
			// sort in ascending order of score
			sort.Slice(caseScores,
				func(i, j int) bool {
					return caseScores[i].Score < caseScores[j].Score
				},
			)
		}
		totalRanks := make(map[string]float64)
		for _, caseScores := range scoresByCase {
			currScore := caseScores[0].Score
			currRank := 0
			for _, score := range caseScores {
				if score.Score > currScore {
					currScore = score.Score
					currRank += 1
				}
				totalRanks[score.Edges.Team.Name] += float64(currRank)
			}
		}
		for teamName, totalRank := range totalRanks {
			scoreboard.Entries = append(scoreboard.Entries, ScoreboardEntry{
				TeamName:  teamName,
				TeamScore: totalRank / float64(len(totalRanks)),
			})
		}
	} else {
		// filter to single case
		query = query.Where(projectscore.CaseID(*case_id))
		scores, err := query.WithTeam().All(ctx)
		if err != nil {
			return nil, err
		}
		for _, score := range scores {
			scoreboard.Entries = append(scoreboard.Entries, ScoreboardEntry{
				TeamName:  score.Edges.Team.Name,
				TeamScore: score.Score,
			})
		}
	}
	// break ties by team name
	sort.Slice(scoreboard.Entries,
		func(i, j int) bool {
			return scoreboard.Entries[i].TeamName < scoreboard.Entries[j].TeamName
		},
	)
	// then sort by increasing score (i.e. cost)
	sort.SliceStable(scoreboard.Entries,
		func(i, j int) bool {
			return scoreboard.Entries[i].TeamScore < scoreboard.Entries[j].TeamScore
		},
	)
	return scoreboard, nil
}
