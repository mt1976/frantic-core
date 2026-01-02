package lookup

import "github.com/mt1976/frantic-core/logHandler"

// Lookup represents a collection of lookup data entries
type Lookup struct {
	Data []LookupData
}

// LookupData represents a single entry in the lookup data
type LookupData struct {
	Key          string
	Value        string
	AltID        string
	Description  string
	ObjectDomain string
	Selected     bool
}

// Spew outputs the contents of the Lookup structure for debugging purposes
func (a *Lookup) Spew() error {
	// Spew the Audit Data
	noAudit := len(a.Data)
	//logger.InfoLogger.Printf(" No Updates=[%v]", noAudit)
	if noAudit > 0 {
		for i := 0; i < noAudit; i++ {
			upd := a.Data[i]
			logHandler.TraceLogger.Printf("[LKP] Lookup Data [%v] Key=[%v] Value=[%v]", i, upd.Key, upd.Value)
		}
	}
	return nil
}
