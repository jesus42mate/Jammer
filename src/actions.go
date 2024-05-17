package main

import (
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

func InstallNeovim() {
  fmt.Println("\nInstalling 'Neovim'...")
  output_in_bytes, err := exec.Command("cd ~/").Output()
  if err != nil {
    fmt.Printf("ERROR: %s", err)
  }
  output := string(output_in_bytes[:])
  fmt.Printf("%s", output)

}
