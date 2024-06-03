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

func WriteTerm(term *term.Terminal, msg string) {
  term.Write([]byte(msg+"\n"))
}

func WriteTermError(term *term.Terminal, err error) {
  term.Write([]byte(fmt.Sprintf("error: %v\n", err)))
}

func WriteTermPanic(term *term.Terminal, err error) {
  term.Write([]byte(fmt.Sprintf("error: %v\n", err)))
  panic(err)
}
