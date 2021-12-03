package main

import (
  "log"
  "strconv"
)

type Product struct {
	id string
	name string
	price int
}

type OrderItem struct {
	id string
	cnt int
	discount int
	discount_percent float32
		total int
}

type Order struct {
	   items []OrderItem
	discount int  // 0 < discount < total
	total int
}

//helper
func (this Order) print() {
	log.Printf("Order: \n")
	log.Printf("---------------------\n")
	log.Printf("   discount: $%d\n", this.discount)
	log.Printf("      total: $%d\n", this.total)
	for _,item := range this.items {
		item.print()
	}
	log.Printf("---------------------\n\n\n")
}

func (this OrderItem) print() {
	log.Printf("Item: \n")
	log.Printf("-------\n")
	log.Printf("         id: %s\n", this.id)
	log.Printf("        cnt: %d\n", this.cnt)
	log.Printf("   discount: $%d\n", this.discount)
	log.Printf("dis percent: %d%%\n", int(this.discount_percent*100))
	log.Printf("      total: $%d\n", this.total)
}

// store
type Store struct {
	g_list []Product
	g_map map[string]Product
}

// global singleton
var g_store *Store
func get_store() Store {
	if (g_store == nil) {
		g_store = &Store{}
		g_store.g_list = make([]Product, 0, 10)
		g_store.g_map = make(map[string]Product)
	}
	return *g_store
}

// get product list
func (this Store) list_products() (_list []Product, _e error) {
	return this.g_list, nil
}

// get product by id
func (this Store) get_product(_id string) (prod Product,ok bool) {
	prod,ok = this.g_map[_id]
	return
}

// create product
func (this Store) create_product(_id string, _name string, _price int) Product {
	// check id
	id := _id
	// check dup id
	_,ok := this.g_map[_id]	
	if ok { // id exists, or id<=0, XXX
		id = strconv.Itoa(len(this.g_list)) // allocate an id  XXX
	}

	// check name: empty, len>256, encoding/escape...
	// XXX

	// check name: duplicate
	// XXX

	p := Product{id, _name, _price}
	this.g_list = append(this.g_list, p)
	this.g_map[id] = p
	return p
}

// new order
func (this Store) proc_order (_request Order) (_response Order, _e error) {
	order := &_request
	for i,_:= range order.items {
		item := &order.items[i]
		prod,ok := this.get_product(item.id)
		if !ok {
			panic("product does not exist")
		}
		item.total = prod.price * item.cnt
		order.total += item.total
	}

	// invalide discount
	if (order.discount < 0) {
		panic("discount < 0") // XXX
	}
	if (order.discount > order.total) { // == means gift ?
		panic("discount > total") // XXX
	}
	if (order.total == 0) {  // ???
		panic("toal: 0") // XXX
	}

	var unit float64
	curr_discount := 0
	unit = float64(order.discount) / float64(order.total)
	// log.Printf("unit: %.3f. (%d/%d)\n", unit, order.discount, order.total)

	for i,_:= range order.items {
		item := &order.items[i]

		if i+1<len(order.items) {
			item.discount = int(unit*float64(item.total))
			item.print()
		} else { // last item: fix the float/multi error
	 		item.discount = order.discount - curr_discount
			item.print()
		}
		item.discount_percent = float32(item.discount) / (float32)(order.discount)

		curr_discount += item.discount
	}

	// check discount
	if (curr_discount != order.discount) {
		panic("wrong curr_discount")
	}

	order.total -= order.discount
	return *order, nil
}

func main() {
}
