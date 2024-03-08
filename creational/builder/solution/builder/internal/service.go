package internal

import "log"

type Logger interface {
	Log(...any)
}

type StdLogger struct{}

func (StdLogger) Log(...any) {}

type FileLogger struct{}

func (FileLogger) Log(...any) {}

type Notifier interface {
	Send(msg string)
}

type EmailNotify struct{}

func (EmailNotify) Send(msg string) {}

type SMSNotify struct{}

func (SMSNotify) Send(msg string) {}

type DataLayer interface {
	Save()
}

type MysqlDB struct{}

func (MysqlDB) Save() {}

type MongoDB struct{}

func (MongoDB) Save() {}

type Uploader interface {
	Upload()
}

type AwsUploader struct{}

func (AwsUploader) Upload() {}

type GoogleDriveUploader struct{}

func (GoogleDriveUploader) Upload() {}

type complexSvc struct {
	name      string
	logger    Logger
	notifier  Notifier
	DataLayer DataLayer
	uploader  Uploader
}

func (s *complexSvc) setName(name string) {
	s.name = name
}

func (s *complexSvc) setLogger(log Logger) {
	s.logger = log
}

func (s *complexSvc) setNotifier(n Notifier) {
	s.notifier = n
}

func (s *complexSvc) setDatalayer(db DataLayer) {
	s.DataLayer = db
}

func (s *complexSvc) setUploader(u Uploader) {
	s.uploader = u
}

func (s *complexSvc) DoBusiness() {
	s.logger.Log(s.name)
	s.uploader.Upload()
	s.DataLayer.Save()
	s.notifier.Send("hello world!")

	log.Println("service do business normally")
}
