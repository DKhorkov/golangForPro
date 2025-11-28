// Package handlers for the RESTful Server
//
// Documentation for REST API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"time"
)

// @termsOfService http://swagger.io/terms/

// Entry defines the structure for a Full Entry Record
//
// swagger:model
type Entry struct {
	// The Name of the user
	//
	// required: true
	Name string `json:"name"`

	// The Surname of the User
	//
	// required: true
	Surname string `json:"surname"`

	// The Phone of the User
	//
	// required: true
	Phone string `json:"phone"`

	// The Last time of the appeal
	//
	// required: true
	LastAccess time.Time `json:"lastAccess"`
}

// OK message returned as an HTTP Status Code
// swagger:response OK
type OK struct {
	// Description of the situation
	// in: body
	Body int
}

// BadRequest message returned as an HTTP Status Code
// swagger:response BadRequest
type BadRequest struct {
	// Description of the situation
	// in: body
	Body int
}

// NotFound message returned as an HTTP Status Code
// swagger:response NotFound
type NotFound struct {
	// Description of the situation
	// in: body
	Body int
}

// InternalServerError message returned as an HTTP Status Code
// swagger:response InternalServerError
type InternalServerError struct {
	// Description of the situation
	// in: body
	Body int
}

// Entries is a list of all entries
// swagger:response Entries
type Entries struct {
	// A list of entries
	// in: body
	Body []Entry
}

// swagger:parameters deleteKey
type deleteKey struct {
	// The key of Entry to be deleted
	// in: path
	// required: true
	Key string `json:"key"`
}

// swagger:parameters searchKey
type searchKey struct {
	// The key of Entry to be searched
	// in: path
	// required: true
	Key string `json:"key"`
}

// swagger:parameters reverseKey
type reverseKey struct {
	// Whether to reverse all entries or not
	// in: query
	// required: false
	Reverse bool `json:"reverse"`
}

// InsertEntry
// swagger:parameters insertEntryInput
type InsertEntry struct {
	// New Entry to insert
	// in: body
	Body Entry
}
