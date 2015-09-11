

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

var (
	name, location, age = "Prine Oberyn", "Dorne", 32
)

// A variable can contain any type, including functions:
func main() {
	action := func() {
                  //doing something
	}
	action ()

	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)
}


// Outside a function, every construct begins with a keyword (var, func, and so on) and the := construct is not available.


// things are read from left to right with the type declaration being on the right

// Constants are declared like variables, but with the const keyword.

// Constants can only be character, string, boolean, or numeric values and cannot be declared using the := syntax.
// An untyped constant takes the type needed by its context.

const (
	Pi = 3.14
	Truth = false
	Big = 1 << 100
	Small = Big >> 99
)
func main() {
	const Greeting = ""
	fmt.Println(Greeting)
	fmt.Println(Pi)
}




func main() { 
    name := "Caprica-Six" 
    aka := fmt.Sprintf("Number	%d", 6) 
    fmt.Printf("%s is also known as %s", name, aka) 
}


// Every Go program is made up of packages.  Programs starting running in package main
