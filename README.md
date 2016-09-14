[![GoDoc](https://godoc.org/github.com/steakknife/hamming?status.png)](https://godoc.org/github.com/steakknife/hamming)


# hamming distance calculations in Go

Copyright © 2014, 2015, 2016 Barry Allard

[MIT license](MIT-LICENSE.txt)

## Performance

```
$ go test -bench=.
BenchmarkInt8PopCnt-4               	500000000	         3.97 ns/op
BenchmarkInt16PopCnt-4              	500000000	         3.22 ns/op
BenchmarkInt32PopCnt-4              	500000000	         3.33 ns/op
BenchmarkInt64PopCnt-4              	500000000	         3.36 ns/op
BenchmarkIntPopCnt-4                	300000000	         5.67 ns/op
BenchmarkUint8PopCnt-4              	1000000000	         2.78 ns/op
BenchmarkUint16PopCnt-4             	1000000000	         2.82 ns/op
BenchmarkUint32PopCnt-4             	1000000000	         2.67 ns/op
BenchmarkUint64PopCnt-4             	1000000000	         2.68 ns/op
BenchmarkUintPopCnt-4               	300000000	         4.87 ns/op
BenchmarkBytePopCnt-4               	500000000	         4.07 ns/op
BenchmarkRunePopCnt-4               	500000000	         3.69 ns/op
BenchmarkCountBitsInt8-4            	2000000000	         0.41 ns/op
BenchmarkCountBitsInt16-4           	2000000000	         0.40 ns/op
BenchmarkCountBitsInt32-4           	2000000000	         0.38 ns/op
BenchmarkCountBitsInt64-4           	2000000000	         0.36 ns/op
BenchmarkCountBitsInt-4             	200000000	         6.20 ns/op
BenchmarkCountBitsUint16-4          	2000000000	         0.37 ns/op
BenchmarkCountBitsUint32-4          	2000000000	         0.37 ns/op
BenchmarkCountBitsUint64-4          	2000000000	         0.36 ns/op
BenchmarkCountBitsUint64Alt-4       	200000000	         7.25 ns/op
BenchmarkCountBitsUint-4            	500000000	         4.08 ns/op
BenchmarkCountBitsUintReference-4   	100000000	        17.1 ns/op
BenchmarkCountBitsByte-4            	2000000000	         0.36 ns/op
BenchmarkCountBitsByteAlt-4         	2000000000	         0.36 ns/op
BenchmarkCountBitsRune-4            	2000000000	         0.35 ns/op
PASS
ok  	github.com/steakknife/hamming	45.092s
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
