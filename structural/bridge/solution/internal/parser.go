package internal

import (
	"fmt"
)

type DataParser interface {
	Parse() (*Data, error)
}

type MysqlParser struct{}

func (MysqlParser) Parse() (*Data, error) {
	fmt.Println("Mysql parse data success")
	return &Data{}, nil
}

type MongoParser struct{}

func (MongoParser) Parse() (*Data, error) {
	fmt.Println("Mongo parse data success")
	return &Data{}, nil
}

type FileParser struct{}

func (FileParser) Parse() (*Data, error) {
	fmt.Println("File parse data success")
	return &Data{}, nil
}
