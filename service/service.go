package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-abtest/sdk"
)

var (
	// DBPath ...
	DBPath string
)

// CreateABTestConfig 创建新的实验配置
func CreateABTestConfig(projectID string, zones []*sdk.Zone) {
	fmt.Println("DBPath", DBPath)

	// 入库
	content, err := json.Marshal(zones)
	if err != nil {
		log.Fatal("CreateABTestConfig call json.Marshal failed , error:", err)
	}
	ioutil.WriteFile(DBPath+projectID, content, 0644)

	// 校验配置正确性
	if !checkProjectABTConfigValid(projectID, zones) {
		log.Fatal("checkProjectABTConfigValid fatal, projectID:", projectID)
	}
}

// CheckProjectABTConfigValid 校验实验配置正确性
func checkProjectABTConfigValid(projectID string, zones []*sdk.Zone) bool {
	layerID2ActualTotal := make(map[string]uint32)
	layerID2ExpectedTotal := make(map[string]uint32)
	for _, zone := range zones {
		if zone.ProjectID == projectID {
			layerID2ActualTotal[zone.LayerID] = layerID2ActualTotal[zone.LayerID] + (zone.Max - zone.Min) + 1
			layerID2ExpectedTotal[zone.LayerID] = zone.TotalWeight
		}
	}

	for layer, actual := range layerID2ActualTotal {
		expected, ok := layerID2ExpectedTotal[layer]
		if !ok || expected == 0 {
			expected = sdk.DefualtTotalWeight
		}
		if actual != expected {
			fmt.Println(fmt.Sprintf("projectID(%s) layer(%s) actual total weight(%d) not equal expected(%d)", projectID, layer, actual, expected))
			return false
		}
	}

	return true
}

// DescribeABTestConfig 返回 json
func DescribeABTestConfig(projectID string) []byte {
	content, err := ioutil.ReadFile(DBPath + projectID)
	if err != nil {
		log.Fatal("doSyncDB call ioutil.ReadFile failed, error:", err)
	}
	zones := make([]*sdk.Zone, 0)
	err = json.Unmarshal(content, &zones)
	if err != nil {
		log.Fatal("doSyncDB call json.Unmarshal failed, error:", err)
	}

	return content
}
