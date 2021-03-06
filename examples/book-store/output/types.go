package executor

import (
	"encoding/json"
	"fmt"
	validator "github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
)

type Validatable interface {
	Validate() error
}

/////////////////////////////////////////////////////////////////////
//BookType

type BookType string

const (
	Magazine BookType = "magazineItem"
	BookItem BookType = "book"
)

func (v BookType) Validate() error {
	switch v {
	case Magazine:
		return nil
	case BookItem:
		return nil

	}

	return fmt.Errorf("no such value: %v", v)
}

/////////////////////////////////////////////////////////////////////
//Book

type Book struct {
	Type      BookType `json:"type"`
	Id        string   `json:"id"`
	AuthorId  string   `json:"authorId"`
	CreatedAt int64    `json:"createdAt"`
	Title     string   `json:"title"`
}

func (v *Book) Validate() error {

	{
		value := v.Type
		isValid := value.Validate() == nil

		if !isValid {
			return fmt.Errorf("Type is invalid")
		}
	}

	{
		value := v.Id
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("Id is invalid")
		}
	}

	{
		value := v.AuthorId
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("AuthorId is invalid")
		}
	}

	{
		value := v.Title
		isValid := (len(value) >= 0 && len(value) <= 255)

		if !isValid {
			return fmt.Errorf("Title is invalid")
		}
	}

	return nil
}

func (v Book) String() string {
	jsonRepresentation, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonRepresentation)
}

/////////////////////////////////////////////////////////////////////
//Author

type Author struct {
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Id         string  `json:"id"`
	Name       string  `json:"name"`
}

func (v *Author) Validate() error {

	{
		value := v.Surname
		isValid := (len(value) >= 0 && len(value) <= 255)

		if !isValid {
			return fmt.Errorf("Surname is invalid")
		}
	}

	{
		value := v.Patronymic
		isValid := value == nil || (len(*value) >= 0 && len(*value) <= 255)

		if !isValid {
			return fmt.Errorf("Patronymic is invalid")
		}
	}

	{
		value := v.Id
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("Id is invalid")
		}
	}

	{
		value := v.Name
		isValid := (len(value) >= 0 && len(value) <= 255)

		if !isValid {
			return fmt.Errorf("Name is invalid")
		}
	}

	return nil
}

func (v Author) String() string {
	jsonRepresentation, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonRepresentation)
}

/////////////////////////////////////////////////////////////////////
//PARAMETERS

/////////////////////////////////////////////////////////////////////
//getBook

type GetBookParams struct {
	Id string `json:"id"`
}

func (v *GetBookParams) Validate() error {

	{
		value := v.Id
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("Id is invalid")
		}
	}

	return nil
}

func (v GetBookParams) String() string {
	jsonRepresentation, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonRepresentation)
}

/////////////////////////////////////////////////////////////////////
//getBooks

type GetBooksParams struct {
	Id string `json:"id"`
}

func (v *GetBooksParams) Validate() error {

	{
		value := v.Id
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("Id is invalid")
		}
	}

	return nil
}

func (v GetBooksParams) String() string {
	jsonRepresentation, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonRepresentation)
}

/////////////////////////////////////////////////////////////////////
//getAuthor

type GetAuthorParams struct {
	Id string `json:"id"`
}

func (v *GetAuthorParams) Validate() error {

	{
		value := v.Id
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("Id is invalid")
		}
	}

	return nil
}

func (v GetAuthorParams) String() string {
	jsonRepresentation, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonRepresentation)
}

/////////////////////////////////////////////////////////////////////
//getAuthors

type GetAuthorsParams struct {
	Id string `json:"id"`
}

func (v *GetAuthorsParams) Validate() error {

	{
		value := v.Id
		isValid := validator.IsUUID(value)

		if !isValid {
			return fmt.Errorf("Id is invalid")
		}
	}

	return nil
}

func (v GetAuthorsParams) String() string {
	jsonRepresentation, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonRepresentation)
}
