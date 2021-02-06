package src

import (
	"fmt"

	"github.com/monaco-io/request"
)

// Translate ...
func Translate(word string) {
	var result Response
	resp := request.
		New().
		POST(BaiduTranslateURL).
		AddURLEncodedForm(Encode(word)).
		Send().
		Scan(&result)
	if resp.OK() && result.ErrorCode == 0 {
		if len(result.TransResult) > 0 {
			fmt.Println(result.TransResult[0].Dst)
		}
		return
	}
	fmt.Println(resp.Error(), result.ErrorCode)
}
