package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	errNodeIsNotEmpty     = "graph %s is already pointing to another graph on its %s"
	errDirectionNotMapped = "direction %s is not mapped"

	nowhereToGo       = "Ops, looks like there's nowhere to go from %s. Trying another way\n"
	endOfTheLine      = "End of the line, nowhere to go from %s. Going back to %s"
	navigatingMessage = "Going %s. Heading to %s\n"
	arrivedMessage    = "Got home after walking %d miles"

	gasolineStationDest = "Gasoline Station"
	policeStationDest   = "Police Station"
	supermarketDest     = "Supermarket"
	shoppingDest        = "Shopping"
	homeDest            = "Home"

	goLeft  = 0
	goRight = 1

	failuresCount = 0
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
		Name:   name,
		Weight: weight,
	}
}

func main() {
	gasolineStation := NewGraph(gasolineStationDest, 0)
	shopping := NewGraph(shoppingDest, 5)
	superMarket := NewGraph(supermarketDest, 3)
	policeStation := NewGraph(policeStationDest, 7)
	home := NewGraph(homeDest, 9)

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

	RandomWayHome(gasolineStation)
}

func RandomWayHome(node *BidirectionalGraph) {
	var ok bool
	rand.Seed(time.Now().UnixNano())

	totalDistance := node.Weight
	startingPoint := node
	currentPoint := node
	for currentPoint.Name != homeDest {
		randomNumber := rand.Intn(2)
		switch randomNumber {
		case goLeft:
			currentPoint, ok = navigate(startingPoint, currentPoint, left)
			if ok {
				totalDistance += currentPoint.Weight
			}
		case goRight:
			currentPoint, ok = navigate(startingPoint, currentPoint, right)
			if ok {
				totalDistance += currentPoint.Weight
			}
		}
	}
	log.Printf(arrivedMessage, totalDistance)
}

func navigate(startPoint, currentPoint *BidirectionalGraph, direction nodeDirection) (*BidirectionalGraph, bool) {
	if currentPoint.LeftNode == nil && currentPoint.RightNode == nil {
		log.Printf(endOfTheLine, currentPoint.Name, startPoint.Name)
		return startPoint, false
	}
	nextPoint := currentPoint
	switch direction {
	case right:
		if currentPoint.RightNode == nil {
			log.Printf(nowhereToGo, direction)
			return navigate(startPoint, currentPoint, left)
		}
		log.Printf(navigatingMessage, direction, currentPoint.RightNode.Name)
		startPoint = currentPoint
		nextPoint = currentPoint.RightNode
	case left:
		if currentPoint.LeftNode == nil {
			log.Printf(nowhereToGo, direction)
			return navigate(startPoint, currentPoint, right)
		}
		log.Printf(navigatingMessage, direction, currentPoint.RightNode.Name)
		startPoint = currentPoint
		nextPoint = currentPoint.LeftNode
	default:
		log.Fatalf(errDirectionNotMapped, direction)
	}

	return nextPoint, true
}
