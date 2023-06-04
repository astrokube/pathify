package main

import (
	"fmt"
	"github.com/astrokube/pathify"
)

func ExampleArrayWithOneElement() {
	fmt.Println(pathify.Array().Set("[0]", "Jane").JSON())
	fmt.Println(pathify.Array().Set("[0].firstname", "Jane").JSON())
	fmt.Println(pathify.Array().Set("[2].firstname", "Jane").JSON())
	fmt.Println(pathify.Array().Set("[0]", "David", "[2].firstname", "Jane").JSON())
	// Output:
	// ["Jane"]
	// [{"firstname":"Jane"}]
	// [null,null,{"firstname":"Jane"}]
	// ["David",null,{"firstname":"Jane"}]

}
