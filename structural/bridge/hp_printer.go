package bridge

import "fmt"

type HpPrinter struct {
}

func (p *HpPrinter) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
