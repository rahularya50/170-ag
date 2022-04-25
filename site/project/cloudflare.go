package project

import (
	"170-ag/ent/generated/projectscore"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var cloudflareZone = os.Getenv("CLOUDFLARE_ZONE")
var leaderboardURL = fmt.Sprintf("%s/scoreboard", os.Getenv("LEADERBOARD_URL"))
var teamURL = fmt.Sprintf("%s/team", os.Getenv("LEADERBOARD_URL"))

const maxFineGrainedInvalidations = 20

type invalidationPayload struct {
	PurgeEverything bool   `json:"purge_everything,omitempty"`
	Files           []file `json:"files,omitempty"`
}

type file struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func invalidateWorker(payload invalidationPayload) error {
	endpoint := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/purge_cache", cloudflareZone)
	for _, file := range payload.Files {
		file.Headers = make(map[string]string)
		file.Headers["Origin"] = os.Getenv("LEADERBOARD_ORIGIN")
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CLOUDFLARE_TOKEN")))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("cloudflare replied with failure error code: %s", data)
	}

	return nil
}

func urlOfCase(testCase caseKey) string {
	return fmt.Sprintf("%s/%s/%d/", leaderboardURL, testCase.CaseType, testCase.CaseID)
}

func invalidateList(cases []caseKey, team string) error {
	files := make([]file, len(cases))
	for i, testCase := range cases {
		files[i] = file{Url: urlOfCase(testCase)}
	}
	for _, size := range []projectscore.Type{projectscore.TypeSmall, projectscore.TypeSmall, projectscore.TypeSmall} {
		files = append(files, file{Url: fmt.Sprintf("%s/%s/", leaderboardURL, size)})
	}
	files = append(files, file{Url: fmt.Sprintf("%s/", leaderboardURL)})
	files = append(files, file{Url: fmt.Sprintf("%s/%s/", teamURL, url.QueryEscape(team))})
	return invalidateWorker(invalidationPayload{Files: files})
}

func invalidateAll() error {
	return invalidateWorker(invalidationPayload{PurgeEverything: true})
}

func invalidateCases(cases []caseKey, team string) error {
	switch {
	case len(cases) == 0:
		return nil
	case len(cases) <= maxFineGrainedInvalidations:
		return invalidateList(cases, team)
	default:
		return invalidateAll()
	}
}
