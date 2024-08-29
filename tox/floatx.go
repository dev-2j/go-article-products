package tox

func Float2(s *float64) float64 {
	if s == nil {
		return 0
	}
	return *s
}
