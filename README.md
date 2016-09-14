[![GoDoc](https://godoc.org/github.com/steakknife/hamming?status.png)](https://godoc.org/github.com/steakknife/hamming)


# hamming distance calculations in Go

Copyright © 2014, 2015, 2016 Barry Allard

[MIT license](MIT-LICENSE.txt)

## Performance

```
$ go test -bench=.
BenchmarkCountBitsInt8PopCnt-4      	500000000	         3.96 ns/op
BenchmarkCountBitsInt16PopCnt-4     	500000000	         3.24 ns/op
BenchmarkCountBitsInt32PopCnt-4     	500000000	         3.36 ns/op
BenchmarkCountBitsInt64PopCnt-4     	500000000	         3.44 ns/op
BenchmarkCountBitsIntPopCnt-4       	300000000	         5.42 ns/op
BenchmarkCountBitsUint8PopCnt-4     	1000000000	         2.60 ns/op
BenchmarkCountBitsUint16PopCnt-4    	1000000000	         2.59 ns/op
BenchmarkCountBitsUint32PopCnt-4    	1000000000	         2.55 ns/op
BenchmarkCountBitsUint64PopCnt-4    	1000000000	         2.51 ns/op
BenchmarkCountBitsUintPopCnt-4      	300000000	         4.38 ns/op
BenchmarkCountBitsBytePopCnt-4      	500000000	         3.21 ns/op
BenchmarkCountBitsRunePopCnt-4      	500000000	         3.29 ns/op
BenchmarkCountBitsInt8-4            	2000000000	         0.38 ns/op
BenchmarkCountBitsInt16-4           	2000000000	         0.41 ns/op
BenchmarkCountBitsInt32-4           	2000000000	         0.36 ns/op
BenchmarkCountBitsInt64-4           	2000000000	         0.37 ns/op
BenchmarkCountBitsInt-4             	200000000	         6.36 ns/op
BenchmarkCountBitsUint16-4          	2000000000	         0.36 ns/op
BenchmarkCountBitsUint32-4          	2000000000	         0.35 ns/op
BenchmarkCountBitsUint64-4          	2000000000	         0.37 ns/op
BenchmarkCountBitsUint64Alt-4       	200000000	         7.06 ns/op
BenchmarkCountBitsUint-4            	300000000	         4.16 ns/op
BenchmarkCountBitsUintReference-4   	100000000	        16.9 ns/op
BenchmarkCountBitsByte-4            	2000000000	         0.36 ns/op
BenchmarkCountBitsByteAlt-4         	2000000000	         0.36 ns/op
BenchmarkCountBitsRune-4            	2000000000	         0.37 ns/op
PASS
ok  	github.com/steakknife/hamming	42.730s
$
```

## Usage

```go
import 'github.com/steakknife/hamming'

// ...

// hamming distance between values
hamming.Byte(0xFF, 0x00) // 8
hamming.Byte(0x00, 0x00) // 0

// just count bits in a byte
hamming.CountBitsByte(0xA5), // 4
```

See help in the [docs](https://godoc.org/github.com/steakknife/hamming)

## Get

    go get -u github.com/steakknife/hamming  # master is always stable

## Source

- On the web: https://github.com/steakknife/hamming

- Git: `git clone https://github.com/steakknife/hamming`

## Contact

- [Feedback](mailto:barry.allard@gmail.com)

- [Issues](https://github.com/steakknife/hamming/issues)

## License 

[MIT license](MIT-LICENSE.txt)

Copyright © 2014, 2015, 2016 Barry Allard
