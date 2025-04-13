package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
)

func copyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("clip")
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = stdin.Write([]byte(text))
	if err != nil {
		return err
	}

	err = stdin.Close()
	if err != nil {
		return err
	}

	cmd.Wait()
	return nil
}

func makeUrlSafe(prompt string) string {
	return url.QueryEscape(prompt)
}

func usage() {
	fmt.Printf("Usage: lmgt \"prompt here\" [options]\nExample: lmgt \"when is presidents day\"\nExample (copy to clipboard): lmgt \"when is presidents day\" copy\n")
	os.Exit(0)
}

func parseArgs(args []string) (map[string]any, error) {
	if len(args) == 0 {
		return nil, errors.New("arguments must be provided")
	}
	var results map[string]any = map[string]any{
		"prompt": makeUrlSafe(args[0]),
		"copy":   false,
	}
	for _, arg := range args {
		if arg == "copy" {
			results["copy"] = true
		}
		if arg == "help" || arg == "-h" || arg == "--help" {
			usage()
		}
	}
	return results, nil
}

func main() {
	var args []string = os.Args[1:]
	options, err := parseArgs(args)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	var baseUrl string = "https://letmegooglethat.com/?q="
	fullUrl := baseUrl + options["prompt"].(string)
	var copyMessage string = ""
	if options["copy"] == true {
		err := copyToClipboard(fullUrl)
		if err != nil {
			fmt.Printf("Error copying to clipboard: %v\n", err)
			os.Exit(2)
		}
		copyMessage = "(copied to clipboard)\n"
	}
	fmt.Printf("%s%s\n", copyMessage, fullUrl)
}
