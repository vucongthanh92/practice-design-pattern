package main

import "design-pattern/structural/bridge/solution/internal"

func main() {
	_ = internal.ParseAndSaveData(internal.MysqlParser{}, internal.JSONFilePersistent{})
	_ = internal.ParseAndSaveData(internal.MongoParser{}, internal.RPCServicePersistent{})
	_ = internal.ParseAndSaveData(internal.FileParser{}, internal.AWSS3Persistent{})
}
