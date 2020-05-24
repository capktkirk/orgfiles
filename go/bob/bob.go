package bob

import (
	"strings"
)

func Hey(remark string) string {
	var numeric bool
	remark = strings.ReplaceAll(remark, " ", "")
	numeric = strings.ToUpper(remark) == remark && strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	silence := strings.Fields(remark)
	if strings.HasSuffix(remark, "?") && numeric {
		return "Calm down, I know what I'm doing!"
	}
	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}
	if remark == "" || len(silence) == 0 {
		return "Fine. Be that way!"
	}
	if strings.ToUpper(remark) == remark && numeric {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
