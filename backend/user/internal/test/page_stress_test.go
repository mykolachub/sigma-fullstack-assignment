package test

import (
	"net/http"
	"net/http/httptest"
	"sigma-user/config"
	"sigma-user/internal/app"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrackPage(t *testing.T) {
	t.Run("stress test for tracking pages", func(t *testing.T) {
		env := config.ConfigEnv()
		router := app.SetupRouter(env)

		numRequests := 200
		url := "/api/page/track?page=boss"

		var wg sync.WaitGroup
		wg.Add(numRequests)

		for i := 0; i < numRequests; i++ {
			go func() {
				defer wg.Done()
				res := httptest.NewRecorder()
				req, err := http.NewRequest("POST", url, nil)
				if err != nil {
					require.NoError(t, err)
				}

				router.ServeHTTP(res, req)

				if assert.Equal(t, http.StatusOK, res.Code) {
				} else {
					t.Errorf("Request failed with status code: %d", res.Code)
					t.Errorf("Request body: %v", res.Body)
				}
			}()
		}

		wg.Wait()
	})
}
