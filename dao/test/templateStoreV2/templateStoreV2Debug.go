package templateStoreV2

import (
	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/logHandler"
)

// Spew outputs the contents of the TemplateStore record to the Trace log.
func (record *TemplateStore) Spew() {
	logHandler.TraceLogger.Printf("[%v] Record=[%+v]", tableName, godump.DumpStr(record))
}
