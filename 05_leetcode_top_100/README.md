### 一、golang中int类型的最大值和最小值

####1. 无符号整型uint

```go
package main

//最大值
const UINT_MAX = ^uint(0) //18446744073709551615
//最小值
const UINT_MIN = 0
```

####2. 有符号整型int

```go
package main

//根据补码，其最大值二进制表示，首位0，其余1
const INT_MAX = int(^uint(0) >> 1) //-9223372036854775808
//根据补码，其最小值二进制表示，首位1，其余0
const INT_MIN = ^INT_MAX //9223372036854775807
```