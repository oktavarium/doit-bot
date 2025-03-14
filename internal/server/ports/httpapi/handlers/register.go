package handlers

// func (h *Handlers) Register(c *gin.Context) {
// 	initData, ok := common.InitDataFromContext(c)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, ""))
// 		return
// 	}

// 	if err := h.app.CreateUser(
// 		c,
// 		initData.User.ID,
// 		initData.Chat.ID,
// 		initData.User.FirstName,
// 		initData.User.LastName,
// 		initData.User.Username,
// 	); err != nil {
// 		c.JSON(http.StatusInternalServerError, common.NewStatusResponse(http.StatusInternalServerError, err.Error()))
// 		return
// 	}

// 	c.JSON(http.StatusOK, common.NewStatusResponse(http.StatusOK, "welcome!"))
// }
