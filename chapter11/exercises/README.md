## Chapter 11 exercises

#### Exercise 11.1

Write tests for the `charcount` program in Section 4.3.

#### Exercise 11.2

Write a set of tests for `IntSet (§6.5)` that checks that its behavior after each operation is equivalent to a set bsed on built-in maps.
Save your implementation for benchmarking in Exercise 11.7.

#### Exercise 11.3

`TestRandomPalindromes` only tests palindromes.
Write a randomized test that generates and verifies *non-palindromes.*

#### Exercise 11.4

Modify `randomPalindrome` to exercise `IsPalindrome's` handling of punctuation and spaces.

#### Exercise 11.5

Extend `TestSplit` to use a table of inputs and expected outputs.

#### Exercise 11.6

Write benchmarks to compare the `PopCount` implementation in Section 2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5.
At what point does the table-based approach break even?

```text
BenchmarkPopCountTable1-4        20000       121925    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable10-4       10000       567430    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable100-4      10000       5678211   ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable1000-4     3000        16778123  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable10000-4    300         17531809  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountTable100000-4   100         54971520  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift1-4        20000000    59.1      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift10-4       3000000     620       ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift100-4      200000      8679      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift1000-4     10000       113931    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift10000-4    2000        1130659   ns/op  0  B/op  0  allocs/op
BenchmarkPopCountShift100000-4   100         11198696  ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears1-4       1000000000  3.12      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears10-4      30000000    39.0      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears100-4     3000000     528       ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears1000-4    200000      7529      ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears10000-4   10000       100799    ns/op  0  B/op  0  allocs/op
BenchmarkPopCountClears100000-4  1000        1257457   ns/op  0  B/op  0  allocs/op
```

#### Exercise 11.7

Write benchmarks for `Add`, `UnionWith`, and other methods of `*IntSet` (6.5) using large pseudo-random inputs.
How fast can you make these methods run?
How does the choice of word size affect performance?
How fast is `IntSet` compared to a set implementation based on the built-in map type?

```text
BenchmarkMapIntSetAdd10-4       1000000   1411    ns/op  323    B/op  3     allocs/op
BenchmarkMapIntSetAdd100-4      100000    16217   ns/op  3474   B/op  20    allocs/op
BenchmarkMapIntSetAdd1000-4     10000     194432  ns/op  55611  B/op  98    allocs/op
BenchmarkMapIntSetHas10-4       20000000  59.8    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetHas100-4      20000000  68.4    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetHas1000-4     20000000  63.7    ns/op  0      B/op  0     allocs/op
BenchmarkMapIntSetAddAll10-4    2000000   985     ns/op  323    B/op  3     allocs/op
BenchmarkMapIntSetAddAll100-4   100000    13628   ns/op  3475   B/op  20    allocs/op
BenchmarkMapIntSetAddAll1000-4  5000      206743  ns/op  55627  B/op  98    allocs/op
BenchmarkMapIntSetString10-4    500000    2916    ns/op  368    B/op  14    allocs/op
BenchmarkMapIntSetString100-4   50000     33491   ns/op  4577   B/op  108   allocs/op
BenchmarkMapIntSetString1000-4  5000      336538  ns/op  41164  B/op  994   allocs/op
BenchmarkBitIntSetAdd10-4       3000000   536     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAdd100-4      300000    4428    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAdd1000-4     30000     45583   ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas10-4       30000000  42.8    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas100-4      30000000  41.3    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetHas1000-4     30000000  41.9    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll10-4    20000000  106     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll100-4   2000000   609     ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetAddAll1000-4  300000    5464    ns/op  0      B/op  0     allocs/op
BenchmarkBitIntSetString10-4    500000    3065    ns/op  256    B/op  12    allocs/op
BenchmarkBitIntSetString100-4   50000     27027   ns/op  3649   B/op  106   allocs/op
BenchmarkBitIntSetString1000-4  5000      209314  ns/op  33009  B/op  1002  allocs/op
```

