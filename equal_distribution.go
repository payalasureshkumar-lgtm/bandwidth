
package main

import (
	"fmt"
)

type Node struct {
	ID        int
	Allocated float64
}

type Network struct {
	Nodes          []Node
	TotalBandwidth float64
}

func NewNetwork(count int, bandwidth float64) *Network {
	nodes := make([]Node, count)

	for i := range nodes {
		nodes[i].ID = i + 1
	}

	return &Network{
		Nodes:          nodes,
		TotalBandwidth: bandwidth,
	}
}

func (n *Network) CalculateAllocation() {
	count := len(n.Nodes)

	if count == 0 || n.TotalBandwidth <= 0 {
		return
	}

	share := n.TotalBandwidth / float64(count)

	for i := range n.Nodes {
		n.Nodes[i].Allocated = share
	}
}

func (n *Network) CalculateUtilization() float64 {
	if n.TotalBandwidth <= 0 {
		return 0
	}

	var total float64

	for _, node := range n.Nodes {
		total += node.Allocated
	}

	return (total / n.TotalBandwidth) * 100
}

func (n *Network) DisplayAllocation() {
	fmt.Println("\nEqual Distribution Results")
	fmt.Println("--------------------------")

	for _, node := range n.Nodes {
		fmt.Printf("Node %d: %.2f Mbps\n",
			node.ID, node.Allocated)
	}

	fmt.Printf("Bandwidth Utilization: %.2f%%\n",
		n.CalculateUtilization())
}

func (n *Network) ResetAllocation() {
	for i := range n.Nodes {
		n.Nodes[i].Allocated = 0
	}
}

func main() {
	var count int
	var bandwidth float64

	fmt.Print("Enter number of nodes: ")
	fmt.Scan(&count)

	fmt.Print("Enter total bandwidth: ")
	fmt.Scan(&bandwidth)

	if count <= 0 || bandwidth <= 0 {
		fmt.Println("Invalid network configuration")
		return
	}

	network := NewNetwork(count, bandwidth)

	network.CalculateAllocation()
	network.DisplayAllocation()
}
