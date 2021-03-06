// Code generated by protoc-gen-go.
// source: gps.proto
// DO NOT EDIT!

package mrn

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// * NMEA mode: %d, 0=no mode value yet seen, 1=no fix, 2=2D, 3=3D.
type TPV_NEMA int32

const (
	TPV_NO_MODE_VALUE TPV_NEMA = 0
	TPV_NO_FIX        TPV_NEMA = 1
	TPV_FIX_2D        TPV_NEMA = 2
	TPV_FIX_3D        TPV_NEMA = 3
)

var TPV_NEMA_name = map[int32]string{
	0: "NO_MODE_VALUE",
	1: "NO_FIX",
	2: "FIX_2D",
	3: "FIX_3D",
}
var TPV_NEMA_value = map[string]int32{
	"NO_MODE_VALUE": 0,
	"NO_FIX":        1,
	"FIX_2D":        2,
	"FIX_3D":        3,
}

func (x TPV_NEMA) String() string {
	return proto.EnumName(TPV_NEMA_name, int32(x))
}
func (TPV_NEMA) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type TPV struct {
	Mode TPV_NEMA `protobuf:"varint,1,opt,name=mode,enum=mrn.TPV_NEMA" json:"mode,omitempty"`
	// * Time (UTC) in ms since epoch. May be absent if mode is not 2 or 3.
	Time *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	// * Estimated timestamp error (seconds, 95% confidence). Present if time is present.
	Ept *google_protobuf2.Duration `protobuf:"bytes,3,opt,name=ept" json:"ept,omitempty"`
	// * Latitude in degrees: +/- signifies North/South. Present when mode is 2 or 3.
	Lat *google_protobuf.DoubleValue `protobuf:"bytes,4,opt,name=lat" json:"lat,omitempty"`
	// * Longitude in degrees: +/- signifies East/West. Present when mode is 2 or 3.
	Lon *google_protobuf.DoubleValue `protobuf:"bytes,5,opt,name=lon" json:"lon,omitempty"`
	// * Altitude in meters. Present if mode is 3.
	Alt *google_protobuf.DoubleValue `protobuf:"bytes,6,opt,name=alt" json:"alt,omitempty"`
	// * Longitude error estimate in meters, 95% confidence. Present if mode is 2 or 3 and DOPs can be calculated from the satellite view.
	Epx *google_protobuf.DoubleValue `protobuf:"bytes,7,opt,name=epx" json:"epx,omitempty"`
	// * Latitude error estimate in meters, 95% confidence. Present if mode is 2 or 3 and DOPs can be calculated from the satellite view.
	Epy *google_protobuf.DoubleValue `protobuf:"bytes,8,opt,name=epy" json:"epy,omitempty"`
	// * Estimated vertical error in meters, 95% confidence. Present if mode is 3 and DOPs can be calculated from the satellite view.
	Epv *google_protobuf.DoubleValue `protobuf:"bytes,9,opt,name=epv" json:"epv,omitempty"`
	// * Course over ground, degrees from true north.
	Track *google_protobuf.DoubleValue `protobuf:"bytes,10,opt,name=track" json:"track,omitempty"`
	// * Speed over ground, meters per second.
	Speed *google_protobuf.DoubleValue `protobuf:"bytes,11,opt,name=speed" json:"speed,omitempty"`
	// * Climb (positive) or sink (negative) rate, meters per second.
	Climb *google_protobuf.DoubleValue `protobuf:"bytes,12,opt,name=climb" json:"climb,omitempty"`
	// * Direction error estimate in degrees, 95% confidence.
	Epd *google_protobuf.DoubleValue `protobuf:"bytes,13,opt,name=epd" json:"epd,omitempty"`
	// * Speed error estinmate in meters/sec, 95% confidence.
	Eps *google_protobuf.DoubleValue `protobuf:"bytes,14,opt,name=eps" json:"eps,omitempty"`
	// * Climb/sink error estimate in meters/sec, 95% confidence.
	Epc *google_protobuf.DoubleValue `protobuf:"bytes,15,opt,name=epc" json:"epc,omitempty"`
}

func (m *TPV) Reset()                    { *m = TPV{} }
func (m *TPV) String() string            { return proto.CompactTextString(m) }
func (*TPV) ProtoMessage()               {}
func (*TPV) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TPV) GetMode() TPV_NEMA {
	if m != nil {
		return m.Mode
	}
	return TPV_NO_MODE_VALUE
}

func (m *TPV) GetTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TPV) GetEpt() *google_protobuf2.Duration {
	if m != nil {
		return m.Ept
	}
	return nil
}

func (m *TPV) GetLat() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Lat
	}
	return nil
}

func (m *TPV) GetLon() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Lon
	}
	return nil
}

func (m *TPV) GetAlt() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Alt
	}
	return nil
}

func (m *TPV) GetEpx() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Epx
	}
	return nil
}

func (m *TPV) GetEpy() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Epy
	}
	return nil
}

func (m *TPV) GetEpv() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Epv
	}
	return nil
}

func (m *TPV) GetTrack() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Track
	}
	return nil
}

func (m *TPV) GetSpeed() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Speed
	}
	return nil
}

func (m *TPV) GetClimb() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Climb
	}
	return nil
}

func (m *TPV) GetEpd() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Epd
	}
	return nil
}

