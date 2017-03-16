//go:generate protoc --go_out=plugins=grpc:. receiver.proto gps.proto modes.proto mrn.proto
package mrn
