package main

import(
  "github.com/fatih/color"
  "fmt"
)

func main() {
  red := color.New(color.FgRed).SprintFunc()
  amir := "salammmm   " + red("arminnn")
  fmt.Println(amir)
}
