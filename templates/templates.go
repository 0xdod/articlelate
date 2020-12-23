package templates

import (
	"fmt"
	"html"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"truncate":     Truncate,
		"pluralize":    Pluralize,
		"decr":         Decrement,
		"linebreaks":   LineBreaks,
		"mark":         mark,
		"contains":     strings.Contains,
		"icontains":    icontains,
		"isProduction": isProduction,
	}
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	partials, err := filepath.Glob(templatesDir + "/partials/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(err.Error())
	}
	funcs := GetFuncMap()
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		files = append(files, partials...)
		r.AddFromFilesFuncs(filepath.Base(include), funcs, files...)
	}
	return r
}

func Truncate(limit int, content string) string {
	if len(content) <= limit {
		return content
	}
	oldContentSlice := strings.Fields(content)
	var newContentSlice []string
	var newContent string
	for _, v := range oldContentSlice {
		if len(newContent) >= limit {
			break
		}
		newContent += v + " "
		newContentSlice = append(newContentSlice, v)
	}
	return strings.Join(newContentSlice, " ") + "..."
}

func Pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

func Decrement(value int) int {
	return value - 1
}

func LineBreaks(content string) template.HTML {
	content = html.EscapeString(content)
	content = strings.ReplaceAll(content, "\r", "")
	content = strings.ReplaceAll(content, "\n", "<br>")
	cs := strings.Split(content, "<br><br>")
	newContent := ""
	for _, c := range cs {
		newContent += "<p>" + c + "</p>"
	}
	return template.HTML(newContent)
}

func mark(content, search string) template.HTML {
	content = html.EscapeString(content)
	if search != "" {
		exp := fmt.Sprintf(`(?mi)(?P<key>%s)`, search)
		re := regexp.MustCompile(exp)
		temp := "<mark>$key</mark>"
		content = re.ReplaceAllString(content, temp)
	}
	return template.HTML(content)
}

func icontains(str, substr string) bool {
	str = strings.ToLower(str)
	substr = strings.ToLower(substr)
	return strings.Contains(str, substr)
}

func isProduction() bool {
	return os.Getenv("GIN_MODE") == "release"
}