func (m *TPV) GetEps() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Eps
	}
	return nil
}

func (m *TPV) GetEpc() *google_protobuf.DoubleValue {
	if m != nil {
		return m.Epc
	}
	return nil
}

func init() {
	proto.RegisterType((*TPV)(nil), "mrn.TPV")
	proto.RegisterEnum("mrn.TPV_NEMA", TPV_NEMA_name, TPV_NEMA_value)
}

func init() { proto.RegisterFile("gps.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0xd2, 0x4f, 0x8b, 0xe2, 0x30,
	0x18, 0x06, 0xf0, 0xad, 0xad, 0xdd, 0x35, 0xae, 0x6e, 0x37, 0xa7, 0xac, 0x2c, 0xbb, 0xae, 0x27,
	0x61, 0x21, 0x42, 0x3d, 0xcf, 0x41, 0xa8, 0x82, 0x30, 0xfe, 0xa1, 0x38, 0xc5, 0x5b, 0x49, 0xdb,
	0x8c, 0xc8, 0xb4, 0x4d, 0x68, 0x53, 0x47, 0xbf, 0xc5, 0x7c, 0xe4, 0x21, 0x89, 0x5e, 0x46, 0x06,
	0x72, 0x0b, 0xcd, 0xef, 0x79, 0xf2, 0xbe, 0x50, 0xd0, 0x39, 0xf0, 0x1a, 0xf3, 0x8a, 0x09, 0x06,
	0xed, 0xa2, 0x2a, 0x07, 0x7f, 0x0f, 0x8c, 0x1d, 0x72, 0x3a, 0x51, 0x9f, 0x92, 0xe6, 0x79, 0x22,
	0x8e, 0x05, 0xad, 0x05, 0x29, 0xb8, 0x56, 0x83, 0x3f, 0x1f, 0x41, 0xd6, 0x54, 0x44, 0x1c, 0x59,
	0xf9, 0xd9, 0xfd, 0x6b, 0x45, 0x38, 0xa7, 0xd5, 0xf5, 0x95, 0xd1, 0x9b, 0x0b, 0xec, 0xdd, 0x36,
	0x82, 0xff, 0x80, 0x53, 0xb0, 0x8c, 0x22, 0x6b, 0x68, 0x8d, 0xfb, 0x7e, 0x0f, 0x17, 0x55, 0x89,
	0x77, 0xdb, 0x08, 0xaf, 0xe7, 0xab, 0x59, 0xa8, 0xae, 0x20, 0x06, 0x8e, 0x7c, 0x1d, 0xb5, 0x86,
	0xd6, 0xb8, 0xeb, 0x0f, 0xb0, 0x6e, 0xc6, 0xb7, 0x66, 0xbc, 0xbb, 0x8d, 0x16, 0x2a, 0x07, 0xff,
	0x03, 0x9b, 0x72, 0x81, 0x6c, 0xc5, 0x7f, 0xdd, 0xf1, 0xe0, 0x3a, 0x68, 0x28, 0x15, 0xc4, 0xc0,
	0xce, 0x89, 0x40, 0x8e, 0xc2, 0xbf, 0xef, 0x31, 0x6b, 0x92, 0x9c, 0x46, 0x24, 0x6f, 0x68, 0x28,
	0xa1, 0xf2, 0xac, 0x44, 0x6d, 0x23, 0xcf, 0x4a, 0xe9, 0x49, 0x2e, 0x90, 0x6b, 0xe2, 0x49, 0xae,
	0xfa, 0x29, 0x3f, 0xa3, 0xaf, 0x26, 0x9e, 0xf2, 0xb3, 0xf6, 0x17, 0xf4, 0xcd, 0xcc, 0x5f, 0xb4,
	0x3f, 0xa1, 0x8e, 0x99, 0x3f, 0x41, 0x1f, 0xb4, 0x45, 0x45, 0xd2, 0x17, 0x04, 0x0c, 0x12, 0x9a,
	0xca, 0x4c, 0xcd, 0x29, 0xcd, 0x50, 0xd7, 0x24, 0xa3, 0xa8, 0xcc, 0xa4, 0xf9, 0xb1, 0x48, 0xd0,
	0x77, 0x93, 0x8c, 0xa2, 0x7a, 0x97, 0x0c, 0xf5, 0xcc, 0x76, 0xc9, 0xb4, 0xaf, 0x51, 0xdf, 0xcc,
	0xd7, 0xda, 0xa7, 0xe8, 0x87, 0x99, 0x4f, 0x47, 0x0f, 0xc0, 0x91, 0xbf, 0x2d, 0xfc, 0x09, 0x7a,
	0xeb, 0x4d, 0xbc, 0xda, 0x04, 0xf3, 0x38, 0x9a, 0x3d, 0x3e, 0xcd, 0xbd, 0x2f, 0x10, 0x00, 0x77,
	0xbd, 0x89, 0x17, 0xcb, 0xbd, 0x67, 0xc9, 0xf3, 0x62, 0xb9, 0x8f, 0xfd, 0xc0, 0x6b, 0xdd, 0xce,
	0xd3, 0xc0, 0xb3, 0x13, 0x57, 0x35, 0x4f, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x8f, 0x43,
	0x96, 0x8c, 0x03, 0x00, 0x00,
}
