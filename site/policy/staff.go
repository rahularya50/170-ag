package policy

import (
	"170-ag/ent/generated/privacy"
	"170-ag/site/web"
	"context"
)

func AllowIfViewerIsStaff() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(c context.Context) error {
		viewer, ok := web.ViewerFromContext(c)
		if !ok {
			return privacy.Skip
		}
		if viewer.IsStaff {
			return privacy.Allow
		}
		return privacy.Skip
	})
}
