package db

import "context"

type RecipeInterface interface {
	CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error)
	DeleteRecipeByID(ctx context.Context, id int64) (int64, error)
	GetRecipeByID(ctx context.Context, id int64) (Recipe, error)
	GetRecipeByIngredients(ctx context.Context, ingredients []string) ([]Recipe, error)
	GetRecipes(ctx context.Context) ([]Recipe, error)
	GetRecipesByTime(ctx context.Context) ([]Recipe, error)
	UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error)
}

type UserInterface interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

// var _ Querier = (*Queries)(nil)
