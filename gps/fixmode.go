package gps

import "encoding/json"

type FixMode uint8

const (
	// no mode value yet seen
	FixNotSeen FixMode = 0
	// no fix
	FixNone FixMode = 1
	// 2D
	Fix2D FixMode = 2
	// 3D
	Fix3D FixMode = 3
)

func (f *FixMode) MarshalJSON() ([]byte, error) {
	if nil != f {
		return json.Marshal(uint8(*f))
	}
	return []byte{}, nil
}

func (f *FixMode) UnmarshalJSON(b []byte) error {
	var tmp *uint8
	err := json.Unmarshal(b, tmp)
	if nil != err {
		if nil != tmp {
			tmp2 := *tmp
			// No mode
			if tmp2 == 0 {
				tmp3 := FixNotSeen
				f = &tmp3
				return nil
			}
			// No fix
			if tmp2 == 1 {
				tmp3 := FixNone
				f = &tmp3
				return nil
			}
			// 2D
			if tmp2 == 2 {
				tmp3 := Fix2D
				f = &tmp3
				return nil
			}
			// 3D
			if tmp2 == 3 {
				tmp3 := Fix3D
				f = &tmp3
				return nil
			}
		}
	} else {
		f = nil
		return err
	}
	f = nil
	return nil
}
