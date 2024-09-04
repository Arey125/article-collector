package article

import (
	"fmt"

	"github.com/xyproto/ollamaclient/v2"
)

var oc = (func() *ollamaclient.Config {
	client := ollamaclient.New("deepseek-coder-v2")
	client.Verbose = true

	return client
})()

func getCodeBlockLanguage(codeBlock string) string {
	const mainPrompt = `
    Please output language name for the code blocks in the markdown file. Do not include any additional comments. Only output the appropriate language identifier that should be at the begining of this block
    `
	prompt := fmt.Sprintf("%s\n\n%s", mainPrompt, codeBlock)
    language, _ := oc.GetOutput(prompt)
    fmt.Println(language)

	return language
}
