package q2

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 { // se strs = []
		return "" // output = ""
	}
	// strs = ["flower","flow","flight"]
	prefix := strs[0] // prefix = "flower"

	for i := 1; i < len(strs) && len(prefix) > 0; i++ { // loop percorrendo ["flow", "flight"]
		for !strings.HasPrefix(strs[i], prefix) {
			// enquanto a palavra atual nao tiver o prefixo, exemplo: flow nao con
			// tem flower, nem flowe, mas contem flow. Depois : flight nao contem
			// flow, nem flo, mas contem fl, que eh o maior prefixo. Ele vai diminuindo na linha debaixo
			prefix = prefix[:len(prefix)-1] // flower -> flowe -> flow, depois compara de novo: flow -> flo -> fl.
		}
	}

	return prefix // output = "fl"
}
