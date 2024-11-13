package datastream_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/datastream"
	datastreamMock "github.com/es-debug/backend-academy-2024-go-template/pkg/datastream/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestStruct struct {
	Field1 string
	Field2 int
}

func TestProcessFromFile(t *testing.T) {
	mockParser := datastreamMock.NewParser[TestStruct](t)
	mockUpdater := datastreamMock.NewUpdater[TestStruct](t)

	mockParser.On("Parse", "line1").Return(&TestStruct{Field1: "value1", Field2: 1}, nil).Once()
	mockParser.On("Parse", "line2").Return(&TestStruct{Field1: "value2", Field2: 2}, nil).Once()
	mockParser.On("Parse", "line3").Return(&TestStruct{Field1: "value3", Field2: 3}, nil).Once()

	mockUpdater.On("Update", mock.AnythingOfType("*datastream_test.TestStruct")).Times(3)

	file, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.WriteString("line1\nline2\nline3\n")
	assert.NoError(t, err)
	file.Close()

	paths := []string{file.Name()}
	err = datastream.ProcessFromFile(paths, mockParser, mockUpdater)
	assert.NoError(t, err)

	mockParser.AssertExpectations(t)
	mockUpdater.AssertExpectations(t)
}

func TestProcessFromURL(t *testing.T) {
	mockParser := datastreamMock.NewParser[TestStruct](t)
	mockUpdater := datastreamMock.NewUpdater[TestStruct](t)

	mockParser.On("Parse", "line1").Return(&TestStruct{Field1: "value1", Field2: 1}, nil).Once()
	mockParser.On("Parse", "line2").Return(&TestStruct{Field1: "value2", Field2: 2}, nil).Once()
	mockParser.On("Parse", "line3").Return(&TestStruct{Field1: "value3", Field2: 3}, nil).Once()

	mockUpdater.On("Update", mock.AnythingOfType("*datastream_test.TestStruct")).Times(3)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("line1\nline2\nline3\n"))
		assert.NoError(t, err)
	}))
	defer server.Close()

	err := datastream.ProcessFromURL(server.URL, mockParser, mockUpdater)
	assert.NoError(t, err)

	mockParser.AssertExpectations(t)
	mockUpdater.AssertExpectations(t)
}
