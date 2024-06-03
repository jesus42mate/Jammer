package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
  WriteTerm(term, "Upgrading 'apt'...")
  out, err := exec.Command("sudo", "apt", "upgrade", "-y").Output()
  if err != nil {
    WriteTermError(term, err)
  }
  fmt.Println("Upgrade successfull!")
  WriteTerm(term, string(out[:]))
}

func InstallNeovim(
  scn *bufio.Scanner,
  term *term.Terminal,
  homeDir string,
) error {
  WriteTerm(term, "\nInstalling 'Neovim' latest stable...\n")
  ChangeDir(homeDir)
  if BinaryChoice(term, "Would you like to install xclip? (Recommended)") {
    ShellExec(term, "sudo apt install xclip")
  }
  
  ShellExec(term, "sudo apt-get install clang -y")
  ShellExec(term, "sudo apt-get install cmake -y")
  ShellExec(term, "sudo apt-get install ninja-build gettext cmake unzip curl build-essential -y")
  /* Dangerous */
  ShellExec(term, "git clone https://github.com/neovim/neovim")
  err := os.Chdir("neovim")
  if err != nil {
    WriteTermError(term, err)
    return errors.New("Failed to cd into the neovim folder")
  } else {
    ShellExec(term, "make CMAKE_BUILD_TYPE=RelWithDebInfo" )
    ShellExec(term, "git checkout stable")
    ShellExec(term, "cd build && cpack -G DEB && sudo dpkg -i nvim-linux64.deb")
    ShellExec(term, "sudo make install")
    ShellExec(term, "sudo apt-get install ripgrep")
  }

  //if !BinaryChoice(term, "Would you like to install Jesus' neovim configuration? (Recommended!)") {
  //  return
  //}

  ChangeDir(homeDir)

  file, err := os.Open(".config")
  if err != nil {
    WriteTermError(term, err)
    err := os.Mkdir(".config", 700)
    if err != nil {
      WriteTermPanic(term, err)
    } else {
      file, _ = os.Open(".config")
    }
  } else {
    WriteTerm(term, "# Found .config")
  }
  err = file.Chdir()
  if err != nil {
    WriteTermError(term, err)
    return errors.New("Something failed when trying to change directories")
  }

  ShellExec(term, "pwd")
  ShellExec(term, "git clone https://github.com/jesus42mate/neovim-config")
  PrintDir(term , "neovim-config")

  _, err = os.Open("nvim")
  if err != nil {
    WriteTerm(term, "# nvim folder not found. Great!")
    ShellExec(term, "mv neovim-config nvim")
    WriteTerm(term, "# Finished setting up neovim, try it with: $ nvim 'example.txt'")
  } else {
    WriteTerm(term, "# Found a nvim folder, renaming to nvim-old to swap configs")
    ShellExec(term, "mv nvim nvim-old") 
  }
  return nil
}

// Purge removes all the files that may have been created with Jammer
func Purge() {
}




