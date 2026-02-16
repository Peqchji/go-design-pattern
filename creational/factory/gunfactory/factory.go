package gunfactory

import "errors"

type GunType = string

const (
	GunTypeAK47  GunType = "AK47"
	GunTypeM16   GunType = "M16"
	GunTypeG36   GunType = "G36"
	GunTypeGlock GunType = "Glock"
)

var ErrInvalidGunType = errors.New("Invalid Gun Type")

type Gun interface {
	Name() string
}

type GunFactory struct {
	guns map[GunType]Gun
}

func NewGunFactory() *GunFactory {
	return &GunFactory{
		guns: map[GunType]Gun{
			GunTypeAK47:  &AK47{},
			GunTypeM16:   &M16{},
			GunTypeG36:   &G36{},
			GunTypeGlock: &Glock{},
		},
	}
}

func (gf *GunFactory) CreateGun(gunType GunType) (Gun, error) {
	gun, ok := gf.guns[gunType]
	if !ok {
		return nil, ErrInvalidGunType
	}

	return gun, nil
}

type AK47 struct {}

func (ak *AK47) Name() string {
	return GunTypeAK47
}

type M16 struct {}

func (m *M16) Name() string {
	return GunTypeM16
}

type G36 struct {}

func (g *G36) Name() string {
	return GunTypeG36
}

type Glock struct {}

func (g *Glock) Name() string {
	return GunTypeGlock
}


