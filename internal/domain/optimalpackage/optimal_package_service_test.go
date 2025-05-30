package optimalpackage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"order-package/internal/domain/optimalpackage/dto"
	"order-package/internal/domain/optimalpackage/entity"
	"order-package/internal/infra/repository/mock"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOptimalPackageService_Find(t *testing.T) {
	tests := []struct {
		name           string
		mockAvailable  []int64
		requestBody    dto.PackageAmount
		expectedStatus int
		expectedResult dto.PackCombination
	}{
		{
			name:           "Exact match with available packs",
			mockAvailable:  []int64{250, 500, 1000},
			requestBody:    dto.PackageAmount{Amount: 250},
			expectedStatus: http.StatusOK,
			expectedResult: dto.PackCombination{Packs: []dto.Pack{{Size: 250, Amount: 1}}},
		},
		{
			name:           "Best fit with combination",
			mockAvailable:  []int64{250, 500, 1000},
			requestBody:    dto.PackageAmount{Amount: 750},
			expectedStatus: http.StatusOK,
			expectedResult: dto.PackCombination{Packs: []dto.Pack{{Size: 500, Amount: 1}, {Size: 250, Amount: 1}}},
		},
		{
			name:           "Empty input body",
			mockAvailable:  []int64{250, 500, 1000},
			requestBody:    dto.PackageAmount{},
			expectedStatus: http.StatusOK,
			expectedResult: dto.PackCombination{Packs: []dto.Pack{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare mock repo
			mockRepo := &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 {
					return tt.mockAvailable
				},
			}

			// Setup use case and service
			useCase := NewPackageUseCase(mockRepo)
			service := NewOptimalPackageService(useCase)

			// Setup Gin router
			router := gin.Default()
			router.POST("/find", service.Find)

			// Marshal request body
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/find", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, w.Code)

			var result dto.PackCombination
			err := json.Unmarshal(w.Body.Bytes(), &result)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestOptimalPackageService_Create(t *testing.T) {
	tests := []struct {
		name       string
		input      dto.Package
		mockErr    error
		wantStatus int
	}{
		{
			name:       "successful create",
			input:      dto.Package{Size: 500},
			mockErr:    nil,
			wantStatus: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mock.MockPackRepository{
				AddPacksMock: func(ctx context.Context, pack []entity.PackDocument) {

				},
			}

			useCase := NewPackageUseCase(mockRepo)
			service := NewOptimalPackageService(useCase)

			router := gin.Default()
			router.POST("/packs", service.Create)

			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/packs", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestOptimalPackageService_Delete(t *testing.T) {
	tests := []struct {
		name       string
		param      int64
		mockErr    error
		wantStatus int
	}{
		{
			name:       "success delete",
			param:      500,
			mockErr:    nil,
			wantStatus: http.StatusOK,
		},
		{
			name:       "repository error",
			param:      1000,
			mockErr:    errors.New("delete error"),
			wantStatus: http.StatusOK, 
		},
		{
			name:       "invalid param",
			param:      -1,
			mockErr:    nil,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mock.MockPackRepository{
				RemovePackMock: func(ctx context.Context, pack entity.PackDocument) {
				},
			}

			useCase := NewPackageUseCase(mockRepo)
			service := NewOptimalPackageService(useCase)

			router := gin.Default()
			router.DELETE("/packs/:size", service.Delete)

			var path string
			if tt.param >= 0 {
				path = "/packs/" + strconv.FormatInt(tt.param, 10)
			} else {
				path = "/packs/invalid"
			}

			req := httptest.NewRequest(http.MethodDelete, path, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
