package network

import (
    "fmt"
)

// Node represents a peer in the P2P network.
type Node struct {
    ID      string
    Address string
}

// Network represents the P2P network.
type Network struct {
    Nodes map[string]*Node
}

// NewNetwork creates a new P2P network.
func NewNetwork() *Network {
    return &Network{
        Nodes: make(map[string]*Node),
    }
}

// AddNode adds a new node to the network.
func (n *Network) AddNode(id, address string) {
    n.Nodes[id] = &Node{ID: id, Address: address}
    fmt.Printf("Node added: %s at %s\n", id, address)
}

// ConnectNodes connects two nodes in the network (dummy implementation).
func (n *Network) ConnectNodes(id1, id2 string) {
    if _, exists := n.Nodes[id1]; !exists {
        fmt.Printf("Node %s does not exist.\n", id1)
        return
    }
    if _, exists := n.Nodes[id2]; !exists {
        fmt.Printf("Node %s does not exist.\n", id2)
        return
    }
    fmt.Printf("Connected %s and %s.\n", id1, id2)
}

// BroadcastMessage sends a message to all nodes in the network.
func (n *Network) BroadcastMessage(message string) {
    fmt.Printf("Broadcasting message: %s\n", message)
    for _, node := range n.Nodes {
        // Here you would add the logic to send the message to each node.
        fmt.Printf("Message sent to node %s at %s\n", node.ID, node.Address)
    }
}
