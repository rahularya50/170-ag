package project

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/projectscore"
	"170-ag/privacyrules"
	"crypto/hmac"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func HandlerCheckingAuthorizationToken(h http.Handler, accessToken string) http.HandlerFunc {
	if os.Getenv("ENV") == "dev" {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			r = r.WithContext(privacyrules.NewContextWithAccessToken(
				r.Context(), privacyrules.CloudflareCacheAccessToken,
			))
			h.ServeHTTP(rw, r)
		})
	} else {
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
}

func getRequestArgs(r *http.Request) []string {
	parts := strings.Split(r.URL.Path, "/")

	// path begins with /scoreboard/
	parts = parts[2:]
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	return parts
}

type scoreboardHandler struct {
	client *ent.Client
}

func ScoreboardHandler(client *ent.Client) *scoreboardHandler {
	return &scoreboardHandler{client: client}
}

func (s *scoreboardHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	parts := getRequestArgs(r)

	var raw_case_id, raw_case_type string
	sort_by_rank := true

	switch len(parts) {
	case 0:
		/* no-op */
	case 1:
		raw_case_type = parts[0]
	case 2:
		raw_case_type = parts[0]
		raw_case_id = parts[1]
		sort_by_rank = false
	default:
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	filter := scoreboardFilter{}

	if raw_case_id != "" {
		parsed_case_id, err := strconv.Atoi(raw_case_id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		parsed_case_id_i32 := int32(parsed_case_id)
		filter.case_id = &parsed_case_id_i32
	}

	if raw_case_type != "" {
		parsed_case_type := projectscore.Type(raw_case_type)
		if err := projectscore.TypeValidator(parsed_case_type); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		filter.case_type = &parsed_case_type
	}

	scores, err := queryScores(r.Context(), s.client, filter)
	if err != nil {
		panic(err)
	}

	var scoreboard *Scoreboard
	if sort_by_rank {
		scoreboard = scoreByRank(scores)
	} else {
		scoreboard = scoreByPoints(scores)
	}

	out, err := json.Marshal(scoreboard)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}

type teamScoreboardHandler struct {
	client *ent.Client
}

func TeamScoreboardHandler(client *ent.Client) *teamScoreboardHandler {
	return &teamScoreboardHandler{client: client}
}

func (s *teamScoreboardHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	parts := getRequestArgs(r)

	var raw_team_name string
	switch len(parts) {
	case 1:
		raw_team_name = parts[0]
	default:
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var team_name *string
	if raw_team_name != "" {
		parsed_team_name, err := url.QueryUnescape(raw_team_name)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		team_name = &parsed_team_name
	}

	scores, err := queryScores(r.Context(), s.client, scoreboardFilter{})
	if err != nil {
		panic(err)
	}

	ranks := getAllRanks(scores)
	scoreboard := scoreTeamPointsAndRank(scores, ranks, *team_name)

	out, err := json.Marshal(scoreboard)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}
