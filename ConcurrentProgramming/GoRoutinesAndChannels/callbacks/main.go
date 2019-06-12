package main

import (
	"time"
	"fmt"
)

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234

	time.Sleep(5 * time.Second)
	callback <- po
}

func main() {

	po := new(PurchaseOrder)
	po.Value = 42.27

	ch := make(chan *PurchaseOrder)

	go SavePO(po, ch)

	newPo := <-ch

	fmt.Printf("PO : %v \n", newPo)
}
