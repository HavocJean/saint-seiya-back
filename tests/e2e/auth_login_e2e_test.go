package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"saint-seiya-back/internal/infrastructure/database/entities"
	"saint-seiya-back/tests/e2e/setup"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginE2E(t *testing.T) {
	testDB, err := setup.ConnectTestDB()
	require.NoError(t, err, "deve conectar ao banco de teste")

	err = setup.RunMigrations(testDB)
	require.NoError(t, err, "deve executar migrations")

	err = setup.CleanDatabase(testDB)
	require.NoError(t, err, "deve limpar o banco de teste")

	testUser := &entities.UserEntity{
		Name:     "Test User E2E",
		Nickname: "teste2e",
		Email:    "teste2e@example.com",
		Password: "password123",
	}

	err = testDB.Create(testUser).Error
	require.NoError(t, err, "deve criar usuário no banco")
	require.NotZero(t, testUser.ID, "usuário deve ter ID gerado")

	router := setup.SetupTestServer(testDB)

	loginRequest := map[string]string{
		"email":    "teste2e@example.com",
		"password": "password123",
	}

	jsonBody, err := json.Marshal(loginRequest)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// ============================================
	// ASSERT: Validar resposta HTTP completa
	// ============================================

	// 1. Status HTTP
	assert.Equal(t, http.StatusOK, w.Code, "deve retornar status 200")

	// 2. Content-Type
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json", "deve retornar JSON")

	// 3. Estrutura do JSON
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err, "resposta deve ser JSON válido")

	// 4. Campos obrigatórios presentes
	assert.Contains(t, response, "success", "resposta deve conter campo 'success'")
	assert.Contains(t, response, "message", "resposta deve conter campo 'message'")
	assert.Contains(t, response, "data", "resposta deve conter campo 'data'")

	// 5. Valores
	assert.True(t, response["success"].(bool), "success deve ser true")
	assert.Equal(t, "Login successful", response["message"], "mensagem deve ser correta")

	// 6. Validar estrutura do data
	data, ok := response["data"].(map[string]interface{})
	require.True(t, ok, "data deve ser um objeto")

	// 7. Token JWT presente e não vazio
	assert.Contains(t, data, "token", "data deve conter campo 'token'")
	token, ok := data["token"].(string)
	require.True(t, ok, "token deve ser string")
	assert.NotEmpty(t, token, "token não deve estar vazio")
	assert.Greater(t, len(token), 20, "token deve ter tamanho razoável (JWT)")

	// ============================================
	// CLEANUP: Limpar banco após teste
	// ============================================
	err = setup.CleanDatabase(testDB)
	require.NoError(t, err, "deve limpar banco após teste")
}

func TestLoginE2E_InvalidCredentials(t *testing.T) {
	testDB, err := setup.ConnectTestDB()
	require.NoError(t, err)

	err = setup.RunMigrations(testDB)
	require.NoError(t, err)

	err = setup.CleanDatabase(testDB)
	require.NoError(t, err)

	router := setup.SetupTestServer(testDB)

	loginRequest := map[string]string{
		"email":    "naoexiste@example.com",
		"password": "senhaerrada",
	}

	jsonBody, _ := json.Marshal(loginRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/user/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Retornar 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.False(t, response["success"].(bool))
	assert.Equal(t, "Invalid credentials", response["message"])

	setup.CleanDatabase(testDB)
}
