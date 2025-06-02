package routers

import (
	"github.com/destroyxiety/CourseWorkSallary/internal/handlers"
	"github.com/destroyxiety/CourseWorkSallary/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, svcs *services.ServicesFactory) {
	customers := e.Group("/customers")
	customers.GET("", handlers.GetAllCustomers(svcs.CustomersService))
	customers.POST("/filter/amount", handlers.GetCustomersByAmount(svcs.CustomersService))
	customers.POST("", handlers.AddCustomer(svcs.CustomersService))
	customers.DELETE("/:customerID", handlers.DeleteCustomer(svcs.CustomersService))
	customers.PUT("/:customerID", handlers.UpdateCutomer(svcs.CustomersService))

	accruals := e.Group("/accruals")
	accruals.GET("", handlers.GetAllAccruals(svcs.AccrualsService))
	accruals.POST("/employee/:employeeID/payment/:paymentID", handlers.AddAccrual(svcs.AccrualsService))
	accruals.DELETE("/employee/:employeeID/payment/:paymentID", handlers.DeleteAccrual(svcs.AccrualsService))

	deals := e.Group("/deals")
	deals.GET("", handlers.GetAllDeals(svcs.DealsService))
	deals.POST("/filter/date", handlers.GetDealsByDate(svcs.DealsService))
	deals.POST("/customer/:customerID", handlers.AddDeals(svcs.DealsService))
	deals.DELETE("/:dealID", handlers.DeleteDeal(svcs.DealsService))

	employees := e.Group("/employees")
	employees.GET("", handlers.GetAllEmployees(svcs.EmploeesService))
	employees.POST("/total-deal", handlers.GetEmployeesByTotalDeal(svcs.EmploeesService))
	employees.POST("/profit", handlers.GetEmployeesByProfit(svcs.EmploeesService))
	employees.POST("/filter/monthly-salary", handlers.GetEmployeesBySalary(svcs.EmploeesService))
	employees.POST("/filter/amount", handlers.GetEmployeesByAmount(svcs.EmploeesService))
	employees.POST("/filter/deal", handlers.GetEmployeesByDeal(svcs.EmploeesService))
	employees.POST("/position/:positionID", handlers.AddEmployees(svcs.EmploeesService))
	employees.PUT("/position/:employeeID", handlers.UpdatePositionEmployee(svcs.EmploeesService))
	employees.DELETE("/:employeeID", handlers.DeleteEmployees(svcs.EmploeesService))

	payments := e.Group("/payments")
	payments.GET("", handlers.GetAllPayments(svcs.PaymentsService))

	paymentsTaxes := e.Group("/payments-taxes")
	paymentsTaxes.GET("", handlers.GetAllPaymentsTaxes(svcs.PaymentsTaxesService))
	paymentsTaxes.POST("/payment/:paymentID/tax/:taxID", handlers.AddPaymentTax(svcs.PaymentsTaxesService))
	paymentsTaxes.DELETE("/payment/:paymentID/tax/:taxID", handlers.DeletePaymentTax(svcs.PaymentsTaxesService))

	percentages := e.Group("/percentages")
	percentages.GET("", handlers.GetAllPercentages(svcs.PercentagesService))
	percentages.POST("/employee/:employeeID/deal/:dealID", handlers.AddPercent(svcs.PercentagesService))
	percentages.DELETE("/employee/:employeeID/deal/:dealID", handlers.DeletePercent(svcs.PercentagesService))

	positions := e.Group("/positions")
	positions.GET("", handlers.GetAllPositions(svcs.PositionsService))
	positions.GET("/count", handlers.GetCountPositions(svcs.PositionsService))
	positions.POST("", handlers.AddPosition(svcs.PositionsService))
	positions.PUT("/:positionID", handlers.UpdatePositionSalary(svcs.PositionsService))
	positions.DELETE("/:positionID", handlers.DeletePosition(svcs.PositionsService))

	taxes := e.Group("/taxes")
	taxes.GET("", handlers.GetAllTaxes(svcs.TaxesService))
}
