package main

import (
	"fmt"

	"github.com/caoyingjunz/pixiulib/exec"
)

func main() {
	exec := exec.New()

	// 确认命令行是否存在
	if _, err := exec.LookPath("ls"); err != nil {
		panic(err)
	}
	// 属性
	out, err := exec.Command("ls", "-al").CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
