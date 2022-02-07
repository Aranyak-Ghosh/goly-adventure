package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/Aranyak-Ghosh/spotigo/media_storage/models/entity"
)

func TestCreateQuery(t *testing.T) {
	var data entity.TrackEntity = entity.TrackEntity{
		Title:       "Achiles Come Down",
		Description: "Track by Gangs of Youth",
		CreatedAt:   time.Now(),
	}

	query, val := GenerateCreateQuery(entity.Track, data)

	fmt.Printf("%s %+v", query, val)
}
