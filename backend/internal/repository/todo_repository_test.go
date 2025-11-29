package repository

import (
	"backend/internal/model"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sql.DB {
	testDBConn := os.Getenv("TEST_DB_CONN")
	if testDBConn == "" {
		t.Skip("TEST_DB_CONN not set, skipping integration tests")
	}

	db, err := sql.Open("postgres", testDBConn)
	require.NoError(t, err)

	// テーブルをクリーンアップ
	_, err = db.Exec("DELETE FROM todos WHERE is_deleted = false")
	require.NoError(t, err)
	_, err = db.Exec("DELETE FROM sprints WHERE is_deleted = false")
	require.NoError(t, err)

	return db
}

func TestTodoRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	todo, err := repo.Create("Test Todo", "Test Description", nil)

	assert.NoError(t, err)
	assert.NotNil(t, todo)
	assert.Greater(t, todo.ID, 0)
	assert.Equal(t, "Test Todo", todo.Title)
	assert.Equal(t, "Test Description", todo.Description)
	assert.False(t, todo.Completed)
	assert.Nil(t, todo.SprintID)
}

func TestTodoRepository_FindAll(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// テスト用のTODOを作成
	_, err := repo.Create("Todo 1", "Description 1", nil)
	require.NoError(t, err)
	_, err = repo.Create("Todo 2", "Description 2", nil)
	require.NoError(t, err)

	todos, err := repo.FindAll()

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(todos), 2)
}

func TestTodoRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// テスト用のTODOを作成
	todo, err := repo.Create("Original Title", "Original Description", nil)
	require.NoError(t, err)

	// 更新
	rowsAffected, message, err := repo.Update("Updated Title", true, todo.ID)

	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)
	assert.NotEmpty(t, message)
}

func TestTodoRepository_Update_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// 存在しないIDで更新
	rowsAffected, _, err := repo.Update("Updated Title", true, 99999)

	assert.NoError(t, err)
	assert.Equal(t, 0, rowsAffected)
}

func TestTodoRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// テスト用のTODOを作成
	todo, err := repo.Create("To Be Deleted", "Description", nil)
	require.NoError(t, err)

	// 削除
	err = repo.Delete(todo.ID)

	assert.NoError(t, err)

	// 削除されたことを確認（論理削除なのでis_deleted=trueになっているはず）
	var isDeleted bool
	err = db.QueryRow("SELECT is_deleted FROM todos WHERE id = $1", todo.ID).Scan(&isDeleted)
	require.NoError(t, err)
	assert.True(t, isDeleted)
}

func TestTodoRepository_Search_ByTitle(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// テスト用のTODOを作成
	_, err := repo.Create("Search Test Todo", "Description", nil)
	require.NoError(t, err)
	_, err = repo.Create("Another Todo", "Description", nil)
	require.NoError(t, err)

	// タイトルで検索
	title := "Search"
	req := &model.TodoSearchRequest{
		Title: &title,
	}
	todos, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(todos), 1)
	// 少なくとも1つは"Search"を含むタイトルがあるはず
	found := false
	for _, todo := range todos {
		if todo.Title == "Search Test Todo" {
			found = true
			break
		}
	}
	assert.True(t, found)
}

func TestTodoRepository_Search_ByCompleted(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// テスト用のTODOを作成
	todo, err := repo.Create("Completed Todo", "Description", nil)
	require.NoError(t, err)

	// 完了状態に更新
	_, _, err = repo.Update("Completed Todo", true, todo.ID)
	require.NoError(t, err)

	// 未完了のTODOも作成
	_, err = repo.Create("Incomplete Todo", "Description", nil)
	require.NoError(t, err)

	// 完了状態で検索
	completed := true
	req := &model.TodoSearchRequest{
		Completed: &completed,
	}
	todos, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(todos), 1)
	// すべて完了状態のはず
	for _, todo := range todos {
		assert.True(t, todo.Completed)
	}
}

func TestTodoRepository_Search_MultipleConditions(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewTodoRepository(db)

	// テスト用のTODOを作成
	_, err := repo.Create("Multi Search Todo", "Special Description", nil)
	require.NoError(t, err)

	// 複数条件で検索
	title := "Multi"
	description := "Special"
	req := &model.TodoSearchRequest{
		Title:       &title,
		Description: &description,
	}
	todos, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(todos), 1)
}
