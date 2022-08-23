package utils

var (
	IdVerify               = Rules{"ID": {NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
	CourseInfoVerify       = Rules{"Abstract": {NotEmpty()}, "Desc": {NotEmpty()}, "Title": {NotEmpty()}}
	UserLevelPriceVerify   = Rules{"Name": {NotEmpty()}, "Grade": {NotEmpty()}, "Price": {NotEmpty()}, "Type": {NotEmpty()}}
	MemberInterestsVerify  = Rules{"Title": {NotEmpty()}, "Linked": {NotEmpty()}}
	CourseSectionVerify    = Rules{"CourseId": {NotEmpty()}, "CoverImage": {NotEmpty()}, "Lessons": {NotEmpty()}, "LongTime": {NotEmpty()}, "Title": {NotEmpty()}, "VideoUrl": {NotEmpty()}}
	ContactUsVerify        = Rules{"Account": {NotEmpty()}, "AccountName": {NotEmpty()}, "Icon": {NotEmpty()}}
	PlatformRegisterVerify = Rules{"BackgroundImage": {NotEmpty()}, "Linked": {NotEmpty()}}
)
