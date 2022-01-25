package policy

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/site/web"
	"context"
)

func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(c context.Context) error {
		_, ok := web.ViewerFromContext(c)
		if !ok {
			return privacy.Deny
		}
		return privacy.Skip
	})
}

func AllowQueryIfIDsMatchViewer(get_ids func(context.Context, ent.Query) ([]int, error)) privacy.QueryRule {
	return privacy.QueryRuleFunc(func(c context.Context, q ent.Query) error {
		viewer, ok := web.ViewerFromContext(c)
		if !ok {
			return privacy.Skip
		}
		allow_ctx := privacy.DecisionContext(c, privacy.Allow)
		ids, err := get_ids(allow_ctx, q)
		if err != nil {
			return privacy.Skip
		}
		for _, id := range ids {
			if viewer.ID != id {
				return privacy.Skip
			}
		}
		return privacy.Allow
	})
}

func FilterToViewerID(filter_id func(context.Context, privacy.Filter, int) error) privacy.MutationRule {
	return privacy.FilterFunc(func(c context.Context, f privacy.Filter) error {
		viewer, ok := web.ViewerFromContext(c)
		if !ok {
			return privacy.Deny
		}
		return filter_id(c, f, viewer.ID)
	})
}
