package commonConfig

import "strconv"

func (s *Settings) GetDatabase_Version() string {
	return importedVersionInfo
}

func (s *Settings) GetDatabase_Type() string {
	return s.Database.Type
}

func (s *Settings) GetDatabase_Name() string {
	return s.Database.Name
}

func (s *Settings) GetDatabase_Path() string {
	return s.Database.Path
}

func (s *Settings) GetDatabase_Host() string {
	return s.Database.Host
}

func (s *Settings) GetDatabase_Port() int {
	return s.Database.Port
}

func (s *Settings) GetDatabase_User() string {
	return s.Database.User
}

func (s *Settings) GetDatabase_Password() string {
	return s.Database.Pass
}

func (s *Settings) GetDatabase_PortString() string {
	return strconv.Itoa(s.Database.Port)
}

func (s *Settings) GetDatabase_PoolSize() int {
	if s.Database.PoolSize == 0 {
		return 10 // Default to 10 connections
	}
	return s.Database.PoolSize
}

func (s *Settings) GetDatabase_Timeout() int {
	if s.Database.Timeout == 0 {
		return 30 // Default to 30 seconds
	}
	return s.Database.Timeout
}
