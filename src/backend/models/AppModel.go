package models

import (
	"strings"
)

type AppModel struct{}

func splitQExt(qExt string) (splitQExt []string) {
	if qExt != "" {
		splitQExt = strings.Split(qExt, "&")
		if len(splitQExt) > 0 {
			return splitQExt
		}
	}

	return nil
}
