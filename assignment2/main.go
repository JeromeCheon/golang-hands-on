package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

// assignment2 : Hard Mode Assignment
func main() {
	args := os.Args
	if len(args) != 2 {
		os.Exit(0)
	}
	filename := args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)

	/* 아래는 정답이 아닌, 내 풀이 */
	// lw := logWriter{}
	// io.Copy(lw, file)

}

func (logWriter) Write(p []byte) (n int, err error) {
	content := string(p)
	fmt.Println(content)
	return len(content), nil
}
