package main

import (
	"fmt"
	"strconv"

	"k8s.io/klog/v2"
)

func main() {
	var hexInput string
	klog.Info("Starting hex to base64 converter. Insert your hex number:")
	fmt.Scanln(&hexInput)

	int, err := strconv.ParseUint(hexInput, 16, 64)
	if err != nil {
		// s is not a valid
		klog.Error("hex input is not a valid")
	} else {
		klog.Info("Converted to intiger:", int)
	}

}
