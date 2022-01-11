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

type clonableEntQuery interface {
	ent.Query
	Clone() ent.Query
}

func AllowIfIDMatchesViewer(get_id func(context.Context, *ent.Query) (int, error)) privacy.QueryRule {
	return privacy.QueryRuleFunc(func(c context.Context, q ent.Query) error {
		viewer, ok := site.ViewerFromContext(c)
		if !ok {
			return privacy.Skip
		}
		cloned := q.(clonableEntQuery).Clone()
		id, err := get_id(c, &cloned)
		if err != nil {
			return privacy.Skip
		}
		if viewer.ID != id {
			return privacy.Skip
		}
		return privacy.Allow
	})
}
