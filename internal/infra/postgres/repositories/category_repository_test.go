package repositories

import (
	"database/sql"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/infra/postgres/migrations"
	"github.com/intwone/catalog/internal/tests/factories"
	"github.com/intwone/catalog/internal/tests/utils"
	"github.com/stretchr/testify/require"
)

var sqlDB *sql.DB

func TestMain(m *testing.M) {
	testT := &testing.T{}
	db := utils.NewTestDatabase(testT)
	defer db.Close(testT)
	connStr := db.ConnectionString(testT)
	sqlDB, _ = sql.Open("postgres", connStr)
	defer sqlDB.Close()
	migrations.RunMigration(sqlDB)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func cleanup() {
	_, err := sqlDB.Exec("delete from categories")
	if err != nil {
		panic(err)
	}
}

func TestCategoryRepository_Save(t *testing.T) {
	t.Run("should save a category in the database", func(t *testing.T) {
		// Arange
		categories := *factories.CategoryFactory(1, true)
		categoryRepository := NewCategoryRepository(sqlDB)

		// Act
		err := categoryRepository.Save(categories[0])

		// Assert
		require.Nil(t, err)
		defer cleanup()
	})
}

func TestCategoryRepository_Get(t *testing.T) {
	t.Run("should get all active categories in the database", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description test", true)
		category2, _ := entities.NewCategory("horror category", "description test", false)
		category3, _ := entities.NewCategory("drama category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)
		categoryRepository.Save(*category3)

		// Act
		result, err := categoryRepository.Get()

		// Assert
		require.Nil(t, err)
		require.Equal(t, 2, len(*result))
		defer cleanup()
	})

	t.Run("should return empty array when not have categories in the database", func(t *testing.T) {
		// Arrange
		categoryRepository := NewCategoryRepository(sqlDB)

		// Act
		result, err := categoryRepository.Get()

		// Assert
		require.Nil(t, err)
		require.Equal(t, 0, len(*result))
		defer cleanup()
	})
}

func TestCategoryRepository_GetByID(t *testing.T) {
	t.Run("should get a category by id in the database", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description test", true)
		category2, _ := entities.NewCategory("horror category", "description test", true)
		category3, _ := entities.NewCategory("drama category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)
		categoryRepository.Save(*category3)

		// Act
		result, err := categoryRepository.GetByID(category1.GetID().String())

		// Assert
		require.Nil(t, err)
		require.Equal(t, category1.GetID().String(), result.GetID().String())
		require.Equal(t, category1.GetName(), result.GetName())
		require.Equal(t, category1.GetDescription(), result.GetDescription())
		defer cleanup()
	})

	t.Run("should not get a category when category not found", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description test", true)
		category2, _ := entities.NewCategory("horror category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)

		// Act
		categoryIDNotRegistered := uuid.New().String()
		_, err := categoryRepository.GetByID(categoryIDNotRegistered)

		// Assert
		require.NotNil(t, err)
		require.Equal(t, errs.ResourceNotFound.Error(), err.Error())
		defer cleanup()
	})
}

func TestCategoryRepository_DeleteByID(t *testing.T) {
	t.Run("should deactive a category by id in the database", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description test", true)
		category2, _ := entities.NewCategory("horror category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)

		// Act
		err := categoryRepository.DeleteByID(category1.GetID().String())
		categoriesActivated, _ := categoryRepository.Get()
		category2Activated, _ := categoryRepository.GetByID(category2.GetID().String())

		// Assert
		require.Nil(t, err)
		require.Equal(t, 1, len(*categoriesActivated))
		require.Equal(t, category2.GetID().String(), category2Activated.GetID().String())
		defer cleanup()
	})

	t.Run("should not deactive a category when category not found", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description test", true)
		category2, _ := entities.NewCategory("horror category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)

		// Act
		categoryIDNotRegistered := uuid.New().String()
		err := categoryRepository.DeleteByID(categoryIDNotRegistered)

		// Assert
		require.NotNil(t, err)
		require.Equal(t, errs.ResourceNotFound.Error(), err.Error())
		defer cleanup()
	})
}

func TestCategoryRepository_Update(t *testing.T) {
	t.Run("should update a category in the database", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description test", true)
		category2, _ := entities.NewCategory("horror category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)

		category1.SetName("other name category")
		category1.SetDescription("other description category")

		// Act
		err := categoryRepository.Update(*category1)
		category1Updated, _ := categoryRepository.GetByID(category1.GetID().String())

		// Assert
		require.Nil(t, err)
		require.Equal(t, "other name category", category1Updated.GetName())
		require.Equal(t, "other description category", category1Updated.GetDescription())
		require.Equal(t, "horror category", category2.GetName())
		require.Equal(t, "description test", category2.GetDescription())
		defer cleanup()
	})

	t.Run("should not update a category when category not found", func(t *testing.T) {
		// Arrange
		category, _ := entities.NewCategory("comedy category", "description test", true)
		categoryRepository := NewCategoryRepository(sqlDB)

		// Act
		err := categoryRepository.Update(*category)

		// Assert
		require.NotNil(t, err)
		require.Equal(t, errs.ResourceNotFound.Error(), err.Error())
		defer cleanup()
	})
}

func TestCategoryRepository_Search(t *testing.T) {
	t.Run("should search a category", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description comedy", true)
		category2, _ := entities.NewCategory("horror category", "description horror", true)
		category3, _ := entities.NewCategory("drama category", "description drama", true)
		category4, _ := entities.NewCategory("science fiction category", "description science", true)
		category5, _ := entities.NewCategory("war category", "description war", true)
		category6, _ := entities.NewCategory("action category", "description action", true)
		category7, _ := entities.NewCategory("romantic comedy category", "description romantic comedy", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)
		categoryRepository.Save(*category3)
		categoryRepository.Save(*category4)
		categoryRepository.Save(*category5)
		categoryRepository.Save(*category6)
		categoryRepository.Save(*category7)

		// Act
		result, err := categoryRepository.Search(0, 3, "")

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		require.Equal(t, 3, len(*result))
		require.Equal(t, category1.GetName(), (*result)[0].GetName())
		require.Equal(t, category2.GetName(), (*result)[1].GetName())
		require.Equal(t, category3.GetName(), (*result)[2].GetName())
		defer cleanup()
	})

	t.Run("should search a category by name", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description comedy", true)
		category2, _ := entities.NewCategory("horror category", "description horror", true)
		category3, _ := entities.NewCategory("drama category", "description drama", true)
		category4, _ := entities.NewCategory("science fiction category", "description science", true)
		category5, _ := entities.NewCategory("war category", "description war", true)
		category6, _ := entities.NewCategory("action category", "description action", true)
		category7, _ := entities.NewCategory("romantic comedy category", "description romantic comedy", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)
		categoryRepository.Save(*category3)
		categoryRepository.Save(*category4)
		categoryRepository.Save(*category5)
		categoryRepository.Save(*category6)
		categoryRepository.Save(*category7)

		// Act
		result, err := categoryRepository.Search(0, 3, "comedy")

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		require.Equal(t, 2, len(*result))
		require.Equal(t, category1.GetName(), (*result)[0].GetName())
		require.Equal(t, category7.GetName(), (*result)[1].GetName())
		defer cleanup()
	})

	t.Run("should search a category by description and get only 3 result", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description comedy test", true)
		category2, _ := entities.NewCategory("horror category", "description horror test", true)
		category3, _ := entities.NewCategory("drama category", "description drama test", true)
		category4, _ := entities.NewCategory("science fiction category", "description science test", true)
		category5, _ := entities.NewCategory("war category", "description war test", true)
		category6, _ := entities.NewCategory("action category", "description action test", true)
		category7, _ := entities.NewCategory("romantic comedy category", "description romantic comedy test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)
		categoryRepository.Save(*category3)
		categoryRepository.Save(*category4)
		categoryRepository.Save(*category5)
		categoryRepository.Save(*category6)
		categoryRepository.Save(*category7)

		// Act
		result, err := categoryRepository.Search(0, 3, "test")

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		require.Equal(t, 3, len(*result))
		require.Equal(t, category1.GetName(), (*result)[0].GetName())
		require.Equal(t, category2.GetName(), (*result)[1].GetName())
		require.Equal(t, category3.GetName(), (*result)[2].GetName())
		defer cleanup()
	})

	t.Run("should search a category by name and skip first 3 results", func(t *testing.T) {
		// Arrange
		category1, _ := entities.NewCategory("comedy category", "description comedy test", true)
		category2, _ := entities.NewCategory("horror category", "description horror test", true)
		category3, _ := entities.NewCategory("drama category", "description drama test", true)
		category4, _ := entities.NewCategory("science fiction category", "description science test", true)
		category5, _ := entities.NewCategory("war category", "description war test", true)
		category6, _ := entities.NewCategory("action category", "description action test", true)
		category7, _ := entities.NewCategory("romantic comedy category", "description romantic comedy test", true)
		categoryRepository := NewCategoryRepository(sqlDB)
		categoryRepository.Save(*category1)
		categoryRepository.Save(*category2)
		categoryRepository.Save(*category3)
		categoryRepository.Save(*category4)
		categoryRepository.Save(*category5)
		categoryRepository.Save(*category6)
		categoryRepository.Save(*category7)

		// Act
		result, err := categoryRepository.Search(3, 3, "test")

		// Assert
		require.Nil(t, err)
		require.NotNil(t, result)
		require.Equal(t, 3, len(*result))
		require.Equal(t, category4.GetName(), (*result)[0].GetName())
		require.Equal(t, category5.GetName(), (*result)[1].GetName())
		require.Equal(t, category6.GetName(), (*result)[2].GetName())
		defer cleanup()
	})
}
