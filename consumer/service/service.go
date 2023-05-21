package service

import (
	"consumer/config"
	"consumer/database"
	"consumer/logger"
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
	split := strings.Split(s.msgq, "#")
	if len(split) == 0 {
		return fmt.Errorf("%s", "recuice from mq not found")
	}
	s.env = split[0]
	logger.DebugLn(s.env)
	s.table = split[1]
	logger.DebugLn(s.table)
	s.value = split[2]
	logger.DebugLn(s.value)
	return nil
}

func (s *Service) setKey() error {
	// ^UTBL("aaa","bbb")=1|2|3|4
	begin := strings.Index(s.value, "(")
	end := strings.Index(s.value, "=")
	sub := s.value[begin+1 : end-1]
	del := strings.ReplaceAll(sub, "\"", "")
	s.key = strings.ReplaceAll(del, ",", "-")
	logger.DebugLn(s.key)
	return nil
}

func (s *Service) cmpKey() error {
	// first in same
	var same models.Same
	result := database.DB.Where("gtm_table = ? and key = ?", s.table, s.key).Find(&same)
	if result.RowsAffected <= 0 {
		//insert diff
		tmp := models.Diff{
			GtmEnv:   s.env,
			GtmTable: s.table,
			Key:      s.key,
			Value:    s.value,
		}
		if err := database.DB.Create(&tmp).Error; err != nil {
			return err
		}

	}
	return nil
}

func (s *Service) save() error {
	var t models.Same
	result := database.DB.Where("gtm_table = ? and key = ?", s.table, s.key).Find(&t)
	if result.RowsAffected <= 0 {
		// insert
		tmp := models.Same{
			GtmTable: s.table,
			Key:      s.key,
			Value:    s.value,
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
	msgDebug := fmt.Sprintf("read config scmp[%s] cmps[%s] env[%s]", scmp, cmps, s.env)
	logger.DebugLn(msgDebug)
	if s.env == scmp {
		if err := s.save(); err != nil {
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
