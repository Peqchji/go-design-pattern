package vending_machine

import (
	"errors"
)

var (
	ErrNoItem        = errors.New("no item selected")
	ErrOutOfStock    = errors.New("item out of stock")
	ErrInsufficient  = errors.New("insufficient money")
	ErrAlreadyHas    = errors.New("already has selection")
	ErrInvalidAction = errors.New("invalid action")
)

// VendingMachine is the context that maintains the current
type VendingMachine struct {
	idleState         State
	hasSelectionState State
	outOfStockState   State
	itemRequested     State
	dispensedState    State

	currentState State
	itemCount    int
	itemPrice    int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	v.idleState = &IdleState{vendingMachine: v}
	v.outOfStockState = &OutOfStockState{vendingMachine: v}
	v.hasSelectionState = &HasSelectionState{vendingMachine: v}
	v.itemRequested = &ItemRequestedState{vendingMachine: v}

	if itemCount > 0 {
		v.currentState = v.idleState
	} else {
		v.currentState = v.outOfStockState
	}

	return v
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) SelectItem() error {
	return v.currentState.SelectItem()
}

func (v *VendingMachine) InsertMoney(amount int) error {
	return v.currentState.InsertMoney(amount)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}
