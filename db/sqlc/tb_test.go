package db

import (
	"context"
	"database/sql"
	"testing"
	"tpl/utils"

	"github.com/stretchr/testify/require"
)

func CreateRandomTb(t *testing.T, pros Project) Tb {

	arg := CreateTbParams{
		Name:      sql.NullString{String: utils.RandomString(5), Valid: true},
		ProjectID: sql.NullInt32{Int32: pros.ID, Valid: true},
		Describe:  sql.NullString{String: utils.RandomString(20), Valid: true},
	}

	tb, err := testQueries.CreateTb(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tb)
	require.Equal(t, arg.Name, tb.Name)
	require.NotZero(t, tb.ID)
	require.NotZero(t, tb.Name)

	return tb
}

// test list project
func TestCreateTb(t *testing.T) {

	pros := ListProjects(t)
	for _, p := range pros {
		CreateRandomTb(t, p)
	}
}

func TestWhereTb(t *testing.T) {
	proList := ListProjects(t)
	for _, p := range proList {
		projectID := sql.NullInt32{
			Int32: p.ID,
			Valid: true,
		}
		arg := WhereTbByPIDParams{
			Limit:     int32(5),
			Offset:    int32(0),
			ProjectID: projectID, // 代入 Column4 的值
		}
		_, err := testQueries.WhereTbByPID(context.Background(), arg)
		require.NoError(t, err)
	}
}

func GetTbList(t *testing.T) []Tb {
	arg := ListTbParams{
		Limit:  int32(10),
		Offset: int32(0),
	}
	tb, err := testQueries.ListTb(context.Background(), arg)
	require.NoError(t, err)
	return tb
}

// test tb list
func TestTbList(t *testing.T) {
	tb := GetTbList(t)
	require.NotEmpty(t, tb)
}

// del list project
func TestDeleteTb(t *testing.T) {
	tbs := GetTbList(t)
	var gg bool = true
	for _, tb := range tbs {
		tid := sql.NullInt32{
			Int32: tb.ID,
			Valid: true,
		}
		tfbyid, err := testQueries.GetTFBytID(context.Background(), tid)
		require.NoError(t, err)

		if tfbyid == nil && gg == true {
			err := testQueries.DeleteTb(context.Background(), tb.ID)
			require.NoError(t, err)
			gg = false
		}
	}
}
