package handler

import (
	"book-app/config"
	"book-app/entity"
	"book-app/formatter"
	"book-app/service"
	"book-app/transport"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service service.ServiceBook
}

func NewBookHandler(service service.ServiceBook) *bookHandler {
	return &bookHandler{service}
}

func(h *bookHandler) CreateBook(c *gin.Context) {
	var input transport.InputDataBook

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := transport.FormatValidationError(err)
		response := transport.ApiResponse("Failed Create Book", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(entity.User)
	input.User = currentUser

	newBook, err := h.service.AddBook(input)
	if err != nil {
		response := transport.ApiResponse("Failed Create Book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := transport.ApiResponse("Success Create Book", http.StatusOK, "success", formatter.FormatBook(newBook))

	c.JSON(http.StatusOK, response)
}

func(h *bookHandler) GetBook(c *gin.Context) {
	idBook := c.Param("id")

	convertIdBook, err := strconv.Atoi(idBook)
	if err != nil {
		fmt.Println("masuk1")
		response := transport.ApiResponse("Failed to get detail book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	bookDetail, err := h.service.FindByID(convertIdBook)
	if err != nil {
		fmt.Println("masuk2")
		response := transport.ApiResponse("Failed to get detail book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	} 

	response := transport.ApiResponse("Success get book detail", http.StatusOK, "success", formatter.FormatBookDetail(bookDetail))
	c.JSON(http.StatusOK, response)
}

func(h *bookHandler) GetBooks(c *gin.Context) {
	genre := c.Query("genre")
	isRead := c.Query("is_read")
	startYear, _ := strconv.Atoi(c.Query("start_year"))
	endYear, _ := strconv.Atoi(c.Query("end_year"))
	userID, _ := strconv.Atoi(c.Query("user_id"))
	limitQuery := c.DefaultQuery("limit", config.DefaultLimit)
	pageQuery := c.DefaultQuery("page", config.DefaultPage)

	limit, _ := strconv.Atoi(limitQuery)
	page, _ := strconv.Atoi(pageQuery)
	
	filterBook := transport.FilterBook{}
	filterBook.Genre = genre
	filterBook.IsRead = isRead
	filterBook.StartYear = startYear
	filterBook.EndYear = endYear
	fmt.Println("start year", startYear)
	fmt.Println("end year", endYear)

	books, count, err := h.service.GetBooks(userID, filterBook, limit, page)
	if err != nil {
		fmt.Println(err.Error())
		response := transport.ApiResponse("Error to get books", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(count)

	response := transport.ApiResponseGetListBook("List of books", http.StatusOK, "success", formatter.FormatBooks(books), count)
	
	c.JSON(http.StatusOK, response)
}

func(h *bookHandler) UpdateBook(c *gin.Context) {

	var inputID transport.InputDetailIdBook
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := transport.ApiResponse("Failed to update book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	var inputData transport.InputDataBookUpdate
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := transport.FormatValidationError(err)
		response := transport.ApiResponse("Failed Update book", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	currentUser := c.MustGet("currentUser").(entity.User) //! melakukan auth user, hanya user yang memiliki item tsb bisa melakukabn update
	inputData.User = currentUser

	updateBook, err := h.service.UpdateBook(inputID, inputData)
	if err != nil {
		response := transport.ApiResponse("Failed Update book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := transport.ApiResponse("Success Update book", http.StatusOK, "success", formatter.FormatBook(updateBook))

	c.JSON(http.StatusOK, response)
}

func(h *bookHandler) DeleteBook(c *gin.Context) {
	var inputID transport.InputDetailIdBook

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		fmt.Println("masuk1")
		response := transport.ApiResponse("Failed to Delete book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	var inputData transport.InputDetailIdBook

	currentUser := c.MustGet("currentUser").(entity.User) //! melakukan auth user, hanya user yang memiliki item tsb bisa melakukabn update
	inputData.User = currentUser

	_, err = h.service.DeleteBook(inputID, inputData)
	if err != nil {
		fmt.Println("masuk3")
		response := transport.ApiResponse("Failed Delete book", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return;
	}

	response := transport.ApiResponse("Success Delete book", http.StatusOK, "success", "")

	c.JSON(http.StatusOK, response)
}

