# Fibonacci Number
Calculate the fibonacci number for the index.

Indexes grater than `92` overflows `int64` so
 `BigFib` can be used for the big index or big output using `big.Int` type.

However, two functions with alternative algorithms are available as `RFib` and `RBigFib` using recursive function. Check benchmark to see the functions performances. 