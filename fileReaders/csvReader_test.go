package csvReaders

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReadCsv(t *testing.T) {
	dbConfig := ReadCsv("dbConfigForTest.csv")
	assert.Equal(t,"todoMaker",dbConfig["DB_SCHEMA"])
	assert.Equal(t,"postgres",dbConfig["DB_USER"])
	assert.Equal(t,"postgres",dbConfig["DB_PASSWORD"])
	assert.Equal(t,"postgres",dbConfig["DB_NAME"])
}