package dice

import (
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func roll(size int, number int) []int {
	var wg sync.WaitGroup
	roll := make([]int, number)
	for d := 0; d < number; d++ {
		d2 := d
		wg.Add(1)
		go func() {
			roll[d2] = rand.Intn(size) + 1
			defer wg.Done()
		}()
	}
	wg.Wait()
	return roll
}

type DicePool struct {
	Number []int
	Size   []int
}

func (pool *DicePool) RollPool() [][]int {
	var wg sync.WaitGroup
	rolled := make([][]int, len(pool.Number))
	for p := 0; p < len(pool.Number); p++ {
		p2 := p
		wg.Add(1)
		go func() {
			rolled[p2] = roll(pool.Size[p2], pool.Number[p2])
			defer wg.Done()
		}()
	}
	wg.Wait()
	return rolled
}

func (pool *DicePool) PoolProbSum(target int) float32 {

}

func (pool *DicePool) PoolProbSuccess(success int, target int) float32 {

}
