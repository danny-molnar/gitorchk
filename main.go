package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Verify Git repository
	if _, err := exec.Command("git", "rev-parse").Output(); err != nil {
		fmt.Println("Error: Not inside a Git repository.")
		return
	}

	// Fetch latest remote information
	fmt.Println("Fetching latest information from remote...")
	if _, err := exec.Command("git", "fetch").Output(); err != nil {
		fmt.Println("Error: Unable to fetch latest information from remote.")
		return
	}

	// Determine branch status
	output, err := exec.Command("git", "rev-list", "--left-right", "--count", "main...origin/main").Output()
	if err != nil {
		fmt.Println("Error: Unable to determine branch status.")
		return
	}

	// Parse output
	counts := strings.Fields(string(output))
	if len(counts) != 2 {
		fmt.Println("Error: Unable to process comparison results.")
	}

	// Prompt for action
	if counts[1] != "0" {
		fmt.Printf("Your local 'main' branch iis behind by %s commits.\n", counts[1])
		fmt.Println("Consider rebasig your work with the latest 'main' branch using 'git rebase' or merging updates using 'git merge'.")
	} else {
		fmt.Println("Your local 'main' branch is up-to-date with 'origin/main'.")
	}
}
