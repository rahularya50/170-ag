package site

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var endpoint = os.Getenv("ENDPOINT")

type okpyRosterReply struct {
	Code    int
	Message string
	Data    map[string][]struct {
		Id    string
		Email string
	}
}

func IsStaff(email string) (bool, error) {
	if os.Getenv("ENV") == "dev" {
		return true, nil
	}
	okpy, _ := url.Parse(fmt.Sprintf("https://okpy.org/api/v3/course/%s/export/enrollment", endpoint))
	query := okpy.Query()
	query.Set("export_token", os.Getenv("EXPORT_TOKEN"))
	okpy.RawQuery = query.Encode()
	resp, err := http.Get(okpy.String())
	if err != nil {
		return false, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var rosterReply okpyRosterReply
	err = json.Unmarshal(data, &rosterReply)
	if err != nil {
		return false, err
	}
	for _, role := range []string{"staff", "grader", "instructor"} {
		for _, user := range rosterReply.Data[role] {
			if user.Email == email {
				return true, nil
			}
		}
	}
	return false, nil
}
