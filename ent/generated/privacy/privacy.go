// Code generated by entc, DO NOT EDIT.

package privacy

import (
	"170-ag/ent/generated"
	"context"
	"fmt"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with an allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with an deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, generated.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	return f(ctx, q)
}

type (
	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, generated.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	return f(ctx, m)
}

// Policy groups query and mutation policies.
type Policy struct {
	Query    QueryPolicy
	Mutation MutationPolicy
}

// EvalQuery forwards evaluation to query a policy.
func (policy Policy) EvalQuery(ctx context.Context, q generated.Query) error {
	return policy.Query.EvalQuery(ctx, q)
}

// EvalMutation forwards evaluation to mutate a  policy.
func (policy Policy) EvalMutation(ctx context.Context, m generated.Mutation) error {
	return policy.Mutation.EvalMutation(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, generated.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, generated.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ generated.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ generated.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op generated.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m generated.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op generated.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m generated.Mutation) error {
		return Denyf("generated/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The CodingDraftQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CodingDraftQueryRuleFunc func(context.Context, *generated.CodingDraftQuery) error

// EvalQuery return f(ctx, q).
func (f CodingDraftQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.CodingDraftQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.CodingDraftQuery", q)
}

// The CodingDraftMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CodingDraftMutationRuleFunc func(context.Context, *generated.CodingDraftMutation) error

// EvalMutation calls f(ctx, m).
func (f CodingDraftMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.CodingDraftMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.CodingDraftMutation", m)
}

// The CodingProblemQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CodingProblemQueryRuleFunc func(context.Context, *generated.CodingProblemQuery) error

// EvalQuery return f(ctx, q).
func (f CodingProblemQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.CodingProblemQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.CodingProblemQuery", q)
}

// The CodingProblemMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CodingProblemMutationRuleFunc func(context.Context, *generated.CodingProblemMutation) error

// EvalMutation calls f(ctx, m).
func (f CodingProblemMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.CodingProblemMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.CodingProblemMutation", m)
}

// The CodingSubmissionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CodingSubmissionQueryRuleFunc func(context.Context, *generated.CodingSubmissionQuery) error

// EvalQuery return f(ctx, q).
func (f CodingSubmissionQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.CodingSubmissionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.CodingSubmissionQuery", q)
}

// The CodingSubmissionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CodingSubmissionMutationRuleFunc func(context.Context, *generated.CodingSubmissionMutation) error

// EvalMutation calls f(ctx, m).
func (f CodingSubmissionMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.CodingSubmissionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.CodingSubmissionMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *generated.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *generated.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.UserMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q generated.Query) (Filter, error) {
	switch q := q.(type) {
	case *generated.CodingDraftQuery:
		return q.Filter(), nil
	case *generated.CodingProblemQuery:
		return q.Filter(), nil
	case *generated.CodingSubmissionQuery:
		return q.Filter(), nil
	case *generated.UserQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m generated.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *generated.CodingDraftMutation:
		return m.Filter(), nil
	case *generated.CodingProblemMutation:
		return m.Filter(), nil
	case *generated.CodingSubmissionMutation:
		return m.Filter(), nil
	case *generated.UserMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
