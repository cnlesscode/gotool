package gstring

import (
	"io/ioutil"
	"strings"
	"unicode/utf8"

	"github.com/go-ego/gse"
)

type NLP struct {
	Content string
}

type NLPKwd struct {
	KeyWord string
	Count   int
}

func (m *NLP) SearchNLPKwd(slice []NLPKwd, str string) int {
	for idx, v := range slice {
		if v.KeyWord == str {
			return idx
		}
	}
	return -1
}

func (m *NLP) Cut() []NLPKwd {
	seg := gse.Segmenter{}
	seg.LoadDict("./resources/dict/zh/s_1.txt")
	kwds := make([]NLPKwd, 0)
	NPLNotWords, err := ioutil.ReadFile("./resources/dict/nplNotWords.txt")
	NPLNotWordsString := string(NPLNotWords) + "\t\n"
	for _, v := range seg.Slice(m.Content, true) {
		idxForNotRec := -1
		if utf8.RuneCountInString(v) > 1 {
			if err != nil {
				idxForNotRec = -1
			} else {
				idxForNotRec = strings.Index(string(NPLNotWordsString), v)
			}
			if idxForNotRec == -1 {
				idx := m.SearchNLPKwd(kwds, v)
				if idx == -1 {
					kwds = append(kwds, NLPKwd{KeyWord: v, Count: 1})
				} else {
					kwds[idx].Count++
				}
			}
		}
	}
	return kwds
}
