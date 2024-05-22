package main

import (
	"bufio"
	"fmt"
	"os/exec"

	"golang.org/x/term"
)

func AptUpdate(term *term.Terminal) {
  WriteTerm(term, "Updating 'apt'...")
  out, err := exec.Command("sudo", "apt", "update").Output()
  if err != nil {
    WriteTermError(term, err)
  }
  WriteTerm(term, string(out[:]))
}

func AptUpgrade(term *term.Terminal) {
  WriteTerm(term, "Updating 'apt'...")
  out, err := exec.Command("sudo", "apt", "upgrade", "-y").Output()
  if err != nil {
    WriteTermError(term, err)
  }
  fmt.Println("Upgrade successfull!")
  WriteTerm(term, string(out[:]))
}

func InstallNeovim(scn *bufio.Scanner, term *term.Terminal, homeDir string) {
  WriteTerm(term, "Installing 'Neovim'...\n")
  ChangeDir(homeDir)
  if BinaryChoice(term, "Would you like to install xclip? (Recommended)") {
    ShellExec(term, "sudo", "apt", "install", "xclip")
  }
}

// Purge removes all the files that may have been created with Jammer
func Purge() {
}




