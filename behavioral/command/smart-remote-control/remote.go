package smart_remote_control

import "errors"

var ErrSlotNotFound = errors.New("slot not found")

type Command interface {
	Name() string
	Execute()
	Undo()
}

type RemoteControl struct {
	slots map[string]Command
}

func NewRemoteControl(commands ...Command) *RemoteControl {
	commandsMap := make(map[string]Command)
	for _, cmd := range commands {
		commandsMap[cmd.Name()] = cmd
	}

	return &RemoteControl{
		slots: commandsMap,
	}
}

func (r *RemoteControl) PressButton(slot string) error {
	if cmd, exists := r.slots[slot]; exists {
		cmd.Execute()
		return nil
	}

	return ErrSlotNotFound
}

func (r *RemoteControl) PressUndoButton(slot string) error {
	if cmd, exists := r.slots[slot]; exists {
		cmd.Undo()
		return nil
	}

	return ErrSlotNotFound
}
