package abstract_factory

type ModernChair struct{}

func (c ModernChair) GetInfo() string {
	return "This is a modern chair."
}
