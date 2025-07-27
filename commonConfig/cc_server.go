package commonConfig

import "strconv"

func (s *Settings) GetServer_Port() int {
	if s.Server.Port == 0 {
		return 80
	}
	return s.Server.Port
}

func (s *Settings) GetServer_PortString() string {
	a := s.Server.Port
	return strconv.Itoa(a)
}

func (s *Settings) GetServer_Protocol() string {
	if s.Server.Protocol == "" {
		return "http"
	}
	return s.Server.Protocol
}

func (s *Settings) GetServer_Host() string {
	if s.Server.Host == "" {
		return "localhost"
	}
	return s.Server.Host
}

func (s *Settings) GetServer_Environment() string {
	if s.Server.Environment == "" {
		return "production"
	}
	return s.Server.Environment
}

func (s *Settings) GetServer_Compression() bool {
	return s.Server.Compression
}
func (s *Settings) GetServer_CompressionString() string {
	if s.Server.Compression {
		return "enabled"
	}
	return "disabled"
}
