package models

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"my-ganji-app/common"
	rds_conn "my-ganji-app/redis"
	"my-ganji-app/types"
	type_user "my-ganji-app/types/user"
)

type User struct {
	BaseModel
	Id             int64  `json:"id"`
	Phone          string `orm:"size(64);index" json:"phone"`
	UserName       string `orm:"size(128)" json:"user_name"`
	Avatar         string `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"`
	Password       string `orm:"size(128)" json:"password"`
	FundPassword   string `orm:"size(128)" json:"fund_password"` // 钱包资金密码
	Email          string `orm:"size(128);index" json:"email"`
	LoginCount     int64  `orm:"default(0);index" json:"login_count"`
	Token          string `orm:"size(128)" json:"token"`
	IsAuth         int8   `orm:"default(0);index" json:"is_auth"`         // 0 未实名认证，1: 实名认证中；2:实名认证成功；3实名认证失败
	MemberLevel    int8   `orm:"default(1);index" json:"member_level"`    // 0:v0:普通会员 1:V1:白银会员，2:V2:白金会员，3:V3:黄金会员; 4:V4:砖石会有; 5:V5:皇冠会员
	MyInviteCode   string `orm:"size(128)" json:"my_invite_code"`         // 用户自己网体邀请码
	InviteMeUserId int64  `orm:"size(64);index" json:"invite_me_user_id"` // 网体上级用户id
}

func (u *User) TableName() string {
	return common.TableName("users")
}

// SearchField 返回用户表可模糊搜索字段
func (u *User) SearchField() []string {
	return []string{"user_name"}
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}

func (u *User) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(u); err != nil {
		return err, 0
	}
	return nil, id
}

func (u *User) ExistByPhone(phone string) bool {
	return orm.NewOrm().QueryTable(u).Filter("phone", phone).Exist()
}

func (u *User) ExistByEmail(email string) bool {
	return orm.NewOrm().QueryTable(u).Filter("email", email).Exist()
}

func (u *User) GetInviteMeUser(inviteCode string) (*User, error) {
	var user User
	err := user.Query().Filter("my_invite_code", inviteCode).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id int64) (User, error) {
	var queryUser User
	err := queryUser.Query().Filter("Id", id).Limit(1).One(&queryUser)
	if err != nil {
		return queryUser, errors.New("user is not exist")
	}
	return queryUser, nil
}

func (u *User) GetUserByEmail(email string) (User, error) {
	var queryUser User
	err := queryUser.Query().Filter("email", email).Limit(1).One(&queryUser)
	if err != nil {
		return queryUser, errors.New("user is not exist")
	}
	return queryUser, nil
}

func (u *User) GetUserByPhone(phone string) (User, error) {
	var queryUser User
	err := queryUser.Query().Filter("phone", phone).Limit(1).One(&queryUser)
	if err != nil {
		return queryUser, errors.New("user is not exist")
	}
	return queryUser, nil
}

// AddLoginCount 用户登录次数自增，并更新修改时间
func (u *User) AddLoginCount() error {
	u.LoginCount += 1
	u.UpdatedAt = time.Now()
	// 仅更新 LoginCount 和 UpdatedAt 两个字段
	return u.Update("LoginCount", "UpdatedAt")
}

func GetUserByToken(token string) (*User, error) {
	u := User{}
	if err := orm.NewOrm().QueryTable(u.TableName()).RelatedSel().Filter("token", token).One(&u); err != nil {
		return nil, errors.Wrap(err, "error in GetUserByToken")
	}
	return &u, nil
}

