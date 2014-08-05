package weixin

import (
	"crypto/sha1"
	"io"
	
"fmt"
)

func getSha1(data string) string {
	m := sha1.New()
	io.WriteString(m, data)
	return fmt.Sprintf("%x", m.Sum(nil));
}
