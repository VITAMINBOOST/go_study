package main

import (
	"fmt"
	"math"
	"unsafe"
)

func HelloWorld() {
	// go build <filename>
	fmt.Println("Hello, world!")
}

func VarTypes() {
	var i int
	var s string

	var uage int = 10
	var name string = "Maria"

	// var address // it makes compile error

	fmt.Println(i, s, uage, name)

	tage := 20
	tname := "Marrrrria"

	fmt.Println(tage, tname)

	var x, y int = 30, 40
	var ttage, ttname = 10, "mariaaaaaa"

	fmt.Println(x, y, ttage, ttname)

	a, b, c, d := 1, 3.4, "HELLO, WORLD", false

	fmt.Println(a, b, c, d)

	var tx, ty int
	var tttage int
	tx, ty, tttage = 10, 20, 5

	fmt.Println(tx, ty, tttage)

	aa := 1
	bb := 2
	_ = bb // it prevent to make error
	fmt.Println(aa)

	// 자료형							설명																												범위
	// uint8						부호 없는(unsigned) 8비트, 1바이트 정수														0 ~ 255
	// uint16						부호 없는 16비트, 2바이트 정수																		0 ~ 65535
	// uint32						부호 없는 32비트, 4바이트 정수																		0 ~ 4294967295
	// uint64						부호 없는 64비트, 8바이트 정수																		0 ~ 18446744073709551615
	// int8							부호 있는(signed) 8비트, 1바이트 정수															-128 ~ 127
	// int16						부호 있는 16비트, 2바이트 정수																		-32768 ~ 32767
	// int32						부호 있는 32비트, 4바이트 정수																		-2147483648 ~ 2147483647
	// int64						부호 있는 64비트, 8바이트 정수																		-9223372036854775808 ~ 9223372036854775807
	// uint							32비트 시스템에서는 uint32, 64비트 시스템에서는 uint64
	// int							32비트 시스템에서는 int32, 64비트 시스템에서는 int64
	// uintptr					uint와 크기가 동일하며 포인터를 저장할 때 사용
	// float32					IEEE-754 32비트 단정밀도 부동소수점, 7자리 정밀도 보장
	// float64					IEEE-754 64비트 배정밀도 부동소수점, 15자리 정밀도 보장
	// complex64				float32 크기의 실수부와 허수부로 된 복소수
	// complex128				float64 크기의 실수부와 허수부로 된 복소수
	// byte							uint8과 크기가 동일, 바이트 단위로 저장할 때 사용
	// rune							int32와 크기가 동일, 유니코드 문자 코드(Code point)를 저장할 때 사용
}

func RealNumber() {

	// 소수점 사용
	var num1 float32 = 0.1
	var num2 float32 = .35
	var num3 float32 = 132.73287

	// 지수 표기법
	var num4 float32 = 1e7
	var num5 float64 = .12345E+2
	var num6 float64 = 5.32521e-10

	fmt.Println(num1, num2, num3, num4, num5, num6)

	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	const epsilon = 1e-14                   // Go 언어 머신 입실론
	fmt.Println(math.Abs(a-9.0) <= epsilon) // true: 연산한 값과 비교할 값의 차이를 구한 뒤
}

func complexNumber() {
	var num1 complex64 = 1 + 2i           // 실수부 1, 허수부 2
	var num2 complex64 = 4.2342 + 2.3527i // 실수부 소수점 사용 4.2342,
	// 허수부 지수 표기법 2.3527
	var num3 complex64 = 1.e+3i       // 실수부 지수 표기법 1.e, 허수부 3
	var num4 complex64 = 7.27151e-10i // 실수부 없음, 허수부 지수 표기법 7.27151e-10

	var num5 complex128 = 1 + 2i                   // 실수부 1, 허수부 2
	var num6 complex128 = 5.32521e-10 + .12345E+2i // 실수부 지수 표기법 5.32521e-10,
	// 허수부 지수 표기법 .12345E+2

	var r1 float32 = real(num1) // num1의 실수부 1
	var i1 float32 = imag(num1) // num1의 허수부 2

	var r2 float64 = real(num5) // mum5의 실수부 1
	var i2 float64 = imag(num5) // num5의 허수부 2

	var num11 complex64 = complex(1, 2)                    // 실수부 1, 허수부 2
	var num22 complex128 = complex(5.32521e-10, .12345E+2) // 실수부 지수 표기법 5.32521e-10,
	// 허수부 지수 표기법 .12345E+2

	fmt.Println(num1, num2, num3, num4, num5, num6, r1, i1, r2, i2, num11, num22)
}

func ByteWorld() {
	var num1 byte = 10   // 10진수 저장
	var num2 byte = 0x32 // 16진수 저장
	var num3 byte = 'a'  // 문자 저장

	fmt.Println(num1, num2, num3)

	// var num11 byte = "a"  // 컴파일 에러
	// var num22 byte = 'ab' // 컴파일 에러
	// var num33 byte = "ab" // 컴파일 에러

	// fmt.Println(num11, num22, num33)
}

func RuneWorld() {
	var r1 rune = '한'
	var r2 rune = '\ud55c'     // 한
	var r3 rune = '\U0000d55c' // 한

	fmt.Println(r1, r2, r3)

	// var r1 rune = "한"   // 컴파일 에러
	// var r2 rune = '한글' // 컴파일 에러
	// var r3 rune = "한글" // 컴파일 에러

	// 전체 유니코드 목록: http://ko.wikipedia.org/wiki/유니코드_블록
	// 한글 U+AC00~U+D7AF
	// http://ko.wikipedia.org/wiki/유니코드_A000~AFFF
	// http://ko.wikipedia.org/wiki/유니코드_B000~BFFF
	// http://ko.wikipedia.org/wiki/유니코드_C000~CFFF
	// http://ko.wikipedia.org/wiki/유니코드_D000~DFFF
}

func CalcWorld() {
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)  // 5
	fmt.Println(num1 - num2)  // 1
	fmt.Println(num1 * num2)  // 6
	fmt.Println(num1 / num2)  // 1
	fmt.Println(num1 % num2)  // 1
	fmt.Println(num1 << num2) // 12
	fmt.Println(num1 >> num2) // 0
	fmt.Println(^num1)        // 252: 비트 반전 연산자
}

func FlowTest() {

	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64

	// overflow compile error
	// var num1 uint8 = math.MaxUint8 + 1
	// var num2 uint16 = math.MaxUint16 + 1
	// var num3 uint32 = math.MaxUint32 + 1
	// var num4 uint64 = math.MaxUint64 + 1

	fmt.Println(num1) // 255
	fmt.Println(num2) // 65535
	fmt.Println(num3) // 4294967295
	fmt.Println(num4) // 18446744073709551615

	// underflow compile error
	// var num1 uint8 = 0 - 1
	// var num2 uint16 = 0 - 1
	// var num3 uint32 = 0 - 1
	// var num4 uint64 = 0 - 1
}

func CheckByteSize() {
	var num1 int8 = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	fmt.Println(unsafe.Sizeof(num1)) // 1
	fmt.Println(unsafe.Sizeof(num2)) // 2
	fmt.Println(unsafe.Sizeof(num3)) // 4
	fmt.Println(unsafe.Sizeof(num4)) // 8

	// http://pyrasis.com/book/GoForTheReallyImpatient/Unit08/08
}

func main() {
	CheckByteSize()
}
