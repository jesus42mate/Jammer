package main

import (
  "fmt"
)

func main() {
  fmt.Println("Welcome, what would you like to do?")

  choices, err := ChooseFrom([]string{
    "1) Update and Upgrade apt packet manager.",
    "2) Install Neovim from source(LTS).",
    "3) Install NVM (Node Version Manager).",
  })

  if err != nil {
    fmt.Println(err)
    return
  }

  if choices[0] {
    //AptUpdate()
    //AptUpgrade()
  }

  if (choices[1]) {
    InstallNeovim()
  }


  

}
