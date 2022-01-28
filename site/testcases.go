package site

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/models"
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GenerateInput(test_case_data []*ent.CodingTestCaseData) (string, error) {
	totalCases := 0
	cases := []string{}
	for _, data := range test_case_data {
		scanner := bufio.NewScanner(strings.NewReader(data.Input))
		if !scanner.Scan() {
			return "", errors.New("empty input test case")
		}
		count, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "", err
		}
		totalCases += count
		for scanner.Scan() {
			cases = append(cases, scanner.Text())
		}
	}
	return strconv.Itoa(totalCases) + "\n" + strings.Join(cases, "\n") + "\n", nil
}

func ScoreOutput(
	test_case_data []*ent.CodingTestCase,
	output string,
	defaultFailResult models.CodingSubmissionResult,
) models.CodingSubmissionResults {
	outputs := bufio.NewScanner(strings.NewReader(output))
	maxSize := len(output) + 5
	buffer := make([]byte, 0, maxSize)
	outputs.Buffer(buffer, maxSize)
	results := models.CodingSubmissionResults{}
	for i, test_case := range test_case_data {
		expected := bufio.NewScanner(strings.NewReader(test_case.Edges.Data.Output))
		maxSize := len(test_case.Edges.Data.Output) + 5
		buffer := make([]byte, 0, maxSize)
		expected.Buffer(buffer, maxSize)
		result := models.ResultAccepted
		points := test_case.Points
		for expected.Scan() {
			if !outputs.Scan() {
				result = defaultFailResult
				points = 0
				break
			}
			if strings.Trim(expected.Text(), " ") != strings.Trim(outputs.Text(), " ") {
				result = models.ResultWrongAnswer
				points = 0
			}
		}
		results.CaseResults = append(
			results.CaseResults,
			models.CodingSubmissionCaseResults{
				CaseName: fmt.Sprintf("Case %d", i+1),
				Result:   result,
				Points:   points,
			},
		)
	}
	return results
}
