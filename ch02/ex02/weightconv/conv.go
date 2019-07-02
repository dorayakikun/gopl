package weightconv

func PToK(p Pound) Kilogram { return Kilogram(p / 2.2046) }
func KToP(k Kilogram) Pound { return Pound(k * 2.2046) }