package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func ConnectionMysql(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// 读写分离
func ConnectionGroupMysql(masterUrl, replicaUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(masterUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dbResolverCfg := dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(masterUrl)},
		Replicas: []gorm.Dialector{mysql.Open(replicaUrl)},
		Policy:   dbresolver.RandomPolicy{},
	}
	readWitePlugin := dbresolver.Register(dbResolverCfg)
	db.Use(readWitePlugin)
	return db, nil
}
