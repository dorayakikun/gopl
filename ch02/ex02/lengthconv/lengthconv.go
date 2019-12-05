package lengthconv

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
