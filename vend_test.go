package main

import "testing"
import "fmt"

func TestVend(_t *testing.T) {
	//
	store := get_store()

	// create new product
	// ---------
	store.create_product("0", "bowl", 10)
	store.create_product("1", "table", 100)
	store.create_product("2", "chair", 300)
	store.create_product("3", "cup", 5)

  // list product 
	// ---------
	fmt.Printf("product list\n")
	fmt.Printf("----------- \n")
	list,e := store.list_products()
	if e!=nil {
		panic(e)
	}
	for i,v := range list {
		fmt.Printf("%d: %v\n", i, v)
	}
	fmt.Printf("\n")

	// new order
	// ---------
	fmt.Printf("create new order\n")
	fmt.Printf("----------- \n")
	items := []OrderItem{ {"1", 3, 0, 0.0, 0}, {"2", 2, 0, 0.0, 0} }
	req := Order{items:items, discount:15}
	// fmt.Printf("order: %v\n", req)
	// req.print()
	response,e := store.proc_order(req)
	if e!=nil {
		panic(e)
	}
	// fmt.Printf("respone: %v", response)
	response.print()
}

