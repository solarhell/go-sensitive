package sensitive_test

import (
	"fmt"
	"github.com/solarhell/go-sensitive"
	"testing"
)

func TestSensitive(t *testing.T) {
	filter := sensitive.New()
	filter.AddWord("ChatGPT")
	filter.AddWord("GPT")
	content := `gPt3.5对\比/chat|Gpt4`

	fmt.Println(filter.Validate(content))
	fmt.Println(filter.FindAll(content))
	fmt.Println(filter.RemoveNoise(content))
}
