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

type scoreboardHandler struct {
	client  *ent.Client
	forTeam bool
}

func ScoreboardHandler(client *ent.Client, forTeam bool) *scoreboardHandler {
	return &scoreboardHandler{client: client, forTeam: forTeam}
}

func (s *scoreboardHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if os.Getenv("ENV") == "dev" {
		r = r.WithContext(privacyrules.NewContextWithAccessToken(
			r.Context(), privacyrules.CloudflareCacheAccessToken,
		))
	}
	parts := strings.Split(r.URL.Path, "/")

	// path begins with /scoreboard/
	parts = parts[2:]
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	var raw_case_id, raw_case_type, raw_team_name string
	sort_by_rank := true

	if s.forTeam {
		switch len(parts) {
		case 1:
			raw_team_name = parts[0]
			sort_by_rank = false
		default:
			rw.WriteHeader(http.StatusNotFound)
			return
		}
	} else {
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
	}

	var filter scoreboardFilter

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

	if raw_team_name != "" {
		parsed_team_name, err := url.QueryUnescape(raw_team_name)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		filter.team_name = &parsed_team_name
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
	sortScoreboard(scoreboard)

	out, err := json.Marshal(scoreboard)
	if err != nil {
		panic(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}
