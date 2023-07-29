package utils

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Command Linux 命令执行
func Command(cmd string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	return string(output), err
}
func CommandByDir(cmd, dir string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	c.Dir = dir
	output, err := c.CombinedOutput()
	return string(output), err
}
func KillProcess(processName string) {
	psCmd := exec.Command("ps", "-eo", "pid,comm")
	output, err := psCmd.Output()
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(output), "\n")
	var targetPID int
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 && fields[1] == processName {
			pid, err := strconv.Atoi(fields[0])
			if err == nil {
				targetPID = pid
				break
			}
		}
	}
	if targetPID == 0 {
		return
	}
	killCmd := exec.Command("kill", strconv.Itoa(targetPID))
	if err := killCmd.Run(); err != nil {
		log.Println(err)
	}
}
