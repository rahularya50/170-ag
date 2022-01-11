package privacyrules

import (
	"170-ag/ent/generated/privacy"
	"170-ag/site"
	"context"
)

func AllowIfViewerIsStaff() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(c context.Context) error {
		viewer, ok := site.ViewerFromContext(c)
		if !ok {
			return privacy.Skip
		}
		if viewer.IsStaff {
			return privacy.Allow
		}
		return privacy.Skip
	})
}
