package vending_machine

type State interface {
	AddItem(count int) error
	SelectItem() error
	InsertMoney(amount int) error
	DispenseItem() error
}


// --- IdleState ---

type IdleState struct {
	vendingMachine *VendingMachine
}

func (i *IdleState) AddItem(count int) error {
	i.vendingMachine.itemCount += count
	return nil
}

func (i *IdleState) SelectItem() error {
	i.vendingMachine.SetState(i.vendingMachine.hasSelectionState)
	return nil
}

func (i *IdleState) InsertMoney(amount int) error {
	return ErrNoItem
}

func (i *IdleState) DispenseItem() error {
	return ErrNoItem
}

// --- HasSelectionState ---

type HasSelectionState struct {
	vendingMachine *VendingMachine
}

func (s *HasSelectionState) AddItem(count int) error {
	return ErrAlreadyHas
}

func (s *HasSelectionState) SelectItem() error {
	return ErrAlreadyHas
}

func (s *HasSelectionState) InsertMoney(amount int) error {
	if amount < s.vendingMachine.itemPrice {
		return ErrInsufficient
	}

	s.vendingMachine.SetState(s.vendingMachine.itemRequested)
	return nil
}

func (s *HasSelectionState) DispenseItem() error {
	return ErrInsufficient
}

// --- OutOfStockState ---

type OutOfStockState struct {
	vendingMachine *VendingMachine
}

func (s *OutOfStockState) AddItem(count int) error {
	if count <= 0 {
		return nil
	}

	s.vendingMachine.itemCount += count
	s.vendingMachine.SetState(s.vendingMachine.idleState)
	return nil
}

func (s *OutOfStockState) SelectItem() error {
	return ErrOutOfStock
}

func (s *OutOfStockState) InsertMoney(amount int) error {
	return ErrOutOfStock
}

func (s *OutOfStockState) DispenseItem() error {
	return ErrOutOfStock
}

// --- ItemRequestedState ---

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (s *ItemRequestedState) AddItem(count int) error {
	return ErrInvalidAction
}

func (s *ItemRequestedState) SelectItem() error {
	return ErrAlreadyHas
}

func (s *ItemRequestedState) InsertMoney(amount int) error {
	return ErrAlreadyHas
}

func (s *ItemRequestedState) DispenseItem() error {
	if s.vendingMachine.itemCount <= 0 {
		s.vendingMachine.SetState(s.vendingMachine.outOfStockState)
		return ErrOutOfStock
	}

	s.vendingMachine.itemCount -= 1

	if s.vendingMachine.itemCount <= 0 {
		s.vendingMachine.SetState(s.vendingMachine.outOfStockState)
	} else {
		s.vendingMachine.SetState(s.vendingMachine.idleState)
	}

	return nil
}
