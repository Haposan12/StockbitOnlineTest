package uuid

import uuid "github.com/satori/go.uuid"

func GenerateId() string {
	return uuid.NewV4().String()
}
