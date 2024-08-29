package tox

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func UuidString(s string) (*uuid.UUID, error) {
	if strings.TrimSpace(s) == "" {
		return &uuid.Nil, nil
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return nil, fmt.Errorf(`uuid.Parse : %v`, err.Error())
	}
	return &id, nil
}
