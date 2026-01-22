## GO CONCURRENCY PATTERN

1. **_Producer-Consumer_**: One goroutine creates numbers, another consumes them. Uses context for clean shutdown.
2. **_Alternate Print_**: Two goroutines take turns printing numbers using ping-pong channels.
3. **_Ring Pattern_**: Three goroutines pass control in a circle to print "ABCABC..."
4. **_Worker Pool_**: 3 workers grab jobs from a queue, process them, and send results back.
5. **_Deadlock Prevention_**: Two goroutines safely acquire two locks by always locking in the same order.
6. **_Rate Limiter_**: Uses a ticker to limit how fast requests are processed (2 per second here).
7. **_Fan-Out Fan-In_**: One source splits work to many workers, then merges all results back together (like a tree - branches out then comes back to trunk)
8. **_Pipeline_**: Data flows through stages like a factory assembly line - each stage does one transformation
9. **_Semaphore_**: Limit how many goroutines run at once (like a parking lot with limited spaces)
10. **_Select with Timeout_**: Don't wait forever - give up after a time limit
11. **_Future/Promise_**: Start work now, get result later when you need it
12. **_Or-Done_**: Combine context cancellation with channel reading for clean shutdowns
13. **_Error Group_**: Run multiple tasks and stop if any fails
14. **_Broadcast_**: Send one signal to many listeners at once (using close() as the broadcast mechanism)
15. **_Tee_**: Split one input stream to multiple identical outputs (like a T-pipe in plumbing)
16. **_Bridge_**: Connect multiple channels end-to-end into one continuous stream
17. **_Buffering_**: Use buffers to handle burst traffic - fast producer, slow consumer
18. **_Dropping_**: Drop requests when system is overloaded (better than crashing!)
19. **_Barrier_**: Make all goroutines wait at a checkpoint before proceeding together
20. **_Phased Shutdown_**: Graceful shutdown in stages - stop accepting work, finish current work, then exit
21. **_Retry with Backoff_**: Retry failed operations with increasing delays (100ms, 200ms, 400ms...)
22. **_Circuit Breaker_**: Stop trying after too many failures to protect the system
23. **_Debounce_**: Wait for "quiet period" before executing (like search suggestions)
24. **_Throttle_**: Limit how often something executes (max once per 300ms)
25. **_Or-Channel_**: Return when ANY channel closes (whichever is fastest)
26. **_Replicated Requests_**: Send same request to multiple servers, use fastest response
27. **_Bounded Parallelism_**: Process many items but limit concurrency (like semaphore but for a list)
28. **_Context Propagation_**: Pass context through multiple layers for cancellation
29. **_Confinement_**: Only one goroutine owns/modifies data (prevents races)
30. **_For-Select Loop_**: Standard pattern for processing channels with cancellation
31. **_Singleton_**: Create exactly one instance, thread-safe (using sync.Once)
32. **_Map-Reduce_**: Parallel map operation, then reduce to single result
33. **_Preemption_**: Interrupt long-running tasks
34. **_Heartbeat_**: Periodic signals proving goroutine is still alive
35. **_Leak Prevention_**: Ensure goroutines clean up (use context!)
36. **_Staged Pipeline with Errors_**: Pipeline that propagates errors through stages
37. **_Resource Pool_**: Reuse expensive resources (like database connections)
38. **_Atomic Operations_**: Lock-free operations using atomic package
39. **_Mailbox_**: Each goroutine has own message queue (Actor model)
40. **_Scatter-Gather_**: Distribute work, collect ALL results (vs fan-out which streams)
41. **_Reactive Streams_**: Filter, map, transform event streams
42. **_Split-Merge_**: Split by condition, process differently, merge back
43. **_Cancellation Cascade_**: Cancel parent → all children cancel automatically
44. **_Token Passing_**: Goroutines pass token in circle for coordination
45. **_Active Object_**: Object with its own execution thread
46. **_Guarded Suspension_**: Wait/suspend until condition is met (using sync.Cond)
47. **_Two-Phase Termination_**: Stop request → cleanup → stop confirmation
48. **_Balking_**: Return immediately if precondition not met (don't wait)
49. **_Scheduler_**: Schedule tasks for future execution
50. **_Read-Write Lock_**: Multiple readers OR one writer (sync.RWMutex)
51. **_Double-Checked Locking_**: Optimize initialization checks (atomic + mutex)
52. **_Lock Striping_**: Multiple locks to reduce contention (like ConcurrentHashMap in Java)
53. **_Copy-on-Write_**: Share reads, copy on write (used in filesystems)
54. **_Version Vector_**: Track causality in distributed systems
55. **_Optimistic Locking_**: Use versions to detect conflicts (database pattern)
56. **_Lease_**: Time-limited exclusive access (distributed locks)
57. **_Immutable_**: Thread-safe by design - never mutate!
58. **_Thread-Local Storage_**: Each goroutine has own copy
59. **_Continuation Passing_**: Pass "what to do next" as callback
60. **_Monitor Object_**: Encapsulate all synchronization in one object
