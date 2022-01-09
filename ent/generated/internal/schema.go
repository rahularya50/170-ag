// Code generated by entc, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"170-ag/ent/schema","Package":"170-ag/ent/generated","Schemas":[{"name":"User","config":{"Table":""},"fields":[{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":128,"unique":true,"immutable":true,"validators":2,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":128,"optional":true,"validators":2,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["email"]}],"policy":[{"Index":0,"MixedIn":false,"MixinIndex":0}]}],"Features":["privacy","schema/snapshot"]}`
