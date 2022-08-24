package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type albumTest struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albumsTest = []albumTest{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func Test_getAlbums(t *testing.T) {
	t.Run("Test album lengths", func(t *testing.T) {
		if len(albums) != len(albumsTest) {
			t.Fail()
		}

		for i, v := range albums {
			if v.ID != albumsTest[i].ID {
				t.Fail()
			}
		}
	})
}

func Test_getAlbumsWithRequest(t *testing.T) {
	var newAlbumsList []albumTest
	t.Run("Test get albums with request", func(t *testing.T) {
		r := gin.Default()
		r.GET("/", getAlbums)

		req, _ := http.NewRequest("GET", "/", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		response := w.Result().Body
		responseCode := w.Result().StatusCode

		err := json.NewDecoder(response).Decode(&newAlbumsList)
		if err != nil {
			t.Fail()
		}

		if len(albums) != len(newAlbumsList) {
			t.Fail()
		}

		for i, v := range newAlbumsList {
			if v.ID != albums[i].ID {
				t.Fail()
			}
		}

		assert.NotEmpty(t, response)
		assert.Equal(t, http.StatusOK, responseCode)
	})
}

func Test_postAlbums(t *testing.T) {
	t.Run("Test create an album with request", func(t *testing.T) {
		r := gin.Default()
		r.POST("/albums", postAlbums)

		w := httptest.NewRecorder()

		body, _ := json.Marshal(&album{
			ID:     "4",
			Title:  "Test Album",
			Artist: "Test Artist",
			Price:  20,
		})
		req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)

		response := w.Result().Body
		responseCode := w.Result().StatusCode

		if len(albums) != 4 {
			t.Fail()
		}

		if albums[3].ID == "" {
			t.Fail()
		}

		assert.NotEmpty(t, response)
		assert.Equal(t, http.StatusCreated, responseCode)
	})
}

func Test_postAlbums_fail(t *testing.T) {
	t.Run("Fail to create an album with request", func(t *testing.T) {
		r := gin.Default()
		r.POST("/albums", postAlbums)

		w := httptest.NewRecorder()

		body := []byte("")
		req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(body))
		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode

		assert.Equal(t, http.StatusBadRequest, responseCode)
	})
}

func Test_getAlbumsByID(t *testing.T) {
	t.Run("Test get album by ID", func(t *testing.T) {
		var obtainedAlbum album
		r := gin.Default()
		r.GET("/:id", getAlbumsByID)

		req, _ := http.NewRequest("GET", "/3", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		response := w.Result().Body
		responseCode := w.Result().StatusCode

		err := json.NewDecoder(response).Decode(&obtainedAlbum)
		if err != nil {
			t.Fail()
		}

		assert.NotEmpty(t, response)
		assert.NotNil(t, obtainedAlbum)
		assert.Equal(t, "Sarah Vaughan", obtainedAlbum.Artist)
		assert.Equal(t, http.StatusOK, responseCode)
	})
}

func Test_getAlbumsByID_fail(t *testing.T) {
	t.Run("Fail to get album by ID", func(t *testing.T) {
		r := gin.Default()
		r.GET("/:id", getAlbumsByID)

		req, _ := http.NewRequest("GET", "/10", nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseCode := w.Result().StatusCode
		assert.Equal(t, http.StatusNotFound, responseCode)
	})
}
