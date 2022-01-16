package format

import "strings"

func Key(biz string, indexes ...string) string {
	return biz + ":" + strings.Join(indexes, ":")
}
