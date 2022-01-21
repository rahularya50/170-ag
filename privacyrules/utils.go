package privacyrules

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/site"
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

func DenyQueryIfSubPolicyFails(rules ...privacy.QueryRule) privacy.QueryRule {
	return privacy.QueryRuleFunc(func(c context.Context, q ent.Query) error {
		for _, rule := range rules {
			switch ret := rule.EvalQuery(c, q); {
			case errors.Is(ret, privacy.Allow):
				return privacy.Skip // if the subpolicy passes, skip
			case errors.Is(ret, privacy.Skip):
				continue // skip any indeterminate subpolicy rules
			case errors.Is(ret, privacy.Deny):
				return privacy.Deny // deny if the subpolicy fails
			default:
				return privacy.Deny // if anything else happens, fail-safe and Deny
			}
		}
		return privacy.Skip
	})
}

func AllowQueryIfDelegatesVisible(
	getDelegateIDs func(allow_ctx context.Context, q ent.Query) ([]int, error),
	verifyDelegateIDs func(client *ent.Client, ctx context.Context, ids []int) ([]int, error),
) privacy.QueryRule {
	return privacy.QueryRuleFunc(func(c context.Context, q ent.Query) error {
		allow_ctx := privacy.DecisionContext(c, privacy.Allow)
		delegateIDs, err := getDelegateIDs(allow_ctx, q)
		if err != nil {
			return privacy.Skip
		}
		verifiedIDs, err := verifyDelegateIDs(site.EntClientFromContext(c), c, delegateIDs)
		if err != nil || len(delegateIDs) != len(verifiedIDs) /* make sure nothing got filtered */ {
			return privacy.Skip
		}
		return privacy.Allow
	})
}
