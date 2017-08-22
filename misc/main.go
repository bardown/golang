package main

import (
	"fmt"
	"log"
	"os/exec"
)

func pipedCMD() {
	cmd := "echo | systemctl status mariadb.service"
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(output))
}

func waitForMe(stressApp string) (results string) {
	cmd := exec.Command(stressAppp)

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error not nil on cmd.Start() for sleep 5: %v\n", err)
	}
	log.Printf("Waiting for command to finish...")

	if err := cmd.Wait(); err != nil {
		log.Fatal(fmt.Sprintf("Command finished with error: %v\n", err))
	}
	fmt.Printf("Command finsihed with great success!\n")
}

func main() {
	waitForMe()
}
