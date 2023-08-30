package main

import (
	"fmt"
	abstractfactory "gitlab.com/pinvest/design-pattern/creational/abstract-factory"
	"gitlab.com/pinvest/design-pattern/creational/builder"
	factorymethod "gitlab.com/pinvest/design-pattern/creational/factory-method"
	"gitlab.com/pinvest/design-pattern/creational/prototype"
	"gitlab.com/pinvest/design-pattern/creational/singleton"
	"gitlab.com/pinvest/design-pattern/structural/adapter"
	"gitlab.com/pinvest/design-pattern/structural/bridge"
)

func main() {

	fmt.Println("==============================================================")
	fmt.Println("==Test Design Pattern Builder==")
	normalBuilder := builder.NewNormalHouse()
	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	iglooBuilder := builder.NewIglooHouse()
	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

	fmt.Println("==============================================================")
	fmt.Println("==Test Design Pattern Factory Method==")

	carFactory := factorymethod.CarFactory{}
	car := carFactory.Produce()
	fmt.Println("Vehicle type: ", car.GetName())

	motorcycleFactory := factorymethod.MotorcycleFactory{}
	motorcycle := motorcycleFactory.Produce()
	fmt.Println("Vehicle type: ", motorcycle.GetName())

	fmt.Println("==============================================================")

	fmt.Println("==Test Design Pattern Abstract Factory==")

	modernFactory := abstractfactory.ModernFurnitureFactory{}
	modernChair := modernFactory.CreateChair()
	modernTable := modernFactory.CreateTable()
	fmt.Println(modernChair.GetInfo())
	fmt.Println(modernTable.GetInfo())

	victorianFactory := abstractfactory.VictorianFurnitureFactory{}
	victorianChair := victorianFactory.CreateChair()
	victorianTable := victorianFactory.CreateTable()
	fmt.Println(victorianChair.GetInfo())
	fmt.Println(victorianTable.GetInfo())

	fmt.Println("==============================================================")
	fmt.Println("==Test Design Pattern Prototype==")

	file1 := &prototype.File{Name: "File1"}
	file2 := &prototype.File{Name: "File2"}
	file3 := &prototype.File{Name: "File3"}

	folder1 := &prototype.Folder{
		Inners: []prototype.NodeInterface{file1},
		Name:   "Folder1",
	}

	folder2 := &prototype.Folder{
		Inners: []prototype.NodeInterface{folder1, file2, file3},
		Name:   "Folder2",
	}
	folder2.PrintDetails("  ")

	cloneFolder := folder2.Clone()
	cloneFolder.PrintDetails("  ")

	fmt.Println("==============================================================")
	fmt.Println("==Test Design Pattern Singleton==")
	config1 := singleton.GetInstance()
	config2 := singleton.GetInstance()

	// Both config1 and config2 point to the same instance
	fmt.Println("Config 1 API Key:", config1.APIKey)
	fmt.Println("Config 2 API Key:", config2.APIKey)

	// Modifying config through one instance affects the other instance
	config1.APIKey = "new-api-key"
	fmt.Println("Config 2 API Key after modification:", config2.APIKey)

	fmt.Println("==============================================================")

	fmt.Println("==Test Design Pattern Adapter==")
	client := &adapter.Client{}

	mac := &adapter.Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowsMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	fmt.Println("==============================================================")

	fmt.Println("==Test Design Pattern Bridge==")
	hpPrinter := &bridge.HpPrinter{}
	epsonPrinter := &bridge.EpsonPrinter{}

	macComputer := &bridge.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &bridge.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()

	fmt.Println("==============================================================")
}
