package testin

type CheckAddInp struct {
	Name string `json:"name" v:"required#名称不能为空"`
}
