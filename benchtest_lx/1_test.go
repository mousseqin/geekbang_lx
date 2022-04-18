package benchtest_lx

import "testing"

// 单元测试覆盖率 	go test -v --coverprofile=benchtest_lx.cover ../benchtest_lx
// 单元测试覆盖率报告  go tool cover -html=benchtest_lx.cover -o=benchtest_lx.html

func TestFibonacci(t *testing.T) {
	//预先定义的一组斐波那契数列作为测试用例
	fsMap := map[int]int{}
	fsMap[-1] = 0
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 1
	fsMap[3] = 2
	fsMap[4] = 3
	fsMap[5] = 5
	fsMap[6] = 8
	fsMap[7] = 13
	fsMap[8] = 21
	fsMap[9] = 34
	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)
		}
	}
}

func BenchmarkFibonacci(b *testing.B){
	for i:=0;i<b.N;i++{
		Fibonacci(10)
	}
}
