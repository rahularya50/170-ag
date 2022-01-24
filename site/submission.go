package site

import (
	ent "170-ag/ent/generated"
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func GenerateInput(test_case_data []*ent.CodingTestCaseData) (string, error) {
	totalCases := 0
	cases := []string{}
	for _, data := range test_case_data {
		reader := bufio.NewScanner(strings.NewReader(data.Input))
		if !reader.Scan() {
			return "", errors.New("empty input test case")
		}
		count, err := strconv.Atoi(reader.Text())
		if err != nil {
			return "", err
		}
		totalCases += count
		for reader.Scan() {
			cases = append(cases, reader.Text())
		}
	}
	return strconv.Itoa(totalCases) + "\n" + strings.Join(cases, "\n") + "\n", nil
}
