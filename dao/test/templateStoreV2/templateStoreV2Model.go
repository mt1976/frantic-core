package templateStoreV2

import (
	"time"

	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/entities"
)

var TableName = entities.Table("TemplateStore")
var tableName = TableName.String()

// TemplateStore represents a sample entity for demonstrating reduced DAO boilerplate.
// Replace this struct and Fields as needed for your real tableName entity.
type TemplateStore struct {
	// The primary key field(s), managed by the framework, DO NOT MODIFY
	ID  int    `storm:"id,increment=100"`
	Key string `storm:"index,unique"`
	Raw string `storm:"index,unique"`
	// Audit information, managed by the framework, DO NOT MODIFY
	Audit audit.Audit `csv:"-"`

	// Domain specific fields
	UID           string `validate:"required"`
	GID           string `storm:"index" validate:"required"`
	RealName      string `validate:"required,min=5"`
	UserName      string `validate:"required,min=5"`
	UserCode      string `storm:"index" validate:"required,min=5"`
	Email         string
	Notes         string `validate:"max=75"`
	Active        entities.Bool
	ExampleInt    entities.Int
	ExampleFloat  entities.Float
	ExampleBool   entities.Bool
	ExampleDate   time.Time
	ExampleString string
	LastLogin     time.Time
	LastHost      string `storm:"index"`
}

type fieldNames struct {
	// The primary key field(s), managed by the framework, DO NOT MODIFY
	ID  entities.Field
	Key entities.Field
	Raw entities.Field
	// The audit information, managed by the framework, DO NOT MODIFY
	Audit entities.Field
	// Domain specific fields, please modify as required
	UID           entities.Field
	GID           entities.Field
	RealName      entities.Field
	UserName      entities.Field
	UserCode      entities.Field
	Email         entities.Field
	Notes         entities.Field
	Active        entities.Field
	ExampleInt    entities.Field
	ExampleFloat  entities.Field
	ExampleBool   entities.Field
	ExampleDate   entities.Field
	ExampleString entities.Field
	LastLogin     entities.Field
	LastHost      entities.Field
}

var Fields = fieldNames{
	// The primary key field(s), managed by the framework, DO NOT MODIFY
	ID:  "ID",
	Key: "Key",
	Raw: "Raw",
	// The audit information, managed by the framework, DO NOT MODIFY
	Audit: "Audit",
	// tableName-specific fields, please modify as required
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
