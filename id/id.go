package id

import (
	"crypto/md5"
	"encoding/binary"
	"github.com/sony/sonyflake"
	"os"
)

func machineID() (uint16, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return 0, err
	}
	sum := md5.Sum([]byte(hostname))
	machineID := binary.BigEndian.Uint16(sum[len(sum)-2:]) // extract last 2 bytes
	return uint16(machineID), nil
}

type ID interface {
	Generate() (int64, error)
	Must() int64
}

type SonyflakeID struct {
	ID

	sf *sonyflake.Sonyflake
}

func New() *SonyflakeID {
	st := sonyflake.Settings{
		MachineID: machineID,
	}
	sf := sonyflake.NewSonyflake(st)

	return &SonyflakeID{
		sf: sf,
	}
}

func (i *SonyflakeID) Generate() (int64, error) {
	id, err := i.sf.NextID()
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func (i *SonyflakeID) Must() int64 {
	id, _ := i.Generate()
	return id
}
