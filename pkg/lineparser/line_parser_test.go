package lineparser_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/lineparser"
	mock "github.com/es-debug/backend-academy-2024-go-template/pkg/lineparser/mocks"
	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	mockParser := mock.NewLineParser[string](t)
	mockParser.On("Parse", "line1").Return(pointer("parsed data 1"), nil).Once()
	mockParser.On("Parse", "line2").Return(pointer("parsed data 2"), nil).Once()
	mockParser.On("Parse", "line3").Return(pointer("parsed data 3"), nil).Once()

	file, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.WriteString("line1\nline2\nline3\n")
	assert.NoError(t, err)
	file.Close()

	paths := []string{file.Name()}
	data, err := lineparser.ReadFromFile(paths, mockParser)
	assert.NoError(t, err)
	assert.Equal(t, []*string{pointer("parsed data 1"), pointer("parsed data 2"), pointer("parsed data 3")}, data)

	mockParser.AssertExpectations(t)
}

func TestReadFromURL(t *testing.T) {
	mockParser := mock.NewLineParser[string](t)
	mockParser.On("Parse", "line1").Return(pointer("parsed data 1"), nil).Once()
	mockParser.On("Parse", "line2").Return(pointer("parsed data 2"), nil).Once()
	mockParser.On("Parse", "line3").Return(pointer("parsed data 3"), nil).Once()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("line1\nline2\nline3\n"))
		assert.NoError(t, err)
	}))
	defer server.Close()

	data, err := lineparser.ReadFromURL(server.URL, mockParser)
	assert.NoError(t, err)
	assert.Equal(t, []*string{pointer("parsed data 1"), pointer("parsed data 2"), pointer("parsed data 3")}, data)

	mockParser.AssertExpectations(t)
}

func pointer(s string) *string {
	return &s
}
