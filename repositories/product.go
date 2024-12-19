package repositories

type ProductRepository interface {
}

type productRepository struct {
}

func ProvideProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetProductByID(id int) error {
	return nil
}
