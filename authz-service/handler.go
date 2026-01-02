package authzservice


import (
	"net/http"

)

type handler struct {

}

type payLoad struct {
	Sub string 
	Email string
}

func (h *handler) authorizationHandler {
	
}