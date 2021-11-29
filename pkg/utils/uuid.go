package utils

import "github.com/google/uuid"

func UUID() string {
	v4 := uuid.New()
	return v4.String()
}
