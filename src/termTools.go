package main

import (
  "fmt"
  "golang.org/x/term"
)

type FormalPanicNeeds struct {
  FD int;
  prevState *term.State;
}

func FormalPanic(needs FormalPanicNeeds, err error) {
  term.Restore(needs.FD, needs.prevState)
  panic(err)
}

func ClearScreen() {
  fmt.Print("\033[2J")
}

func ResetCaret() {
  fmt.Print("\033[H")
}
