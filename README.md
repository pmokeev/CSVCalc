# CSVCalc

CSVCalc allows you to calculate simple integer expressions, in tables, with the analogy as in Microsoft Excel.

# Example

![Example](assets/example.gif)

Before running the application, you must install [Golang 1.19+](https://go.dev/dl/).

**Example of usage:**
```bash
$ go build -o csvcalc cmd/main.go
$ ./csvcalc ./test/data/example.csv
,A,B,Cell
1,1,0,1
2,2,6,0
30,0,1,5
```

# Information

- Zero external library dependency
- Implemented using queue data structure

# What next?

- The check for cyclic dependence of table cells is not implemented, for example:
    ```yaml
    ,A
    1,A1
    ```
- Add parallelism in calculations
