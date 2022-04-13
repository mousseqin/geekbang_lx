package rolling

import (
	"sync"
	"time"
)

// Number 跟踪一个numberBucket在一个有界数的时间桶。目前，这些桶的长度为一秒，只保留最后10秒
type Number struct {
	Buckets map[int64]*numberBucket
	Mutex *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

func NewNumber() *Number {
	r := &Number{
		Buckets:make(map[int64]*numberBucket),
		Mutex: &sync.RWMutex{},
	}
	return r
}

func (r *Number) getCurrentBucket() *numberBucket{
	now := time.Now().Unix()
	var bucket *numberBucket
	var flag bool

	if bucket, flag = r.Buckets[now];!flag {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}
	return bucket
}

func (r *Number) RemoveOldBucket() {
	now := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets,timestamp)
		}
	}
}

// Increment 递增timeBucket中的数字
func (r *Number) Increment (i float64) {
	if i == 0 {
		return
	}
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b:= r.getCurrentBucket()
	b.Value+=i
	r.RemoveOldBucket()
}

// UpdateMax 更新当前桶中的最大值
func (r  *Number) UpdateMax(n float64) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b:= r.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	r.RemoveOldBucket()
}

// Sum 对桶内在过去10秒的值进行求和
func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for t , bucket := range r.Buckets {
		if t >= now.Unix() - 10 {
			sum+=bucket.Value
		}
	}
	return sum
}

// Max 返回过去10秒内看到的最大值
func (r *Number) Max(now time.Time) float64 {
	var max float64
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for t , bucket := range r.Buckets {
		if t >= now.Unix() - 10 {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}
	return max
}

// Avg 过去10秒的平均值
func (r *Number) Avg(now time.Time) float64 {
	return r.Sum(now) / 10
}











































