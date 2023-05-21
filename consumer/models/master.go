package models

import (
	"consumer/config"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Master struct {
	gorm.Model
	GtmTable string `json:"gtmTable"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

func (Master) TableName() string {
	cfg := config.Appconfig
	runMode := os.Getenv("RUN_MODE")
	return fmt.Sprintf("%s.masters", cfg.GetString(fmt.Sprintf("%s.database.schema", runMode)))
}
