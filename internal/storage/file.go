package storage

import (
	"encoding/json"
	"fmt"
	"os"
	models "tidybeaver/pkg/models"
)

func WriteSampleLogsToFile(sampleLogs models.SampleLogs) {
	encodedSampleLogs, err := json.Marshal(sampleLogs)
	if err != nil {
		return
	} else {
		os.WriteFile((`.\Logs\TidyBeaverSampleLogs.json`), encodedSampleLogs, 0644)
		fmt.Println("Sample Logs saved as Json")
	}
}

func WriteLogsToFile(Logs interface{}) {
	encodedLogs, err := json.Marshal(Logs)
	if err != nil {
		return
	} else {
		os.WriteFile((`.\Logs\TidyBeaverAdaptedLogs.json`), encodedLogs, 0644)
		fmt.Println("Logs saved as Json")
	}
}
