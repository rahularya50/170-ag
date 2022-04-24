package privacyrules

import (
	"170-ag/ent/generated/privacy"
	"context"
)

type accessTokenSentinel struct{}

type privacyAccessToken int

var JudgeScalingServerAccessToken privacyAccessToken = 0
var SubmissionEnqueuingAccessToken privacyAccessToken = 1
var FullSubmissionTestCaseAccessToken privacyAccessToken = 2
var GradescopeProjectSubmissionAccessToken privacyAccessToken = 3
var CloudflareCacheAccessToken privacyAccessToken = 4

func AllowWithPrivacyAccessToken(token privacyAccessToken) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(c context.Context) error {
		_, ok := c.Value(token).(accessTokenSentinel)
		if ok {
			return privacy.Allow
		} else {
			return privacy.Skip
		}
	})
}

func NewContextWithAccessToken(ctx context.Context, token privacyAccessToken) context.Context {
	return context.WithValue(ctx, token, accessTokenSentinel{})
}
