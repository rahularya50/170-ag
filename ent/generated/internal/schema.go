// Code generated by entc, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"170-ag/ent/schema","Package":"170-ag/ent/generated","Schemas":[{"name":"CodingDraft","config":{"Table":""},"edges":[{"name":"author","type":"User","unique":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"coding_problem","type":"CodingProblem","unique":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"code","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"edges":["author","coding_problem"]}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"CodingProblem","config":{"Table":""},"edges":[{"name":"drafts","type":"CodingDraft","ref_name":"coding_problem","inverse":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"test_cases","type":"CodingTestCase","annotations":{"EntGQL":{"Bind":true},"EntSQL":{"on_delete":"CASCADE"}}},{"name":"submissions","type":"CodingSubmission","ref_name":"coding_problem","inverse":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":128,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"statement","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":2147483647,"default":true,"default_value":"This is the problem statement","default_kind":24,"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"released","type":{"Type":1,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"CodingSubmission","config":{"Table":""},"edges":[{"name":"author","type":"User","unique":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"coding_problem","type":"CodingProblem","unique":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"staff_data","type":"CodingSubmissionStaffData","ref_name":"coding_submission","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"code","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"status","type":{"Type":6,"Ident":"codingsubmission.Status","PkgPath":"","Nillable":false,"RType":null},"enums":[{"N":"Queued","V":"QUEUED"},{"N":"Running","V":"RUNNING"},{"N":"Completed","V":"COMPLETED"}],"default":true,"default_value":"QUEUED","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"edges":["author","coding_problem"]},{"edges":["coding_problem"]},{"fields":["status"]}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"CodingSubmissionStaffData","config":{"Table":""},"edges":[{"name":"coding_submission","type":"CodingSubmission","unique":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"execution_id","type":{"Type":13,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"input","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"output","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":65535,"nillable":true,"optional":true,"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"stderr","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":65535,"nillable":true,"optional":true,"validators":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"exit_error","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":65535,"nillable":true,"optional":true,"validators":1,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["execution_id"]}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"CodingTestCase","config":{"Table":""},"edges":[{"name":"coding_problem","type":"CodingProblem","ref_name":"test_cases","unique":true,"inverse":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}},{"name":"data","type":"CodingTestCaseData","ref_name":"test_case","unique":true,"inverse":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"points","type":{"Type":12,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"public","type":{"Type":1,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"CodingTestCaseData","config":{"Table":""},"edges":[{"name":"test_case","type":"CodingTestCase","unique":true,"required":true,"annotations":{"EntGQL":{"Bind":true}}}],"fields":[{"name":"input","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"output","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"User","config":{"Table":""},"edges":[{"name":"drafts","type":"CodingDraft","ref_name":"author","inverse":true}],"fields":[{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":128,"unique":true,"immutable":true,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":128,"optional":true,"validators":2,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"is_staff","type":{"Type":1,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["email"]}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]}],"Features":["privacy","schema/snapshot","sql/upsert","entql"]}`
