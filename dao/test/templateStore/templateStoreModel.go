package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

//TODO: RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Update the template_Store struct to match the domain entity
//TODO: Update the Fields. constants to match the domain entity

import (
	"time"

	"github.com/mt1976/frantic-core/dao"
	audit "github.com/mt1976/frantic-core/dao/audit"
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
	Active        dao.StormBool
	ExampleInt    dao.Int
	ExampleFloat  dao.Float
	ExampleBool   dao.Bool
	ExampleDate   time.Time
	ExampleString string
	LastLogin     time.Time
	LastHost      string
	// Last field is mandatory for all DAO entities
	Audit audit.Audit `csv:"-"` // audit data
}

// Fields provides a structured way to reference model field names.
type fieldNames struct {
	// First four fields are mandatory for all DAO entities
	ID    dao.Field
	Key   dao.Field
	Raw   dao.Field
	Audit dao.Field
	// Add your domain entity fields below
	UID           dao.Field
	GID           dao.Field
	RealName      dao.Field
	UserName      dao.Field
	UserCode      dao.Field
	Email         dao.Field
	Notes         dao.Field
	Active        dao.Field
	ExampleInt    dao.Field
	ExampleFloat  dao.Field
	ExampleBool   dao.Field
	ExampleDate   dao.Field
	ExampleString dao.Field
	LastLogin     dao.Field
	LastHost      dao.Field
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
