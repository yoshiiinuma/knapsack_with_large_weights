
package knapsack

import "fmt"

const DEBUG = true
const INF = 999999999

func showDP(DP [][]int) {
  for _, dpi := range(DP) {
    r := make([]string, len(dpi))
    for j, e :=  range(dpi) {
      if e == INF {
        r[j] = " *"
      } else {
        r[j] = fmt.Sprintf("%2d", e)
      }
    }
    fmt.Println(r)
  }
}


func setupDP(V []int) ([][]int, int) {
  N := len(V)
  M := 1
  for i := 0; i < N; i++ {
    M += V[i]
  }
  DP := make([][]int, N + 1)
  for i := 0; i < N + 1; i++ {
    DP[i] = make([]int, M)
  }
  for i := 1; i < M; i++ {
    DP[0][i] = INF
  }
  DP[0][0] = 0
  return DP, M
}

func Knapsack(K int, V []int, W []int) int {
  DP, M := setupDP(V)
  return knapsack(K, V, W, DP, M)
}

func min(a int, b int) int {
  if a < b {
    return a
  }
  return b
}

/**
 *
 * Suppose DP[i][j] represents the maximum value that can be obtained
 * from a combination of items 1 through i selected so that the total
 * weight does not exceed j 
 *
 * if j < W[i]
 *   DP[i+1][j] = DP[i][j]
 * else
 *   DP[i+1][j] = MIN DP[i][j]
 *                    DP[i][j-V[i]] + W[i]
 *
 */
func knapsack(K int, V []int, W []int, DP [][]int, M int) int {
  N := len(V)
  max := 0

  if DEBUG {
    fmt.Println("--------------------------------")
    fmt.Printf("N: %d, K: %d, M: %d\n", N, K, M)
    fmt.Printf("V: %v\n", V)
    fmt.Printf("W: %v\n", W)
    fmt.Println("--------------------------------")
  }
  for i := 1; i <= N; i++ {
    for j := 1; j < M; j++ {
      DP[i][j] = DP[i-1][j]
      if j >= V[i-1] {
        if DP[i][j] > DP[i-1][j - V[i-1]] + W[i-1] {
          DP[i][j] = DP[i-1][j - V[i-1]] + W[i-1]
        }
      }
      if DP[i][j] > 0 && DP[i][j] <= K {
        if max < j {
          max = j
        }
      }
    }
  }
  if DEBUG { showDP(DP) }
  return max
}

