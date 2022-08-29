package booking

import "time"

type Limiter struct {
	Queue map[string][]int64
}

func Constructor() Limiter {
	limiter := Limiter{}
	limiter.Queue = make(map[string][]int64)
	return limiter
}
func (l Limiter) Next(ip, port string, duration, n int64) bool {
	key := ip + ":" + port
	curTime := time.Now().Unix()
	if v, ok := l.Queue[key]; ok {
		count := int64(len(v))
		if count < n {
			return true
		} else {
			for len(v) > 0 {
				earlyTime := v[0]
				if curTime-earlyTime > duration {
					v = v[1:]
				}
			}
			if int64(len(v)) < n {
				return true
			} else {
				return false
			}
		}
	} else {
		l.Queue[key] = make([]int64, 0)
		l.Queue[key] = append(l.Queue[key], curTime)
	}
	return true
}
