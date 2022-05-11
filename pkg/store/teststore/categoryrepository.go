package teststore

import (
	"database/sql"
	"github.com/morozvol/money_manager_api/pkg/model"
	"github.com/morozvol/money_manager_api/pkg/model/category_tree"
)

type CategoryRepository struct {
	store      *Store
	categories map[int]*model.Category
}

type category struct {
	Id       int                        `db:"id"`
	Name     string                     `db:"name"`
	Type     model.OperationPaymentType `db:"type"`
	IdOwner  sql.NullInt64              `db:"id_owner"`
	IdParent sql.NullInt64              `db:"id_parent_category"`
	IsEnd    bool                       `db:"is_end"`
	IsSystem bool                       `db:"is_system"`
}

func (c category) toModel() model.Category {
	return model.Category{Id: c.Id, Name: c.Name, Type: c.Type, IdOwner: int(c.IdOwner.Int64), IdParent: int(c.IdParent.Int64), IsEnd: c.IsEnd}
}

func (r *CategoryRepository) Create(c *model.Category) error {
	c.Id = len(r.categories) + 1
	r.categories[c.Id] = c
	return nil
}

func (r *CategoryRepository) Get(userId int) (*category_tree.CategoryTree, error) {
	res := make([]model.Category, 0)
	var tree *category_tree.CategoryTree
	var node *category_tree.Node
	for _, c := range r.categories {

		res = append(res, *c)
	}

	for _, rc := range res {
		if rc.Id == 1 {
			tree = category_tree.NewCategoryTree(category_tree.NewNode(rc, tree))
			node = tree.Root
		}
	}
	FillNode(node, res, tree)
	return tree, nil
}
func FillNode(node *category_tree.Node, res []model.Category, tree *category_tree.CategoryTree) {
	for _, cc := range res {
		if cc.IdParent == node.Category.Id {
			newNode := category_tree.NewNode(cc, tree)
			node.AddChild(newNode)
			if !cc.IsEnd {
				FillNode(newNode, res, tree)
			}
		}
	}
}

func (r *CategoryRepository) GetSystem() ([]model.Category, error) {
	res := make([]model.Category, 0)
	for _, c := range r.categories {

		if c.IsSystem {
			res = append(res, *c)
		}
	}
	return res, nil
}

func (r *CategoryRepository) Delete(id int) error {
	delete(r.categories, id)
	return nil
}
