package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	ExitOk  = 0
	ExitErr = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	// historyコマンドの実行結果を配列に入れる
	history, err := executeCommand("history")
	if err != nil {
		panic(err)
	}
	historyArr := strings.Split(history, "\n")

	// ユーザーの入力を受け付ける
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter search term: ")
	scanner.Scan()
	searchTerm := scanner.Text()

	// 検索結果を表示し、選択されたコマンドを実行する
	for i := len(historyArr) - 2; i >= 0; i-- {
		if strings.Contains(historyArr[i], searchTerm) {
			fmt.Printf("%d: %s\n", i, historyArr[i])
		}
	}
	fmt.Print("Enter command number: ")
	scanner.Scan()
	commandNumber := scanner.Text()

	cn, err := strconv.Atoi(commandNumber)
	if err != nil {
		panic(err)
	}

	// 選択されたコマンドを実行する
	command := strings.Split(historyArr[cn], " ")[1]
	result, err := executeCommand(command)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return ExitOk
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	status := cli.Run(os.Args)
	os.Exit(status)
}

func executeCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
