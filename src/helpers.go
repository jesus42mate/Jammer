package main

import (
  "bufio"
  "errors"
  "fmt"
  "os/exec"
  "github.com/eiannone/keyboard"
  "golang.org/x/term"
)

func TermChoice(choices []string, isBinary bool, term *term.Terminal) (chosen map[int]bool, err error) {
  position := 0
  m := map[int]bool{}
  exit := false

  ClearScreen()
  ResetCaret()
  term.Write([]byte("\nInstructions:\n\n"))
  term.Write([]byte("'j' and 'k' to move between choices\n\n"))
  term.Write([]byte("'a' to mark a choice\n\n"))
  term.Write([]byte("'q' to exit\n\n"))
  term.Write([]byte("Enter to continue\n"))

  for i := 0; i < len(choices); i++ {
    m[position] = false
    if position == i {
      term.Write([]byte("\n -> [] " + choices[i]))
    } else {
      term.Write([]byte("\n    [] " + choices[i]))
    }
  }

  for (exit != true) {
    char, key, err := keyboard.GetSingleKey()
    ClearScreen()
    ResetCaret()
    if err != nil {
      fmt.Println("ERROR: ", err)
    }
    if char == 'k' && position > 0 {
      position = position - 1
    }
    if char == 'j' && position < len(choices) - 1 {
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
    term.Write([]byte("\nInstructions:\n\n"))
    term.Write([]byte("'j' and 'k' to move between choices\n\n"))
    term.Write([]byte("'a' to mark a choice\n\n"))
    term.Write([]byte("'q' to exit\n\n"))
    term.Write([]byte("Enter to continue\n"))
    for i := 0; i < len(choices); i++ {
      if position == i {
	if m[position] == true {
	  term.Write([]byte("\n -> [x] " + choices[i]))
	} else {
	  term.Write([]byte("\n -> [] " + choices[i]))
	}
      } else {
	if m[i] == true {
	  term.Write([]byte("\n    [x] " + choices[i]))
	} else {
	  term.Write([]byte("\n    [] " + choices[i]))
	}
      }
    }
  }

  return m, nil

}

func ChooseFrom(choices []string, isBinary bool) (map[int]bool, error) {
  position := 0
  m := map[int]bool{}
  exit := false

  fmt.Print("Instructions:\n\n")
  fmt.Print("'j' and 'k' to move between choices\n\n")
  fmt.Print("'a' to mark a choice\n\n")
  fmt.Print("'q' to exit\n\n")
  fmt.Print("Enter to continue\n\n")

  for i := 0; i < len(choices); i++ {
    m[position] = false
    if position == i {
      fmt.Println(" -> []  " + choices[i])
    } else {
      fmt.Println("    []  " + choices[i])
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
    if char == 'j' && position < len(choices) - 1 {
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
    fmt.Print("'j' and 'k' to move between choices\n\n")
    fmt.Print("'a' to mark a choice\n\n")
    fmt.Print("Q to exit\n\n")
    fmt.Print("Enter to continue\n\n")
    for i := 0; i < len(choices); i++ {
      if position == i {
	if m[position] == true {
	  fmt.Println(" -> [x] " + choices[i])
	} else {
	  fmt.Println(" -> []  " + choices[i])
	}
      } else {
	if m[i] == true {
	  fmt.Println("    [x] " + choices[i])
	} else {
	  fmt.Println("    []  " + choices[i])
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
    fmt.Printf("[yes/no]:")
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


