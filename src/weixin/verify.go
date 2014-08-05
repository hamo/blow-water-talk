package weixin

import (
	"sort"
	"strings"
)

const (
	ApiVerifyToken = "VerifyToken"
)

func Verify(signature, timestamp, nonce string) bool {
	tmpArr := [3]string{ApiVerifyToken, timestamp, nonce}
	sort.Strings(tmpArr[:])

	tmpStr := strings.Join(tmpArr[:], "")

	tmpSha1 := getSha1(tmpStr)

	return (tmpSha1 == signature)
}
