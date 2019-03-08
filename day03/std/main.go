// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	addr1 := flag.String("addr1", "127.0.0.1", "usage of addr1")

// 	var addr2 string
// 	flag.StringVar(&addr2, "addr2", "127.0.0.2", "usage of addr2")

// 	flag.Parse()

// 	fmt.Println(addr1)
// 	fmt.Println(*addr1)
// 	fmt.Println(addr2)
// }

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	type flags struct {
		BoolFlag     bool
		DurationFlag time.Duration
		Float64Flag  float64
		Int64Flag    int64
		IntFlag      int
		StringFlag   string
		Uint64Flag   uint64
		UintFlag     uint
	}

	var ff flags

	flag.BoolVar(&ff.BoolFlag, "enableLasers", true, "usage")
	flag.DurationVar(&ff.DurationFlag, "duration", time.Second, "Duration")
	flag.Float64Var(&ff.Float64Flag, "float64", 12.34, "Float64")
	flag.Int64Var(&ff.Int64Flag, "int64", 1234, "Int64")
	flag.IntVar(&ff.IntFlag, "int", 1234, "Int")
	flag.StringVar(&ff.StringFlag, "string", "default", "String")
	flag.Uint64Var(&ff.Uint64Flag, "uint64", 5678, "Uint64")
	flag.UintVar(&ff.UintFlag, "uint", 6789, "Uint")
	flag.Parse()

	fmt.Printf("bool=%v\n", ff.BoolFlag)
	fmt.Printf("duration=%v\n", ff.DurationFlag)
	fmt.Printf("float64=%v\n", ff.Float64Flag)
	fmt.Printf("int64=%v\n", ff.Int64Flag)
	fmt.Printf("int=%v\n", ff.IntFlag)
	fmt.Printf("string=%v\n", ff.StringFlag)
	fmt.Printf("uint64=%v\n", ff.Uint64Flag)
	fmt.Printf("uint=%v\n", ff.UintFlag)
}
