package gps

import "time"

type TPV struct {
	// Fixed: "TPV"
	Class string `json:"class"`

	// Name of originating device
	Device *string `json:"device,omitempty"`

	// GPS status: %d, 2=DGPS fix, otherwise not present.
	Status *uint8 `json:"status,omitempty"`

	// NMEA mode: %d, 0=no mode value yet seen, 1=no fix, 2=2D, 3=3D.
	Mode FixMode `json:"mode"`

	// Time/date stamp in ISO8601 format, UTC. May have a fractional part of up to .001sec precision. May be absent if mode is not 2 or 3.
	Time *time.Time `json:"time,omitempty"`

	// Estimated timestamp error (%f, seconds, 95% confidence). Present if time is present.
	Ept *TpvEpt `json:"ept,omitempty"`

	// Latitude in degrees: +/- signifies North/South. Present when mode is 2 or 3.
	Lat *float64 `json:"lat,omitempty"`

	// Longitude in degrees: +/- signifies East/West. Present when mode is 2 or 3.
	Lon *float64 `json:"lon,omitempty"`

	// Altitude in meters. Present if mode is 3.
	Alt *float64 `json:"alt,omitempty"`

	// Longitude error estimate in meters, 95% confidence. Present if mode is 2 or 3 and DOPs can be calculated from the satellite view.
	Epx *float64 `json:"epx,omitempty"`

	// Latitude error estimate in meters, 95% confidence. Present if mode is 2 or 3 and DOPs can be calculated from the satellite view.
	Epy *float64 `json:"epy,omitempty"`

	// Estimated vertical error in meters, 95% confidence. Present if mode is 3 and DOPs can be calculated from the satellite view.
	Epv *float64 `json:"epv,omitempty"`

	// Course over ground, degrees from true north.
	Track *float64 `json:"track,omitempty"`

	// Speed over ground, meters per second.
	Speed *float64 `json:"speed,omitempty"`

	// Climb (positive) or sink (negative) rate, meters per second.
	Climb *float64 `json:"climb,omitempty"`

	// Direction error estimate in degrees, 95% confidence.<Paste>
	Epd *float64 `json:"epd,omitempty"`

	// Speed error estinmate in meters/sec, 95% confidence.
	Eps *float64 `json:"eps,omitempty"`

	// Climb/sink error estimate in meters/sec, 95% confidence.
	Epc *float64 `json:"epc,omitempty"`
}
