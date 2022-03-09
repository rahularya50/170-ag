// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CodingDraftsColumns holds the columns for the "coding_drafts" table.
	CodingDraftsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Size: 2147483647},
		{Name: "coding_draft_author", Type: field.TypeInt, Nullable: true},
		{Name: "coding_draft_coding_problem", Type: field.TypeInt, Nullable: true},
	}
	// CodingDraftsTable holds the schema information for the "coding_drafts" table.
	CodingDraftsTable = &schema.Table{
		Name:       "coding_drafts",
		Columns:    CodingDraftsColumns,
		PrimaryKey: []*schema.Column{CodingDraftsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coding_drafts_users_author",
				Columns:    []*schema.Column{CodingDraftsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "coding_drafts_coding_problems_coding_problem",
				Columns:    []*schema.Column{CodingDraftsColumns[5]},
				RefColumns: []*schema.Column{CodingProblemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "codingdraft_coding_draft_author_coding_draft_coding_problem",
				Unique:  true,
				Columns: []*schema.Column{CodingDraftsColumns[4], CodingDraftsColumns[5]},
			},
		},
	}
	// CodingExtensionsColumns holds the columns for the "coding_extensions" table.
	CodingExtensionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "deadline", Type: field.TypeTime},
		{Name: "coding_extension_student", Type: field.TypeInt, Nullable: true},
		{Name: "coding_extension_coding_problem", Type: field.TypeInt, Nullable: true},
	}
	// CodingExtensionsTable holds the schema information for the "coding_extensions" table.
	CodingExtensionsTable = &schema.Table{
		Name:       "coding_extensions",
		Columns:    CodingExtensionsColumns,
		PrimaryKey: []*schema.Column{CodingExtensionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coding_extensions_users_student",
				Columns:    []*schema.Column{CodingExtensionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "coding_extensions_coding_problems_coding_problem",
				Columns:    []*schema.Column{CodingExtensionsColumns[5]},
				RefColumns: []*schema.Column{CodingProblemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "codingextension_coding_extension_student_coding_extension_coding_problem",
				Unique:  true,
				Columns: []*schema.Column{CodingExtensionsColumns[4], CodingExtensionsColumns[5]},
			},
			{
				Name:    "codingextension_coding_extension_student",
				Unique:  false,
				Columns: []*schema.Column{CodingExtensionsColumns[4]},
			},
			{
				Name:    "codingextension_coding_extension_coding_problem",
				Unique:  false,
				Columns: []*schema.Column{CodingExtensionsColumns[5]},
			},
		},
	}
	// CodingProblemsColumns holds the columns for the "coding_problems" table.
	CodingProblemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 128, Default: "Untitled Problem"},
		{Name: "statement", Type: field.TypeString, Size: 2147483647, Default: "This is the problem statement"},
		{Name: "skeleton", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "released", Type: field.TypeBool, Default: false},
		{Name: "deadline", Type: field.TypeTime},
	}
	// CodingProblemsTable holds the schema information for the "coding_problems" table.
	CodingProblemsTable = &schema.Table{
		Name:       "coding_problems",
		Columns:    CodingProblemsColumns,
		PrimaryKey: []*schema.Column{CodingProblemsColumns[0]},
	}
	// CodingSubmissionsColumns holds the columns for the "coding_submissions" table.
	CodingSubmissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Size: 65535},
		{Name: "is_validation", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"QUEUED", "RUNNING", "COMPLETED", "INTERNAL_ERROR"}, Default: "QUEUED"},
		{Name: "points", Type: field.TypeInt, Nullable: true},
		{Name: "results", Type: field.TypeJSON, Nullable: true},
		{Name: "coding_submission_author", Type: field.TypeInt, Nullable: true},
		{Name: "coding_submission_coding_problem", Type: field.TypeInt, Nullable: true},
		{Name: "coding_submission_staff_data_coding_submission", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CodingSubmissionsTable holds the schema information for the "coding_submissions" table.
	CodingSubmissionsTable = &schema.Table{
		Name:       "coding_submissions",
		Columns:    CodingSubmissionsColumns,
		PrimaryKey: []*schema.Column{CodingSubmissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coding_submissions_users_author",
				Columns:    []*schema.Column{CodingSubmissionsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "coding_submissions_coding_problems_coding_problem",
				Columns:    []*schema.Column{CodingSubmissionsColumns[9]},
				RefColumns: []*schema.Column{CodingProblemsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "coding_submissions_coding_submission_staff_data_coding_submission",
				Columns:    []*schema.Column{CodingSubmissionsColumns[10]},
				RefColumns: []*schema.Column{CodingSubmissionStaffDataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "codingsubmission_coding_submission_author_coding_submission_coding_problem",
				Unique:  false,
				Columns: []*schema.Column{CodingSubmissionsColumns[8], CodingSubmissionsColumns[9]},
			},
			{
				Name:    "codingsubmission_coding_submission_coding_problem",
				Unique:  false,
				Columns: []*schema.Column{CodingSubmissionsColumns[9]},
			},
			{
				Name:    "codingsubmission_status",
				Unique:  false,
				Columns: []*schema.Column{CodingSubmissionsColumns[5]},
			},
		},
	}
	// CodingSubmissionStaffDataColumns holds the columns for the "coding_submission_staff_data" table.
	CodingSubmissionStaffDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "execution_id", Type: field.TypeInt64, Nullable: true},
		{Name: "input", Type: field.TypeString, Size: 2147483647},
		{Name: "output", Type: field.TypeString, Nullable: true, Size: 8388607},
		{Name: "stderr", Type: field.TypeString, Nullable: true, Size: 262143},
		{Name: "exit_error", Type: field.TypeString, Nullable: true, Size: 262143},
	}
	// CodingSubmissionStaffDataTable holds the schema information for the "coding_submission_staff_data" table.
	CodingSubmissionStaffDataTable = &schema.Table{
		Name:       "coding_submission_staff_data",
		Columns:    CodingSubmissionStaffDataColumns,
		PrimaryKey: []*schema.Column{CodingSubmissionStaffDataColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "codingsubmissionstaffdata_execution_id",
				Unique:  false,
				Columns: []*schema.Column{CodingSubmissionStaffDataColumns[3]},
			},
		},
	}
	// CodingTestCasesColumns holds the columns for the "coding_test_cases" table.
	CodingTestCasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "points", Type: field.TypeInt, Default: 0},
		{Name: "visibility", Type: field.TypeEnum, Enums: []string{"PRIVATE", "COLLAPSED", "EXPANDED"}, Default: "PRIVATE"},
		{Name: "coding_problem_test_cases", Type: field.TypeInt, Nullable: true},
		{Name: "coding_test_case_data_test_case", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CodingTestCasesTable holds the schema information for the "coding_test_cases" table.
	CodingTestCasesTable = &schema.Table{
		Name:       "coding_test_cases",
		Columns:    CodingTestCasesColumns,
		PrimaryKey: []*schema.Column{CodingTestCasesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "coding_test_cases_coding_problems_test_cases",
				Columns:    []*schema.Column{CodingTestCasesColumns[5]},
				RefColumns: []*schema.Column{CodingProblemsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "coding_test_cases_coding_test_case_data_test_case",
				Columns:    []*schema.Column{CodingTestCasesColumns[6]},
				RefColumns: []*schema.Column{CodingTestCaseDataColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CodingTestCaseDataColumns holds the columns for the "coding_test_case_data" table.
	CodingTestCaseDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "input", Type: field.TypeString, Size: 2147483647, Default: "0\n"},
		{Name: "output", Type: field.TypeString, Size: 2147483647, Default: ""},
	}
	// CodingTestCaseDataTable holds the schema information for the "coding_test_case_data" table.
	CodingTestCaseDataTable = &schema.Table{
		Name:       "coding_test_case_data",
		Columns:    CodingTestCaseDataColumns,
		PrimaryKey: []*schema.Column{CodingTestCaseDataColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "is_staff", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CodingDraftsTable,
		CodingExtensionsTable,
		CodingProblemsTable,
		CodingSubmissionsTable,
		CodingSubmissionStaffDataTable,
		CodingTestCasesTable,
		CodingTestCaseDataTable,
		UsersTable,
	}
)

func init() {
	CodingDraftsTable.ForeignKeys[0].RefTable = UsersTable
	CodingDraftsTable.ForeignKeys[1].RefTable = CodingProblemsTable
	CodingExtensionsTable.ForeignKeys[0].RefTable = UsersTable
	CodingExtensionsTable.ForeignKeys[1].RefTable = CodingProblemsTable
	CodingSubmissionsTable.ForeignKeys[0].RefTable = UsersTable
	CodingSubmissionsTable.ForeignKeys[1].RefTable = CodingProblemsTable
	CodingSubmissionsTable.ForeignKeys[2].RefTable = CodingSubmissionStaffDataTable
	CodingTestCasesTable.ForeignKeys[0].RefTable = CodingProblemsTable
	CodingTestCasesTable.ForeignKeys[1].RefTable = CodingTestCaseDataTable
}
