package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "os/exec"

  "github.com/eiannone/keyboard"
  "golang.org/x/term"
)

// Binary choice will return true for "yes" and false for "no",
func BinaryChoice(term *term.Terminal, question string) bool {
  ResetCaret()
  ClearScreen()
  position := false
  sure := false
  exit := false
  WriteInstructions(term)
  WriteTerm(term, question)
  if position && !sure {
    WriteTerm(term, " -> yes")
    WriteTerm(term, "    no")
  }
  if !position && !sure {
    WriteTerm(term, "    yes")
    WriteTerm(term, " -> no")
  }
  if position && sure {
    WriteTerm(term, " -> yes   sure?")
    WriteTerm(term, "    no")
  } 
  if !position && sure {
    WriteTerm(term, "    yes")
    WriteTerm(term, " -> no   sure?")
  }
  for (!exit) {
    char, _, err := keyboard.GetSingleKey()
    if err != nil {
      fmt.Println("ERROR: ", err)
    }
    ClearScreen()
    ResetCaret()
    WriteInstructions(term)
    if char == 'k' && !position {
      position = true
      sure = false
    }
    if char == 'j' && position {
      position = false
      sure = false
    }
    if char == 'a' {
      if sure {
	return position
      } else {
	sure = true
      }
    }
    WriteTerm(term, question)
    if position && !sure {
      WriteTerm(term, " -> yes")
      WriteTerm(term, "    no")
    }
    if !position && !sure {
      WriteTerm(term, "    yes")
      WriteTerm(term, " -> no")
    }
    if position && sure {
      WriteTerm(term, " -> yes   sure?")
      WriteTerm(term, "    no")
    } 
    if !position && sure {
      WriteTerm(term, "    yes")
      WriteTerm(term, " -> no   sure?")
    }
  }
  return false
}




//if err != nil {
//  WriteTermError(term, err)
//}
//if choices[0] { // if true
//  return true
//} else {
//  return false
//}


// TermChoice is meant to present a set of choices into a raw-moded terminal
//
// @param choices is a list of strings that the user can choose from
//
// @param isBinary is a boolean that determines if the user can choose multiple choices
//
// @param term is the terminal object that the function writes to
//
// @return chosen is a map of integers to booleans, where the integer is the index of the choice
// and the boolean is true if the choice is selected, and false otherwise
//
// @return err is an error that is returned if the user exits the function
func TermChoice(choices []string, term *term.Terminal, legend string) (chosen map[int]bool, err error) {
  position := 0
  m := map[int]bool{}
  exit := false

  ClearScreen()
  ResetCaret()
  WriteInstructions(term)

  for i := 0; i < len(choices); i++ {
    m[position] = false
    if position == i {
      opt := fmt.Sprintf("\n -> []  %s", choices[i])
      term.Write([]byte(opt))
    } else {
      opt := fmt.Sprintf("\n    [] %s", choices[i])
      term.Write([]byte(opt))
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
    WriteInstructions(term)
    for i := 0; i < len(choices); i++ {
      if position == i {
	if m[position] == true {
	  opt := fmt.Sprintf("\n -> [x]  %s", choices[i])
	  term.Write([]byte(opt))
	} else {
	  opt := fmt.Sprintf("\n -> []  %s", choices[i])
	  term.Write([]byte(opt))
	}
      } else {
	if m[i] == true {
	  opt := fmt.Sprintf("\n    [x]  %s", choices[i])
	  term.Write([]byte(opt))
	} else {
	  opt := fmt.Sprintf("\n    [] %s", choices[i])
	  term.Write([]byte(opt))
	}
      }
    }
  }
  WriteTerm(term, "\n")
  return m, nil
}

// Executes a shell command
// 
// @param name is the command to be executed; But it
// might also be 'sudo', while the actual command is inside args. 
func ShellExec(term *term.Terminal, name string, args ...string) {
  out, err := exec.Command(name, args...).Output()
  if err != nil {
    WriteTermError(term, err)
  }
  output := string(out[:])
  WriteTerm(term, output)
}

func ReadYesOrNo(scn *bufio.Scanner, term *term.Terminal) bool {
  var exit bool = false
  for (exit == false) {
    WriteTerm(term, "[yes/no]:")
    inp, err := term.ReadLine()
    if err != nil {
      WriteTermError(term, err)
    } 
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

func PrintByteInput(info []byte) {
  fmt.Printf(string(info[:]))
}

func ChangeDir(where string) {
  err := os.Chdir(where)
  if err != nil {
    fmt.Printf("%s", err)
  }
}

func WriteInstructions(term *term.Terminal) {
  WriteTerm(term, "\nInstructions:\n")
  WriteTerm(term, "'j' and 'k' to move between options\n")
  WriteTerm(term, "'a' to mark an option\n")
  WriteTerm(term, "'q' to exit\n")
  WriteTerm(term, "Enter to continue")
}


