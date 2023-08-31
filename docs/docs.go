// Package docs School Management  API.
//
// 学生信息管理系统.
//
// Version: 1.0.0
// Schemes: http
// Host: localhost:8080
// BasePath: /schoolManagement
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
//			SecurityDefinitions:
//		 BearerAuth:
//		      type: apiKey
//		      name: Authorization
//	       in: header
//
// swagger:meta
package docs

import "GolandProjects/School-Management/bean"

// swagger:route POST /user/login Public login
// 学生或管理员登录接口
// responses:
//	  default: response
//    201: withTokenResponse

// swagger:route POST /user/createUser Public createUser
// 学生注册接口
// responses:
//	  default: response

// swagger:route PUT /user/updateUser User updateUser
// 学生修改个人信息接口
// Security:
//   - BearerAuth: []
// responses:
//	  default: response
//    201: withTokenResponse

// swagger:route DELETE /user/deleteUser User deleteUser
// 学生删除个人信息接口
// Security:
//   - BearerAuth: []
// responses:
//	  default: response

// swagger:route GET /user/getUser User getUser
// 获取学生/管理员信息接口
// responses:
//	  default: response
// Security:
//   - BearerAuth: []

// swagger:route PUT /admin/updateUser Admin adminUpdateUser
// 管理员修改学生个人信息接口
// Security:
//   - BearerAuth: []
// responses:
//	  default: response

// swagger:route GET /admin/getAllUser Admin getAllUser
// 管理员修改学生个人信息接口
// Security:
//   - BearerAuth: []
// responses:
//	  default: response

// swagger:parameters login
type userLoginRequestWrapper struct {
	// 用于学生和管理员登录 需要提供 学号（学工号） + 密码
	// in:body
	Body bean.UserLoginRequest
}

// swagger:parameters createUser
type registerResponseWrapper struct {
	// 注册需要提供 学号 姓名 密码 年级
	// in:body
	Body bean.UserCreateRequest
}

// swagger:parameters updateUser
type updateUserInfoRequestWrapper struct {
	// 修改密码则需要提供旧密码和新密码 不允许改变学号
	// in:body
	Body bean.UpdateUserInfoRequest
}

// swagger:parameters deleteUser
type deleteRequestWrapper struct {
	// 学生删除自己账户需要提供密码, 管理员删除学生账户不需要提供密码
	// in:body
	Body bean.UserDeleteRequest
}

// 默认返回的的Response
// swagger:response response
type resultWrapper struct {
	// in:body
	Body bean.CommonResult
}

// 加上token返回的Response
// swagger:response withTokenResponse
type withTokenResult struct {
	Body bean.WithTokenResult
}

// swagger:parameters adminUpdateUser
type adminUpdateUserRequestWrapper struct {
	// 管理员修改用户个人信息
	// in:body
	Body bean.AdminUpdateUserRequest
}
