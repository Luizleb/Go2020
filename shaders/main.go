package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type shaderType int

const (
	none     shaderType = -1
	vertex              = 0
	fragment            = 1
)

func main() {
	result := make(map[shaderType]string)
	content, err := os.Open("./sources/Basic.shader")
	if err != nil {
		fmt.Println(err)
	}
	getShaders(content, result)
}

func getShaders(file *os.File, res map[shaderType]string) {
	var shType shaderType = none
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		if strings.Contains(line, "#vertex") {
			shType = vertex
			res[shType] = "`\n"
		} else if strings.Contains(line, "#fragment") {
			shType = fragment
			res[shType] = "`"
		} else {
			res[shType] += line + "\n"
		}
	}
}
