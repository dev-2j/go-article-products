package config

func (s *ConfigT) SetServiceName(serviceName string) {

	s.lsvn.Lock()
	defer s.lsvn.Unlock()

	s.svn = serviceName
}

func (s *ConfigT) GetServiceName() string {

	s.lsvn.Lock()
	defer s.lsvn.Unlock()

	if s.svn == `` {
		s.svn = GetEnv(`SERVICE_NAME`)
	}
	return s.svn
}
