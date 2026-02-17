package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

var a template.HTML = "$EXP"

const jwtType = jwt.UnsafeAllowNoneSignatureType

var JwtType = jwt.UnsafeAllowNoneSignatureType

func handler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
}

func Ceil(x float64) int {
	return int(math.Ceil(x))
}

func Floor(x float64) int {
	return int(math.Floor(x))
}

func logRequest(r *http.Request) {
	println(r.Method)
}

func server() {
	http.HandleFunc("/hello", handler)
	return
	http.ListenAndServe(":8080", nil)
}

func trycheck(i int) {
	if i == 1 {
		return
	} else {
		return
	}
	fmt.Println("issue")
	return
}

func test() string {
	return "A"
}

func init() {

	fmt.Println(Ceil(2.3))

	x := 10

	// ❌ Violates: $X == $X
	if x == x {
		println("always true")
	}

	// ❌ Violates: $X != $X
	if x != x {
		println("never true")
	}

	if true {
		println("this always runs")
	} else {
		println("this never runs")
	}

	os.OpenFile("go.mod", 0, 0700)

	fmt.Println(jwt.DecodeStrict)

	n, err := strconv.Atoi("40000") // > 32767
	if err != nil {
		return
	}

	v := int16(n) // ❌ overflow risk
	_ = v

}
