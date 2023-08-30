package prototype

import "fmt"

type Folder struct {
	Inners []NodeInterface
	Name   string
}

func (f *Folder) Clone() NodeInterface {
	cloneFolder := &Folder{Name: f.Name + "_copy"}
	var tempInners []NodeInterface
	for _, i := range f.Inners {
		copy := i.Clone()
		tempInners = append(tempInners, copy)
	}
	cloneFolder.Inners = tempInners
	return cloneFolder
}

func (f *Folder) PrintDetails(indentation string) {
	fmt.Println(indentation, f.Name)
	for _, i := range f.Inners {
		i.PrintDetails(indentation + indentation)

	}
}
