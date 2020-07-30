package learning

import "fmt"

// We can define a struct like so. Structs are kind of like Objects (They are "a collection of fields")
// The default value of a struct is nil
type Corgi struct {
  name string
  color string
  age int
}

func StructsExample() {
  var corgi1 = Corgi{ "Jimmy", "Orange", 2 }

  // Or shorthand (with type inference) *Note* that Go does NOT use implicit type conversion
  corgi2 := Corgi{ "Tyler", "Blue", 5 }
  fmt.Println(corgi1)
  fmt.Println(corgi2)

  // We can access/edit properties for a struct with a dot
  fmt.Printf("Name: %s\nColor: %s\nAge: %b\n", corgi1.name, corgi1.color, corgi1.age)

  // There can also be pointers to a struct
  corgiPointerTo1 := &corgi1
  corgiPointerTo1.name = "Jim" // This changes corgi1's name to Jim from Jimmy

  // Struct literals
  var (
    c1 = Corgi{"James", "Blue", 3} // c1 has type Corgi
    c2 = Corgi{ "Will", "Gray", 5}
    cp = &Corgi{"Shadow", "Black", 3} // Pointer to struct (Not really sure what this would be used for.)
  )

  fmt.Println(c1, c2, *cp);
}
