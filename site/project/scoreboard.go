package project

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/projectscore"
	"170-ag/ent/generated/projectteam"
	"context"
	"sort"
)

type Scoreboard struct {
	Entries []ScoreboardEntry
}

type ScoreboardEntry struct {
	TeamName  string
	TeamScore float64
	TestCase  caseKey
}

type caseKey struct {
	CaseID   int32
	CaseType projectscore.Type
}

type scoreboardFilter struct {
	case_id   *int32
	case_type *projectscore.Type
	team_name *string
}

func queryScores(ctx context.Context, client *ent.Client, filter scoreboardFilter) ([]*ent.ProjectScore, error) {
	query := client.ProjectScore.Query()
	if filter.case_type != nil {
		query = query.Where(projectscore.TypeEQ(*filter.case_type))
	}
	if filter.case_id != nil {
		query = query.Where(projectscore.CaseID(*filter.case_id))
	}
	if filter.team_name != nil {
		query = query.Where(projectscore.HasTeamWith(projectteam.Name(*filter.team_name)))
	}
	return query.WithTeam().All(ctx)
}

func scoreByRank(scores []*ent.ProjectScore) *Scoreboard {
	scoreboard := &Scoreboard{Entries: []ScoreboardEntry{}}

	scoresByCase := make(map[caseKey][]*ent.ProjectScore)
	for _, score := range scores {
		key := caseKey{CaseID: score.CaseID, CaseType: score.Type}
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
	return scoreboard
}

func scoreByPoints(scores []*ent.ProjectScore) *Scoreboard {
	scoreboard := &Scoreboard{Entries: []ScoreboardEntry{}}
	for _, score := range scores {
		scoreboard.Entries = append(scoreboard.Entries, ScoreboardEntry{
			TeamName:  score.Edges.Team.Name,
			TeamScore: score.Score,
			TestCase:  caseKey{CaseID: score.CaseID, CaseType: score.Type},
		})
	}
	return scoreboard
}

func sortScoreboard(scoreboard *Scoreboard) {
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
}
