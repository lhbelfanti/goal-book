## Chapter 9 exercises

#### Exercise 9.1

Add a function `Withdraw(amount int) bool` to the `gopl.io/ch9/bank1` program.
The result should indicate whether the transaction succeeded or failed due to insufficient funds.
The message send to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to `Withdraw`.

#### Exercise 9.2

Rewrite the `PopCount` example from Section 2.6.2 so that it initializes the lookup table using `sync.Once` the first time it is needed.
(Realistically, the cost of synchronization whould be prohibitive for a small and highly optimized function like `PopCount`)

#### Exercise 9.3

Extend the Func type and the `(*Memo).Get` method so that callers may provide an optional done channel through which they can cancel the operation (§8.9).
The results of a cancelled Func call should not be cached.

#### Exercise 9.4

Construct a pipeline that connects an arbitrary number of goroutines with channels.
What is the maximum number of pipeline stages you can create with out running out of memory?
How long does a value take to transit the entire pipeline?

## Result

### With ResetTimer after initialize all channels

```text
BenchmarkPipeline10-4       500000  2573       ns/op  0   B/op  0  allocs/op
BenchmarkPipeline100-4      50000   24956      ns/op  0   B/op  0  allocs/op
BenchmarkPipeline1000-4     5000    286223     ns/op  0   B/op  0  allocs/op
BenchmarkPipeline10000-4    300     3980524    ns/op  0   B/op  0  allocs/op
BenchmarkPipeline100000-4   30      39138676   ns/op  26  B/op  0  allocs/op
BenchmarkPipeline1000000-4  3       379899748  ns/op  0   B/op  0  allocs/op
```

### Without ResetTimer after initialize all channels

```text
BenchmarkPipeline10-4       500000  2497        ns/op  0          B/op  0        allocs/op
BenchmarkPipeline100-4      50000   24588       ns/op  0          B/op  0        allocs/op
BenchmarkPipeline1000-4     5000    287139      ns/op  26         B/op  0        allocs/op
BenchmarkPipeline10000-4    300     3697650     ns/op  4825       B/op  41       allocs/op
BenchmarkPipeline100000-4   30      42878492    ns/op  381184     B/op  3657     allocs/op
BenchmarkPipeline1000000-4  1       4149927880  ns/op  621803856  B/op  2799323  allocs/op
```

## With a 16GB OS, ~8GB available mem

When trying to create 5000000, it running out of memory.

#### Exercise 9.5

Write a program with two gotoutines that send messages back and forth over two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?

## Result

```text
2.1975286807808625e+06 rounds per second
```

#### Exercise 9.6

Measure how the performance of a compute-bond parallel program (see Exercise 8.5) varies with GOMAXPROCS.
What is the optimal value of your computer?
How many CPUs does your computer have?

## Result

With Intel I5 6200, 4 Cores, 2 worker threads is economical.

```text
NumCPU: 4
Done no concurrency. Used: 5.101064434s
Done. Worker Number: 1 Used: 5.076399681s
Done. Worker Number: 2 Used: 2.735404646s
Done. Worker Number: 3 Used: 2.56978686s
Done. Worker Number: 4 Used: 2.522902384s
```

