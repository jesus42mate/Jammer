package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  var scn *bufio.Scanner = bufio.NewScanner(os.Stdin)
  fmt.Println("Welcome, what would you like to do?")

  choices, exit := ChooseFrom([]string{
    "1) Update and Upgrade apt packet manager.",
    "2) Install Neovim from source(LTS).",
    "3) Install NVM (Node Version Manager).",
  })
  if exit != nil {
    fmt.Println(exit)
    return
  }

  if ReadYesOrNo(scn) {
    fmt.Println("true!")
  } else {
    fmt.Println("false!")
  }


  if choices[0] {
    //AptUpdate()
    //AptUpgrade()
  }

  if (choices[1]) {
    //InstallNeovim()
  }

  fmt.Println("\n-----------------")
  fmt.Println("\nGraceful shutdown.")
  fmt.Println("Thanks for using Jammer!")

}
