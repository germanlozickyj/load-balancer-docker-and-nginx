package server

type Config struct {
	Port string
	JWTSecret string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
} 

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, error.New("Port is required")
	}
	if config.JWTSecret == "" {
		return nil, error.New("JWTSecret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, error.New("DatabaseUrl is required")
	}
	
	broker := &Broker {
		config: config,
		router: mux.NewRouter()
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
}