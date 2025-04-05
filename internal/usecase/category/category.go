package category

import (
	"context"

	error_message "github.com/jonathanmoreiraa/planejja/internal/domain/error"
	entity "github.com/jonathanmoreiraa/planejja/internal/domain/model"
	"github.com/jonathanmoreiraa/planejja/internal/domain/repository"
	services "github.com/jonathanmoreiraa/planejja/internal/usecase/category/contract"
	"github.com/jonathanmoreiraa/planejja/pkg/util"
)

type categoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUseCase(repo repository.CategoryRepository) services.CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: repo,
	}
}

// TODO: limitar a quantidade de categorias que um usuário pode cadastrar e adicionar isso aos requisitos.
func (useCase *categoryUseCase) Create(ctx context.Context, category entity.Category) (entity.Category, error) {
	category, err := useCase.categoryRepo.Create(ctx, category)
	if err != nil {
		return entity.Category{}, util.ErrorWithMessage(err, error_message.ErrCreateAccount)
	}

	return category, nil
}

// TODO: serão pré-cadastradas categorias para todos os usuários ou criaremos algumas padrão, como alimentação e ids fixos? Porque ele terá que retornar algumas categorias base.
func (useCase *categoryUseCase) GetAllCategories(ctx context.Context, userId int) ([]entity.Category, error) {
	categories, err := useCase.categoryRepo.FindAll(ctx, userId)
	if err != nil {
		return []entity.Category{}, util.ErrorWithMessage(err, error_message.ErrFindCategory)
	}

	return categories, nil
}

func (useCase *categoryUseCase) GetCategory(ctx context.Context, name string, userId int) ([]entity.Category, error) {
	category, err := useCase.categoryRepo.FindByName(ctx, name, userId)
	if err != nil {
		return []entity.Category{}, util.ErrorWithMessage(err, error_message.ErrFindCategory)
	}

	return category, nil
}

func (useCase *categoryUseCase) GetCategoryById(ctx context.Context, id int, userId int) (entity.Category, error) {
	category, err := useCase.categoryRepo.FindById(ctx, id, userId)
	if err != nil {
		return entity.Category{}, util.ErrorWithMessage(err, error_message.ErrFindCategory)
	}

	return category, nil
}

// func (useCase *categoryUseCase) Update(ctx context.Context, category entity.Category) error {
// 	err := useCase.categoryRepo.Update(ctx, category)
// 	if err != nil {
// 		return util.ErrorWithMessage(err, error_message.ErrUpdateCategory)
// 	}

// 	return nil
// }

// func (useCase *categoryUseCase) Delete(ctx context.Context, category entity.Category) error {
// 	err := useCase.categoryRepo.Delete(ctx, category)
// 	if err != nil {
// 		return util.ErrorWithMessage(err, error_message.ErrUpdateCategory)
// 	}

// 	return nil
// }
