package main

import (
	"os"
	"strings"
)

func main() {
	var args []string = os.Args[1:]
	if len(args) != 1 {
		println("Usage: generate_ast <output directory>")
		os.Exit(64)
	}
	outputDir := args[0]

	types := []string{
		"Binary   : left Expr, operator Token, right Expr",
		"Grouping : expression Expr",
		"Literal  : value interface{}",
		"Unary    : operator Token, right Expr",
	}

	DefineAst(outputDir, "Expr", types)
}

func DefineAst(outputDir string, baseName string, types []string) {
	filename := strings.ToLower(baseName)
	path := outputDir + "/" + filename + ".go"
	f, err := os.Create(path)
	if err != nil {
		print("Encountered error: " + err.Error())
		os.Exit(1)
	}
	// remember to close the file
	defer f.Close()

	f.WriteString("package main\n")
	f.WriteString("\n")
	f.WriteString("type " + baseName + " interface {}\n\n")

	for _, t := range types {
		class_and_fields := strings.Split(t, ":")
		class := strings.TrimSpace(class_and_fields[0])
		fields := strings.TrimSpace(class_and_fields[1])
		DefineType(f, baseName, class, fields)
	}
}

func DefineType(f *os.File, baseName string, class string, fields string) {
	TabWrite(f, 0, "type "+class+" struct {\n")
	fieldList := strings.Split(fields, ", ")
	for _, field := range fieldList {
		TabWrite(f, 1, field+"\n")
	}
	TabWrite(f, 0, "}\n")
}

func TabWrite(f *os.File, tabs uint8, content string) {
	prefix := "    "
	pre_content := ""
	for _ = range tabs {
		pre_content += prefix
	}
	f.WriteString(pre_content + content)
}
