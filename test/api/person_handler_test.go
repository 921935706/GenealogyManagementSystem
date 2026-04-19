package api

import (
	"GenealogyManagementSystem/internal/application/person"
	personDomain "GenealogyManagementSystem/internal/domain/person"
	"GenealogyManagementSystem/internal/infrastructure/web/handler"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPersonApplicationService 模拟人员应用服务
type MockPersonApplicationService struct {
	mock.Mock
}

func (m *MockPersonApplicationService) CreatePerson(person *personDomain.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockPersonApplicationService) UpdatePerson(person *personDomain.Person) error {
	args := m.Called(person)
	return args.Error(0)
}

func (m *MockPersonApplicationService) DeletePerson(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPersonApplicationService) GetPersonByID(id int) (*personDomain.Person, error) {
	args := m.Called(id)
	return args.Get(0).(*personDomain.Person), args.Error(1)
}

func (m *MockPersonApplicationService) GetPersonByUUID(uuid string) (*personDomain.Person, error) {
	args := m.Called(uuid)
	return args.Get(0).(*personDomain.Person), args.Error(1)
}

func (m *MockPersonApplicationService) ListPersons(page, pageSize int) ([]*personDomain.Person, int, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*personDomain.Person), args.Int(1), args.Error(2)
}

func (m *MockPersonApplicationService) SearchPersons(query string, page, pageSize int) ([]*personDomain.Person, int, error) {
	args := m.Called(query, page, pageSize)
	return args.Get(0).([]*personDomain.Person), args.Int(1), args.Error(2)
}

func TestPersonHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(*MockPersonApplicationService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "成功创建人员",
			requestBody: map[string]interface{}{
				"name":              "张三",
				"first_name":        "张",
				"last_name":         "三",
				"gender":            "男",
				"is_alive":          true,
				"birth_date":        time.Now().Format(time.RFC3339),
				"birth_place":       "北京",
				"generation":        1,
				"is_public":         true,
				"verification_status": "verified",
				"created_by":        "test",
			},
			setupMock: func(m *MockPersonApplicationService) {
				m.On("CreatePerson", mock.AnythingOfType("*person.Person")).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedError:  false,
		},
		{
			name: "无效的请求体",
			requestBody: map[string]interface{}{
				"name": 123, // 错误的类型
			},
			setupMock: func(m *MockPersonApplicationService) {
				// 不设置任何mock调用，因为请求体解析会失败
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "服务层错误",
			requestBody: map[string]interface{}{
				"name":       "李四",
				"first_name": "李",
				"last_name":  "四",
				"gender":     "女",
				"is_alive":   true,
				"birth_date": time.Now().Format(time.RFC3339),
			},
			setupMock: func(m *MockPersonApplicationService) {
				m.On("CreatePerson", mock.AnythingOfType("*person.Person")).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockPersonApplicationService)
			tt.setupMock(mockService)

			// 创建应用服务实例并设置mock
			appService := &person.ApplicationService{}
			handler := handler.NewPersonHandler(appService)

			router := gin.New()
			router.POST("/person/create", handler.Create)

			jsonBody, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/person/create", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedError {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err == nil {
					assert.Contains(t, response, "error")
				}
			} else {
				var response personDomain.Person
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err == nil {
					assert.NotEmpty(t, response.Name)
				}
			}

			mockService.AssertExpectations(t)
		})
	}
}

func TestPersonHandler_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		personID       string
		setupMock      func(*MockPersonApplicationService)
		expectedStatus int
		expectedError  bool
	}{
		{
			name:     "成功获取人员",
			personID: "1",
			setupMock: func(m *MockPersonApplicationService) {
				expectedPerson := &personDomain.Person{
					ID:         1,
					Name:       "张三",
					FirstName:  "张",
					LastName:   "三",
					Gender:     "男",
					IsAlive:    true,
					BirthDate:  time.Now(),
				}
				m.On("GetPersonByID", 1).Return(expectedPerson, nil)
			},
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:     "无效的ID格式",
			personID: "abc",
			setupMock: func(m *MockPersonApplicationService) {
				// 不设置mock调用，因为参数解析会失败
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:     "人员不存在",
			personID: "999",
			setupMock: func(m *MockPersonApplicationService) {
				m.On("GetPersonByID", 999).Return((*personDomain.Person)(nil), assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockPersonApplicationService)
			tt.setupMock(mockService)

			appService := &person.ApplicationService{}
			handler := handler.NewPersonHandler(appService)

			router := gin.New()
			router.GET("/person/:id", handler.GetByID)

			req, _ := http.NewRequest("GET", "/person/"+tt.personID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedError {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err == nil {
					assert.Contains(t, response, "error")
				}
			}

			mockService.AssertExpectations(t)
		})
	}
}