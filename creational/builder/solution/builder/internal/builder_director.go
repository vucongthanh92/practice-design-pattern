package internal

type SvcDirector interface {
	BuilderService(builder Builder) Service
}

type svcBuilderDirector struct {
}

func (sbd svcBuilderDirector) BuilderService(builder Builder) Service {
	builder.reset()
	builder.setName("complex Service")
	builder.buildLogger(StdLogger{})
	builder.buildNotifier(EmailNotify{})
	builder.buildDataLayer(MysqlDB{})
	builder.buildUploader(AwsUploader{})

	return builder.result()
}

func NewDirector() svcBuilderDirector {
	return svcBuilderDirector{}
}
