
package main

import (
	"fmt"
)

type Node struct {
	ID        int
	Demand    float64
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

func (n *Network) ReadDemands() {
	for i := range n.Nodes {
		fmt.Printf("Enter demand for Node %d: ",
			n.Nodes[i].ID)
		fmt.Scan(&n.Nodes[i].Demand)
	}
}

func (n *Network) CalculateTotalDemand() float64 {
	var total float64

	for _, node := range n.Nodes {
		if node.Demand > 0 {
			total += node.Demand
		}
	}

	return total
}

func (n *Network) CalculateAllocation() {
	totalDemand := n.CalculateTotalDemand()

	if totalDemand <= 0 || n.TotalBandwidth <= 0 {
		return
	}

	for i := range n.Nodes {
		if n.Nodes[i].Demand > 0 {
			n.Nodes[i].Allocated =
				(n.Nodes[i].Demand / totalDemand) *
					n.TotalBandwidth
		}
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
	fmt.Println("\nDemand-Aware Distribution Results")
	fmt.Println("----------------------------------")

	for _, node := range n.Nodes {
		fmt.Printf(
			"Node %d | Demand: %.2f Mbps | Allocation: %.2f Mbps\n",
			node.ID,
			node.Demand,
			node.Allocated,
		)
	}

	fmt.Printf("Bandwidth Utilization: %.2f%%\n",
		n.CalculateUtilization())
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

	network.ReadDemands()
	network.CalculateAllocation()
	network.DisplayAllocation()
}
