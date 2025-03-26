package commonConfig

// / SECURITY SESSION
func (s *Settings) GetSecuritySession_ExpiryPeriod() int {
	return s.Security.Sessions.ExpiryPeriod
}

// / SECURITY SESSION KEYS

func (s *Settings) GetSecuritySessionKey_Session() string {
	return s.Security.Sessions.Keys.Session
}

func (s *Settings) GetSecuritySessionKey_UserKey() string {
	return s.Security.Sessions.Keys.UserKey
}

func (s *Settings) GetSecuritySessionKey_UserCode() string {
	return s.Security.Sessions.Keys.UserCode
}

func (s *Settings) GetSecuritySessionKey_Token() string {
	return s.Security.Sessions.Keys.Token
}

func (s *Settings) GetSecuritySessionKey_ExpiryPeriod() string {
	return s.Security.Sessions.Keys.ExpiryPeriod
}

// / SECURITY SESSION SERVICE USER
func (s *Settings) GetServiceUser_Name() string {
	return s.Security.Service.UserName
}

func (s *Settings) GetServiceUser_UID() string {
	return s.Security.Service.UserUID
}

func (s *Settings) GetServiceUser_UserCode() string {
	return s.GetServiceUser_UID() + "_" + s.GetServiceUser_Name()
}

// locale = "xs-locale"
// theme = "xs-theme"
// timezone = "xs-timezone"

func (s *Settings) GetSecuritySessionKey_Locale() string {
	return s.Security.Sessions.Keys.Locale
}

func (s *Settings) GetSecuritySessionKey_Theme() string {
	return s.Security.Sessions.Keys.Theme
}

func (s *Settings) GetSecuritySessionKey_Timezone() string {
	return s.Security.Sessions.Keys.Timezone
}

func (s *Settings) GetSecuritySessionKey_UserRole() string {
	return s.Security.Sessions.Keys.Role
}
