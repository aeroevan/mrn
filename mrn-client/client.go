package main

import "fmt"
import "net"
import "flag"
import "bufio"
import "strings"
import "strconv"
import mrn "github.com/aeroevan/mrn/proto"

//import proto "github.com/golang/protobuf/proto"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import ptw "github.com/golang/protobuf/ptypes/wrappers"

func adsb(connstr string) <-chan string {
	out := make(chan string)
	con, _ := net.Dial("tcp", connstr)
	reader := bufio.NewScanner(con)
	go func() {
		for {
			if ok := reader.Scan(); !ok {
				fmt.Println("Error reading from dump1090 socket")
				close(out)
				break
			}
			out <- reader.Text()
		}
	}()
	return out
}

func parse_mode_s(msg string) *mrn.ModeSMessage {
	out := new(mrn.ModeSMessage)
	cols := strings.Split(msg, ",")
	ts := strings.Split(cols[0], ".")
	sec, errs := strconv.Atoi(ts[0])
	nano, errn := strconv.Atoi(ts[1])
	if errs == nil && errn == nil {
		out.Time = &timestamp.Timestamp{
			Seconds: int64(sec),
			Nanos:   int32(nano),
		}
	}

	out.Addr = cols[1]

	addr_type := cols[2]
	out.AddrType = mode_s_addr_type(addr_type)

	out.Msg = cols[3]
	out.SignalLevel = parse_double_value(cols[4]).Value
	out.Latitude = parse_double_value(cols[5])
	out.Longitude = parse_double_value(cols[6])
	out.Altitude = parse_int32_value(cols[7])
	out.AltitudeUnit = mode_s_altitude_unit(cols[8])
	out.AltitudeType = mode_s_altitude_type(cols[9])
	out.GnssDelta = parse_int32_value(cols[10])
	out.Heading = parse_int32_value(cols[11])
	out.HeadingType = mode_s_heading_type(cols[12])
	out.Speed = parse_int32_value(cols[13])
	out.SpeedType = mode_s_speed_type(cols[14])
	out.VertRate = parse_int32_value(cols[15])
	out.VertRateSource = mode_s_altitude_type(cols[16])
	out.Sqawk = parse_string_value(cols[17])
	out.Callsign = parse_string_value(cols[18])
	return out
}

func parse_string_value(val string) *ptw.StringValue {
	if len(val) > 0 {
		return &ptw.StringValue{Value: val}
	}
	return nil
}

func parse_int32_value(val string) *ptw.Int32Value {
	i, err := strconv.ParseInt(val, 10, 32)
	if err == nil {
		return &ptw.Int32Value{Value: int32(i)}
	}
	return nil
}

func parse_double_value(val string) *ptw.DoubleValue {
	f, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return &ptw.DoubleValue{Value: f}
	}
	return nil
}

func mode_s_air_ground(air_ground string) mrn.ModeSMessage_AirGround {
	switch air_ground {
	case "I":
		return mrn.ModeSMessage_INVALID
	case "G":
		return mrn.ModeSMessage_GROUND
	case "A":
		return mrn.ModeSMessage_AIRBORNE
	case "U":
		return mrn.ModeSMessage_UNCERTAIN
	default:
		return mrn.ModeSMessage_UNCERTAIN
	}
}

func mode_s_speed_type(speed_type string) mrn.ModeSMessage_SpeedType {
	switch speed_type {
	case "GS":
		return mrn.ModeSMessage_GROUNDSPEED
	case "IAS":
		return mrn.ModeSMessage_IAS
	case "TAS":
		return mrn.ModeSMessage_TAS
	default:
		return mrn.ModeSMessage_NONE
	}
}

func mode_s_heading_type(heading_type string) mrn.ModeSMessage_HeadingType {
	switch heading_type {
	case "TRU":
		return mrn.ModeSMessage_TRUE
	case "MAG":
		return mrn.ModeSMessage_MAGNETIC
	default:
		return mrn.ModeSMessage_NO_HEADING_TYPE
	}
}

func mode_s_altitude_type(altitude_type string) mrn.ModeSMessage_AltitudeType {
	switch altitude_type {
	case "BARO":
		return mrn.ModeSMessage_BAROMETRIC
	case "GNSS":
		return mrn.ModeSMessage_GNSS
	default:
		return mrn.ModeSMessage_NO_ALTITUDE_TYPE
	}
}

func mode_s_altitude_unit(altitude_unit string) mrn.ModeSMessage_AltitudeUnit {
	switch altitude_unit {
	case "f":
		return mrn.ModeSMessage_FEET
	case "m":
		return mrn.ModeSMessage_METERS
	default:
		return mrn.ModeSMessage_NO_ALTITUDE
	}
}

func mode_s_addr_type(addr_type string) mrn.ModeSMessage_AddrType {
	switch addr_type {
	case "UNKNOWN":
		return mrn.ModeSMessage_UNKNOWN
	case "ADSB_ICAO":
		return mrn.ModeSMessage_ADSB_ICAO
	case "ADSB_ICAO_NT":
		return mrn.ModeSMessage_ADSB_ICAO_NT
	case "ADSR_ICAO":
		return mrn.ModeSMessage_ADSR_ICAO
	case "TISB_ICAO":
		return mrn.ModeSMessage_TISB_ICAO
	case "ADSB_OTHER":
		return mrn.ModeSMessage_ADSB_OTHER
	case "ADSR_OTHER":
		return mrn.ModeSMessage_ADSR_OTHER
	case "TISB_TRACKFILE":
		return mrn.ModeSMessage_TISB_TRACKFILE
	case "TISB_OTHER":
		return mrn.ModeSMessage_TISB_OTHER
	default:
		return mrn.ModeSMessage_UNKNOWN
	}
}

func mode_s_channel(in <-chan string) <-chan *mrn.ModeSMessage {
	out := make(chan *mrn.ModeSMessage)
	go func() {
		for msg := range in {
			msm := parse_mode_s(msg)
			out <- msm
		}
	}()
	return out
}

var (
	dump1090 = flag.String("dump1090", "localhost:30002", "The raw dump1090 host:port")
)

func main() {
	flag.Parse()
	msgs := adsb(*dump1090)
	msms := mode_s_channel(msgs)
	for msm := range msms {
		fmt.Println(msm)
	}
}
