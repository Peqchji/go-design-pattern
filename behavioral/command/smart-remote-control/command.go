package smart_remote_control

type LightOnCommand struct {
	light *Light
	lastState bool
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{
		light: light,
		lastState: light.IsOn,
	}
}

func (l *LightOnCommand) Name() string {
	return "LightOnCommand"
}

func (l *LightOnCommand) Execute() {
	l.lastState = l.light.IsOn
	l.light.On()
}

func (l *LightOnCommand) Undo() {
	if !l.lastState {
		l.light.Off()
	}
}

type LightOffCommand struct {
	light *Light
	lastState bool
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{
		light: light,
		lastState: light.IsOn,
	}
}

func (l *LightOffCommand) Name() string {
	return "LightOffCommand"
}

func (l *LightOffCommand) Execute() {
	l.lastState = l.light.IsOn
	l.light.Off()
}

func (l *LightOffCommand) Undo() {
	if l.lastState {
		l.light.On()
	}
}

type ACSetTempCommand struct {
	ac *AirConditioner
	lastTemp int
	temp int
}

func NewACSetTempCommand(ac *AirConditioner, temp int) *ACSetTempCommand {
	return &ACSetTempCommand{
		ac: ac,
		temp: temp,
	}
}

func (a *ACSetTempCommand) Name() string {
	return "ACSetTempCommand"
}

func (a *ACSetTempCommand) Execute() {
	a.lastTemp = a.ac.Temp
	a.ac.SetTemp(a.temp)
}

func (a *ACSetTempCommand) Undo() {
	a.ac.SetTemp(a.lastTemp)
}


type PartyModeCommand struct {
	light *Light
	ac *AirConditioner

	lightLastState bool
	acLastTemp int
}

func NewPartyModeCommand(light *Light, ac *AirConditioner) *PartyModeCommand {
	return &PartyModeCommand{
		light: light,
		ac: ac,
	}
}

func (p *PartyModeCommand) Name() string {
	return "PartyModeCommand"
}

func (p *PartyModeCommand) Execute() {
	p.lightLastState = p.light.IsOn
	p.acLastTemp = p.ac.Temp

	p.light.On()
	p.ac.SetTemp(18)
}

func (p *PartyModeCommand) Undo() {
	if !p.lightLastState {
		p.light.Off()
	}

	p.ac.SetTemp(p.acLastTemp)
}