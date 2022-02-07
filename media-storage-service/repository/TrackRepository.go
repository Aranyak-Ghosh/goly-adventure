package repository

import (
	"github.com/Aranyak-Ghosh/spotigo/media_storage/models/entity"
	"github.com/Aranyak-Ghosh/spotigo/media_storage/utils/db"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func AddTrack(track entity.TrackEntity) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var query, data = db.GenerateCreateQuery(entity.Track, track)
		var result, err = tx.Run(query, data)
		if err != nil {
			return nil, err
		}
		return result.Consume()
	}
}
