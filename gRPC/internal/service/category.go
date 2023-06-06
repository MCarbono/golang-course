package service

import (
	"context"
	"gRPC/internal/database"
	"gRPC/internal/pb"
	"io"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return categoryResponse, nil
}

func (c *CategoryService) ListCategory(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesPB = make([]*pb.Category, len(categories))
	for i := range categoriesPB {
		categoriesPB[i] = &pb.Category{
			Id:          categories[i].ID,
			Name:        categories[i].Name,
			Description: categories[i].Description,
		}
	}
	return &pb.CategoryList{Categories: categoriesPB}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(input.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirection(stream pb.CategoryService_CreateCategoryStreamBidirectionServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		newCategory, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Category{
			Id:          newCategory.ID,
			Name:        newCategory.Name,
			Description: newCategory.Description,
		})
		if err != nil {
			return err
		}
	}
}
