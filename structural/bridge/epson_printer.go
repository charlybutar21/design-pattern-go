package bridge

import "fmt"

type EpsonPrinter struct {
}

func (p *EpsonPrinter) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
