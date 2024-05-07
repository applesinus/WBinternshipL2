package main

import (
	"fmt"
	"pattern"
)

func main() {
	// test pattern Facade
	fmt.Printf("\nFacade test:\n")
	pattern.FacadeTest()

	// test pattern Builder
	fmt.Printf("\nBuilder test:\n")
	pattern.BuilderTest()

	// test pattern Visitor
	fmt.Printf("\nVisitor test:\n")
	pattern.VisitorTest()

	// test pattern Command
	fmt.Printf("\nCommand test:\n")
	pattern.CommandTest()

	// test pattern Chain of Responsibility
	fmt.Printf("\nChain of Responsibility test:\n")
	pattern.ChainOfRespTest()

	// test pattern Factory Method
	fmt.Printf("\nFactory Method test:\n")
	pattern.FactoryMethodTest()

	// test pattern Strategy
	fmt.Printf("\nStrategy test:\n")
	pattern.StrategyTest()

	// test pattern State
	fmt.Printf("\nState test:\n")
	pattern.StateTest()
}
