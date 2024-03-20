package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mehak2716/sample-manager/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpSampleRepoTest() (sqlmock.Sqlmock, *SampleRepository) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := SampleRepository{DB: gormDB}

	return mock, &repo
}

const (
	testSegment         string = "customer_segment"
	testProductID       string = "aaaaa"
	testSampleProductID string = "aaaa"
)

func TestSampleMappingCreatedSuccessfully(t *testing.T) {
	mock, repo := setUpSampleRepoTest()
	sampleMapping := models.SampleMapping{
		CustomerSegment: testSegment,
		ProductID:       testProductID,
		SampleProductID: testSampleProductID,
	}

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "customer_segment", "product_id", "sample_product_id"}).
		AddRow(1, testSegment, testProductID, testSampleProductID)
	mock.ExpectQuery("INSERT INTO \"sample_mappings\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := repo.Save(&sampleMapping)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if res.ID != 1 || res.ProductID != testProductID || res.CustomerSegment != testSegment {
		t.Fatal("Unexpected Result")
	}
}

func TestIsExistsForExistingMappingSuccessfully(t *testing.T) {
	mock, repo := setUpSampleRepoTest()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	result := repo.IsExists(testSegment, testProductID)

	if !result {
		t.Fatalf("Expected IsExists to return true, but got false")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestIsExistsForNonExistingMappingSuccessfully(t *testing.T) {
	mock, repo := setUpSampleRepoTest()

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	result := repo.IsExists(testSegment, testProductID)

	if result {
		t.Fatalf("Expected IsExists to return false, but got true")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFetchSampleIDSForNoSample(t *testing.T) {
	mock, repo := setUpSampleRepoTest()

	customerSegments := []string{"segment1", "segment2"}
	productIDs := []string{"productid1", "productid2"}

	rows := sqlmock.NewRows([]string{"product_id", "sample_product_id", "customer_segment", "id"})
	mock.ExpectQuery("SELECT product_id,sample_product_id,customer_segment,id FROM (.+) WHERE rn = 1").
		WillReturnRows(rows)
	result, err := repo.FetchSampleIDs(customerSegments, productIDs)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if len(result) != 0 {
		t.Fatalf("Unexpected result")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFetchSampleIDSForSamplesPresent(t *testing.T) {
	mock, repo := setUpSampleRepoTest()

	customerSegments := []string{"segment1", "segment2"}
	productIDs := []string{"productid1", "productid2"}

	rows := sqlmock.NewRows([]string{"product_id", "sample_product_id", "customer_segment", "id"}).
		AddRow("productid1", "sample_product_id_1", "segment1", 1).
		AddRow("productid2", "sample_product_id_2", "segment2", 2)

	mock.ExpectQuery("SELECT product_id,sample_product_id,customer_segment,id FROM (.+) WHERE rn = 1").
		WillReturnRows(rows)
	result, err := repo.FetchSampleIDs(customerSegments, productIDs)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("Unexpected result")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
