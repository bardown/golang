package main

import (
	"fmt"
	"os/exec"
)

type Player struct {
	Name        string
	GamesPlayed int
	Goals       int
	Assists     int
	Points      int
}

type Game struct {
	Vs string
	GF int
	GA int
}

func pipedCMD() {
	cmd := "echo | systemctl status mariadb.service"
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(output))
}

//func waitForMe(stressApp string) (results string) {
//	cmd := exec.Command(stressAppp)
//
//	if err := cmd.Start(); err != nil {
//		fmt.Printf("Error not nil on cmd.Start() for sleep 5: %v\n", err)
//	}
//	log.Printf("Waiting for command to finish...")
//
//	if err := cmd.Wait(); err != nil {
//		log.Fatal(fmt.Sprintf("Command finished with error: %v\n", err))
//	}
//	fmt.Printf("Command finsihed with great success!\n")
//}

func main() {
	Lucas := Player{Name: "Lucas", Goals: 5, Assists: 3, Points: 8, GamesPlayed: 2}

	fmt.Printf("Name: %s, Goals: %d, Assists: %d, Points: %d, GamesPlayed: %d", Lucas.Name, Lucas.Goals, Lucas.Assists, Lucas.Points, Lucas.GamesPlayed)
}
