package vending_machine_test

import (
	vending "design_pattern/behavioral/state/vending-machine"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVendingMachine_TableDriven(t *testing.T) {
	type step struct {
		name        string
		action      func(vm *vending.VendingMachine) error
		expectedErr error
	}

	tests := []struct {
		name         string
		initialItems int
		price        int
		steps        []step
	}{
		{
			name:         "Successful purchase flow",
			initialItems: 1,
			price:        10,
			steps: []step{
				{
					name:   "Select item",
					action: func(vm *vending.VendingMachine) error { return vm.SelectItem() },
				},
				{
					name:   "Insert sufficient money",
					action: func(vm *vending.VendingMachine) error { return vm.InsertMoney(10) },
				},
				{
					name:   "Dispense item",
					action: func(vm *vending.VendingMachine) error { return vm.DispenseItem() },
				},
			},
		},
		{
			name:         "Out of stock error",
			initialItems: 0,
			price:        10,
			steps: []step{
				{
					name:        "Select item when empty",
					action:      func(vm *vending.VendingMachine) error { return vm.SelectItem() },
					expectedErr: vending.ErrOutOfStock,
				},
			},
		},
		{
			name:         "Insufficient money error",
			initialItems: 1,
			price:        10,
			steps: []step{
				{
					name:   "Select item",
					action: func(vm *vending.VendingMachine) error { return vm.SelectItem() },
				},
				{
					name:        "Insert insufficient money",
					action:      func(vm *vending.VendingMachine) error { return vm.InsertMoney(5) },
					expectedErr: vending.ErrInsufficient,
				},
			},
		},
		{
			name:         "Restocking flow",
			initialItems: 0,
			price:        10,
			steps: []step{
				{
					name:        "Select when empty",
					action:      func(vm *vending.VendingMachine) error { return vm.SelectItem() },
					expectedErr: vending.ErrOutOfStock,
				},
				{
					name:   "Add item",
					action: func(vm *vending.VendingMachine) error { return vm.AddItem(1) },
				},
				{
					name:   "Select now should work",
					action: func(vm *vending.VendingMachine) error { return vm.SelectItem() },
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := vending.NewVendingMachine(tt.initialItems, tt.price)

			for _, s := range tt.steps {
				t.Run(s.name, func(t *testing.T) {
					err := s.action(vm)
					if s.expectedErr != nil {
						assert.ErrorIs(t, err, s.expectedErr)
					} else {
						assert.NoError(t, err)
					}
				})
			}
		})
	}
}
