package db

import (
	"context"
	"database/sql"
	"testing"
	"tpl/utils"

	"github.com/stretchr/testify/require"
)

func CreateRandomProject(t *testing.T) Project {

	arg := CreateProjectsParams{
		Name:  sql.NullString{String: utils.RandomString(10), Valid: true},
		IsGen: sql.NullInt32{Int32: 0, Valid: true},
	}

	project, err := testQueries.CreateProjects(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, project)

	require.Equal(t, arg.Name, project.Name)

	require.NotZero(t, project.ID)
	require.NotZero(t, project.Name)

	return project
}

// test list project
func TestCreateproject(t *testing.T) {
	for i := 1; i <= 3; i++ {
		CreateRandomProject(t)
	}
}

// test insert project
func TestListProject(t *testing.T) {
	ListProjects(t)
}

// test delete
func TestDeleteProject(t *testing.T) {
	var err error
	pro := ListProjects(t)
	projectID := sql.NullInt32{
		Int32: pro[0].ID, // 設置要刪除的 projectID
		Valid: true,      // 設置為 true，表示 projectID 有效
	}
	parg := WhereTbByPIDParams{
		Limit:  29,
		Offset: 0,
		ProjectID: sql.NullInt32{
			Int32: pro[0].ID, // 設置要刪除的 projectID
			Valid: true,      // 設置為 true，表示 projectID 有效
		},
	}

	tables, _ := testQueries.WhereTbByPID(context.Background(), parg)

	for _, table := range tables {
		tabnull := sql.NullInt32{Int32: table.ID, Valid: true}
		err = testQueries.DeleteTbFieldByTableID(context.Background(), tabnull)
		require.NoError(t, err)
	}

	err = testQueries.DeleteTbByPID(context.Background(), projectID)
	require.NoError(t, err)

	err = testQueries.DeleteProject(context.Background(), pro[0].ID)
	require.NoError(t, err)
}

// test get
func TestListGetProject(t *testing.T) {
	pro := ListProjects(t)
	project, err := testQueries.GetProject(context.Background(), pro[0].ID)
	require.NoError(t, err)
	require.Equal(t, pro[0].Name, project.Name)
}

func ListProjects(t *testing.T) []Project {
	arg := ListProjectsParams{
		Limit:  int32(3),
		Offset: int32(0),
	}
	project, err := testQueries.ListProjects(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, project)
	require.NotEmpty(t, project[0].ID)
	return project
}
