package smart_remote_control_test

import (
	smartDevice "design_pattern/behavioral/command/smart-remote-control"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteControl_TableDriven(t *testing.T) {
	light := smartDevice.NewLight()
	ac := smartDevice.NewAirConditioner()

	lightOn := smartDevice.NewLightOnCommand(light)
	lightOff := smartDevice.NewLightOffCommand(light)
	acSetTemp := smartDevice.NewACSetTempCommand(ac, 22)
	partyMode := smartDevice.NewPartyModeCommand(light, ac)

	remote := smartDevice.NewRemoteControl(
		lightOn,
		lightOff,
		acSetTemp,
		partyMode,
	)

	type testCase struct {
		name          string
		slot          string
		expectedError error
		validate      func(t *testing.T)
	}

	tests := []testCase{
		{
			name: "Light On",
			slot: lightOn.Name(),
			validate: func(t *testing.T) {
				assert.True(t, light.IsOn)
			},
		},
		{
			name: "Light Off",
			slot: lightOff.Name(),
			validate: func(t *testing.T) {
				assert.False(t, light.IsOn)
			},
		},
		{
			name: "AC Temp 22",
			slot: acSetTemp.Name(),
			validate: func(t *testing.T) {
				assert.Equal(t, 22, ac.Temp)
			},
		},
		{
			name: "Party Mode (Light On, AC 18)",
			slot: partyMode.Name(),
			validate: func(t *testing.T) {
				assert.True(t, light.IsOn)
				assert.Equal(t, 18, ac.Temp)
			},
		},
		{
			name:          "Unknown Slot Error",
			slot:          "Invalid",
			expectedError: smartDevice.ErrSlotNotFound,
			validate:      func(t *testing.T) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := remote.PressButton(tt.slot)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
			tt.validate(t)
		})
	}
}

func TestRemoteControl_Undo_TableDriven(t *testing.T) {
	light := smartDevice.NewLight()
	ac := smartDevice.NewAirConditioner()

	light.Off()
	ac.SetTemp(20)

	lightOn := smartDevice.NewLightOnCommand(light)
	acSetTemp := smartDevice.NewACSetTempCommand(ac, 25)
	partyMode := smartDevice.NewPartyModeCommand(light, ac)

	remote := smartDevice.NewRemoteControl(lightOn, acSetTemp, partyMode)

	tests := []struct {
		name     string
		slot     string
		execute  bool
		validate func(t *testing.T)
	}{
		{
			name:    "Undo Light On",
			slot:    lightOn.Name(),
			execute: true,
			validate: func(t *testing.T) {
				assert.False(t, light.IsOn)
			},
		},
		{
			name:    "Undo AC Temp (Back to 20)",
			slot:    acSetTemp.Name(),
			execute: true,
			validate: func(t *testing.T) {
				assert.Equal(t, 20, ac.Temp)
			},
		},
		{
			name:    "Undo Party Mode (Back to default)",
			slot:    partyMode.Name(),
			execute: true,
			validate: func(t *testing.T) {
				assert.False(t, light.IsOn)
				assert.Equal(t, 20, ac.Temp)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.execute {
				_ = remote.PressButton(tt.slot)
			}
			err := remote.PressUndoButton(tt.slot)
			assert.NoError(t, err)
			tt.validate(t)
		})
	}
}
