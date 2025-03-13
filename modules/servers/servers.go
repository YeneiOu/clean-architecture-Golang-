package servers

import (
	"bank/configs"
	"bank/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"log"
)

type Server struct {
	App *fiber.App
	Cfg *configs.Configs
	Db  *sqlx.DB
	//RedisClient *redis.Client
}

func NewServer(cfg *configs.Configs, Db *sqlx.DB) *Server {
	return &Server{
		App: fiber.New(),
		Cfg: cfg,
		Db:  Db,
	}
}

func (s *Server) Start() {
	if err := s.MapHandler(); err != nil {
		panic(err)
	}

	//	Connect Fiber
	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
