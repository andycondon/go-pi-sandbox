package arduino

import (
	"github.com/andycondon/go-pi-sandbox/pkg/ir"
	"periph.io/x/periph/conn/i2c"
)

const (
	StatusAddress=  0x10
)

type Status struct {
	IR ir.Sensor
}

type Arduino struct {
	dev *i2c.Dev
}

func New(bus i2c.Bus) *Arduino {
	return &Arduino{
		dev: &i2c.Dev{Addr: 0x1a, Bus: bus},
	}
}

func (a *Arduino) GetStatus() (Status, error)  {
	read := make([]byte, 3)
	if err := a.dev.Tx([]byte{StatusAddress}, read); err != nil {
		return Status{}, err
	}
	irSensor, err := ir.FromBytes(read)
	if err != nil {
		return Status{}, err
	}
	return Status{
		IR: irSensor,
	}, nil
}
