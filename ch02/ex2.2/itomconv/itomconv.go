package itomconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Foot float64
type Meter float64
type Kilogram float64
type Pound float64

const (
	PoundsInKilogram = 2.20462
	FootsInMeter     = 3.28084
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (ft Foot) String() string      { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (kg Kilogram) String() string  { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string     { return fmt.Sprintf("%glb", lb) }
