package learn

import "fmt"

// this is the main file

func some() {
	fmt.Print("testing this with another package name, another main in same folder")
}

func ReturnLibraryName() string {
	return "you have to define the mod file, each folder is considered as a module then, this function Some inside a module learn"
}
