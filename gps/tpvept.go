package gps

import "time"
import "encoding/json"

type TpvEpt time.Duration

func (t *TpvEpt) MarshalJSON() ([]byte, error) {
	if nil != t {
		tmp := *t
		return json.Marshal(time.Duration(tmp).Seconds())
	}
	return []byte{}, nil
}

func (t *TpvEpt) UnmarshalJSON(b []byte) error {
	var tmp *float64
	err := json.Unmarshal(b, tmp)
	if nil != err {
		if nil != tmp {
			tmp2 := *tmp
			tmp3 := time.Duration(int64(tmp2 * 1000000000))
			tmp4 := TpvEpt(tmp3)
			t = &tmp4
			return nil
		}
	} else {
		t = nil
		return err
	}
	t = nil
	return nil
}
