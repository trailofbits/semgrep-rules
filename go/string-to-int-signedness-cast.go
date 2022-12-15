package main

import (
	"fmt"
	"strconv"
)

const BigNum1 = "-4394967295"
const BigNum2 = "4294967294"
const BigNum3 = "-32766"
const BigNum4 = "18446744073709551615"
const BigNum5 = "253"

func main(){
	conv_atoi1(BigNum1)
	conv_atoi2(BigNum1)
	conv_atoi3(BigNum1)
	conv_atoi4(BigNum1)
	conv_atoi5(BigNum1)
	conv_atoi6(BigNum1)
	conv_atoi7(BigNum1)
	conv_atoi8(BigNum1)

	conv_parse_int1(BigNum1)
	conv_parse_int2(BigNum1)
	conv_parse_int3(BigNum1)
	conv_parse_int4(BigNum1)
	conv_parse_int5(BigNum1)
	conv_parse_int6(BigNum1)
	conv_parse_int7(BigNum1)
	conv_parse_int8(BigNum1)

	conv_parse_uint1(BigNum4)
	conv_parse_uint2(BigNum4)
	conv_parse_uint3(BigNum4)
	conv_parse_uint4(BigNum4)
	conv_parse_uint5(BigNum4)
	conv_parse_uint6(BigNum4)
	conv_parse_uint7(BigNum4)
	conv_parse_uint8(BigNum4)

	conv32_parse_uint(BigNum2)
	conv16_parse_int(BigNum3)
	conv8_parse_uint8(BigNum5)
}

/* ATOI */
func conv_atoi1(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv_atoi1 result: ", numStr, "->", num)
}

func conv_atoi2(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv_atoi2 result: ", numStr, "->", num)
}

func conv_atoi3(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv_atoi3 result: ", numStr, "->", num)
}

func conv_atoi4(numStr string){
	// it is ok on 32-bit systems
	// todook: string-to-int-signedness-cast
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := uint64(res)
	fmt.Println("conv_atoi4 result: ", numStr, "->", num)
}

func conv_atoi5(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int8(res)
	fmt.Println("conv_atoi5 result: ", numStr, "->", num)
}

func conv_atoi6(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int16(res)
	fmt.Println("conv_atoi6 result: ", numStr, "->", num)
}

func conv_atoi7(numStr string){
	// it is ok on 32-bit systems
	// todook: string-to-int-signedness-cast
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int32(res)
	fmt.Println("conv_atoi7 result: ", numStr, "->", num)
}

func conv_atoi8(numStr string){
	// it is not ok on 32-bit systems
	// todoruleid: string-to-int-signedness-cast
	// ok: string-to-int-signedness-cast
	res, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	num := int64(res)
	fmt.Println("conv_atoi8 result: ", numStr, "->", num)
}

/* PARSE_INT */
func conv_parse_int1(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv_parse_int1 result: ", numStr, "->", num)
}

func conv_parse_int2(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv_parse_int2 result: ", numStr, "->", num)
}

func conv_parse_int3(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv_parse_int3 result: ", numStr, "->", num)
}

func conv_parse_int4(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint64(res)
	fmt.Println("conv_parse_int4 result: ", numStr, "->", num)
}

func conv_parse_int5(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int8(res)
	fmt.Println("conv_parse_int5 result: ", numStr, "->", num)
}

func conv_parse_int6(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int16(res)
	fmt.Println("conv_parse_int6 result: ", numStr, "->", num)
}

func conv_parse_int7(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int32(res)
	fmt.Println("conv_parse_int7 result: ", numStr, "->", num)
}

func conv_parse_int8(numStr string){
	// ok: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int64(res)
	fmt.Println("conv_parse_int8 result: ", numStr, "->", num)
}

/* PARSE_UINT */
func conv_parse_uint1(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint8(res)
	fmt.Println("conv_parse_uint1 result: ", numStr, "->", num)
}

func conv_parse_uint2(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv_parse_uint2 result: ", numStr, "->", num)
}

func conv_parse_uint3(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint32(res)
	fmt.Println("conv_parse_uint3 result: ", numStr, "->", num)
}

func conv_parse_uint4(numStr string){
	// ok: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := uint64(res)
	fmt.Println("conv_parse_uint4 result: ", numStr, "->", num)
}

func conv_parse_uint5(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int8(res)
	fmt.Println("conv_parse_uint5 result: ", numStr, "->", num)
}

func conv_parse_uint6(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int16(res)
	fmt.Println("conv_parse_uint6 result: ", numStr, "->", num)
}

func conv_parse_uint7(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int32(res)
	fmt.Println("conv_parse_uint7 result: ", numStr, "->", num)
}

func conv_parse_uint8(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 64)
	if err != nil {
		return
	}
	num := int64(res)
	fmt.Println("conv_parse_uint8 result: ", numStr, "->", num)
}


/* 32 bits*/
func conv32_parse_uint(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 32)
	if err != nil {
		return
	}
	num := int32(res)
	fmt.Println("conv32_parse_uint result: ", numStr, "->", num)
}


/* 16 bits*/
func conv16_parse_int(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseInt(numStr, 0, 16)
	if err != nil {
		return
	}
	num := uint16(res)
	fmt.Println("conv16_parse_int result: ", numStr, "->", num)
}



/* 8 bits*/
func conv8_parse_uint8(numStr string){
	// ruleid: string-to-int-signedness-cast
	res, err := strconv.ParseUint(numStr, 0, 8)
	if err != nil {
		return
	}
	num := int8(res)
	fmt.Println("conv8_parse_uint8 result: ", numStr, "->", num)
}

