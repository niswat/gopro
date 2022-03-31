package suite

import "fmt"

func main() {
  fmt.Println("This is suite")
  res := suite()
  fmt.Println(res)
}

func suite() string {
	return "suite"
}
