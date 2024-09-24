package db_test

import (
	"fmt"
	"github.com/pwinning1991/pjw-swag/db"
	"os"
	"testing"
	"time"
)

const defaultURL = "postgres://postgres@127.0.0.1:5432/swag_test?sslmode=disable"

var (
	testURL string
)

func init() {
	testURL = os.Getenv("PSQL_URL")
	if testURL == "" {
		testURL = defaultURL
	}
	if db.DB != nil {
		db.DB.Close()
	}
	db.Open(testURL)
}

func TestCampaigns(t *testing.T) {
	setup := reset
	teardown := reset

	t.Run("Create", func(t *testing.T) {
		setup(t)
		testCreateCampaign(t)
		teardown(t)
	})
}

func reset(t *testing.T) {
	_, err := db.DB.Exec("DELETE FROM orders")
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	_, err = db.DB.Exec("DELETE FROM campaigns")
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

}

func count(t *testing.T, table string) int {
	var beforeCount int
	err := db.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&beforeCount)
	if err != nil {
		t.Fatalf("count failed: %v", err)
	}
	return beforeCount
}

func testCreateCampaign(t *testing.T) {
	beforeCount := count(t, "campaigns")
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
	afterCount := count(t, "campaigns")
	diff := afterCount - beforeCount
	if diff != 1 {
		t.Fatalf("AfterCount = %d; want %d", diff, 1)
	}

	got, err := db.GetCampaign(campaign.ID)
	if err != nil {
		t.Fatalf("GetCampaign(%d) err = %v; want nil", campaign.ID, err)
	}
	if got.ID <= 0 {
		t.Errorf("ID = %d; want > 0", got.ID)
	}
	//if !got.StartsAt.Equal(start) {
	//	t.Errorf("StartsAt = %v; want %v", got.StartsAt, start)
	//}
}
