package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode"
)

func main() {
	toInstall := getInput()
	fmt.Println(toInstall)
	runInstallation()
}

func getInput() []string {
	toInstall := []string{}
	input := prompt("Do you want to install zshell? (Y|n)")
	if unicode.ToLower([]rune(input)[0]) == 'y' {
		toInstall = append(toInstall, "zshell")
	}
	input = prompt("Do you want to install spacemacs? (Y|n)")
	if unicode.ToLower([]rune(input)[0]) == 'y' {
		toInstall = append(toInstall, "spacemacs")
	}
	input = prompt("Do you want to install deskterm? (Y|n)")
	if unicode.ToLower([]rune(input)[0]) == 'y' {
		toInstall = append(toInstall, "deskterm")
	}
	input = prompt("Do you want to install tmux? (Y|n)")
	if unicode.ToLower([]rune(input)[0]) == 'y' {
		toInstall = append(toInstall, "tmux")
	}
	return toInstall
}

func runInstallation() {
	prepareSystem()
	installZshell()
	installTmux()
	installSpacemacs()
	activate()
}

func prepareSystem() {
	verifyInstallation("git")
	verifyInstallation("wget")
	verifyInstallation("curl")
}

func installZshell() {
	verifyInstallation("zsh")
	fmt.Println("Installing zshell for user . . .")
	run("bash "+getDir()+"/zshell/install.sh", os.Getenv("SUDO_USER"))
	fmt.Println("Installing zshell for root . . .")
	run("bash "+getDir()+"/zshell/install.sh", "root")
}

func installTmux() {
	verifyInstallation("tmux")
	fmt.Println("Installing tmux for user . . .")
	run("bash "+getDir()+"/tmux/install.sh", os.Getenv("SUDO_USER"))
	fmt.Println("Installing tmux for root . . .")
	run("bash "+getDir()+"/tmux/install.sh", "root")
}

func installSpacemacs() {
	verifyInstallation("emacs")
	fmt.Println("Installing spacemacs for user . . .")
	run("bash "+getDir()+"/spacemacs/install.sh", os.Getenv("SUDO_USER"))
	fmt.Println("Installing spacemacs for root . . .")
	run("bash "+getDir()+"/spacemacs/install.sh", "root")
}

func activate() {
	run("reset")
}

func prompt(params ...string) string {
	prompt := ": "
	if len(params) > 0 {
		prompt = params[0]
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return input
}

func getDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func verifyInstallation(program string) {
	_, err := exec.LookPath(program)
	if err != nil {
		fmt.Println("Installing " + program + " . . .")
		run("apt-get install -y " + program)
	}
}

func run(parameters ...string) error {
	command := parameters[0]
	username := ""
	verbose := "1"
	if len(parameters) > 1 {
		username = parameters[1]
	}
	if len(parameters) > 2 {
		verbose = parameters[2]
	}
	arguments := strings.Fields(command)
	program, err := exec.LookPath(arguments[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(program, arguments[1:]...)
	realhome, _ := syscall.Getenv("HOME")
	if username != "" {
		user, err := user.Lookup(username)
		if err != nil {
			return err
		}
		home := "/home/" + username
		if username == "root" {
			home = "/root"
		}
		err = syscall.Setenv("HOME", home)
		if err != nil {
			return err
		}
		uid, err := strconv.Atoi(user.Uid)
		if err != nil {
			return err
		}
		gid, err := strconv.Atoi(user.Gid)
		if err != nil {
			return err
		}
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}
	if verbose == "1" {
		cmd.Stdout = os.Stdout
	}
	if verbose == "2" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.Start()
	cmd.Wait()
	time.Sleep(100 * time.Millisecond)
	err = syscall.Setenv("HOME", realhome)
	if err != nil {
		return err
	}
	return nil
}
