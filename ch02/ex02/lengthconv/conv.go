package lengthconv

func FToM(f Feet) Meters { return Meters(f / 3.281) }
func MToF(m Meters) Feet { return Feet(m * 3.281) }
