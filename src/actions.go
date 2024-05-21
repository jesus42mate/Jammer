package main

import (
  "bufio"
  "fmt"
  "os/exec"
)

func AptUpdate() {
  fmt.Println("\nUpdating 'apt'...")
  output_in_bytes, err := exec.Command("sudo", "apt", "update").Output()
  if err != nil {
    fmt.Printf("ERROR: %s", err)
  }
  
  output := string(output_in_bytes[:])
  fmt.Printf("%s", output)
}

func AptUpgrade() {
  fmt.Println("\nUpgrading 'apt'...")
  out, err := exec.Command("sudo", "apt", "upgrade", "-y").Output()
  if err != nil {
    fmt.Println("ERROR: ", err)
  }
  fmt.Println("Upgrade successfull!")
  output := string(out[:])
  fmt.Printf("%s", output) 
}

func InstallNeovim(scn *bufio.Scanner) {

  fmt.Println("\nInstalling 'Neovim'...")

  output_in_bytes, err := exec.Command("cd ~/").Output()
  if err != nil {
    fmt.Printf("ERROR: %s", err)
  }
  fmt.Println("\nWould you like to install xclip? (Recommended)")

  ChooseFrom([]string{"yes/no"}, false)

  if ReadYesOrNo(scn) {
    ShellExec("sudo", "apt", "install", "xclip")
  }

  output := string(output_in_bytes[:])
  fmt.Printf("%s", output)

}
