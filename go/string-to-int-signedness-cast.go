package main

import (
	"fmt"
	"strconv"
)

const BigNum1 = "-4394967295"
const BigNum2 = "2000"
const BigNum3 = "200000"
const BigNum4 = "20000000000"

func main(){
	conv1(BigNum1)
	conv2(BigNum1)
	conv3(BigNum1)
	conv4(BigNum1)
	conv5(BigNum1)
	conv6(BigNum1)
	conv7(BigNum1)
	conv8(BigNum1)
	conv9(BigNum1)
	conv10(BigNum1)
	conv11(BigNum1)
	conv12(BigNum1)
	conv13(BigNum1)
	conv14(BigNum2)
	conv15(BigNum3)
	conv16(BigNum4)
	conv17(BigNum3)
	conv18(BigNum4)
	conv19(BigNum4)
}

func conv1(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv1 result: ", num)
}

func conv2(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv2 result: ", num)
}

func conv3(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv3 result: ", num)
}

func conv4(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint64(res)
	fmt.Println("conv4 result: ", num)
}

func conv5(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int8(res)
	fmt.Println("conv5 result: ", num)
}

func conv6(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int16(res)
	fmt.Println("conv6 result: ", num)
}

func conv7(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int32(res)
	fmt.Println("conv7 result: ", num)
}

func conv8(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv8 result: ", num)
}

func conv9(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv9 result: ", num)
}

func conv10(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv10 result: ", num)
}

func conv11(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv11 result: ", num)
}

func conv12(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv12 result: ", num)
}

func conv13(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv13 result: ", num)
}

func conv14(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv14 result: ", num)
}

func conv15(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv15 result: ", num)
}

func conv16(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv16 result: ", num)
}

func conv17(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv17 result: ", num)
}

func conv18(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv18 result: ", num)
}

func conv19(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv19 result: ", num)
}


