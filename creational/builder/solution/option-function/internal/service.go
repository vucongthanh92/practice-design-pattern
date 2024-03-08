package internal

import "log"

type Logger interface {
	Log(...any)
}

// logger service

type StdLogger struct{}

func (StdLogger) Log(...any) {}

type FileLogger struct{}

func (FileLogger) Log(...any) {}

// notifier service

type Notifier interface {
	Send(msg string)
}

type EmailNotify struct{}

func (EmailNotify) Send(msg string) {}

type SMSNotifier struct{}

func (SMSNotifier) Send(msg string) {}

// datalayer service

type Datalayer interface {
	Save()
}

type MysqlDB struct{}

func (MysqlDB) Save() {}

type MongoDB struct{}

func (MongoDB) Save() {}

// uploader service

type Uploader interface {
	Upload()
}

type AwsS3Uploader struct{}

func (AwsS3Uploader) Upload() {}

type GgDriveUploader struct{}

func (GgDriveUploader) Upload() {}

// implement sub service into complex service
type complexSvc struct {
	name      string
	logger    Logger
	notifier  Notifier
	DataLayer Datalayer
	uploader  Uploader
}

func (s complexSvc) DoBusiness() {
	s.logger.Log(s.name)
	s.uploader.Upload()
	s.DataLayer.Save()
	s.notifier.Send("hello world")

	log.Println("complex service do business normally")
}

type Option func(*complexSvc)

func NewService(opts ...Option) complexSvc {
	service := complexSvc{
		name:      "Service",
		logger:    StdLogger{},
		notifier:  SMSNotifier{},
		DataLayer: MongoDB{},
		uploader:  AwsS3Uploader{},
	}

	for i := range opts {
		opts[i](&service)
	}

	return service
}
