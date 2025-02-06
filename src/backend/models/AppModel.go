package models

import "strings"

type AppModel struct{}

func splitQExt(qExt string) (splitQExt []string) {
	if strings.Contains(qExt, "=") {
		splitQExt = strings.Split(qExt, "=")
		if len(splitQExt) > 0 {
			return splitQExt
		}
	}

	return nil
}
