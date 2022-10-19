package response

type LoginResp struct {
	ResultObj  ResultObj `json:"ResultObj"`
	Status     int       `json:"Status"`
	StatusCode int       `json:"StatusCode"`
	Msg        string    `json:"Msg"`
	ErrorObj   ErrorObj  `json:"ErrorObj"`
}
type ResultObj struct {
	UserID             int    `json:"UserID"`
	UserName           string `json:"UserName"`
	Email              string `json:"Email"`
	Telphone           string `json:"Telphone"`
	Gender             bool   `json:"Gender"`
	CollegeID          int    `json:"CollegeID"`
	CollegeName        string `json:"CollegeName"`
	RoleName           string `json:"RoleName"`
	RoleID             int    `json:"RoleID"`
	AccessToken        string `json:"AccessToken"`
	AccessTokenErrCode int    `json:"AccessTokenErrCode"`
	ReturnURL          string `json:"ReturnUrl"`
	DataToken          string `json:"DataToken"`
}
type ErrorObj struct {
}

// cmd resp status
const (
	Success = iota
	Failure
	Exception
	Unknown
)

type CmdResp struct {
	ResultObj  string      `json:"ResultObj"`
	Status     int         `json:"Status"`
	StatusCode int         `json:"StatusCode"`
	Msg        string      `json:"Msg"`
	ErrorObj   interface{} `json:"ErrorObj"`
}

//使用了https://mholt.github.io/json-to-go/
//快速生成resp
