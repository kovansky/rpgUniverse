package config

import "encoding/json"

type Config struct {
	Global Global `json:"global"`
}
