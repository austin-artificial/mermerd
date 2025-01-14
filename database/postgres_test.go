package database

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPostgresEnums(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// Arrange
	var enumValues string

	// Act
	connector, _ := NewConnectorFactory().NewConnector(testConnectionPostgres.connectionString)
	if err := connector.Connect(); err != nil {
		logrus.Error(err)
		t.FailNow()
	}
	columns, err := connector.GetColumns(TableDetail{Schema: "public", Name: "test_2_enum"})

	// Assert
	for _, column := range columns {
		if column.Name == "fruit" {
			enumValues = column.EnumValues
		}
	}

	assert.Nil(t, err)
	assert.Equal(t, "apple,banana", enumValues)
}
