package project

import (
	"170-ag/ent/generated/projectscore"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var cloudflareZone = os.Getenv("CLOUDFLARE_ZONE")
var leaderboardURL = os.Getenv("LEADERBOARD_URL")

const maxFineGrainedInvalidations = 20

type invalidationPayload struct {
	PurgeEverything bool     `json:"purge_everything,omitempty"`
	Files           []string `json:"files,omitempty"`
}

func invalidateWorker(payload invalidationPayload) error {
	endpoint := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/purge_cache", cloudflareZone)
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
	return fmt.Sprintf("%s/%s/%d", leaderboardURL, testCase.CaseType, testCase.CaseID)
}

func invalidateList(cases []caseKey) error {
	files := make([]string, len(cases))
	for i, testCase := range cases {
		files[i] = urlOfCase(testCase)
	}
	for _, size := range []projectscore.Type{projectscore.TypeSmall, projectscore.TypeSmall, projectscore.TypeSmall} {
		files = append(files, fmt.Sprintf("%s/%s", leaderboardURL, size))
	}
	files = append(files, leaderboardURL)
	return invalidateWorker(invalidationPayload{Files: files})
}

func invalidateAll() error {
	return invalidateWorker(invalidationPayload{PurgeEverything: true})
}

func invalidateCases(cases []caseKey) error {
	switch {
	case len(cases) == 0:
		return nil
	case len(cases) <= maxFineGrainedInvalidations:
		return invalidateList(cases)
	default:
		return invalidateAll()
	}
}
