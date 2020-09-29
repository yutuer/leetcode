package _63

// 更好的策略是假设外围有一圈障碍. 然后在使用统一的方法来计算
// 即将obstacleGrid 整体向下向右移动1
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					if i == 0 && j > 0 {
						dp[i][j] = dp[i][j-1]
					} else if j == 0 && i > 0 {
						dp[i][j] = dp[i-1][j]
					}
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
