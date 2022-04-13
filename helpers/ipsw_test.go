package helpers

import (
	"fmt"
	"os"
	"testing"

	"github.com/groob/plist"
	"github.com/stretchr/testify/assert"
)

// // type Version map[]

// type DownloadVersions map[string]interface{}

// type SoftwareVersion map[string]DownloadVersions

// type SoftwareVersions map[string]SoftwareVersion

// type SoftwareVersionsByVersion map[string]SoftwareVersions

func TestGetIpswData(t *testing.T) {
	f := "/tmp/com_apple_macOSIPSW.xml"

	content, err := os.Open(f)
	assert.NoError(t, err)

	var data struct {
		ProductTypes              map[string]map[string]string                 `plist:"MobileDeviceProductTypes"`
		SoftwareVersionsByVersion map[string]map[string]map[string]interface{} `plist:"MobileDeviceSoftwareVersionsByVersion"`
	}

	err = plist.NewXMLDecoder(content).Decode(&data)
	assert.NoError(t, err)

	fmt.Printf("%+v", data)

	for typeName, values := range data.ProductTypes {
		for id, value := range values {
			fmt.Printf("typeName: %s id: %s value: %s\n", typeName, id, value)
		}
	}

	// for version, software := range data.SoftwareVersionsByVersion {
	// 	for model, := range software {

	// 	}
	// }
}
