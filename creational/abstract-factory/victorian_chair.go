package abstract_factory

type VictorianChair struct{}

func (c VictorianChair) GetInfo() string {
	return "This is a Victorian chair."
}
