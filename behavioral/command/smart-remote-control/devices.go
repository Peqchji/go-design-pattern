package smart_remote_control


import "fmt"

type Light struct {
	IsOn bool
}

func NewLight() *Light {
	return &Light{
		IsOn: false,
	}
}

func (l *Light) On() {
	l.IsOn = true
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	l.IsOn = false
	fmt.Println("Light is OFF")
}

type AirConditioner struct {
	Temp int
}

func NewAirConditioner() *AirConditioner {
	return &AirConditioner{
		Temp: 20,
	}
}

func (ac *AirConditioner) SetTemp(t int) {
	ac.Temp = t
	fmt.Printf("AC temperature set to %d\n", t)
}

func (ac *AirConditioner) Off() {
	fmt.Println("AC is OFF")
}
