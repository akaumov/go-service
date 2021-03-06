//!!!GENERATED BY "GO-SERVICE" DON'T CHANGE THIS FILE!!!
package executor

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"

	"github.com/akaumov/go-service/exchange"
)

type Executor struct {
	handler HandlerInterface
}

type SessionInterface interface {
	GetUserId() string
	GetSessionId() string
}

func NewExecutor(handler HandlerInterface) *Executor {
	return &Executor{
		handler: handler,
	}
}

func (e *Executor) Execute(session SessionInterface, packedMessage *[]byte) (*[]byte, error) {
	if packedMessage == nil {
		return nil, errors.New("message text is required")
	}

	response := e.execute(session, packedMessage)
	packed, _ := json.Marshal(response)
	return &packed, nil
}

func (e *Executor) execute(session SessionInterface, packedMessage *[]byte) exchange.ResponseMessage {

	var requestMessage exchange.RequestMessage
	err := json.Unmarshal(*packedMessage, &requestMessage)
	if err != nil {
		return exchange.NewErrorResponse("", "WrongRequest", "can't parse message")
	}

	requestId := requestMessage.Id

	switch requestMessage.Method {

	case "getBook":
		var params GetBookParams
		err := json.Unmarshal(requestMessage.Params, &params)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't parse params: %v", err))
		}

		err = params.Validate()
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't wrong params: %v", err))
		}

		result, err := e.handler.GetBook(session, params.Id)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "ServerError", err.Error())
		}

		return exchange.NewResultResponse(requestId, result)

	case "getBooks":
		var params GetBooksParams
		err := json.Unmarshal(requestMessage.Params, &params)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't parse params: %v", err))
		}

		err = params.Validate()
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't wrong params: %v", err))
		}

		result, err := e.handler.GetBooks(session, params.Id)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "ServerError", err.Error())
		}

		return exchange.NewResultResponse(requestId, result)

	case "getAuthor":
		var params GetAuthorParams
		err := json.Unmarshal(requestMessage.Params, &params)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't parse params: %v", err))
		}

		err = params.Validate()
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't wrong params: %v", err))
		}

		result, err := e.handler.GetAuthor(session, params.Id)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "ServerError", err.Error())
		}

		return exchange.NewResultResponse(requestId, result)

	case "getAuthors":
		var params GetAuthorsParams
		err := json.Unmarshal(requestMessage.Params, &params)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't parse params: %v", err))
		}

		err = params.Validate()
		if err != nil {
			return exchange.NewErrorResponse(requestId, "WrongRequest", fmt.Sprintf("can't wrong params: %v", err))
		}

		result, err := e.handler.GetAuthors(session, params.Id)
		if err != nil {
			return exchange.NewErrorResponse(requestId, "ServerError", err.Error())
		}

		return exchange.NewResultResponse(requestId, result)

	}

	return exchange.NewErrorResponse("", "WrongRequest", "no such method")
}
