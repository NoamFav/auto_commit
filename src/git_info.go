package src

import (
	"fmt"
	"os/exec"
	"strings"
)

func GitDiff() string {
	cmd := exec.Command("git", "diff")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}

func GitStatus() string {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	status := string(out)
	if strings.TrimSpace(status) == "" {
		fmt.Println("Nothing to commit.")
	}
	return status
}

func Summary() string {
	diff := GitDiff()
	status := GitStatus()

	// Check if both are empty
	if strings.TrimSpace(status) == "" && strings.TrimSpace(diff) == "" {
		return ""
	}

	return fmt.Sprintf("Git Diff: %s\nGit Status: %s", diff, status)
}

func GitAdd() {
	cmd := exec.Command("git", "add", ".")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func GitCommit(message string) {
	cmd := exec.Command("git", "commit", "-m", message)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func GitPush() {
	cmd := exec.Command("git", "push")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func AddCommitPush(message string) {
	GitAdd()
	GitCommit(message)
	GitPush()
}
