package Utility

import (
	"github.com/AlecAivazis/survey"
	"os/exec"
	"strings"
)

type Branch struct {
	name string
}

func SelectBranch(path string) Branch {
	cmd := exec.Command("git", "branch", "-a")
	cmd.Dir = path
	output, err := cmd.Output()

	if nil != err {
		panic(err)
	}

	branches := strings.Split(string(output), "\n")

	branch := ""
	prompt := &survey.Select{
		Message: "Which branch?",
		Options: branches,
	}
	err2 := survey.AskOne(prompt, &branch)
	if err2 != nil {
		panic(err2)
	}

	return Branch{name: branch}
}
