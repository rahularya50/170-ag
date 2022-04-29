package project

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/projectscore"
	"context"
	"sort"
)

type Scoreboard struct {
	Entries []ScoreboardEntry
}

type ScoreboardEntry struct {
	TeamName     string
	TeamScore    roundedFloat
	TeamScoreStr string
}

func NewScoreboardEntry(name string, score roundedFloat) ScoreboardEntry {
	return ScoreboardEntry{
		TeamName:     name,
		TeamScore:    score,
		TeamScoreStr: score.String(),
	}
}

type TeamScoreboard struct {
	Entries []TeamScoreboardEntry
}

type TeamScoreboardEntry struct {
	TeamScore    roundedFloat
	TeamScoreStr string
	TeamRank     int
	TestCase     caseKey
}

func NewTeamScoreboardEntry(score roundedFloat, rank int, testCase caseKey) TeamScoreboardEntry {
	return TeamScoreboardEntry{
		TeamScore:    score,
		TeamScoreStr: score.String(),
		TeamRank:     rank,
		TestCase:     testCase,
	}
}

type caseKey struct {
	CaseID   int32
	CaseType projectscore.Type
}

type scoreboardFilter struct {
	case_id   *int32
	case_type *projectscore.Type
}

func queryScores(ctx context.Context, client *ent.Client, filter scoreboardFilter) ([]*ent.ProjectScore, error) {
	query := client.ProjectScore.Query()
	if filter.case_type != nil {
		query = query.Where(projectscore.TypeEQ(*filter.case_type))
	}
	if filter.case_id != nil {
		query = query.Where(projectscore.CaseID(*filter.case_id))
	}
	return query.WithTeam().All(ctx)
}

func getAllRanks(scores []*ent.ProjectScore) map[string]map[caseKey]int {
	scoresByCase := make(map[caseKey][]*ent.ProjectScore)
	for _, score := range scores {
		key := caseKey{CaseID: score.CaseID, CaseType: score.Type}
		scoresByCase[key] = append(scoresByCase[key], score)
	}
	for _, caseScores := range scoresByCase {
		// sort in ascending order of score
		sort.Slice(caseScores,
			func(i, j int) bool {
				return roundFloat(caseScores[i].Score) < roundFloat(caseScores[j].Score)
			},
		)
	}
	ranks := make(map[string]map[caseKey]int) // teamName -> case -> rank
	for testCase, caseScores := range scoresByCase {
		currScore := roundFloat(caseScores[0].Score)
		currRank := 0
		for i, score := range caseScores {
			if roundFloat(score.Score) > currScore {
				currScore = roundFloat(score.Score)
				currRank = i
			}
			if ranks[score.Edges.Team.Name] == nil {
				ranks[score.Edges.Team.Name] = make(map[caseKey]int)
			}
			ranks[score.Edges.Team.Name][testCase] = currRank
		}
	}
	return ranks
}

func scoreByRank(scores []*ent.ProjectScore) *Scoreboard {
	scoreboard := &Scoreboard{Entries: []ScoreboardEntry{}}
	ranks := getAllRanks(scores)
	for teamName, allRanks := range ranks {
		totalRank := 0
		for _, rank := range allRanks {
			totalRank += rank
		}
		scoreboard.Entries = append(scoreboard.Entries, NewScoreboardEntry(
			teamName,
			roundFloat(float64(totalRank)/float64(len(allRanks))),
		))
	}
	return scoreboard
}

func scoreByPoints(scores []*ent.ProjectScore) *Scoreboard {
	scoreboard := &Scoreboard{Entries: []ScoreboardEntry{}}
	for _, score := range scores {
		scoreboard.Entries = append(scoreboard.Entries, NewScoreboardEntry(
			score.Edges.Team.Name,
			roundFloat(score.Score),
		))
	}
	return scoreboard
}

func scoreTeamPointsAndRank(scores []*ent.ProjectScore, ranks map[string]map[caseKey]int, team string) *TeamScoreboard {
	scoreboard := &TeamScoreboard{Entries: []TeamScoreboardEntry{}}
	for _, score := range scores {
		if score.Edges.Team.Name != team {
			continue
		}
		testCase := caseKey{CaseID: score.CaseID, CaseType: score.Type}
		scoreboard.Entries = append(
			scoreboard.Entries,
			NewTeamScoreboardEntry(
				roundedFloat(score.Score),
				ranks[team][testCase],
				testCase,
			))
	}
	return scoreboard
}
