# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to a repository to complete the task. Remember to also submit your answers to Blackboard

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 >  Parallelism is when tasks run at the same time, where as concurrency is when two tasks can start, run and complete in overlapping time periods
 
 ### Why have machines become increasingly multicore in the past decade?
 > To run mulitple tasks at the same time and therefor run faster.
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Input/output-intensive programs mostly wait for input or output operations to complete. Concurrent programming allows the time that would be spent waiting to be used for another task.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > *Your answer here*
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > - A Fiber is a lightweight thread that uses cooperative multitasking instead of preemptive multitasking. 
   - A Coroutine is a component that generalizes a subroutine to allow multiple entry points for suspending and resuming            execution at certain locations. 
   - A Green Thread is a thread that is scheduled by a virtual machine (VM) instead of natively by the underlying operating        system. 
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > Pthread_create() and Threading.thread() creates a thread. Go creates a coroutine.
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > Lock GIL only allows one thread to be in execution at the same time
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > Numpy e.g.
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > Changes the number of processors in use. 
