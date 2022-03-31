package suite

import "fmt"

func main() {
  fmt.Println("This is suite")
  op:= add(2,3)
  fmt.Println(op)
}

func add(a, b int) int {
  res:= a+b
  return res
}
