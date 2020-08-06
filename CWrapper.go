package main

import "C"
import (
	"./GoChest"
	"encoding/binary"
	"math"
	"unsafe"
)

//export FindChangepoints
func FindChangepoints(inpSequence unsafe.Pointer, inpLength C.int, inpMinimumDistance C.float) uintptr {

	byteSequence := C.GoBytes(inpSequence, inpLength)
	length := int(inpLength)
	minimumDistance := float64(inpMinimumDistance)

	goSequence := make([]float64, length/8)
	for i := 0; i < length; i += 8 {
		bits := binary.LittleEndian.Uint64(byteSequence[i : i+8])
		goSequence[i/8] = math.Float64frombits(bits)
	}

	changepoints := GoChest.FindChangePoints(goSequence, minimumDistance)

	// First value of the output is the length of the array
	output := make([]byte, (len(changepoints)+1)*8)

	bytes := make([]byte, 8)
	for index, value := range append([]int{len(changepoints)}, changepoints...) {
		binary.LittleEndian.PutUint32(bytes, uint32(value))
		copy(output[index*8:(index+1)*8], bytes)
	}

	return uintptr(unsafe.Pointer(&output[0]))
}

func main() {

}
