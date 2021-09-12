package voice

import (
	"os"
	"os/exec"
	"bufio"
	"fmt"
)

func Run(words_to_say string, out_file_name string) {
	name := "tts"
	args := []string{
		"--text",
		words_to_say,
		"--model_name",
		"tts_models/en/ek1/tacotron2",
		"--out_path",
		out_file_name,
	}
	cmd := exec.Command(name, args...)
	pipe, err := cmd.StdoutPipe()

	// create a pipe to connect the cmd and the cmp
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	// start the cmd
	if err := cmd.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	// read the cmd output line by line
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line[:8])
	}

	// wait until the cmd exits
	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}

}
