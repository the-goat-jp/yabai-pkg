package rbac

import (
	"github.com/the-goat-jp/yabai-pkg/constants"
	"github.com/the-goat-jp/yabai-pkg/model"
)

func IsPaymentTesterByUser(user model.User) bool {
	return IsPaymentTesterByRoleID(user.RoleID)
}

func IsPaymentTesterByJWTData(jwtData model.JWTData) bool {
	return IsPaymentTesterByRoleID(jwtData.RoleID)
}

func IsPaymentTesterByRoleID(roleID int32) bool {
	return roleID == int32(constants.RoleTester)
}
