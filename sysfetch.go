package main

import (
	"fmt"
 	"os"
	"runtime"
  "syscall"
  "path/filepath" 
  "time"

  // "github.com/shirou/gopsutil/v3/cpu"
	// "github.com/shirou/gopsutil/v3/mem"	
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procGetTickCount = modkernel32.NewProc("GetTickCount64")
)
func GetTickCount() uint64 {
	ret, _, _ := procGetTickCount.Call()
	return uint64(ret)
}
func main() {
  start := time.Now()

	fmt.Println("\033[31m _____     _ _      ")
  fmt.Println("\033[31m|   __|___|_|_|_____")
  fmt.Println("\033[31m|   __|  _| | |     |")
  fmt.Println("\033[31m|_____|_| |_|_|_|_|_|")
  //fmt.Println("")
  //fmt.Println("")


  fmt.Println("\033[32m--------")

	fmt.Println("\033[33mOS:  ", runtime.GOOS)
	fmt.Println("\033[34mArch:", runtime.GOARCH)

	hostname, _ := os.Hostname()
	fmt.Println("\033[35mHostname: ", hostname)
  
	
  uptimeMillis := GetTickCount()
	uptimeDuration := time.Duration(uptimeMillis) * time.Millisecond

	uptimeHours := int(uptimeDuration.Hours())
	uptimeMinutes := int(uptimeDuration.Minutes()) % 60  
  fmt.Printf("\033[36mUptime: 󰥔 %d hours, %d minutes\n", uptimeHours, uptimeMinutes)

  //fmt.Println("\033[34mUptime: 󰥔", GetTickCount()/1000, "seconds")

	shellPath := os.Getenv("ComSpec") // replace with "ComSpec" for Windows
	shellName := filepath.Base(shellPath)
	fmt.Println("\033[31mShell: ", shellName)
  fmt.Println("\033[32mTerminal:   WezTerm")

  // Get CPU usage
	// cpuPercent, _ := cpu.Percent(time.Second, false)
  // fmt.Printf("\033[33mCPU Usage:  %.2f%%\n", cpuPercent[0])

  // Get memory usage
	// memStat, _ := mem.VirtualMemory()
	// fmt.Printf("\033[34mMemory Usage: 󰻠 %.2f%%\n", memStat.UsedPercent)

  elapsed := time.Since(start)
  fmt.Printf("\033[35mExecution Time: 󱎫 %s\n", elapsed)
 
 
 


}
