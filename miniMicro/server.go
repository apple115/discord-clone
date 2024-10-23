package minimicro

type service struct{
	opts Options
	once sync.Once
}

func newService(opts ...Option) Service {
	return &service{
		opts: options,
	}
}

func (s *service) Name() string {
	return s.opts.Server.Options().Name
}

func (s *service) Init(opts ...Option) {
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) Client() client.Client {
	return s.opts.Client
}

func (s *service) Server() server.Server {
	return s.opts.Server
}

func (s *service) String() string {
	return "micro"
}

func (s *service) Start() error {
	return nil
}

func (s *service) Stop() error {
	return err
}

func (s *service) Run() (err error) {
	return s.Stop()
}
