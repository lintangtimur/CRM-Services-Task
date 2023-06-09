package accounts

import (
	"CRM-Services-Task/accounts/dto"
	"CRM-Services-Task/accounts/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	ctrl *Controller
}

func (h RequestHandler) ActorLogin(c *gin.Context) {
	loginRequest := dto.LoginRequest{}
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
	admin := dto.AdminRegisterRequest{}

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
	approveReq := dto.ApproveRequest{}
	if err := c.ShouldBindJSON(&approveReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing param"})
		return
	}

	actors := entity.Actor{}
	regAp := entity.RegisterApproval{}

	res, err := h.ctrl.ApproveAdmin(&actors, &approveReq, &regAp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Reject(c *gin.Context) {
	rejectRequest := dto.ApproveRequest{}
	if err := c.ShouldBindJSON(&rejectRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing param"})
		return
	}

	actors := entity.Actor{}
	ra := entity.RegisterApproval{}
	res, err := h.ctrl.RejectAdmin(&actors, &rejectRequest, &ra)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//update status di approval request

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) Activate(c *gin.Context) {
	req := dto.ActivateAdminRequest{}

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

func (h RequestHandler) Deactivate(c *gin.Context) {
	req := dto.DeActivateAdminRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ctrl.DeactivateAdmin(&req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) DeleteAdmin(c *gin.Context) {
	dar := dto.DeleteAdminRequest{}

	if err := c.ShouldBindJSON(&dar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.ctrl.DeleteAdmin(&dar)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h RequestHandler) GetAllAdmins(con *gin.Context) {
	res, err := h.ctrl.GetAllAdmin(con)

	if err != nil {
		con.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	con.JSON(http.StatusOK, res)
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