// RegisterByPhoneOrEmail 通过手机号或者邮箱注册
func RegisterByPhoneOrEmail(ctx context.Context, registerParam type_user.UserRegisterCheck) (success bool, code int, err error) {
	u := User{}
	var phone, email string
	if registerParam.VerifyWay == 1 { // 1: 手机号验证
		if u.ExistByPhone(registerParam.PhoneEmail) {
			return false, types.UserIsExist, errors.New("用户已经注册")
		}
		// 这里 PhoneEmail 不是拼接字段 而是 方式一的时候 传纯手机号
		phone = registerParam.PhoneEmail
	} else if registerParam.VerifyWay == 2 { // 2: 邮箱验证
		if u.ExistByEmail(registerParam.PhoneEmail) {
			return false, types.UserIsExist, errors.New("用户已经注册")
		}
		email = registerParam.PhoneEmail
	} else {
		return false, types.InvalidVerifyWay, errors.New("无效的校验方式")
	}

	var inviteMeUserID int64
	if len(registerParam.InviteCode) > 0 {
		inviteMeUser, err := u.GetInviteMeUser(registerParam.InviteCode)
		if err != nil {
			return false, types.InviteCodeNotExist, errors.New("对不起，没有这个邀请码，请核对后输入")
		}
		inviteMeUserID = inviteMeUser.Id
		itGral, _ := GetIntegralByUserId(inviteMeUserID)
		itGral.TotalIg += 50
		if err := itGral.Update("TotalIg"); err != nil {
			return false, types.InsertIntegralFail, errors.New("更新上级用户积分失败")
		}
	} else {
		inviteMeUserID = 0
	}

	token, inviteCode := uidCode()
	passwordHash, err := common.HashPassword(registerParam.Password1)
	if err != nil {
		return false, types.CreateUserFail, errors.New("密码加密失败")
	}
	userRegData := User{
		Phone:          phone,
		UserName:       "小鱼儿",
		Password:       passwordHash,
		Email:          email,
		Token:          token,
		InviteMeUserId: inviteMeUserID,
		MyInviteCode:   inviteCode,
	}

	err, userId := userRegData.Insert()
	if err != nil {
		return false, types.CreateUserFail, errors.New("创建用户失败")
	}

	userInfoQuery := UserInfo{
		UserId: userId,
	}
	if err := userInfoQuery.Insert(); err != nil {
		return false, types.CreateUserFail, errors.New("创建用户信息失败")
	}

	userWallet := UserWallet{
		UserId:      userId,
		AssetName:   "人民币",
		TotalAmount: 0,
	}
	if err := userWallet.Insert(); err != nil {
		return false, types.CreateUserWalletFail, errors.New("创建用户钱包失败")
	}

	crfrIntegral := UserIntegral{
		UserId:       userId,
		IntegralName: "商场积分",
		TotalIg:      0,
		UsedIg:       0,
		TodayIg:      0,
	}
	if err := crfrIntegral.Insert(); err != nil {
		return false, types.InsertIntegralFail, errors.New("插入积分失败")
	}
	currentTime := time.Now()
	endTime := currentTime.AddDate(0, 0, 30)
	userCoupon := UserCoupon{
		UserId:      userId,
		CouponName:  "优惠券",
		IsUsed:      0,
		TotalAmount: 10,
		StartTime:   &currentTime,
		EndTime:     &endTime,
		IsInvalid:   0,
	}
	if err := userCoupon.Insert(); err != nil {
		return false, types.InsertIntegralFail, errors.New("插入优惠券失败")
	}
	crfrTree := CrfrUserTree{
		UserId:       userId,
		FatherUserId: inviteMeUserID,
	}
	if err := crfrTree.Insert(); err != nil {
		return false, types.InsertIntegralFail, errors.New("构建User Tree错误")
	}
	// 删除 redis 中的验证码 防止同一个验证码在过期前还能再注册一次
	_ = rds_conn.RdsConn.Del(ctx, registerParam.PhoneEmail).Err()
	return true, types.ReturnSuccess, nil
}

// LoginByPhoneOrEmail 通过手机号码，邮箱登陆
func LoginByPhoneOrEmail(login_param type_user.UserLoginCheck) (bool, *type_user.UserLoginRet, int, error) {
	u := User{}
	if login_param.VerifyWay == 1 { // 1: 手机号码验证
		ret_user, err := u.GetUserByPhone(login_param.PhoneEmail)
		if err != nil {
			return false, nil, types.UserNotRegister, errors.New("用户没有注册")
		}
		if !common.CheckPassword(ret_user.Password, login_param.Password) {
			return false, nil, types.PasswordError, errors.New("输入密码错误")
		}
		if err := u.AddLoginCount(); err != nil {
			return false, nil, types.AddLoginTimesError, errors.New("添加登陆次数错误")
		}
		return true, &type_user.UserLoginRet{
			Id:       ret_user.Id,
			UserName: ret_user.UserName,
			Token:    ret_user.Token,
			Phone:    ret_user.Phone,
		}, types.ReturnSuccess, nil
	} else if login_param.VerifyWay == 2 { // 2: 邮箱验证
		ret_user, err := u.GetUserByEmail(login_param.PhoneEmail)
		if err != nil {
			return false, nil, types.UserNotRegister, errors.New("用户没有注册")
		}
		if !common.CheckPassword(ret_user.Password, login_param.Password) {
			return false, nil, types.PasswordError, errors.New("输入密码")
		}
		if err := u.AddLoginCount(); err != nil {
			return false, nil, types.AddLoginTimesError, errors.New("添加登陆次数错误")
		}
		return true, &type_user.UserLoginRet{
			Id:    ret_user.Id,
			Token: ret_user.Token,
			Phone: ret_user.Phone,
		}, types.ReturnSuccess, nil
	} else {
		return false, nil, types.InvalidVerifyWay, errors.New("没有这种验证方式")
	}
}

func uidCode() (string, string) {
	tokenID := uuid.New()
	inviteID := uuid.New()
	b, _ := inviteID.MarshalBinary() // []byte, len=16
	return tokenID.String(), base64.RawURLEncoding.EncodeToString(b)
}
