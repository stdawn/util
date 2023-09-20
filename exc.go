/**
 * @Time: 2023/4/24 9:52
 * @Author: LiuKun
 * @File: util.go
 * @Description:
 */

package util

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// GetProgramRootDir 获取应用程序所在根目录
func GetProgramRootDir() (string, error) {
	p, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(p)
	return dir, nil
}

// Pause 按回车键继续
func Pause() {
	fmt.Printf("按回车键继续...\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
	}
}

// ExcFuncAndCountTime  采用defer后无法捕获异常
func ExcFuncAndCountTime(f func()) {
	startTime := time.Now()
	f()
	diff := time.Now().Sub(startTime)
	s := DurationDes(diff)
	fmt.Printf("任务执行完毕, 耗时%s\n", s)
	Pause()
}

// IsExeRunning 程序是否运行，exeName:进程名称
func IsExeRunning(exeName string) bool {
	buf := bytes.Buffer{}
	cmd := exec.Command("wmic", "process", "get", "name,executablepath")
	cmd.Stdout = &buf
	_ = cmd.Run()

	cmd2 := exec.Command("findstr", exeName)
	cmd2.Stdin = &buf
	data, err := cmd2.CombinedOutput()
	if err != nil && err.Error() != "exit status 1" {
		//XBLog.LogF("ServerMonitor", "IsExeRunning CombinedOutput error, err:%s", err.Fail())
		return false
	}

	strData := string(data)
	if strings.Contains(strData, exeName) {
		return true
	} else {
		return false
	}
}

// CloseProgram 关闭程序
func CloseProgram(programName string) error {
	if !IsExeRunning(programName) {
		return nil
	}
	cmd := exec.Command("taskkill", "/f", "/t", "/im", programName)
	return cmd.Start()
}
