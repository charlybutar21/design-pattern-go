package prototype

type NodeInterface interface {
	Clone() NodeInterface
	PrintDetails(indentation string)
}
