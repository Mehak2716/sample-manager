package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
	"github.com/Mehak2716/sample-manager/internal/repository"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpSampleServiceTest() (sqlmock.Sqlmock, *SampleService) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := repository.SampleRepository{DB: gormDB}

	service := SampleService{repo}
	return mock, &service
}

const (
	testSegment         string = "customer_segment"
	testProductID       string = "testProductID"
	testSampleProductID string = "testSampleProductID"
)

func TestCreatedSampleMappingSuccessfully(t *testing.T) {
	mock, service := setUpSampleServiceTest()
	req := &samplev1.SampleMappingRequest{
		CustomerSegment: testSegment,
		ProductID:       testProductID,
		SampleProductID: testSampleProductID,
	}

	expectedResponse := &samplev1.SampleMappingResponse{
		ID:              1,
		CustomerSegment: testSegment,
		ProductID:       testProductID,
		SampleProductID: testSampleProductID,
	}

	rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rowsCount)
	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "customer_segment", "product_id", "sample_product_id"}).
		AddRow(1, testSegment, testProductID, testSampleProductID)
	mock.ExpectQuery("INSERT INTO \"sample_mappings\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := service.CreateMapping(req)

	assert.Equal(t, res, expectedResponse)
	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got %", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCreateDuplicateMappingExpectAlreadyExistError(t *testing.T) {
	mock, service := setUpSampleServiceTest()
	req := &samplev1.SampleMappingRequest{
		CustomerSegment: testSegment,
		ProductID:       testProductID,
		SampleProductID: testSampleProductID,
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	res, err := service.CreateMapping(req)

	if res != nil {
		t.Fatalf("Expected response to be nil")
	}
	if err != nil {
		gRPCStatus, ok := status.FromError(err)
		if !ok {
			t.Fatal("Expected gRPC status error but got a different type of error")
		}
		expectedStatusCode := codes.AlreadyExists
		if gRPCStatus.Code() != expectedStatusCode {
			t.Fatalf("Expected error code: %v, but got: %v", expectedStatusCode, gRPCStatus.Code())
		}
	} else {
		t.Fatal("Expected an error, but got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetSampleIDsSuccessfully(t *testing.T) {
	mock, service := setUpSampleServiceTest()

	req := &samplev1.GetSampleIDsRequest{
		CustomerSegments: []string{"segment1", "segment2"},
		ProductIDs:       []string{"productid1", "productid2"},
	}
	expectedResponse := &samplev1.GetSampleIDsResponse{
		SampleIDs: []string{"sample_product_id_1", "sample_product_id_2"},
	}

	rows := sqlmock.NewRows([]string{"product_id", "sample_product_id", "customer_segment", "id"}).
		AddRow("productid1", "sample_product_id_1", "segment1", 1).
		AddRow("productid2", "sample_product_id_2", "segment2", 2)

	mock.ExpectQuery("SELECT product_id,sample_product_id,customer_segment,id FROM (.+) WHERE rn = 1").
		WillReturnRows(rows)

	res, err := service.GetSampleIDs(req)

	assert.Equal(t, res, expectedResponse)
	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got %", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

}
