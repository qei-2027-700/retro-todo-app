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

func setupSprintTestDB(t *testing.T) *sql.DB {
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

func TestSprintRepository_Create(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	sprint, err := repo.Create("Test Sprint")

	assert.NoError(t, err)
	assert.NotNil(t, sprint)
	assert.Greater(t, sprint.ID, 0)
	assert.Equal(t, "Test Sprint", sprint.Title)
	assert.False(t, sprint.Completed)
}

func TestSprintRepository_FindAll(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	_, err := repo.Create("Sprint 1")
	require.NoError(t, err)
	_, err = repo.Create("Sprint 2")
	require.NoError(t, err)

	sprints, err := repo.FindAll()

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 2)
}

func TestSprintRepository_Update(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	sprint, err := repo.Create("Original Sprint")
	require.NoError(t, err)

	// 更新
	rowsAffected, message, err := repo.Update(sprint.ID, "Updated Sprint", true)

	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)
	assert.NotEmpty(t, message)
}

func TestSprintRepository_Update_NotFound(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// 存在しないIDで更新
	rowsAffected, _, err := repo.Update(99999, "Updated Sprint", true)

	assert.NoError(t, err)
	assert.Equal(t, 0, rowsAffected)
}

func TestSprintRepository_Delete(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	sprint, err := repo.Create("To Be Deleted")
	require.NoError(t, err)

	// 削除
	err = repo.Delete(sprint.ID)

	assert.NoError(t, err)

	// 削除されたことを確認（論理削除なのでis_deleted=trueになっているはず）
	var isDeleted bool
	err = db.QueryRow("SELECT is_deleted FROM sprints WHERE id = $1", sprint.ID).Scan(&isDeleted)
	require.NoError(t, err)
	assert.True(t, isDeleted)
}

func TestSprintRepository_Search_ByTitle(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	_, err := repo.Create("Search Test Sprint")
	require.NoError(t, err)
	_, err = repo.Create("Another Sprint")
	require.NoError(t, err)

	// タイトルで検索
	title := "Search"
	req := &model.SprintSearchRequest{
		Title: &title,
	}
	sprints, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 1)
	// 少なくとも1つは"Search"を含むタイトルがあるはず
	found := false
	for _, sprint := range sprints {
		if sprint.Title == "Search Test Sprint" {
			found = true
			break
		}
	}
	assert.True(t, found)
}

func TestSprintRepository_Search_ByCompleted(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	sprint, err := repo.Create("Completed Sprint")
	require.NoError(t, err)

	// 完了状態に更新
	_, _, err = repo.Update(sprint.ID, "Completed Sprint", true)
	require.NoError(t, err)

	// 未完了のスプリントも作成
	_, err = repo.Create("Incomplete Sprint")
	require.NoError(t, err)

	// 完了状態で検索
	completed := true
	req := &model.SprintSearchRequest{
		Completed: &completed,
	}
	sprints, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 1)
	// すべて完了状態のはず
	for _, sprint := range sprints {
		assert.True(t, sprint.Completed)
	}
}

func TestSprintRepository_Search_MultipleConditions(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	sprint, err := repo.Create("Multi Search Sprint")
	require.NoError(t, err)

	// 完了状態に更新
	_, _, err = repo.Update(sprint.ID, "Multi Search Sprint", true)
	require.NoError(t, err)

	// 複数条件で検索
	title := "Multi"
	completed := true
	req := &model.SprintSearchRequest{
		Title:     &title,
		Completed: &completed,
	}
	sprints, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 1)
	// 見つかったスプリントはすべて完了状態のはず
	for _, sprint := range sprints {
		assert.True(t, sprint.Completed)
	}
}
