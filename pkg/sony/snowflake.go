package main

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake // 类似于node
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(startTime string, machineID uint16) (err error) {
	sonyMachineID = machineID
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GenID 生成id
func GenID() (id uint64, err error) {
	if err != nil {
		fmt.Errorf("sonyflake init failed, err:%v\n", err)
		return
	}
	id, err = sonyFlake.NextID()
	return
}

func main() {
	if err := Init("2022-05-03", 1); err != nil {
		fmt.Printf("Init failed, err:%v\n", err)
		return
	}
	id, _ := GenID()
	fmt.Println(id)
}
