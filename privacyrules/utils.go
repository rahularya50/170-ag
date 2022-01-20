package privacyrules

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"context"
	"errors"
)

func AllowQueryIfSubPolicyPasses(rules ...privacy.QueryRule) privacy.QueryRule {
	return privacy.QueryRuleFunc(func(c context.Context, q ent.Query) error {
		for _, rule := range rules {
			switch ret := rule.EvalQuery(c, q); {
			case errors.Is(ret, privacy.Allow):
				return privacy.Allow // allow if the subpolicy passes
			case errors.Is(ret, privacy.Skip):
				continue // skip any indeterminate subpolicy rules
			case errors.Is(ret, privacy.Deny):
				return privacy.Skip // if the subpolicy fails, skip
			default:
				return privacy.Deny // if anything else happens, fail-safe and Deny
			}
		}
		return privacy.Skip
	})
}
