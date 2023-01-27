package main

import (
	"log"
)

const (
	errNodeIsNotEmpty     = "graph %s is already pointing to another graph on its %s"
	errDirectionNotMapped = "direction %s is not mapped"
)

type nodeDirection string

type BidirectionalGraph struct {
	Weight    int
	Name      string
	LeftNode  *BidirectionalGraph
	RightNode *BidirectionalGraph
}

var (
	left  nodeDirection = "LEFT"
	right nodeDirection = "RIGHT"
)

func (g *BidirectionalGraph) LinkNode(node *BidirectionalGraph, direction nodeDirection) {
	switch direction {
	case right:
		if g.RightNode != nil {
			log.Fatalf(errNodeIsNotEmpty, g.Name, direction)
		}
		g.RightNode = node
	case left:
		if g.LeftNode != nil {
			log.Fatalf(errNodeIsNotEmpty, g.Name, direction)
		}
		g.LeftNode = node
	default:
		log.Fatalf(errDirectionNotMapped, direction)
	}
}

func NewGraph(name string, weight int) *BidirectionalGraph {
	return &BidirectionalGraph{
		Weight: weight,
	}
}

func main() {
	gasolineStation := NewGraph("Gasoline Station", 0)
	shopping := NewGraph("Shopping", 5)
	superMarket := NewGraph("Supermarket", 3)
	policeStation := NewGraph("Police Station", 7)
	home := NewGraph("Home", 9)

	/* 		    (Gasoline Station)
				/				\
	           /		 	 	 \
			(Shopping)		  (SuperMarket)
			/ 		\		 	/		\
	       /		 \	   	   / 		 \
		(Home)	(Police)  (Shopping)	(Home) */
	gasolineStation.LinkNode(shopping, left)
	shopping.LinkNode(policeStation, right)
	shopping.LinkNode(home, left)
	gasolineStation.LinkNode(superMarket, right)
	superMarket.LinkNode(shopping, right)
	superMarket.LinkNode(home, left)
}
