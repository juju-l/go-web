package repository

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"go-web/internal/domain/dto"
	"go-web/util/ext"
	// ...
)

// var db *gorm.DB = nil

type PsqlRepository struct { Orm *gorm.DB }

func (p *PsqlRepository) New() (*PsqlRepository, error) {
	// var p = &PsqlRepository{}
	cfg := &dto.ConfigDto{}; yaml.Unmarshal(ext.Must(ioutil.ReadFile("config/app_configs.yml")).([]byte), cfg); // get config .
	p.Orm = ext.Must(gorm.Open(postgres.New(postgres.Config{ DSN: cfg.DbConnStr/*, PreferSimpleProtocol: false, WithoutReturning: false*/ }), &gorm.Config{ NamingStrategy: schema.NamingStrategy{ SingularTable: true } })).(*gorm.DB).
	Clauses(clause.Returning{}).Debug() // add 'RETURNING *'
	return p, nil
}

// func Add() bool {
// 	return true
// }

// func Delete() bool {
// 	return true
// }

// func Update() bool {
// 	return true
// }

// func Select() []interface{} {
// 	return nil
// }

// func (p *PsqlRepository) GetQueryableByPage() interface{} { // 分页处理...
// 	return new(interface{})
// }

//
///
/*
 *
 */
//
// func ConvertPages() interface{} {
// 	return new(interface{})
// }

// func f() {}

func init() {
	//
}