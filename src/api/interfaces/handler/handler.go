
package handler

type AppHandler interface {
	GetUserHandler	
}

type appHandler struct {
	GetUserHandler
}

func NewAppHandler(
	h1 GetUserHandler,
) AppHandler {
	return &appHandler{
		h1,
	}
} 