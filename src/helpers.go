package main

import (
  "bufio"
  "errors"
  "fmt"
  "os/exec"

  "github.com/eiannone/keyboard"
)

func ChooseFrom(options []string) (map[int]bool, error) {
  ShellExec("clear")

  position := 0
  m := map[int]bool{}
  exit := false

  fmt.Print("Instructions:\n\n")
  fmt.Print("'j' and 'k' to move between options\n\n")
  fmt.Print("'a' to mark an option\n\n")
  fmt.Print("Q to exit\n\n")
  fmt.Print("Enter to continue\n\n")

  for i := 0; i < len(options); i++ {
    m[position] = false
    if position == i {
      fmt.Println(" -> []  " + options[i])
    } else {
      fmt.Println("    []  " + options[i])
    }
  }

  for (exit != true) {
    char, key, err := keyboard.GetSingleKey()
    ShellExec("clear")
    if err != nil {
      fmt.Println("ERROR: ", err)
    }
    if char == 'k' && position > 0 {
      position = position - 1
    }
    if char == 'j' && position < len(options) - 1 {
      position = position + 1
    }
    if char == 'a' {
      if m[position] == true {
	m[position] = false
      } else {
	m[position] = true
      }
    }
    if char == 'q' {
      return m, errors.New("Canceled the operation, exiting.")
    }
    if key == keyboard.KeyEnter {
      exit = true
    }
    fmt.Print("Instructions:\n\n")
    fmt.Print("'j' and 'k' to move between options\n\n")
    fmt.Print("'a' to mark an option\n\n")
    fmt.Print("Q to exit\n\n")
    fmt.Print("Enter to continue\n\n")
    for i := 0; i < len(options); i++ {
      if position == i {
	if m[position] == true {
	  fmt.Println(" -> [x] " + options[i])
	} else {
	  fmt.Println(" -> []  " + options[i])
	}
      } else {
	if m[i] == true {
	  fmt.Println("    [x] " + options[i])
	} else {
	  fmt.Println("    []  " + options[i])
	}
      }
    }
  }

  return m, nil
}

// Executes a shell command
// 
// @param name is the command to be executed; But it
// might also be 'sudo', while the actual command is inside args. 
func ShellExec(name string, args ...string) {
  out, err := exec.Command(name, args...).Output()
  if err != nil {
    fmt.Println("ERROR:", err)
  }
  output := string(out[:])
  fmt.Println("output: ", output)
}

func ReadYesOrNo(scn *bufio.Scanner) bool {
  var exit bool = false
  for (exit == false) {
    fmt.Printf("(yes/no):")
    inp := ReadWord(scn)
    if inp == "yes" || inp == "y" {
      return true
    } 
    if inp == "no" || inp == "n" {
      return false
    } 
  }
  return true
}

func ReadWord(scn *bufio.Scanner) string {
  scn.Scan()
  var inp string = string(scn.Bytes()[:])
  return inp
}

func PrintInput(info []byte) {
  fmt.Printf(string(info[:]))
}

