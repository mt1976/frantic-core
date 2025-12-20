package commonConfig

import "strconv"

func (s *Settings) IsGeneralLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.General)

}

func (s *Settings) IsTimingLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Timing)
}

func (s *Settings) IsServiceLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Service)
}

func (s *Settings) IsAuditLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Audit)
}

func (s *Settings) IsTranslationLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Translation)
}

func (s *Settings) IsTraceLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Trace)
}

func (s *Settings) IsWarningLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Warning)
}

func (s *Settings) IsEventLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Event)
}

func (s *Settings) IsSecurityLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Security)
}

func (s *Settings) IsDatabaseLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Database)
}

func (s *Settings) IsApiLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Api)
}

func (s *Settings) IsImportLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Import)
}

func (s *Settings) IsExportLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Export)
}

func (s *Settings) IsCommunicationsLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Communications)
}

func (s *Settings) IsLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.All)
}

func (s *Settings) IsLockLoggingDisabled() bool {
	return isTrueFalse(s.Logging.Disable.Lock)
}

func (s *Settings) GetLogging_MaxSize() int {
	a, _ := strconv.Atoi(s.Logging.Defaults.MaxSize)
	if a == 0 {
		a = 10
	}
	return a
}

func (s *Settings) GetLogging_MaxBackups() int {
	a, _ := strconv.Atoi(s.Logging.Defaults.MaxBackups)
	return a
}

func (s *Settings) GetLogging_MaxAge() int {
	a, _ := strconv.Atoi(s.Logging.Defaults.MaxAge)
	if a == 0 {
		a = 30
	}
	return a
}

func (s *Settings) IsLogCompressionEnabled() bool {
	return isTrueFalse(s.Logging.Defaults.Compress)
}
