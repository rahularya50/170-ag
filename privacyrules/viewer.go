package privacyrules

import (
	ent "170-ag/ent/generated"
	"170-ag/ent/generated/privacy"
	"170-ag/site"
	"context"
)

func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(c context.Context) error {
		_, ok := site.ViewerFromContext(c)
		if !ok {
			return privacy.Deny
		}
		return privacy.Skip
	})
}

func AllowQueryIfIDsMatchViewer(get_ids func(context.Context, ent.Query) ([]int, error)) privacy.QueryRule {
	return privacy.QueryRuleFunc(func(c context.Context, q ent.Query) error {
		viewer, ok := site.ViewerFromContext(c)
		if !ok {
			return privacy.Skip
		}
		ids, err := get_ids(c, q)
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
		viewer, ok := site.ViewerFromContext(c)
		if !ok {
			return privacy.Deny
		}
		return filter_id(c, f, viewer.ID)
	})
}
