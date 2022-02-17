// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"170-ag/ent/generated/codingdraft"
	"170-ag/ent/generated/codingextension"
	"170-ag/ent/generated/codingproblem"
	"170-ag/ent/generated/codingsubmission"
	"170-ag/ent/generated/codingsubmissionstaffdata"
	"170-ag/ent/generated/codingtestcase"
	"170-ag/ent/generated/codingtestcasedata"
	"170-ag/ent/generated/user"
	"170-ag/ent/schema"
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	codingdraftMixin := schema.CodingDraft{}.Mixin()
	codingdraft.Policy = privacy.NewPolicies(schema.CodingDraft{})
	codingdraft.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingdraft.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingdraftMixinFields0 := codingdraftMixin[0].Fields()
	_ = codingdraftMixinFields0
	codingdraftFields := schema.CodingDraft{}.Fields()
	_ = codingdraftFields
	// codingdraftDescCreateTime is the schema descriptor for create_time field.
	codingdraftDescCreateTime := codingdraftMixinFields0[0].Descriptor()
	// codingdraft.DefaultCreateTime holds the default value on creation for the create_time field.
	codingdraft.DefaultCreateTime = codingdraftDescCreateTime.Default.(func() time.Time)
	// codingdraftDescUpdateTime is the schema descriptor for update_time field.
	codingdraftDescUpdateTime := codingdraftMixinFields0[1].Descriptor()
	// codingdraft.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingdraft.DefaultUpdateTime = codingdraftDescUpdateTime.Default.(func() time.Time)
	// codingdraft.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingdraft.UpdateDefaultUpdateTime = codingdraftDescUpdateTime.UpdateDefault.(func() time.Time)
	codingextensionMixin := schema.CodingExtension{}.Mixin()
	codingextension.Policy = privacy.NewPolicies(schema.CodingExtension{})
	codingextension.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingextension.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingextensionMixinFields0 := codingextensionMixin[0].Fields()
	_ = codingextensionMixinFields0
	codingextensionFields := schema.CodingExtension{}.Fields()
	_ = codingextensionFields
	// codingextensionDescCreateTime is the schema descriptor for create_time field.
	codingextensionDescCreateTime := codingextensionMixinFields0[0].Descriptor()
	// codingextension.DefaultCreateTime holds the default value on creation for the create_time field.
	codingextension.DefaultCreateTime = codingextensionDescCreateTime.Default.(func() time.Time)
	// codingextensionDescUpdateTime is the schema descriptor for update_time field.
	codingextensionDescUpdateTime := codingextensionMixinFields0[1].Descriptor()
	// codingextension.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingextension.DefaultUpdateTime = codingextensionDescUpdateTime.Default.(func() time.Time)
	// codingextension.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingextension.UpdateDefaultUpdateTime = codingextensionDescUpdateTime.UpdateDefault.(func() time.Time)
	codingproblemMixin := schema.CodingProblem{}.Mixin()
	codingproblem.Policy = privacy.NewPolicies(schema.CodingProblem{})
	codingproblem.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingproblem.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingproblemMixinFields0 := codingproblemMixin[0].Fields()
	_ = codingproblemMixinFields0
	codingproblemFields := schema.CodingProblem{}.Fields()
	_ = codingproblemFields
	// codingproblemDescCreateTime is the schema descriptor for create_time field.
	codingproblemDescCreateTime := codingproblemMixinFields0[0].Descriptor()
	// codingproblem.DefaultCreateTime holds the default value on creation for the create_time field.
	codingproblem.DefaultCreateTime = codingproblemDescCreateTime.Default.(func() time.Time)
	// codingproblemDescUpdateTime is the schema descriptor for update_time field.
	codingproblemDescUpdateTime := codingproblemMixinFields0[1].Descriptor()
	// codingproblem.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingproblem.DefaultUpdateTime = codingproblemDescUpdateTime.Default.(func() time.Time)
	// codingproblem.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingproblem.UpdateDefaultUpdateTime = codingproblemDescUpdateTime.UpdateDefault.(func() time.Time)
	// codingproblemDescName is the schema descriptor for name field.
	codingproblemDescName := codingproblemFields[0].Descriptor()
	// codingproblem.DefaultName holds the default value on creation for the name field.
	codingproblem.DefaultName = codingproblemDescName.Default.(string)
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
	// codingproblemDescSkeleton is the schema descriptor for skeleton field.
	codingproblemDescSkeleton := codingproblemFields[2].Descriptor()
	// codingproblem.DefaultSkeleton holds the default value on creation for the skeleton field.
	codingproblem.DefaultSkeleton = codingproblemDescSkeleton.Default.(string)
	// codingproblemDescReleased is the schema descriptor for released field.
	codingproblemDescReleased := codingproblemFields[3].Descriptor()
	// codingproblem.DefaultReleased holds the default value on creation for the released field.
	codingproblem.DefaultReleased = codingproblemDescReleased.Default.(bool)
	// codingproblemDescDeadline is the schema descriptor for deadline field.
	codingproblemDescDeadline := codingproblemFields[4].Descriptor()
	// codingproblem.DefaultDeadline holds the default value on creation for the deadline field.
	codingproblem.DefaultDeadline = codingproblemDescDeadline.Default.(func() time.Time)
	codingsubmissionMixin := schema.CodingSubmission{}.Mixin()
	codingsubmission.Policy = privacy.NewPolicies(schema.CodingSubmission{})
	codingsubmission.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingsubmission.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingsubmissionMixinFields0 := codingsubmissionMixin[0].Fields()
	_ = codingsubmissionMixinFields0
	codingsubmissionFields := schema.CodingSubmission{}.Fields()
	_ = codingsubmissionFields
	// codingsubmissionDescCreateTime is the schema descriptor for create_time field.
	codingsubmissionDescCreateTime := codingsubmissionMixinFields0[0].Descriptor()
	// codingsubmission.DefaultCreateTime holds the default value on creation for the create_time field.
	codingsubmission.DefaultCreateTime = codingsubmissionDescCreateTime.Default.(func() time.Time)
	// codingsubmissionDescUpdateTime is the schema descriptor for update_time field.
	codingsubmissionDescUpdateTime := codingsubmissionMixinFields0[1].Descriptor()
	// codingsubmission.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingsubmission.DefaultUpdateTime = codingsubmissionDescUpdateTime.Default.(func() time.Time)
	// codingsubmission.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingsubmission.UpdateDefaultUpdateTime = codingsubmissionDescUpdateTime.UpdateDefault.(func() time.Time)
	// codingsubmissionDescCode is the schema descriptor for code field.
	codingsubmissionDescCode := codingsubmissionFields[0].Descriptor()
	// codingsubmission.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	codingsubmission.CodeValidator = codingsubmissionDescCode.Validators[0].(func(string) error)
	codingsubmissionstaffdataMixin := schema.CodingSubmissionStaffData{}.Mixin()
	codingsubmissionstaffdata.Policy = privacy.NewPolicies(schema.CodingSubmissionStaffData{})
	codingsubmissionstaffdata.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingsubmissionstaffdata.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingsubmissionstaffdataMixinFields0 := codingsubmissionstaffdataMixin[0].Fields()
	_ = codingsubmissionstaffdataMixinFields0
	codingsubmissionstaffdataFields := schema.CodingSubmissionStaffData{}.Fields()
	_ = codingsubmissionstaffdataFields
	// codingsubmissionstaffdataDescCreateTime is the schema descriptor for create_time field.
	codingsubmissionstaffdataDescCreateTime := codingsubmissionstaffdataMixinFields0[0].Descriptor()
	// codingsubmissionstaffdata.DefaultCreateTime holds the default value on creation for the create_time field.
	codingsubmissionstaffdata.DefaultCreateTime = codingsubmissionstaffdataDescCreateTime.Default.(func() time.Time)
	// codingsubmissionstaffdataDescUpdateTime is the schema descriptor for update_time field.
	codingsubmissionstaffdataDescUpdateTime := codingsubmissionstaffdataMixinFields0[1].Descriptor()
	// codingsubmissionstaffdata.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingsubmissionstaffdata.DefaultUpdateTime = codingsubmissionstaffdataDescUpdateTime.Default.(func() time.Time)
	// codingsubmissionstaffdata.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingsubmissionstaffdata.UpdateDefaultUpdateTime = codingsubmissionstaffdataDescUpdateTime.UpdateDefault.(func() time.Time)
	// codingsubmissionstaffdataDescOutput is the schema descriptor for output field.
	codingsubmissionstaffdataDescOutput := codingsubmissionstaffdataFields[2].Descriptor()
	// codingsubmissionstaffdata.OutputValidator is a validator for the "output" field. It is called by the builders before save.
	codingsubmissionstaffdata.OutputValidator = codingsubmissionstaffdataDescOutput.Validators[0].(func(string) error)
	// codingsubmissionstaffdataDescStderr is the schema descriptor for stderr field.
	codingsubmissionstaffdataDescStderr := codingsubmissionstaffdataFields[3].Descriptor()
	// codingsubmissionstaffdata.StderrValidator is a validator for the "stderr" field. It is called by the builders before save.
	codingsubmissionstaffdata.StderrValidator = codingsubmissionstaffdataDescStderr.Validators[0].(func(string) error)
	// codingsubmissionstaffdataDescExitError is the schema descriptor for exit_error field.
	codingsubmissionstaffdataDescExitError := codingsubmissionstaffdataFields[4].Descriptor()
	// codingsubmissionstaffdata.ExitErrorValidator is a validator for the "exit_error" field. It is called by the builders before save.
	codingsubmissionstaffdata.ExitErrorValidator = codingsubmissionstaffdataDescExitError.Validators[0].(func(string) error)
	codingtestcaseMixin := schema.CodingTestCase{}.Mixin()
	codingtestcase.Policy = privacy.NewPolicies(schema.CodingTestCase{})
	codingtestcase.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingtestcase.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingtestcaseMixinFields0 := codingtestcaseMixin[0].Fields()
	_ = codingtestcaseMixinFields0
	codingtestcaseFields := schema.CodingTestCase{}.Fields()
	_ = codingtestcaseFields
	// codingtestcaseDescCreateTime is the schema descriptor for create_time field.
	codingtestcaseDescCreateTime := codingtestcaseMixinFields0[0].Descriptor()
	// codingtestcase.DefaultCreateTime holds the default value on creation for the create_time field.
	codingtestcase.DefaultCreateTime = codingtestcaseDescCreateTime.Default.(func() time.Time)
	// codingtestcaseDescUpdateTime is the schema descriptor for update_time field.
	codingtestcaseDescUpdateTime := codingtestcaseMixinFields0[1].Descriptor()
	// codingtestcase.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingtestcase.DefaultUpdateTime = codingtestcaseDescUpdateTime.Default.(func() time.Time)
	// codingtestcase.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingtestcase.UpdateDefaultUpdateTime = codingtestcaseDescUpdateTime.UpdateDefault.(func() time.Time)
	// codingtestcaseDescPoints is the schema descriptor for points field.
	codingtestcaseDescPoints := codingtestcaseFields[0].Descriptor()
	// codingtestcase.DefaultPoints holds the default value on creation for the points field.
	codingtestcase.DefaultPoints = codingtestcaseDescPoints.Default.(int)
	// codingtestcase.PointsValidator is a validator for the "points" field. It is called by the builders before save.
	codingtestcase.PointsValidator = codingtestcaseDescPoints.Validators[0].(func(int) error)
	codingtestcasedataMixin := schema.CodingTestCaseData{}.Mixin()
	codingtestcasedata.Policy = privacy.NewPolicies(schema.CodingTestCaseData{})
	codingtestcasedata.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := codingtestcasedata.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	codingtestcasedataMixinFields0 := codingtestcasedataMixin[0].Fields()
	_ = codingtestcasedataMixinFields0
	codingtestcasedataFields := schema.CodingTestCaseData{}.Fields()
	_ = codingtestcasedataFields
	// codingtestcasedataDescCreateTime is the schema descriptor for create_time field.
	codingtestcasedataDescCreateTime := codingtestcasedataMixinFields0[0].Descriptor()
	// codingtestcasedata.DefaultCreateTime holds the default value on creation for the create_time field.
	codingtestcasedata.DefaultCreateTime = codingtestcasedataDescCreateTime.Default.(func() time.Time)
	// codingtestcasedataDescUpdateTime is the schema descriptor for update_time field.
	codingtestcasedataDescUpdateTime := codingtestcasedataMixinFields0[1].Descriptor()
	// codingtestcasedata.DefaultUpdateTime holds the default value on creation for the update_time field.
	codingtestcasedata.DefaultUpdateTime = codingtestcasedataDescUpdateTime.Default.(func() time.Time)
	// codingtestcasedata.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	codingtestcasedata.UpdateDefaultUpdateTime = codingtestcasedataDescUpdateTime.UpdateDefault.(func() time.Time)
	// codingtestcasedataDescInput is the schema descriptor for input field.
	codingtestcasedataDescInput := codingtestcasedataFields[0].Descriptor()
	// codingtestcasedata.DefaultInput holds the default value on creation for the input field.
	codingtestcasedata.DefaultInput = codingtestcasedataDescInput.Default.(string)
	// codingtestcasedataDescOutput is the schema descriptor for output field.
	codingtestcasedataDescOutput := codingtestcasedataFields[1].Descriptor()
	// codingtestcasedata.DefaultOutput holds the default value on creation for the output field.
	codingtestcasedata.DefaultOutput = codingtestcasedataDescOutput.Default.(string)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
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
	Version = "v0.10.0"                                         // Version of ent codegen.
	Sum     = "h1:9cBomE1fh+WX34DPYQL7tDNAIvhKa3tXvwxuLyhYCMo=" // Sum of ent codegen.
)
