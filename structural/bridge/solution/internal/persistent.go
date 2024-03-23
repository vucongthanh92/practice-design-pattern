package internal

import (
	"fmt"
)

type DataPersistent interface {
	Save(*Data) error
}

type JSONFilePersistent struct{}

func (JSONFilePersistent) Save(*Data) error {
	fmt.Println("Save json data success")
	return nil
}

type RPCServicePersistent struct{}

func (RPCServicePersistent) Save(*Data) error {
	fmt.Println("Save RPC data success")
	return nil
}

type AWSS3Persistent struct{}

func (AWSS3Persistent) Save(*Data) error {
	fmt.Println("Save File data success")
	return nil
}
