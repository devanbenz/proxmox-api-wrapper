package main

import (
	"github.com/devanbenz/go-proxmox-rest/internal/pkg"
	"net/http"
)

func main() {
	client := http.Client{}
	ep := pkg.GenerateBaseEndpoints("https://<ip>:8006")

	t := &pkg.TicketInput{
		Username: "<username>",
		Password: "<password>",
	}

	cookie, csrf := t.GetTokens(&client, ep.Ticket)
	nodesJson := pkg.SendApiGetRequest(cookie, csrf, client, ep.Node)
	nodes := pkg.GetNodeIds(nodesJson)

	vm := &pkg.VirtualMachineInput{
		Node:    nodes[0],
		Vmid:    300,
		Memory:  1200,
		Sockets: 2,
		Cores:   4,
		Sata:    "",
	}

    vm_raw := vm.CreateVirtualMachine()

    pkg.SendApiPostRequest(cookie, csrf, client, ep.Node, vm_raw)
}
