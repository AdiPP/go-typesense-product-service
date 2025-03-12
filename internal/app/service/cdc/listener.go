package cdc

import (
	"context"
	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgproto3"
	"log"
	"time"
)

func (s *Service) listenForChanges(ctx context.Context, conn *pgconn.PgConn) {
	var msgPointer pglogrepl.LSN
	pluginArgs := []string{"proto_version '1'", "publication_names 'pub'"}

	err := pglogrepl.StartReplication(ctx, conn, slotName, msgPointer, pglogrepl.StartReplicationOptions{PluginArgs: pluginArgs})
	if err != nil {
		log.Fatalf("failed to start replication: %v", err)
	}

	pingTime := time.Now().Add(10 * time.Second)
	for {
		if ctx.Err() != nil {
			break
		}

		if time.Now().After(pingTime) {
			err := pglogrepl.SendStandbyStatusUpdate(ctx, conn, pglogrepl.StandbyStatusUpdate{
				WALWritePosition: msgPointer,
			})
			if err != nil {
				log.Fatalf("Failed to send standby update: %v", err)
			}
			pingTime = time.Now().Add(10 * time.Second)
		}

		recvCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		msg, err := conn.ReceiveMessage(recvCtx)
		cancel()

		if err != nil {
			if ctx.Err() != nil {
				log.Fatalf("Error receiving replication message: %v", err)
			}
			continue
		}

		switch msg := msg.(type) {
		case *pgproto3.CopyData:
			s.handleReplicationMessage(msg.Data)
		default:
			log.Printf("Unexpected message type: %T", msg)
		}
	}
}
