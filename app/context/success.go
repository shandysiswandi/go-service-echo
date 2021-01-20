package context

type responseSuccess struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success is | 200, 201, 204
func (c *CustomContext) Success(s int, m string, d interface{}) error {
	return c.JSON(s, responseSuccess{
		Status:  true,
		Message: m,
		Data:    d,
	})
}
