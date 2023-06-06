package usecase

import (
	"database/sql"
	"intensivego/internal/order/entity"
	"intensivego/internal/order/infra/database"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CalculatePriceUseCaseSuite struct {
	suite.Suite
	OrderRepository entity.OrderRepositoryInterface
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *CalculatePriceUseCaseSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseSuite))
}

func (suite *CalculatePriceUseCaseSuite) TestCalculateFinalPrice() {
	calculateFinalPriceInput := OrderInputDTO{
		ID:    "1",
		Price: 10,
		Tax:   2,
	}
	orderRepository := database.NewOrderRepository(suite.Db)
	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(orderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)
	suite.Equal(calculateFinalPriceInput.ID, output.ID)
	suite.Equal(calculateFinalPriceInput.Price, output.Price)
	suite.Equal(calculateFinalPriceInput.Tax, output.Tax)
	suite.Equal(12.0, output.FinalPrice)
}
