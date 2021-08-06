# terraform-provider-uptends

Register a new API Account
curl -X POST "https://api.uptrends.com/v4/Register" -H  "accept: application/json"


{
    "MonitorGuid": "e4751047-7a43-4b00-be88-acf3a2d9a2e8",
    "Hash": "P4yqf7Ne3DBcYZQp/2CT7Q==",
    "Name": "acctest-uptrends-monitor-ping-549744",
    "IsActive": true,
    "GenerateAlert": true,
    "IsLocked": false,
    "CheckInterval": 5,
    "MonitorMode": "Production",
    "CustomFields": [],
    "SelectedCheckpoints": {},
    "UsePrimaryCheckpointsOnly": true,
    "MonitorType": "Ping",
    "Notes": "",
    "AlertOnLoadTimeLimit1": true,
    "LoadTimeLimit1": 2500,
    "AlertOnLoadTimeLimit2": true,
    "LoadTimeLimit2": 5000,
    "IpVersion": "IpV4",
    "NativeIPv6Only": false,
    "NetworkAddress": "8.8.8.8"
}