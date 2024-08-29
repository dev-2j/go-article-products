package config

import "os"

func GetEnv(key string) string {
	return _Config.getEnv(key)
}

func GetEnvs() map[string]string {
	_Config.lenv.RLock()
	defer _Config.lenv.RUnlock()
	return _Config.envs
}

func (s *ConfigT) getEnv(key string) string {

	// get from env first
	v, ok := func() (string, bool) {
		s.lenv.RLock()
		defer s.lenv.RUnlock()
		v, ok := s.envs[key]
		return v, ok
	}()
	if ok {
		return v
	}

	// get from os and set to env
	if v = os.Getenv(key); v != `` {
		SetEnv(key, v)
	}
	return v
}
