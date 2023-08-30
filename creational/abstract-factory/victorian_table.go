package abstract_factory

type VictorianTable struct{}

func (t VictorianTable) GetInfo() string {
	return "This is a Victorian table."
}
