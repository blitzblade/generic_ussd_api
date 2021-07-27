package utils
import (
	"strings"
)

type HTMLString string

func (htmlString *HTMLString) MakeHTMLReady() {
	*htmlString = HTMLString(strings.Replace(string(*htmlString), "\n", "<br/>", -1))
}