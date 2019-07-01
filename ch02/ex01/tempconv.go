package tempconv

import "fmt"

type Celsius float64
type Fathrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroKlvin             = AbsoluteZeroC
)

func (c Celsius) String() string     { return fmt.Sprintf("%g℃", c) }
func (f Fathrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string      { return fmt.Sprintf("%gK", k) }
