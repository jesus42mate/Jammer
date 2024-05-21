package main

import (
  //"bufio"
  "fmt"
  "os"
  "golang.org/x/term"
)

func main() {
  var stdin *os.File = os.Stdin
  //var scanner *bufio.Scanner = bufio.NewScanner(stdin)
  var terminal *term.Terminal = term.NewTerminal(stdin, "λ ")
  var FD int = int(stdin.Fd())

  terminal.Write([]byte("\nJammer: Helloo world! Im alive!\n"))

  fmt.Println("\nWelcome to Jammer, what would you like to do?")

  prevState, error := term.MakeRaw(FD)
  formalPanicNeeds := FormalPanicNeeds{ FD, prevState }
  if error != nil {
    FormalPanic(formalPanicNeeds, error)
  } else {
    terminal.Write([]byte("\nJammer: Terminal RAW MODE ACTIVATED (in a robot voice)\n"))
  }

  choices, exit := TermChoice([]string{
    "1) Update and Upgrade apt packet manager.",
    "2) Install Neovim from source(LTS).",
    "3) Install NVM (Node Version Manager).",
  }, false, terminal)
  if exit != nil {
    FormalPanic(formalPanicNeeds, exit)
  }

  if choices[0] {
    AptUpdate()
    AptUpgrade()
  }

  //if (choices[1]) {
  //  //InstallNeovim()
  //}

  defer term.Restore(FD, prevState)
  terminal.Write([]byte("\nGraceful shutdown.\n"))
  terminal.Write([]byte("Thanks for using Jammer!\n"))
  terminal.Write([]byte("----------------------\n"))

}


