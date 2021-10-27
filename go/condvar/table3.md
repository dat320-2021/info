# Figure 30.11: Thread Trace: Broken Solution (Version 2)

| Tc1 | State   | Tc2 | State   | Tp | State   | Count | Comment           |
|-----|---------|-----|---------|----|---------|-------|-------------------|
| c1  | Running |     | Ready   |    | Ready   | 0     |                   |
| c2  | Running |     | Ready   |    | Ready   | 0     |                   |
| c3  | Sleep   |     | Ready   |    | Ready   | 0     | Nothing to get    |
|     | Sleep   | c1  | Running |    | Ready   | 0     |                   |
|     | Sleep   | c2  | Running |    | Ready   | 0     |                   |
|     | Sleep   | c3  | Sleep   |    | Ready   | 0     | Nothing to get    |
|     | Sleep   |     | Sleep   | p1 | Running | 0     |                   |
|     | Sleep   |     | Sleep   | p2 | Running | 0     |                   |
|     | Sleep   |     | Sleep   | p4 | Running | 1     | Buffer now full   |
|     | Ready   |     | Sleep   | p5 | Running | 1     | Tc1 awoken        |
|     | Ready   |     | Sleep   | p6 | Running | 1     |                   |
|     | Ready   |     | Sleep   | p1 | Running | 1     |                   |
|     | Ready   |     | Sleep   | p2 | Running | 1     |                   |
|     | Ready   |     | Sleep   | p3 | Sleep   | 1     | Must sleep (full) |
| c2  | Running |     | Sleep   |    | Sleep   | 1     | Recheck condition |
| c4  | Running |     | Sleep   |    | Sleep   | 0     | Tc1 grabs data    |
| c5  | Running |     | Ready   |    | Sleep   | 0     | Oops! Woke Tc2    |
| c6  | Running |     | Ready   |    | Sleep   | 0     |                   |
| c1  | Running |     | Ready   |    | Sleep   | 0     |                   |
| c2  | Running |     | Ready   |    | Sleep   | 0     |                   |
| c3  | Sleep   |     | Ready   |    | Sleep   | 0     | Nothing to get    |
