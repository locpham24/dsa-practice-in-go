## WeightedQuickUnion

### Find
O(N)

### Union
O(lg N)

Explain:
Depth of element X only increase 1 when:
- The tree of X <= larger tree => The tree of X will grow at least double
=> It takes at most lgN times