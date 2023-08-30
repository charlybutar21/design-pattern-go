package adapter

import "fmt"

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com ComputerInterface) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
