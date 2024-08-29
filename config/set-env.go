package config

func SetEnv(key, val string) {

	_Config.lenv.Lock()
	defer _Config.lenv.Unlock()

	_Config.envs[key] = val

}

func SetEnvs(valx map[string]string) {

	_Config.lenv.Lock()
	defer _Config.lenv.Unlock()

	for k, v := range valx {
		_Config.envs[k] = v
	}
}
