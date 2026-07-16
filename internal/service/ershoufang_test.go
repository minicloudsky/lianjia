package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"bou.ke/monkey"
	v1 "github.com/minicloudsky/lianjia/api/lianjia/v1"
	"github.com/minicloudsky/lianjia/internal/biz"
	"github.com/minicloudsky/lianjia/pkg/pagination"
)

// Define a mock struct that implements the biz.LianjiaRepo interface.
type mockRepo struct {
	biz.LianjiaRepo
}

// Implement the ListErshoufang method.
func (m *mockRepo) ListErshoufang(ctx context.Context, p pagination.Pagination, query string) ([]*biz.ErShouFangInfo, int64, error) {
	return nil, 0, nil
}

func TestListErshoufang(t *testing.T) {
	repo := &mockRepo{}
	s := NewService(nil, repo)

	t.Run("success case", func(t *testing.T) {
		// Use monkey patch to mock the ListErshoufang method of mockRepo
		guard := monkey.PatchInstanceMethod(
			reflect.TypeOf(repo),
			"ListErshoufang",
			func(m *mockRepo, ctx context.Context, p pagination.Pagination, query string) ([]*biz.ErShouFangInfo, int64, error) {
				return []*biz.ErShouFangInfo{
					{
						CityID:     110000,
						HouseCode:  "BJ12345",
						ResblockID: 10001,
						Title:      "Beijing Luxury Apartment",
						Desc:       "3 rooms / 120m² / South / Community A",
						TotalPrice: "1000",
						UnitPrice:  "83333",
					},
				}, 1, nil
			},
		)
		defer guard.Restore()

		req := &v1.ListErshoufangRequest{
			Page:  1,
			Size:  10,
			Query: "Beijing",
		}
		resp, err := s.ListErshoufang(context.Background(), req)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}

		if resp.Total != 1 {
			t.Errorf("expected total 1, got %d", resp.Total)
		}
		if len(resp.Data) != 1 {
			t.Fatalf("expected data length 1, got %d", len(resp.Data))
		}
		if resp.Data[0].Title != "Beijing Luxury Apartment" {
			t.Errorf("expected title 'Beijing Luxury Apartment', got '%s'", resp.Data[0].Title)
		}
	})

	t.Run("validation error case", func(t *testing.T) {
		req := &v1.ListErshoufangRequest{
			Page:  -1, // invalid page
			Size:  10,
			Query: "Beijing",
		}
		_, err := s.ListErshoufang(context.Background(), req)
		if err == nil {
			t.Fatal("expected validation error, got nil")
		}
	})

	t.Run("repository error case", func(t *testing.T) {
		// Use monkey patch to return an error from the repository
		guard := monkey.PatchInstanceMethod(
			reflect.TypeOf(repo),
			"ListErshoufang",
			func(m *mockRepo, ctx context.Context, p pagination.Pagination, query string) ([]*biz.ErShouFangInfo, int64, error) {
				return nil, 0, errors.New("database connection failed")
			},
		)
		defer guard.Restore()

		req := &v1.ListErshoufangRequest{
			Page:  1,
			Size:  10,
			Query: "Beijing",
		}
		_, err := s.ListErshoufang(context.Background(), req)
		if err == nil {
			t.Fatal("expected repository error, got nil")
		}
		if err.Error() != "database connection failed" {
			t.Errorf("expected error message 'database connection failed', got '%s'", err.Error())
		}
	})
}
