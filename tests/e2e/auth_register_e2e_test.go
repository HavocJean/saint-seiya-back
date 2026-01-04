package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"saint-seiya-back/tests/e2e/setup"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterE2E_Success(t *testing.T) {
	testDB, err := setup.ConnectTestDB()
	require.NoError(t, err, "deve conectar ao banco de teste")

	err = setup.RunMigrations(testDB)
	require.NoError(t, err, "deve executar migrations")

	err = setup.CleanDatabase(testDB)
	require.NoError(t, err, "deve limpar o banco de teste")

	router := setup.SetupTestServer(testDB)

	registerRequest := map[string]string{
		"name":     "New User E2E",
		"nickname": "newuser",
		"email":    "newuser@example.com",
		"password": "password123",
	}

	jsonBody, err := json.Marshal(registerRequest)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// ============================================
	// ASSERT: Validar resposta HTTP completa
	// ============================================

	// 1. Status HTTP correto (201 Created)
	assert.Equal(t, http.StatusCreated, w.Code, "deve retornar status 201")

	// 2. Content-Type correto
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json", "deve retornar um JSON")

	// 3. Estrutura do JSON
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err, "resposta deve ser JSON válido")

	// 4. Campos obrigatórios presentes
	assert.Contains(t, response, "success", "resposta deve conter campo 'success'")
	assert.Contains(t, response, "message", "resposta deve conter campo 'message'")
	assert.Contains(t, response, "data", "resposta deve conter campo 'data'")

	// 5. Valores corretos
	assert.True(t, response["success"].(bool), "success deve ser true")
	assert.Equal(t, "Registration successful", response["message"], "mensagem deve ser correta")

	// 6. Validar estrutura do data
	data, ok := response["data"].(map[string]interface{})
	require.True(t, ok, "data deve ser um objeto")

	// 7. Token JWT presente e não vazio
	assert.Contains(t, data, "token", "data deve conter campo 'token'")
	token, ok := data["token"].(string)
	require.True(t, ok, "token deve ser string")
	assert.NotEmpty(t, token, "token não deve estar vazio")
	assert.Greater(t, len(token), 20, "token deve ter tamanho razoável (JWT)")

	// 8. Verificar que usuário foi criado no banco
	var userCount int64
	testDB.Table("users").Where("email = ?", "newuser@example.com").Count(&userCount)
	assert.Equal(t, int64(1), userCount, "deve existir exatamente 1 usuário com esse email")

	// ============================================
	// CLEANUP: Limpar banco após teste
	// ============================================
	err = setup.CleanDatabase(testDB)
	require.NoError(t, err, "deve limpar banco após teste")
}

func TestRegisterE2E_DuplicateEmail(t *testing.T) {
	testDB, err := setup.ConnectTestDB()
	require.NoError(t, err)

	err = setup.RunMigrations(testDB)
	require.NoError(t, err)

	err = setup.CleanDatabase(testDB)
	require.NoError(t, err)

	router := setup.SetupTestServer(testDB)

	firstRegister := map[string]string{
		"name":     "First User",
		"nickname": "firstuser",
		"email":    "duplicate@example.com",
		"password": "password123",
	}

	jsonBody1, _ := json.Marshal(firstRegister)
	req1 := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", bytes.NewBuffer(jsonBody1))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	assert.Equal(t, http.StatusCreated, w1.Code, "primeiro registro deve funcionar")

	secondRegister := map[string]string{
		"name":     "Second User",
		"nickname": "seconduser",
		"email":    "duplicate@example.com",
		"password": "password456",
	}

	jsonBody2, _ := json.Marshal(secondRegister)
	req2 := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", bytes.NewBuffer(jsonBody2))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	// Deve retornar 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w2.Code, "deve retornar 400 para email duplicado")

	var response map[string]interface{}
	err = json.Unmarshal(w2.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response["success"].(bool), "success deve ser false")
	assert.Equal(t, "Registration failed", response["message"], "mensagem deve indicar falha")
	assert.Contains(t, response["error"].(string), "email already exists", "deve mencionar email duplicado")

	setup.CleanDatabase(testDB)
}

func TestRegisterE2E_ValidationErrors(t *testing.T) {
	testDB, err := setup.ConnectTestDB()
	require.NoError(t, err)

	err = setup.RunMigrations(testDB)
	require.NoError(t, err)

	err = setup.CleanDatabase(testDB)
	require.NoError(t, err)

	router := setup.SetupTestServer(testDB)

	testCases := []struct {
		name           string
		request        map[string]string
		expectedStatus int
		description    string
	}{
		{
			name: "email inválido",
			request: map[string]string{
				"name":     "Test User",
				"nickname": "testuser",
				"email":    "email-invalido",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "deve rejeitar email inválido",
		},
		{
			name: "senha muito curta",
			request: map[string]string{
				"name":     "Test User",
				"nickname": "testuser",
				"email":    "test@example.com",
				"password": "short",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "deve rejeitar senha muito curta",
		},
		{
			name: "campos obrigatórios faltando",
			request: map[string]string{
				"name":     "Test User",
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
			description:    "deve rejeitar quando campos obrigatórios faltam",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonBody, _ := json.Marshal(tc.request)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code, tc.description)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.False(t, response["success"].(bool), "success deve ser false")
		})
	}

	setup.CleanDatabase(testDB)
}

func TestRegisterE2E_InvalidJSON(t *testing.T) {
	testDB, err := setup.ConnectTestDB()
	require.NoError(t, err)

	router := setup.SetupTestServer(testDB)

	invalidJSON := `{"name": "Test", "email":}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user/register", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Deve retornar 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response["success"].(bool))
}
