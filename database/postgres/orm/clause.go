package orm

import (
	"fmt"

	"github.com/go-zoox/core-utils/regexp"
)

type Clause struct {
	Driver string
	Raw    string
}

func (c *Clause) Get() string {
	switch c.Driver {
	case "postgres":
		re, _ := regexp.New("(\\?|:\\w+)")
		i := 0
		return re.ReplaceFunc(c.Raw, func(s string) string {
			i++
			return fmt.Sprintf("$%d", i)
		})
	case "mysql":
		re, _ := regexp.New("(\\$\\d+|:\\w+)")
		i := 0
		return re.ReplaceFunc(c.Raw, func(s string) string {
			i++
			return fmt.Sprintf("$%d", i)
		})
	case "oracle":
		re, _ := regexp.New("(\\?|:\\w+)")
		i := 0
		return re.ReplaceFunc(c.Raw, func(s string) string {
			i++
			return fmt.Sprintf("$%d", i)
		})
	default:
		return c.Raw
	}
}

func (c *Clause) String() string {
	return c.Get()
}
