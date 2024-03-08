package internal

func WithName(name string) Option {
	return func(s *complexSvc) {
		s.name = name
	}
}

func WithStdLogger() Option {
	return func(s *complexSvc) {
		s.logger = StdLogger{}
	}
}

func WithFileLogger() Option {
	return func(s *complexSvc) {
		s.logger = FileLogger{}
	}
}

func WithCustomLogger(log Logger) Option {
	return func(s *complexSvc) {
		if log == nil {
			WithStdLogger()
			return
		}
		s.logger = log
	}
}

func WithMysqlDB() Option {
	return func(s *complexSvc) {
		s.DataLayer = MysqlDB{}
	}
}

func WithCustomDB(db Datalayer) Option {
	return func(s *complexSvc) {
		if db == nil {
			WithMysqlDB()
			return
		}
		s.DataLayer = db
	}
}

func WithEmailNotifier() Option {
	return func(s *complexSvc) {
		s.notifier = EmailNotify{}
	}
}

func WithAwsS3Uploader() Option {
	return func(s *complexSvc) {
		s.uploader = AwsS3Uploader{}
	}
}

func WithCustomUploader(uploader Uploader) Option {
	return func(s *complexSvc) {
		if uploader == nil {
			WithAwsS3Uploader()
			return
		}
		s.uploader = uploader
	}
}
