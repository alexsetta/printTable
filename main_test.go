package printTable

import (
	"fmt"
	"testing"
)

type Print struct {
	Id        int
	Sistema   string
	Linguagem string
}

func TestNewPrintTable(t *testing.T) {
	var h = []string{"#", "Sistema", "Linguagem"}

	var p = []Print{
		{1, "VPost", "CF/GO"},
		{2, "SNS", "PHP"},
		{3, "eCarta", "Java"},
	}

	s := NewPrintTable(h, "%5s %-10s %-20s").Format(p)
	fmt.Println(s)

}
