# Wait Group

### Steps to using wait groups
    * creating the wait group
    * Calling method add 
    * Running Goroutines with the method done
    * Calling method wait for the execution of goroutines 
     
WaitGroups are passed as pointers/references

## Wait Group Issues
    > .callout_info Without calling wait.Add() the wait group returns immediately without the concurrent code.