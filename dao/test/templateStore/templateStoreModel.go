package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the template_Store struct to match the domain entity
//TODO: Update the Fields. constants to match the domain entity

import (
	"time"

	audit "github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/fields"
)

var Domain = "Template"
var TableName = Domain + "Store"

// TemplateStore represents a User entity.
type TemplateStore struct {
	// First three fields are mandatory for all DAO entities
	ID  int    `storm:"id,increment=100"` // primary key with auto increment
	Key string `storm:"unique"`           // key
	Raw string `storm:"unique"`           // raw ID before encoding
	// Add your domain entity fields below
	UID           string `validate:"required"`
	GID           string `storm:"index" validate:"required"`
	RealName      string `validate:"required,min=5"` // this field will not be indexed
	UserName      string `validate:"required,min=5"`
	UserCode      string `storm:"index" validate:"required,min=5"`
	Email         string
	Notes         string `validate:"max=75"`
	Active        fields.Bool
	ExampleInt    fields.Int
	ExampleFloat  fields.Float
	ExampleBool   fields.Bool
	ExampleDate   time.Time
	ExampleString string
	LastLogin     time.Time
	LastHost      string `storm:"index"`
	// Last field is mandatory for all DAO entities
	Audit audit.Audit `csv:"-"` // audit data
}

// Fields provides a structured way to reference model field names.
type fieldNames struct {
	// First four fields are mandatory for all DAO entities
	ID    fields.Field
	Key   fields.Field
	Raw   fields.Field
	Audit fields.Field
	// Add your domain entity fields below
	UID           fields.Field
	GID           fields.Field
	RealName      fields.Field
	UserName      fields.Field
	UserCode      fields.Field
	Email         fields.Field
	Notes         fields.Field
	Active        fields.Field
	ExampleInt    fields.Field
	ExampleFloat  fields.Field
	ExampleBool   fields.Field
	ExampleDate   fields.Field
	ExampleString fields.Field
	LastLogin     fields.Field
	LastHost      fields.Field
}

// Fields provides a structured way to reference model field names.
var Fields = fieldNames{
	// First four fields are mandatory for all DAO entities
	ID:    "ID",
	Key:   "Key",
	Raw:   "Raw",
	Audit: "Audit",
	// Add your domain entity fields below
	UID:           "UID",
	GID:           "GID",
	RealName:      "RealName",
	UserName:      "UserName",
	UserCode:      "UserCode",
	Email:         "Email",
	Notes:         "Notes",
	Active:        "Active",
	ExampleInt:    "ExampleInt",
	ExampleFloat:  "ExampleFloat",
	ExampleBool:   "ExampleBool",
	ExampleDate:   "ExampleDate",
	ExampleString: "ExampleString",
	LastLogin:     "LastLogin",
	LastHost:      "LastHost",
}
