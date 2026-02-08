package commonConfig

func (s *Settings) GetHistory_MaxHistoryEntries() int {
	return s.History.MaxEntries
}

// GetDisplayDelimiter returns the delimiter used to separate the display elements
// Deprecated: Use GetDefault_Delimiter instead
func (s *Settings) GetDisplayDelimiter() string {
	return s.GetDefault_Delimiter()
}

// GetDefault_Delimiter returns the delimiter used to separate the display elements
func (s *Settings) GetDefault_Delimiter() string {
	if s.Display.Delimiter == "" {
		return "⋮"
	}
	return s.Display.Delimiter
}

// SEP returns the delimiter used to separate the display elements
// Deprecated: Use Delimiter instead
func (s *Settings) SEP() string {
	return s.GetDefault_Delimiter()
}

// Delimiter returns the delimiter used to separate the display elements
// Deprecated: Use GetDefault_Delimiter instead
func (s *Settings) Delimiter() string {
	return s.GetDefault_Delimiter()
}

func (s *Settings) GetWorkerPoolSize() int {
	if s.WorkerPool.Size <= 0 {
		return 10
	}
	return s.WorkerPool.Size
}
