package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgproto3"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/jackc/pglogrepl"
	"github.com/jackc/pgx/v5/pgconn"
)

const SlotName = "replication_slot"
const OutputPlugin = "pgoutput"
const InsertTemplate = "create table rns_dump (id int, name text);"

var Event = struct {
	Relation string
	Columns  []string
}{}

func main() {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?replication=database",
		os.Getenv("PGSQL_DB_USERNAME"),
		os.Getenv("PGSQL_DB_PASSWORD"),
		os.Getenv("PGSQL_DB_HOST"),
		os.Getenv("PGSQL_DB_PORT"),
		os.Getenv("PGSQL_DB_DATABASE"),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	conn, err := pgconn.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	// 1. Create table
	if _, err := conn.Exec(ctx, "DROP TABLE IF EXISTS rns_dump;").ReadAll(); err != nil {
		err = fmt.Errorf("failed to drop table: %v", err)
		log.Fatal(err)
	}

	if _, err := conn.Exec(ctx, InsertTemplate).ReadAll(); err != nil {
		err = fmt.Errorf("failed to create table: %v", err)
		log.Fatal(err)
	}

	// 2. ensure publication exists
	if _, err := conn.Exec(ctx, "DROP PUBLICATION IF EXISTS pub;").ReadAll(); err != nil {
		err = fmt.Errorf("failed to drop publication: %v", err)
		log.Fatal(err)
	}

	if _, err := conn.Exec(ctx, "CREATE PUBLICATION pub FOR TABLE rns_dump;").ReadAll(); err != nil {
		err = fmt.Errorf("failed to create publication: %v", err)
		log.Fatal(err)
	}

	// 3. create temproary replication slot server
	if err := pglogrepl.DropReplicationSlot(ctx, conn, SlotName, pglogrepl.DropReplicationSlotOptions{
		Wait: false,
	}); err != nil {
		err = fmt.Errorf("failed to dorp a replication slot: %v", err)
		log.Fatal(err)
	}

	if _, err = pglogrepl.CreateReplicationSlot(
		ctx,
		conn,
		SlotName,
		OutputPlugin,
		pglogrepl.CreateReplicationSlotOptions{Temporary: false},
	); err != nil {
		err = fmt.Errorf("failed to create a replication slot: %v", err)
		log.Fatal(err)
	}

	var msgPointer pglogrepl.LSN
	pluginArguments := []string{"proto_version '1'", "publication_names 'pub'"}

	// 4. establish connection
	err = pglogrepl.StartReplication(
		ctx,
		conn,
		SlotName,
		msgPointer,
		pglogrepl.StartReplicationOptions{PluginArgs: pluginArguments},
	)
	if err != nil {
		err = fmt.Errorf("failed to establish start replication: %v", err)
		log.Fatal(err)
	}

	pingTime := time.Now().Add(10 * time.Second)
	for !errors.Is(ctx.Err(), context.Canceled) {
		if time.Now().After(pingTime) {
			if err = pglogrepl.SendStandbyStatusUpdate(
				ctx,
				conn,
				pglogrepl.StandbyStatusUpdate{WALWritePosition: msgPointer},
			); err != nil {
				err = fmt.Errorf("failed to send standby update: %v", err)
				log.Fatal(err)
			}
			pingTime = time.Now().Add(10 * time.Second)
			//fmt.Println("client: please standby")
		}

		ctx2, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		msg, err := conn.ReceiveMessage(ctx2)
		if pgconn.Timeout(err) {
			continue
		}
		if err != nil {
			err = fmt.Errorf("something went wrong while listening for message: %v", err)
			log.Fatal(err)
		}

		switch msg := msg.(type) {
		case *pgproto3.CopyData:
			switch msg.Data[0] {
			case pglogrepl.PrimaryKeepaliveMessageByteID:
				fmt.Println("server: confirmed standby")

			case pglogrepl.XLogDataByteID:
				walLog, err := pglogrepl.ParseXLogData(msg.Data[1:])
				if err != nil {
					err = fmt.Errorf("failed to parse logical WAL log: %v", err)
					log.Fatal(err)
				}

				var msg pglogrepl.Message
				if msg, err = pglogrepl.Parse(walLog.WALData); err != nil {
					err = fmt.Errorf("failed to parse logical replication message: %v", err)
					log.Fatal(err)
				}
				switch m := msg.(type) {
				case *pglogrepl.RelationMessage:
					Event.Columns = []string{}
					for _, col := range m.Columns {
						Event.Columns = append(Event.Columns, col.Name)
					}
					Event.Relation = m.RelationName
				case *pglogrepl.InsertMessage:
					var sb strings.Builder
					sb.WriteString(fmt.Sprintf("INSERT %s(", Event.Relation))
					for i := 0; i < len(Event.Columns); i++ {
						sb.WriteString(fmt.Sprintf("%s: %s ", Event.Columns[i], string(m.Tuple.Columns[i].Data)))
					}
					sb.WriteString(")")
					fmt.Println(sb.String())
				case *pglogrepl.UpdateMessage:
					var sb strings.Builder
					sb.WriteString(fmt.Sprintf("UPDATE %s(", Event.Relation))
					for i := 0; i < len(Event.Columns); i++ {
						sb.WriteString(fmt.Sprintf("%s: %s ", Event.Columns[i], string(m.NewTuple.Columns[i].Data)))
					}
					sb.WriteString(")")
					fmt.Println(sb.String())
				case *pglogrepl.DeleteMessage:
					var sb strings.Builder
					sb.WriteString(fmt.Sprintf("DELETE %s(", Event.Relation))
					for i := 0; i < len(Event.Columns); i++ {
						sb.WriteString(fmt.Sprintf("%s: %s ", Event.Columns[i], string(m.OldTuple.Columns[i].Data)))
					}
					sb.WriteString(")")
					fmt.Println(sb.String())
				case *pglogrepl.TruncateMessage:
					fmt.Println("ALL GONE (TRUNCATE)")
				}
			}
		default:
			fmt.Printf("received unexpected message: %T", msg)
		}
	}
}
