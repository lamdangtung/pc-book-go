package sample

import "github.com/lamdangtung/pc-book-go/pb"

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backit: randomBool(),
	}
}

func NewCPU() *pb.CPU{
	brand := randomCPUBrand();
	name := randomCPUName(brand);
	numberCores := randomInt(4,8);
	numberThreads := randomInt(numberCores,12);
	minGHZ := randomFloat64(2.0,3.5);
	maxGHZ := randomFloat64(minGHZ, 5.0);
	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz: minGHZ,
		MaxGhz: maxGHZ,
	};
	return cpu; 
}