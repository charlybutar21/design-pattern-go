package abstract_factory

type ModernTable struct{}

func (t ModernTable) GetInfo() string {
	return "This is a modern table."
}
