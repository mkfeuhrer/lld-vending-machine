package main

import (
	"fmt"
	"log/slog"
	"vendingmachine/app"
	"vendingmachine/balance"
	"vendingmachine/inventory"
	order2 "vendingmachine/order"
	"vendingmachine/payment"
	"vendingmachine/payment/enums"
	"vendingmachine/product"
)

func main() {
	// Create empty inventory
	inv := inventory.NewInventory([]inventory.Stock{})

	// create payment instruments
	paymentInstruments := map[enums.PaymentInstrument]int{
		enums.Coin:  100,
		enums.Notes: 50,
	}

	// Create balance
	bal := balance.NewBalance(paymentInstruments)

	// Create vending machine app
	vendingMachine := app.NewApp(inv, bal)

	// Add payment instruments to the balance
	vendingMachine.Balance.DisplayBalance()
	fmt.Println("Adding balance")
	vendingMachine.AddBalance(enums.Coin, 100)
	vendingMachine.AddBalance(enums.Notes, 50)
	vendingMachine.Balance.DisplayBalance()

	// Add products to the inventory
	fmt.Println("Adding products to inventory")
	vendingMachine.AddProduct(&product.Product{ID: "1", Name: "Coke", Price: 10}, 5)
	vendingMachine.AddProduct(&product.Product{ID: "2", Name: "Pepsi", Price: 10}, 15)
	vendingMachine.AddProduct(&product.Product{ID: "3", Name: "Juice", Price: 30}, 5)
	vendingMachine.AddProduct(&product.Product{ID: "4", Name: "Soda", Price: 15}, 10)
	vendingMachine.Inventory.DisplayInventory()

	fmt.Println("Adding product quantity")
	vendingMachine.IncreaseProductQuantity(&product.Product{ID: "3", Name: "Juice", Price: 30, Quantity: 5})
	vendingMachine.Inventory.DisplayInventory()

	fmt.Println("Reducing product quantity")
	vendingMachine.DecreaseProductQuantity(&product.Product{ID: "3", Name: "Juice", Price: 30, Quantity: 5})
	vendingMachine.Inventory.DisplayInventory()

	// Create order
	fmt.Println("creating order")
	order := order2.NewOrder(
		"1",
		[]product.Product{
			{ID: "3", Name: "Juice", Price: 30, Quantity: 1},
		},
		&payment.Payment{
			TotalAmount: 0,
			Transactions: []payment.Transactions{
				{
					Amount:            30,
					PaymentInstrument: enums.Coin,
				},
			},
		},
	)
	err := vendingMachine.AddOrder(order)
	if err != nil {
		slog.Any("create order error", err)
		return
	}

	vendingMachine.Inventory.DisplayInventory()
	vendingMachine.Balance.DisplayBalance()

}
