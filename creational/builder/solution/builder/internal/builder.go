package internal

type Service interface {
	DoBusiness()
}

type Builder interface {
	reset()
	setName(name string)
	buildLogger(logger Logger)
	buildNotifier(notifier Notifier)
	buildDataLayer(db DataLayer)
	buildUploader(uploader Uploader)
	result() Service
}

type svcBuilder struct {
	service *complexSvc
}

func NewBuilder() *svcBuilder {
	return &svcBuilder{}
}

func (b *svcBuilder) reset() {
	b.service = &complexSvc{}
}

func (b *svcBuilder) setName(name string) {
	b.service.setName(name)
}

func (b *svcBuilder) buildLogger(log Logger) {
	if log == nil {
		log = StdLogger{}
	}

	b.service.setLogger(log)
}

func (b *svcBuilder) buildNotifier(notify Notifier) {
	if notify == nil {
		notify = SMSNotify{}
	}

	b.service.setNotifier(notify)
}

func (b *svcBuilder) buildDataLayer(db DataLayer) {
	if db == nil {
		db = MysqlDB{}
	}
	b.service.setDatalayer(db)
}

func (b *svcBuilder) buildUploader(uploader Uploader) {
	if uploader == nil {
		uploader = AwsUploader{}
	}
	b.service.setUploader(uploader)
}

func (b *svcBuilder) result() Service {
	return b.service
}
