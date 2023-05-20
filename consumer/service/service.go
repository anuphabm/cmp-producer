package service

import (
	"consumer/config"
	"consumer/database"
	"consumer/models"
	"fmt"
	"os"
	"strings"
)

type Service struct {
	msgq  string
	env   string
	table string
	value string
	key   string
}

func (s *Service) setEnvTable() error {
	split := strings.Split(s.msgq, "|")
	if len(split) == 0 {
		return fmt.Errorf("%s", "recuice from mq not found")
	}
	s.env = split[0]
	s.table = split[1]
	s.value = split[2]
	return nil
}

func (s *Service) setKey() error {
	// ^UTBL("aaa","bbb")=1|2|3|4
	begin := strings.Index(s.value, "(")
	end := strings.Index(s.value, "=")
	sub := s.value[begin+1 : end-2]
	del := strings.ReplaceAll(sub, "\"", "")
	s.key = strings.ReplaceAll(del, ",", "-")
	return nil
}

func (s *Service) cmpKey() error {
	// first in same
	var same models.Same
	result := database.DB.Where("table = ?", s.table).Find(&same)
	if result.RowsAffected <= 0 {
		//insert diff
		tmp := models.Diff{
			Env:    s.env,
			Table:  s.table,
			FlgHas: false,
			Key:    s.key,
			Value:  s.value,
		}
		if err := database.DB.Create(&tmp).Error; err != nil {
			return err
		}

	}
	return nil
}

func (s *Service) Process() error {
	// set env and table
	if err := s.setEnvTable(); err != nil {
		return err
	}
	// set key
	if err := s.setKey(); err != nil {
		return err
	}

	// insert or compare
	envMode := os.Getenv("RUN_MODE")
	scmp := config.Appconfig.GetString(fmt.Sprintf("%s.handlers.scmp", envMode)) // insert only
	cmps := config.Appconfig.GetString(fmt.Sprintf("%s.handlers.cmps", envMode)) // compare
	if s.env == scmp {
		// insert
		tmp := models.Same{
			Table: s.table,
			Key:   s.key,
			Value: s.value,
		}
		if err := database.DB.Create(&tmp).Error; err != nil {
			return err
		}
	}

	if s.env == cmps {
		// compare
		if err := s.cmpKey(); err != nil {
			return err
		}
	}
	return nil
}

func NewService(msgq string) *Service {
	return &Service{
		msgq: msgq,
	}
}
