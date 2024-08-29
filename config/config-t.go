package config

import "sync"

type ConfigT struct {
	svn  string
	lsvn *sync.RWMutex

	envs map[string]string
	lenv *sync.RWMutex
}

var _Config = &ConfigT{
	lsvn: &sync.RWMutex{},
	lenv: &sync.RWMutex{},

	envs: map[string]string{},
}
