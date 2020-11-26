package markdown

import (
	"sync"

	"github.com/88250/lute"
)

var once sync.Once

var luteEngine *lute.Lute

func initLuteEngine() {
	if luteEngine == nil {
		once.Do(func() {
			luteEngine = lute.New()
		})
	}
}

//Markdown2HTML convert markdown to HTML
func Markdown2HTML(name, markdown string) string {
	initLuteEngine()
	html := luteEngine.MarkdownStr(name, markdown)
	return html
}
