package factory

import (
	"design_pattern/creational/factory/gunfactory"
	"design_pattern/pkg/result"
	"testing"

	"github.com/stretchr/testify/assert"
)

type GunFactoryTestCase struct {
	name     string
	gunType  gunfactory.GunType
	expected result.Result[gunfactory.Gun]
}

func TestGunFactory(t *testing.T) {
	factory := gunfactory.NewGunFactory()
	table := []GunFactoryTestCase{
		{
			name: "Should create AK47",
			gunType: gunfactory.GunTypeAK47,
			expected: result.Result[gunfactory.Gun]{
				Result: &gunfactory.AK47{},
				Error:  nil,
			},
		},
		{
			name: "Should create M16",
			gunType: gunfactory.GunTypeM16,
			expected: result.Result[gunfactory.Gun]{
				Result: &gunfactory.M16{},
				Error:  nil,
			},
		},
		{
			name: "Should create G36",
			gunType: gunfactory.GunTypeG36,
			expected: result.Result[gunfactory.Gun]{
				Result: &gunfactory.G36{},
				Error:  nil,
			},
		},
		{
			name: "Should create Glock",
			gunType: gunfactory.GunTypeGlock,
			expected: result.Result[gunfactory.Gun]{
				Result: &gunfactory.Glock{},
				Error:  nil,
			},
		},
		{
			name: "Should return error for invalid gun type",
			gunType: gunfactory.GunType("Invalid Gun Type"),
			expected: result.Result[gunfactory.Gun]{
				Result: nil,
				Error:  gunfactory.ErrInvalidGunType,
			},
		},
	}

	for _, tc := range table {
		t.Run(tc.name, func(t *testing.T) {
			gun, err := factory.CreateGun(tc.gunType)

			if err != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expected.Error.Error())
			} else {
			assert.NotNil(t, gun)
			assert.EqualValues(
				t,
				tc.expected.Result,
				gun,
			)
			}
		})
	}
}
