# 01 Knapsack With Large Weights

Given a set of different N items, each one with an associated value and weight, determine the maximum value of items that can be packed into a knapsack of capacity K.
The values and weights of N items are given as integer arrays; V { v1, v2, ..., vn }, W { w1, w2, ..., wn }, respectively.

## Constraints

```
0 < N <= 100
0 < Vi <= 100
0 < Wi <= 10^7
0 < K <= 10^9
```

## Examples

```
INPUT:

  N: 4
  K: 5
  V: [3, 2, 4, 2]
  W: [2, 1, 3, 2]

OUTPUT:

  7 (= 3 + 2 + 2; the total weight is 5 = 2 + 1 + 2)
  
```

```
INPUT:

  N: 3
  K: 50
  V: [40, 80, 100]
  W: [10, 20 ,30]

OUTPUT:

  180 (= 80 + 100; the total weight is 50)
  
```

## Approach

The capacity of the knapsack is too big. To reduce the amount of calculation, let j be the total value that can be obtained from subset of V[1..i], and 
DP[i][j] be the minimum value of the total weight that can be obtained from selected items, then rhe recurrence relation is defined as:

```
i) When j < W[i]

  DP[i+1][j] = DP[i][j]

ii) When j >= W[i]

  DP[i+1][j] = MIN(DP[i][j], D[i][j - V[i]] + W[i])
```

## Complexity Analysis

Let M = Sum of V[1..N]

* Time Complexity: O(N*M)
* Space Complexity: O(N*M)
