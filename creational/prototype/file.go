package prototype

import "fmt"

type File struct {
	Name string
}

func (e *File) Clone() NodeInterface {
	return &File{
		Name: e.Name + "_copy",
	}
}

func (e *File) PrintDetails(indentation string) {
	fmt.Println(indentation, e.Name)
}
