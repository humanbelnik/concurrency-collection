# concurrency-collection

## Patterns

1. [Fan-in](./fan-in/). Merge multiple channels into a single one

   TL;DR

   - Read from each source in a separate goroutine.
   - Close `out` channel when all sources will be closed

2. [Fan-out](./fan-out/). Split channel into K channels

   TL;DR

   - **Example of channel nil-ing**
   - Read from source and Round-Robin (or anything else) traffic on K out channels.
   - Close `outs` when source is closed

3. [Tee](./tee/). Fan-out but we're not Round Robbinning the traffic, we repicate it into K streams

4. [Pipeline](./pipeline/).

5. [Done-channel](./done-ch/). Wait until work is done

## Structures

1. [Pool with gracefull shutdown](./pool/)

   - block/non-blocking select
   - waitgroups
   - range over channel
   - cancellation via context and done channel
