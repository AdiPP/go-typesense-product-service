package cdc

import (
	"fmt"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/jackc/pglogrepl"
	"log"
	"strconv"
)

func (s *Service) handleReplicationMessage(data []byte) {
	switch data[0] {
	case pglogrepl.PrimaryKeepaliveMessageByteID:
		fmt.Println("Server: Confirmed Standby")
	case pglogrepl.XLogDataByteID:
		walLog, err := pglogrepl.ParseXLogData(data[1:])
		if err != nil {
			log.Fatalf("Failed to parse Wal log: %v", err)
		}

		msg, err := pglogrepl.Parse(walLog.WALData)
		if err != nil {
			log.Fatalf("Failed to parse replication message: %v", err)
		}

		s.processEvent(msg)

	}
}

var event = struct {
	Relation string
	Columns  map[string]int
}{}

func (s *Service) processEvent(msg pglogrepl.Message) {
	switch m := msg.(type) {
	case *pglogrepl.RelationMessage:
		event.Columns = map[string]int{}
		for i, col := range m.Columns {
			event.Columns[col.Name] = i
		}
		event.Relation = m.RelationName
	case *pglogrepl.InsertMessage:
	case *pglogrepl.UpdateMessage:
		productIDIndex, exists := event.Columns["product_id"]
		if !exists {
			log.Println("Warning: 'product_id' column not found in update message")
			break
		}

		productIDString := string(m.NewTuple.Columns[productIDIndex].Data)
		productID, err := strconv.ParseInt(productIDString, 10, 64)
		if err != nil {
			log.Printf("Error parsing 'product_id' (%s) to int64: %v", productIDString, err)
			break
		}

		log.Printf("Triggering SyncBatchAsync for product_id: %d", productID)
		s.psService.SyncBatchAsync(&service.SyncBatchProductsParam{
			ProductIDs: []int64{productID},
		})
	case *pglogrepl.DeleteMessage:
	case *pglogrepl.TruncateMessage:
		fmt.Println("ALL GONE (TRUNCATE)")
	}
}
