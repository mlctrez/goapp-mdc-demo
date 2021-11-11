package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/mlctrez/goapp-mdc-demo/demo"
	"github.com/mlctrez/goapp-mdc-demo/demo/markup"
)

func main() {
	var output string
	flag.StringVar(&output, "output", "demo/markup/code.go", "the output file")

	fmt.Println(output)

	getwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths := make(map[string]bool)

	err = filepath.Walk(getwd, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".go") {
			var rel string
			rel, err = filepath.Rel(getwd, path)
			if err != nil {
				return err
			}
			// add root files
			if !strings.Contains(rel, "/") {
				paths[rel] = true
			}
			if strings.HasPrefix(rel, "demo/") &&
				!strings.HasPrefix(rel, "demo/older") &&
				!strings.HasPrefix(rel, "demo/markup") {
				paths[rel] = true
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(output), 0755)
	if err != nil {
		panic(err)
	}

	type CodeOrder struct {
		path  string
		order int
	}

	var outputOrder []CodeOrder

	demo.Routes()
	// navigation items as they appear on the page, match these
	// with the demo code path based on demo/demo_<name>.go
	// where <name> will be the href
	for i, item := range demo.NavigationItems {
		name := strings.TrimPrefix(item.Href, "/")
		if name == "" {
			name = "index"
		}
		codePath := fmt.Sprintf("demo/demo_%s.go", name)
		if paths[codePath] {
			delete(paths, codePath)
			order := CodeOrder{path: codePath, order: i}
			outputOrder = append(outputOrder, order)
		}
	}
	next := len(outputOrder)
	var otherPaths []string
	for k := range paths {
		otherPaths = append(otherPaths, k)
	}
	sort.Strings(otherPaths)
	for i, path := range otherPaths {
		order := CodeOrder{path: path, order: next + i}
		outputOrder = append(outputOrder, order)
	}

	open, err := os.Create(output)
	if err != nil {
		panic(err)
	}

	buff := bytes.Buffer{}
	buff.WriteString("package markup\n\n")

	buff.WriteString("// Generated by scripts/markup - DO NOT EDIT\n\n")

	buff.WriteString("type CodeDetails struct {\n")
	buff.WriteString("	Name string\n")
	buff.WriteString("	Code string\n")
	buff.WriteString("}\n")

	buff.WriteString("\n// Code defines all of the code samples\n")

	buff.WriteString("var Code = []CodeDetails{\n")

	for _, op := range outputOrder {
		path := op.path
		buf := &bytes.Buffer{}
		buf.WriteString("```go\n")
		file, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		buf.Write(file)
		buf.WriteString("```\n")
		p := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs)
		html := markdown.ToHTML(buf.Bytes(), p, nil)

		path = strings.Replace(path, "demo/", "", 1)
		path = strings.Replace(path, "demo_", "", 1)

		encodedCode := markup.Encode(html)

		buff.WriteString(fmt.Sprintf("    {Name:%q,Code:`%s`},\n", path, encodedCode))
	}
	buff.WriteString("}\n")

	_, err = open.Write(buff.Bytes())
	if err != nil {
		panic(err)
	}

}
