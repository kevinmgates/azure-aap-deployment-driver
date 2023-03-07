package model
import (
	"encoding/json"
	"gorm.io/datatypes"
	_ "gorm.io/driver/sqlite"
)

// type Step struct {
// 	BaseModel
// 	Name       string            `gorm:"unique" json:"name"`
// 	Template   datatypes.JSONMap `json:"-"`
// 	Parameters datatypes.JSONMap `json:"-"`
// 	Priority   uint              `json:"order"`
// 	Executions []Execution       `json:"executions" gorm:"constraint:OnUpdate:CASCADE;"`
// }

seedData := `
[
    {
        "ID": 1,
        "name": "VNET__and_Subnets",
        "order": 0,
        "executions": [
            {
                "ID": 1,
                "status": "Succeeded",
                "stepId": 1,
                "error": "",
                "errorDetails": "",
                "code": "",
                "provisioningState": "Succeeded",
                "message": "",
                "details": "",
                "timestamp": "2022-10-11T16:06:58.0259123Z",
                "duration": "16.92 seconds",
                "correlationId": "3db406a1-3d79-4bf9-9200-930785c80d74",
                "executionCount": 1
            }
        ]
    },
    {
        "ID": 2,
        "name": "Private__DNS",
        "order": 1,
        "executions": [
            {
                "ID": 2,
                "status": "Succeeded",
                "stepId": 2,
                "error": "",
                "errorDetails": "",
                "code": "",
                "provisioningState": "Succeeded",
                "message": "",
                "details": "",
                "timestamp": "2022-10-11T16:07:21.2750564Z",
                "duration": "17.75 seconds",
                "correlationId": "1a979ebb-9395-459a-9a5b-9309fe2791fd",
                "executionCount": 1
            }
        ]
    },
    {
        "ID": 3,
        "name": "AAP__Repository",
        "order": 2,
        "executions": [
            {
                "ID": 3,
                "status": "Restarted",
                "stepId": 3,
                "error": "Failed to deploy",
                "errorDetails": "",
                "code": "ERROR_CODE",
                "provisioningState": "Failed",
                "message": "Template failed to deploy",
                "details": "",
                "timestamp": "2022-10-11T16:09:21.2750564Z",
                "duration": "1 minute 32.75 seconds",
                "correlationId": "c89732ab-9395-459a-9a5b-9309fe2791fd",
                "executionCount": 2
            },
            {
                "ID": 4,
                "status": "Failed",
                "stepId": 3,
                "error": "Failed to deploy again",
                "errorDetails": "",
                "code": "ERROR_CODE",
                "provisioningState": "Failed",
                "message": "Template failed to deploy",
                "details": "",
                "timestamp": "2022-10-11T16:13:21.2750564Z",
                "duration": "2 minute 23.98 seconds",
                "correlationId": "d9384cab-9395-459a-9a5b-9309fe2791fd",
                "executionCount": 2
            }
        ]
    },
    {
        "ID": 4,
        "name": "Database__Server__and__Databases",
        "order": 2,
        "executions": [
            {
                "ID": 5,
                "status": "Started",
                "stepId": 4,
                "error": "",
                "errorDetails": "",
                "code": "",
                "provisioningState": "",
                "message": "",
                "details": "",
                "timestamp": "2022-10-11T16:07:21.2750564Z",
                "duration": "17.75 seconds",
                "correlationId": "1a979ebb-9395-459a-9a5b-9309fe2791fd",
                "executionCount": 1
            }
        ]
    },
    {
        "ID": 5,
        "name": "AKS__Cluster",
        "order": 3,
        "executions": []
    },
    {
        "ID": 6,
        "name": "AAP__Operators",
        "order": 4,
        "executions": []
    },
    {
        "ID": 7,
        "name": "AAP_Applications",
        "order": 5,
        "executions": []
    },
    {
        "ID": 8,
        "name": "Application__Ingress",
        "order": 6,
        "executions": []
    },
    {
        "ID": 9,
        "name": "Seeded__Content",
        "order": 7,
        "executions": []
    },
    {
        "ID": 10,
        "name": "Billing",
        "order": 7,
        "executions": []
    }
]`
var mySteps []Step
json.Unmarshal([]byte(seedData), &mySteps)
fmt.Printf("mySteps : %+v", mySteps)