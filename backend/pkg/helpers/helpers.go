package helpers

import "github.com/segmentio/ksuid"

func GetKsuid() string {
	return ksuid.New().String()
}
