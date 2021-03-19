package execcmd

import (
	"io/ioutil"
	"os/exec"
)

const node = "120.25.25.1:8545"

func CallGeth(method, params string) (string, error) {

	cmd := exec.Command("bash", "-c", `curl -H "Content-Type: application/json" -X POST '`+node+`' --data '{"jsonrpc":"2.0","method":"`+method+`","params":`+params+`,"id":1}'`)

	stdout, err := cmd.StdoutPipe() // 创建获取命令输出管道
	if err != nil {
		return "", err
	}

	// 执行命令
	if err := cmd.Start(); err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(stdout) // 读取所有输出
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}
