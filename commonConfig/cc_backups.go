package commonConfig

func (s *Settings) GetBackup_RetainForDays() int {
	return s.Backups.RetainDays
}
