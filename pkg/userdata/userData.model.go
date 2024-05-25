package userdata

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MemberState string

const (
	MEMBER                         MemberState = "MEMBER"
	SET_MEMBER                     MemberState = "SET_MEMBER"
	APPLIED_FOR_SCALINGUP_PLATFORM MemberState = "APPLIED_FOR_SCALINGUP_PLATFORM"
	START_SCALING_UP               MemberState = "START_SCALING_UP"
	PASS_SCALING_UP                MemberState = "PASS_SCALING_UP"
	START_FUND_RAISING             MemberState = "START_FUND_RAISING"
)

type LoginVia string

const (
	SME       LoginVia = "SME"
	KEYCLOAK  LoginVia = "KEYCLOAK"
	SETMEMBER LoginVia = "SETMEMBER"
)

type UserDatas struct {
	Id                            int         `json:"id"`
	FirstName                     string      `json:"firstname"`
	LastName                      string      `json:"lastname"`
	Gender                        string      `json:"gender"`
	PhoneNumber                   string      `json:"phoneNumber,omitempty"`
	Email                         string      `json:"email"`
	EmailBK                       string      `json:"emailBk"`
	LineID                        string      `json:"lineId,omitempty"`
	CompanyName                   string      `json:"companyName"`
	JobTitleID                    int         `json:"jobTitleId,omitempty" gorm:"default:null"`
	JobTitleGroup                 string      `json:"jobTitleGroup,omitempty"`
	JobTitleOthers                string      `json:"jobTitleOthers,omitempty"`
	LoginVia                      LoginVia    `json:"loginVia,omitempty"`
	MemberState                   MemberState `json:"memberState,omitempty"`
	Entrepreneurship              string      `json:"entrepreneurship,omitempty"`
	CorporateNumber               string      `json:"corporateNumber,omitempty"`
	CompanyRegistrationTypeID     int         `json:"companyRegistrationTypeId,omitempty" gorm:"default:null"`
	CompanyRegistrationTypeOthers string      `json:"companyRegistrationTypeOthers,omitempty"`
	IndustryGroupID               int         `json:"industryGroupId,omitempty" gorm:"default:null"`
	IndustryGroupOthers           string      `json:"industryGroupOthers,omitempty"`
	AcceptSubscribe               bool        `json:"acceptSubscribe,omitempty"`
	IsInvitation                  bool        `json:"isInvitation,omitempty"`
	IsPassElearningFinalExam      bool        `json:"isPassElearningFinalExam,omitempty"`
	IsPassAdvancedFinalExam       bool        `json:"isPassAdvancedFinalExam,omitempty"`
	UserRef                       int         `json:"userRef,omitempty"`
	CreatedBy                     string      `json:"createdBy"`
	CreatedAt                     time.Time   `json:"createdAt"`
	UpdatedBy                     string      `json:"updatedBy,omitempty"`
	UpdateProfileAt               time.Time   `json:"updateProfileAt,omitempty"  gorm:"default:null"`
	UpdatedAt                     time.Time   `json:"updatedAt,omitempty"`
}

type UserDatasDbHandler struct {
	DB *gorm.DB
}
