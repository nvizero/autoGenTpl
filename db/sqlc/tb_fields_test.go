package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"tpl/utils"

	"github.com/stretchr/testify/require"
)

func CreateTbFields(t *testing.T) {
	tb := GetTbList(t)
	if tb[0].ID > 0 {
		arg := CreateTbFieldParams{
			TableID:    sql.NullInt32{Int32: tb[0].ID, Valid: true},
			FieldName:  sql.NullString{String: utils.RandomString(10), Valid: true},
			LaravelMap: sql.NullString{String: utils.RandomString(10), Valid: true},
		}
		fmt.Println(arg)
		tf, err := testQueries.CreateTbField(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, tf)
	}

}

// test list project
func TestCreateTbFields(t *testing.T) {
	CreateTbFields(t)
}

func GetTBField(t *testing.T) {

}
