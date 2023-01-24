package clock

import "time"

/*
リアルの現在時間とは別に、テスト用の固定の時間を設定できることを可能にする。
*/
type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

func (rc RealClocker) Now() time.Time {
	return time.Now()
}

type FixedClocker struct{}

func (fc FixedClocker) Now() time.Time {
	return time.Date(2023, 1, 23, 21, 4, 22, 0, time.UTC)
}
