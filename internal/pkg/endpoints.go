package pkg

import "fmt"

type Endpoints struct {
	Ticket  string
	Version string
	Node    string
	Qemu    string
}

func (n *Endpoints) GenerateQemuEndpoints(uri string, node string) {
	n.Qemu = fmt.Sprintf("%s/api2/json/nodes/%s/qemu", uri, node)
}

func GenerateBaseEndpoints(uri string) *Endpoints {
	return &Endpoints{
		Ticket:  fmt.Sprintf("%s/api2/json/access/ticket", uri),
		Version: fmt.Sprintf("%s/api2/json/version", uri),
		Node:    fmt.Sprintf("%s/api2/json/nodes", uri),
	}
}
