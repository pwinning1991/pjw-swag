package db_test

import (
	"github.com/pwinning1991/pjw-swag/db"
	"testing"
	"time"
)

func TestCreateCampaign(t *testing.T) {
	var beforeCount int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM campaigns").Scan(&beforeCount)
	if err != nil {
		t.Fatalf("Scan() err = %v; want nil", err)
	}
	start := time.Now()
	end := time.Now().Add(1 * time.Hour)
	price := 1000
	campaign, err := db.CreateCampaign(start, end, price)
	if err != nil {
		t.Fatalf("Error creating campaign: err = %v; wnat nil", err)
	}
	if campaign.ID <= 0 {
		t.Errorf("ID = %d; want > 0", campaign.ID)
	}
	//if !campaign.StartsAt.Equal(start) {
	//	t.Errorf("StartsAt = %v; want %v", campaign.StartsAt, start)
	//}

	var afterCount int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM campaigns").Scan(&afterCount)
	if err != nil {
		t.Fatalf("Scan() err = %v; want nil", err)
	}
	if afterCount-beforeCount != 1 {
		t.Fatalf("AfterCount = %d; want %d", afterCount-beforeCount, 1)
	}

	got, err := db.GetCampaign(campaign.ID)
	if err != nil {
		t.Fatalf("GetCampaign(%d) err = %v; want nil", campaign.ID, err)
	}
	if got.ID <= 0 {
		t.Errorf("ID = %d; want > 0", got.ID)
	}
	if !got.StartsAt.Equal(start) {
		t.Errorf("StartsAt = %v; want %v", got.StartsAt, start)
	}
}
