package controller

import (
	"fmt"
	"log"
	"strconv"
	booksmodel "tugas_akhir_example/internal/pkg/model"
	booksusecase "tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type BooksController interface {
	GetAllBooks(ctx *fiber.Ctx) error
	GetBooksByID(ctx *fiber.Ctx) error
	CreateBooks(ctx *fiber.Ctx) error
	UpdateBooksByID(ctx *fiber.Ctx) error
	DeleteBooksByID(ctx *fiber.Ctx) error
}

type BooksControllerImpl struct {
	booksusecase booksusecase.BooksUseCase
}

func NewBooksController(booksusecase booksusecase.BooksUseCase) BooksController {
	return &BooksControllerImpl{
		booksusecase: booksusecase,
	}
}

func (uc *BooksControllerImpl) GetAllBooks(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(booksmodel.BooksFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.booksusecase.GetAllBooks(c, booksmodel.BooksFilter{
		Title: filter.Title,
		Limit: filter.Limit,
		Page:  filter.Page,
	})

	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) GetBooksByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	booksid := ctx.Params("id_books")
	if booksid == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.booksusecase.GetBooksByID(c, booksid)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) CreateBooks(ctx *fiber.Ctx) error {
	c := ctx.Context()

	// cara baca context yang diset di middleware
	id := ctx.Locals("userid").(string)
	email := ctx.Locals("useremail").(string)

	fmt.Println("id", id)
	fmt.Println("email", email)

	data := new(booksmodel.BooksReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userid, _ := strconv.Atoi(id)
	data.UserID = uint(userid)
	res, err := uc.booksusecase.CreateBooks(c, *data)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) UpdateBooksByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	booksid := ctx.Params("id_books")
	if booksid == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	data := new(booksmodel.BooksReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.booksusecase.UpdateBooksByID(c, booksid, *data)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func (uc *BooksControllerImpl) DeleteBooksByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	booksid := ctx.Params("id_books")
	if booksid == "" {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	res, err := uc.booksusecase.DeleteBooksByID(c, booksid)
	if err != nil {
		// TODO IMRPOVE FORMAT RESPONSE
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Err.Error(),
		})
	}

	// TODO IMRPOVE FORMAT RESPONSE
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}
