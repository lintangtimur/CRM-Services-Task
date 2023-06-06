package accounts

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	ctrl *Controller
}

func (h RequestHandler) ActorLogin(c *gin.Context) {
	loginRequest := LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ctrl.Login(&loginRequest)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) CreateAdmin(c *gin.Context) {
	admin := AdminRegisterRequest{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rh.ctrl.CreateAdmin(&admin)

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetApprove(c *gin.Context) {

	res, err := h.ctrl.GetApproveList()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h RequestHandler) Approve(c *gin.Context) {
	approveReq := ApproveRequest{}
	if err := c.ShouldBindJSON(&approveReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing param"})
		return
	}

	actors := Actor{}
	regAp := RegisterApproval{}

	res, err := h.ctrl.ApproveAdmin(&actors, &approveReq, &regAp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Reject(c *gin.Context) {
	rejectRequest := ApproveRequest{}
	if err := c.ShouldBindJSON(&rejectRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing param"})
		return
	}

	actors := Actor{}
	ra := RegisterApproval{}
	res, err := h.ctrl.RejectAdmin(&actors, &rejectRequest, &ra)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//update status di approval request

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Activate(c *gin.Context) {
	req := ActivateAdminRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ctrl.ActivateAdmin(&req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{ctrl: ctrl}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}
