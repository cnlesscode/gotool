package nlp

import (
	"os"
	"strings"
	"unicode/utf8"

	"github.com/go-ego/gse"
	"gorm.io/gorm"
)

var Seg gse.Segmenter
var NPLNotWordsString string = ""

func init() {
	Seg.LoadDict("./resources/dict/zh/s_1.txt")
	NPLNotWords, err := os.ReadFile("./resources/dict/nplNotWords.txt")
	if err == nil {
		NPLNotWordsString = string(NPLNotWords) + "\t\n"
	}
}

type NLPKwd struct {
	KeyWord string
	Count   int
}

func SearchKwd(slice []NLPKwd, str string) int {
	for idx, v := range slice {
		if v.KeyWord == str {
			return idx
		}
	}
	return -1
}

// 分词
func Cut(content string) []NLPKwd {
	kwds := make([]NLPKwd, 0)
	for _, v := range Seg.Slice(content, true) {
		idxForNotRec := -1
		if utf8.RuneCountInString(v) > 1 {
			if NPLNotWordsString != "nil" {
				idxForNotRec = strings.Index(string(NPLNotWordsString), v)
			}
			if idxForNotRec == -1 {
				idx := SearchKwd(kwds, v)
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

type KeywordsForDB struct {
	Id    int    `gorm:"column:id;primaryKey"`
	Word  string `gorm:"column:word"`
	Mid   int    `gorm:"column:mid"`
	Score int    `gorm:"column:score"`
}

// 分词并保存到数据库
func CutAndSave(content string, db *gorm.DB, tableName string, mid int) {
	kwds := Cut(content)
	for _, kwd := range kwds {
		InserData := &KeywordsForDB{
			Word:  kwd.KeyWord,
			Mid:   mid,
			Score: kwd.Count,
		}
		db.Table(tableName).Create(InserData)
	}
}
