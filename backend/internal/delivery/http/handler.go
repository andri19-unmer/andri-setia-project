package http

import (
	"net/http"
	"strconv"

	"app-backend/internal/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	userUsecase    domain.UserUsecase
	productUsecase domain.ProductUsecase
}

func NewHandler(e *echo.Echo, uu domain.UserUsecase, pu domain.ProductUsecase) {
	handler := &Handler{
		userUsecase:    uu,
		productUsecase: pu,
	}

	api := e.Group("/api")

	// User Routes
	api.GET("/users", handler.FetchUsers)
	api.GET("/users/:id", handler.GetUserByID)
	api.POST("/users", handler.StoreUser)
	api.PUT("/users/:id", handler.UpdateUser)
	api.DELETE("/users/:id", handler.DeleteUser)

	// Product Routes
	api.GET("/products", handler.FetchProducts)
	api.GET("/products/:id", handler.GetProductByID)
	api.POST("/products", handler.StoreProduct)
	api.PUT("/products/:id", handler.UpdateProduct)
	api.DELETE("/products/:id", handler.DeleteProduct)
}

func (h *Handler) FetchUsers(c echo.Context) error {
	users, err := h.userUsecase.Fetch(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.userUsecase.GetByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) StoreUser(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.userUsecase.Store(c.Request().Context(), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user.ID = uint(id)
	if err := h.userUsecase.Update(c.Request().Context(), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.userUsecase.Delete(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) FetchProducts(c echo.Context) error {
	products, err := h.productUsecase.Fetch(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productUsecase.GetByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func (h *Handler) StoreProduct(c echo.Context) error {
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.productUsecase.Store(c.Request().Context(), &product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, product)
}

func (h *Handler) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product domain.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	product.ID = uint(id)
	if err := h.productUsecase.Update(c.Request().Context(), &product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func (h *Handler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.productUsecase.Delete(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
