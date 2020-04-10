package window

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//ShaderType sets the shaders types
type ShaderType int

const (
	none     ShaderType = -1
	vertex              = 0
	fragment            = 1
)

// ShadersToString reads the shaders file to a string map
func ShadersToString() map[ShaderType]string {
	result := make(map[ShaderType]string)
	content, err := os.Open("./shaders/Basic.shader")
	if err != nil {
		fmt.Println(err)
	}
	getShaders(content, result)
	return result
}

func getShaders(file *os.File, res map[ShaderType]string) {
	var shType ShaderType = none
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		if strings.Contains(line, "#vertex") {
			shType = vertex
			//res[shType] = "`\n"
		} else if strings.Contains(line, "#fragment") {
			shType = fragment
			//res[shType] = "`"
		} else {
			res[shType] += line + "\n"
		}
	}
}
