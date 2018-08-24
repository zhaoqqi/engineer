package car

import "fmt"

type Speed float64

const (
	MPH Speed = 1
	KPH Speed = 1.60934
)

type Color string

const (
    BlueColor Color = "blue"
    GreenColor      = "green"
    RedColor        = "red"
)

type Wheels string

const (
    SportsWheels Wheels = "sports"
    SteelWheels         = "steel"
)

type Builder interface {
    Color(Color) Builder
    Wheels(Wheels) Builder
    TopSpeed(Speed) Builder
    Build() Car
}

type builder struct {
    color Color
    wheels Wheels
    speed Speed
}

func (b *builder) Color(color Color) Builder {
    b.color  = color
    return b
}

func (b *builder) Wheels(wheels Wheels) Builder {
    b.wheels = wheels
    return b
}

func (b *builder) TopSpeed(speed Speed) Builder {
    b.speed = speed
    return b
}

func (b *builder) Build() Car {
    return &car {
            color : b.color,
            wheels : b.wheels,
            speed : b.speed,
    }
}

func NewBuilder() builder {
    return &builder{}
}

type Car interface {
    Drive() error
    Stop() error
}

type car struct {
    color Color
    wheels Wheels
    speed Speed
}


func (car *car) Drive() error {
    return "Driving a " + car.color + " car at speed " + car.speed + " with " + car.wheels + "wheels."
}

func (car *car) Stop() error {
    return "Stopping a " + car.color + " car at speed " + car.speed + " with " + car.wheels + "wheels."
}

func test string {
    return "for test"
}


func main() {
    builder := NewBuilder()
    car := builder.Color(GreenColor).Wheels(SportsWheels).TopSpeed(240*MPH).Build()
    fmt.Println(car.Drive())
    fmt.Println(car.Stop())
}
