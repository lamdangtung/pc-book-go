package sample

import (
	"math/rand"

	"github.com/lamdangtung/pc-book-go/pb"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3){
	case 1: 
		return pb.Keyboard_QWERTY;
	case 2: 
		return pb.Keyboard_AZERTY;
	default:
		return pb.Keyboard_QWERTZ;
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1;
}

func randomCPUBrand() string{
	return randomStringFromSet("Intel", "AMD");
}

func randomCPUName(brand string) string {
	if brand == "Intel"{
		return randomStringFromSet(
			"Intel Core Ultra 7 processor 165H",
			"Intel Core Ultra 5 processor 135U" ,
			"Intel Core i9-14900HX Processor", 
			"Intel Core i7-13700K Processor",
			"Intel Core i9-12900HX Processor" ,
			"Intel Core i7-10710U Processor" ,
			"Intel Core i5-11600KF Processor",
			"Intel Core i7-9850H Processor",
		);
	}else{
		return randomStringFromSet(
			"Ryzen 7 1800X",
			"Ryzen 7 2700X",
			"Ryzen 7 3700X",
			"Ryzen 9 5900X",
			"Ryzen 7 5800X3D",
			"Ryzen 9 7950X",
		);
	}
}

func randomStringFromSet(a ...string) string{
	n := len(a);
	if n == 0{
		return "";
	}
	return a[rand.Intn(n)];
}

func randomInt(min int, max int) int{
	return min + rand.Intn(max - min);
}

func randomFloat64(min float64, max float64) float64{
	return min + rand.Float64() * (max - min);
}