package models

import (
	"consumer/config"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type Diff struct {
	gorm.Model
	Env    string `json:"env"`
	FlgHas bool   `json:"flgHas"`
	Table  string `json:"table"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func (Diff) TableName() string {
	cfg := config.Appconfig
	runMode := os.Getenv("RUN_MODE")
	return fmt.Sprintf("%s.diffs", cfg.GetString(fmt.Sprintf("%s.database.schema", runMode)))
}
