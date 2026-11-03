package main

import (
	"context"
	"fmt"
	"strings"
	"os"
	"os/signal"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func isGzip(data []byte) bool {
    if len(data) < 2 {
        return false
    }
    return data[0] == 0x1f && data[1] == 0x8b
}

func printCharacteristicValue(serviceUuid string, charUuid string, value []byte) {
	isgz := isGzip(value)
	fmt.Printf("Value from %s/%s: %v - %x\n", serviceUuid, charUuid, isgz, value)
}

func monitorCharacteristic(serviceUuid string, charUuid string) func ([]byte) {
	return func (buf []byte) {
		printCharacteristicValue(serviceUuid, charUuid, buf)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	must("enable BLE stack", adapter.Enable())


	println("scanning...")
	var found bluetooth.ScanResult
	adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if strings.HasPrefix(device.LocalName(), "RepOne") {
			found = device
			adapter.StopScan()
		}
	})

	device, err := adapter.Connect(found.Address, bluetooth.ConnectionParams{})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer device.Disconnect()

	services, err := device.DiscoverServices(nil)
	if err != nil {
		panic("failed to discover services: " + err.Error())
	}

	for _, service := range services {
		fmt.Printf("Found service: %s\n", service.UUID().String())
		characteristics, err := service.DiscoverCharacteristics(nil)
		if err != nil {
			panic("failed to discover characteristics: " + err.Error())
		}
		for _, characteristic := range characteristics {
			fmt.Printf("\tFound characteristic: %s\n", characteristic.UUID().String())
			characteristic.EnableNotifications(monitorCharacteristic(
				service.UUID().String(),
				characteristic.UUID().String(),
			))
		}
	}

	select {
	case <-ctx.Done():
		fmt.Println("Exiting...")
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
