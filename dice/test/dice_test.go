package test

import (
	"dice"
	"testing"
)

func TestPoolProbSum(t *testing.T) {
	pool := dice.DicePool{Number: []int{1}, Size: []int{2}}

	probHalf := pool.PoolProbSum(1)

	if probHalf != 0.5 {
		t.Errorf("PoolProbSum failed. Expected to get 0.5, got %f.", probHalf)
	}

}
