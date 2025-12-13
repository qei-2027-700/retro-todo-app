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

	sprint, err := repo.Create("Test Sprint", "bg-purple-500", false)

	assert.NoError(t, err)
	assert.NotNil(t, sprint)
	assert.Greater(t, sprint.ID, 0)
	assert.Equal(t, "Test Sprint", sprint.Name)
	assert.Equal(t, "bg-purple-500", sprint.Color)
	assert.False(t, sprint.IsFavorite)
}

func TestSprintRepository_FindAll(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	_, err := repo.Create("Sprint 1", "bg-purple-500", false)
	require.NoError(t, err)
	_, err = repo.Create("Sprint 2", "bg-blue-500", true)
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
	sprint, err := repo.Create("Original Sprint", "bg-purple-500", false)
	require.NoError(t, err)

	// 更新
	rowsAffected, message, err := repo.Update(sprint.ID, "Updated Sprint", "bg-green-500")

	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)
	assert.NotEmpty(t, message)
}

func TestSprintRepository_Update_NotFound(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// 存在しないIDで更新
	rowsAffected, _, err := repo.Update(99999, "Updated Sprint", "bg-blue-500")

	assert.NoError(t, err)
	assert.Equal(t, 0, rowsAffected)
}

func TestSprintRepository_Delete(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	sprint, err := repo.Create("To Be Deleted", "bg-red-500", false)
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

func TestSprintRepository_Search_ByName(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	_, err := repo.Create("Search Test Sprint", "bg-purple-500", false)
	require.NoError(t, err)
	_, err = repo.Create("Another Sprint", "bg-blue-500", false)
	require.NoError(t, err)

	// 名前で検索
	name := "Search"
	req := &model.SprintSearchRequest{
		Name: &name,
	}
	sprints, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 1)
	// 少なくとも1つは"Search"を含む名前があるはず
	found := false
	for _, sprint := range sprints {
		if sprint.Name == "Search Test Sprint" {
			found = true
			break
		}
	}
	assert.True(t, found)
}

func TestSprintRepository_Search_ByIsFavorite(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成（お気に入り）
	// sprint, err := repo.Create("Favorite Sprint", "bg-purple-500", true)
	_, err := repo.Create("Favorite Sprint", "bg-purple-500", true)
	require.NoError(t, err)

	// お気に入りではないスプリントも作成
	_, err = repo.Create("Non-Favorite Sprint", "bg-blue-500", false)
	require.NoError(t, err)

	// お気に入りで検索
	isFavorite := true
	req := &model.SprintSearchRequest{
		IsFavorite: &isFavorite,
	}
	sprints, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 1)
	// すべてお気に入りのはず
	for _, sprint := range sprints {
		assert.True(t, sprint.IsFavorite)
	}
}

func TestSprintRepository_Search_MultipleConditions(t *testing.T) {
	db := setupSprintTestDB(t)
	defer db.Close()

	repo := NewSprintRepository(db)

	// テスト用のスプリントを作成
	// sprint, err := repo.Create("Multi Search Sprint", "bg-orange-500", true)
	_, err := repo.Create("Multi Search Sprint", "bg-orange-500", true)
	require.NoError(t, err)

	// 複数条件で検索
	name := "Multi"
	isFavorite := true
	req := &model.SprintSearchRequest{
		Name:       &name,
		IsFavorite: &isFavorite,
	}
	sprints, err := repo.Search(req)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(sprints), 1)
	// 見つかったスプリントはすべてお気に入り状態のはず
	for _, sprint := range sprints {
		assert.True(t, sprint.IsFavorite)
	}
}
