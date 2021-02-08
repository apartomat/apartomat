package gravatar

import (
	"crypto/md5"
	"fmt"
)

func Url(email string) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x", md5.Sum([]byte(email)))
}
