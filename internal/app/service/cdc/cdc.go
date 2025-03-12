package cdc

import (
	"context"
	"fmt"
	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service/sync"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
	"os"
	"os/signal"
)

type Service struct {
	cfg       *config.Config
	psService *sync.Service
}

func NewService(cfg *config.Config, psService *sync.Service) *Service {
	return &Service{cfg: cfg, psService: psService}
}

func (s *Service) StartCDC() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	conn, err := pgconn.Connect(ctx, s.cfg.GetPgsqlConnStringCdc())
	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect to pgsql cdc database: %w", err))
	}
	defer conn.Close(ctx)

	log.Print("Success starting cdc")

	s.initConfig(ctx, conn)
	s.listenForChanges(ctx, conn)
	<-ctx.Done()
}

func (s *Service) initConfig(ctx context.Context, conn *pgconn.PgConn) {
	if _, err := conn.Exec(ctx, "DROP PUBLICATION IF EXISTS pub;").ReadAll(); err != nil {
		log.Fatal(fmt.Errorf("failed to drop publication: %v", err))
	}

	if _, err := conn.Exec(ctx, "CREATE PUBLICATION pub FOR TABLE rns_product;").ReadAll(); err != nil {
		log.Fatal(fmt.Errorf("failed to create publication: %v", err))
	}

	result, err := conn.Exec(ctx, fmt.Sprintf("SELECT 1 FROM pg_replication_slots WHERE slot_name = '%s'", slotName)).ReadAll()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to query replication slots: %w", err))
	}

	if len(result) > 0 && len(result[0].Rows) > 0 {
		if err := pglogrepl.DropReplicationSlot(ctx, conn, slotName, pglogrepl.DropReplicationSlotOptions{
			Wait: false,
		}); err != nil {
			log.Fatal(fmt.Errorf("failed to drop a replication slot: %v", err))
		}
	}

	if _, err := pglogrepl.CreateReplicationSlot(ctx, conn, slotName, outputPlugin, pglogrepl.CreateReplicationSlotOptions{Temporary: false}); err != nil {
		log.Fatal(fmt.Errorf("failed to create a replication slot: %v", err))
	}
}
