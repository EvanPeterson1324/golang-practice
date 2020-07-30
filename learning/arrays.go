package learning

import "fmt"

func ArraysExample() {
  var a [2]int // declare a variable as integer array of size 10. Arrays can't be resized
  a[0] = 1
  a[1] = 2

  b := [2]int{3, 4} // or shorthand

  fmt.Println(a, b);

  // Slices - Describes a section of the referenced array.
  var s1 []int = a[0:1]

  // There are slice literals as well.
  var s2 = []bool{true, false} // Creates an array, then build a slice. s2 is a slice!
  fmt.Println(s1,s2);

  // You can leave out high or low bounds when slicing. The below all do the same thing.
  s3 := []int{2, 3, 5, 7, 11, 13}

  s4 := s3[0:6]
  fmt.Println(s4)

  s5 := s3[:6]
  fmt.Println(s5)

  s6 := s3[0:]
  fmt.Println(s6)

  s7 := s3[:]
  fmt.Println(s7)
}
