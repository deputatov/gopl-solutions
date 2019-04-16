package itomconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func PToK(p Pound) Kilogram     { return Kilogram(p / PoundsInKilogram) }
func KToP(k Kilogram) Pound     { return Pound(k * PoundsInKilogram) }
func FToM(f Foot) Meter         { return Meter(f / FootsInMeter) }
func MToF(m Meter) Foot         { return Foot(m * FootsInMeter) }
