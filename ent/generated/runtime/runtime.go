// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/user"
	"170-ag/ent/schema"
	"context"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	codingdraft.Policy = privacy.NewPolicies(schema.CodingDraft{})
	codingdraft.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingdraft.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingproblem.Policy = privacy.NewPolicies(schema.CodingProblem{})
	codingproblem.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingproblem.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingproblemFields := schema.CodingProblem{}.Fields()
	_ = codingproblemFields
	// codingproblemDescName is the schema descriptor for name field.
	codingproblemDescName := codingproblemFields[0].Descriptor()
	// codingproblem.NameValidator is a validator for the "name" field. It is called by the builders before save.
	codingproblem.NameValidator = func() func(string) error {
		validators := codingproblemDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// codingproblemDescStatement is the schema descriptor for statement field.
	codingproblemDescStatement := codingproblemFields[1].Descriptor()
	// codingproblem.DefaultStatement holds the default value on creation for the statement field.
	codingproblem.DefaultStatement = codingproblemDescStatement.Default.(string)
	// codingproblem.StatementValidator is a validator for the "statement" field. It is called by the builders before save.
	codingproblem.StatementValidator = codingproblemDescStatement.Validators[0].(func(string) error)
	// codingproblemDescReleased is the schema descriptor for released field.
	codingproblemDescReleased := codingproblemFields[2].Descriptor()
	// codingproblem.DefaultReleased holds the default value on creation for the released field.
	codingproblem.DefaultReleased = codingproblemDescReleased.Default.(bool)
	codingsubmission.Policy = privacy.NewPolicies(schema.CodingSubmission{})
	codingsubmission.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingsubmission.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingsubmissionFields := schema.CodingSubmission{}.Fields()
	_ = codingsubmissionFields
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescIsStaff is the schema descriptor for is_staff field.
	userDescIsStaff := userFields[2].Descriptor()
	// user.DefaultIsStaff holds the default value on creation for the is_staff field.
	user.DefaultIsStaff = userDescIsStaff.Default.(bool)
}

const (
	Version = "(devel)" // Version of ent codegen.
)