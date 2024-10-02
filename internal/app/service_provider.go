package app

import (
	"context"
	"log"

	"github.com/BelyaevEI/tg-zametker/internal/config"
	"github.com/BelyaevEI/tg-zametker/internal/repository"
	"github.com/BelyaevEI/tg-zametker/internal/service"

	"github.com/BelyaevEI/platform_common/pkg/closer"
	"github.com/BelyaevEI/platform_common/pkg/db"
	"github.com/BelyaevEI/platform_common/pkg/db/pg"
)

type serviceProvider struct {
	tokenConfig     config.TokenConfig
	debugModeConfig config.DebugConfig
	pgConfig        config.PGConfig

	service    service.Servicer
	repository repository.Repositorer

	dbClient db.Client
}

// Создание сервис провайдера
func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Создание обьекта токена
func (s *serviceProvider) TokenConfig() config.TokenConfig {
	if s.tokenConfig == nil {
		cfg, err := config.NewTokenConfig()
		if err != nil {
			log.Fatalf("failed to get token config: %s", err.Error())
		}

		s.tokenConfig = cfg
	}

	return s.tokenConfig
}

// Режим работы бота
func (s *serviceProvider) DebugConfig() config.DebugConfig {
	if s.debugModeConfig == nil {
		cfg, err := config.NewDebugModeConfig()
		if err != nil {
			log.Fatalf("failed to get debug mode config: %s", err.Error())
		}

		s.debugModeConfig = cfg
	}

	return s.debugModeConfig
}

func (s *serviceProvider) ImplementationApp(ctx context.Context) {
	_ = s.newService(ctx)
}

func (s *serviceProvider) newService(ctx context.Context) service.Servicer {
	if s.service == nil {
		s.service = service.NewService(s.newRepository(ctx))
	}

	return s.service
}

func (s *serviceProvider) newRepository(ctx context.Context) repository.Repositorer {
	if s.repository == nil {
		s.repository = repository.NewRepository(s.clientDB(ctx))
	}

	return s.repository
}

func (s *serviceProvider) clientDB(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.configPG().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) configPG() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}
