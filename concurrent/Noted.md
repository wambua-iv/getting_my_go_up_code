# Concurrency 

## Goroutines
- Lightweight thread managed by the Go runtime
  - Launched using go keyword before a function
  - Return values from the functions are ignored
  > Have closures that wraps up business logic to handle concurrent bookkeeping
  
## Channels
Reference type in GO
Goroutines communicate through channels
  - Zero value for a channel is nil.
  - Each value written to a channel can be read only once. 
  -  If multiple goroutines are reading from the same channel, 
  a value written to the channel will be read by only one of them
```go 
//creating a channel
ch := make(chan int)
//writing to a channel
ch <- b
//reading from a channel
a := <- ch

//closing a channel
close(ch)
```