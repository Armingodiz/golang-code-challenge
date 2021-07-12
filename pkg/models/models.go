package models

type BrokerData struct {
	UserName     string `json:"user_name"`
	ID           int    `json:"id"`
	TrafficUsage int    `json:"traffic_usage"`
	Ip           string `json:"ip"`
	Port         string `json:"port"`
}