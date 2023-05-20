package models

import (
	"consumer/config"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Same struct {
	gorm.Model
	Table string `json:"table"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (Same) TableName() string {
	cfg := config.Appconfig
	runMode := os.Getenv("RUN_MODE")
	return fmt.Sprintf("%s.sames", cfg.GetString(fmt.Sprintf("%s.database.schema", runMode)))
}
