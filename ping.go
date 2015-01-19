package main

import (
	"code.google.com/p/go-uuid/uuid"
)

type Ping struct {
	ID int `db:"ping_id"`
	Timestamp time.Time `db:"ping_time"`
	NextPing time.Time `db:"next_ping"`
	SubmitterIP string `db:"submitter_ip"`
	ServerID UUID.uuid `db:"server_id"`
	ServerVersion string `db:"server_version"`
	OSName string `db:"os_name"`
	ProcessorType string `db:"processor_type"`
	ProcessorCores int16 `db:"processor_cores"`
	RAMSize int `db:"ram_amount"` // Counted in megabytes
	Uptime int `db::"uptime"` // Measured in ticks.
	PlayersOnline int `db:"players:online"`
	PlayerCap int `db:"players_cap"`
	PluginList string `db:"plugin_list"`
	locale string `db:"locale"`
}
