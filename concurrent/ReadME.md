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

Closing a Channel is done by the goroutine that is writing to the channel
Unsed channels can be detected by the garbage collector is they are not in use :sparkles:

## Fork - Join model



## Wait Group

### Steps to using wait groups
    * creating the wait group
    * Calling method add 
    * Running Goroutines with the method done
    * Calling method wait for the execution of goroutines 
     
WaitGroups are passed as pointers/references

## Wait Group Issues
    > :green_book: Without calling wait.Add() the wait group returns immediately without the concurrent code.

### Type Assertions using atomic.Load()
    the compiler runs a type assertion against the value passed into the function

    ```go
        type student struct {
	        grade map[string]int
        }

        var val atomic.Value

        //this will lead to a runtime panic 
        // {panic interface conversion : interface {} is nil, not of type student}
        valued := val.Load().(student)
    ```

    Ensuring that type assertion doesnt fail pass the zero value for the type

    ```go
        type student struct {
	        grade map[string]int
        }

        var val atomic.Value
      	val.Store(student{grade: map[string]int {}})
        valued := val.Load().(student)
        valued.grade["intro to prog"] = 60
    ```

## Mutex
  > Mutual Exclusion
    Protecting shared resources from sumultaneous access by multiple goroutines

Internally, the sync.Mutex type in Go utilizes low-level atomic operations provided by the underlying processor architecture. 
These atomic operations ensure that the mutex operations themselves are thread-safe and efficient, allowing one goroutine has access to critial code section at a time

The Lock() to acquires the lock by using the atomic.CompareAndSwapInt32() function. Comparing the state variable's value with 0 and swaps it with 1 if they are equal. If the swap is successful, the lock is acquired without blocking. Otherwise, the Lock() method calls runtime_SemacquireMutex() to wait for the lock to become available.

The Unlock() releases the lock by using the atomic.CompareAndSwapInt32() function comparing the state variable's value with 1 and swaps it with 0 if they are equal. If the swap is successful, indicating that the lock was held, the Unlock() method calls runtime_SemreleaseMutex() to wake up any waiting goroutine.

    > :green_book: 
    > Atomic operations ensures that the lock’s state is updated atomically, without interference from other goroutines. 
    > Atomicity guarantees thread safety and eliminates the need for additional locks or synchronization mechanisms.


## Race Conditions

occur when multiple goroutines access and modify shared data concurrently, leading to unpredictable and erroneous behavior

    > :pushpin: Atomics in the sync package allow for elimination of race conditions
    > :pushpin: Mutexes prevent race conditions by allowing only one goroutine to acquire the lock and access the shared resource.


  