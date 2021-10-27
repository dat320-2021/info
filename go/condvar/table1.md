# Figure 30.11: Thread Trace: Broken Solution (Version 2)

| Tc1 | State   | Tc2 | State   | Tp | State   | Count | Comment           |
|-----|---------|-----|---------|----|---------|-------|-------------------|
| c1  | Running |     | Ready   |    | Ready   | 0     |                   |
| c2  | Running |     | Ready   |    | Ready   | 0     |                   |
| c3  | Sleep   |     | Ready   |    | Ready   | 0     | Nothing to get    |
|     | Sleep   | c1  | Running |    | Ready   | 0     |                   |
|     | Sleep   | c2  | Running |    | Ready   | 0     |                   |
|     | Sleep   | c3  | Sleep   |    | Ready   | 0     | Nothing to get    |
