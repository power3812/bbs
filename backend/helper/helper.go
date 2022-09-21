package helper

import (
	"time"
)

func SetJst() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	time.Local = jst
	return nil
}
