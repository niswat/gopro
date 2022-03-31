package suite

import "fmt"

func main() {
  fmt.Println("This is suite")
  res := suite()
  fmt.Println(res)
  op:= add(2,3)
  fmt.Println(op)
}

func suite() string {
	return "suite"
}

func add(a, b int) int {
  res:= a+b
}
