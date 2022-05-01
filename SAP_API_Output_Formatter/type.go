package sap_api_output_formatter

type BusinessUser struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	APISchema         string `json:"api_schema"`
	BusisnessUserCode string `json:"Business_user_code"`
	Deleted           bool   `json:"deleted"`
}
type BusinessUserCollection struct {
	ObjectID                             string `json:"ObjectID"`
	ETag                                 string `json:"ETag"`
	EmployeeID                           string `json:"EmployeeID"`
	EmployeeUUID                         string `json:"EmployeeUUID"`
	UserID                               string `json:"UserID"`
	TechnicalUserID                      string `json:"TechnicalUserID"`
	IdentityUUID                         string `json:"IdentityUUID"`
	BusinessPartnerID                    string `json:"BusinessPartnerID"`
	BusinessPartnerFormattedName         string `json:"BusinessPartnerFormattedName"`
	DepartmentName                       string `json:"DepartmentName"`
	CompanyName                          string `json:"CompanyName"`
	ManagerName                          string `json:"ManagerName"`
	EmailURI                             string `json:"EmailURI"`
	DecimalFormatCode                    string `json:"DecimalFormatCode"`
	DecimalFormatCodeText                string `json:"DecimalFormatCodeText"`
	DateFormatCode                       string `json:"DateFormatCode"`
	DateFormatCodeText                   string `json:"DateFormatCodeText"`
	TimeFormatCode                       string `json:"TimeFormatCode"`
	TimeFormatCodeText                   string `json:"TimeFormatCodeText"`
	TimeZoneCode                         string `json:"TimeZoneCode"`
	TimeZoneCodeText                     string `json:"TimeZoneCodeText"`
	LogonLanguageCode                    string `json:"LogonLanguageCode"`
	LogonLanguageCodeText                string `json:"LogonLanguageCodeText"`
	UserValidityStartDate                string `json:"UserValidityStartDate"`
	UserValidityEndDate                  string `json:"UserValidityEndDate"`
	UserLockedIndicator                  bool   `json:"UserLockedIndicator"`
	UserCountedIndicator                 bool   `json:"UserCountedIndicator"`
	PasswordPolicyCode                   string `json:"PasswordPolicyCode"`
	PasswordPolicyCodeText               string `json:"PasswordPolicyCodeText"`
	PasswordInactiveIndicator            bool   `json:"PasswordInactiveIndicator"`
	PasswordLockedIndicator              bool   `json:"PasswordLockedIndicator"`
	UserAccountTypeCode                  string `json:"UserAccountTypeCode"`
	UserAccountTypeCodeText              string `json:"UserAccountTypeCodeText"`
	CreatedOn                            string `json:"CreatedOn"`
	UserCreatedBy                        string `json:"UserCreatedBy"`
	EntityLastChangedOn                  string `json:"EntityLastChangedOn"`
	UserChangedBy                        string `json:"UserChangedBy"`
	UserChangedOn                        string `json:"UserChangedOn"`
	ToBusinessUserBusinessRoleAssignment string `json:"BusinessUserBusinessRoleAssignment"`
	ToBusinessUserSubscriptionAssignment string `json:"BusinessUserSubscriptionAssignment"`
	ToEmployeeBasicData                  string `json:"EmployeeBasicData"`
}

type ToBusinessUserBusinessRoleAssignment struct {
	ObjectID            string `json:"ObjectID"`
	ParentObjectID      string `json:"ParentObjectID"`
	EmployeeID          string `json:"EmployeeID"`
	UserID              string `json:"UserID"`
	BusinessRoleID      string `json:"BusinessRoleID"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
}

type ToEmployeeBasicData struct {
	ObjectID                         string `json:"ObjectID"`
	ETag                             string `json:"ETag"`
	EmployeeID                       string `json:"EmployeeID"`
	EmployeeUUID                     string `json:"EmployeeUUID"`
	UserID                           string `json:"UserID"`
	IdentityUUID                     string `json:"IdentityUUID"`
	BusinessPartnerID                string `json:"BusinessPartnerID"`
	CurrentInternalEmployeeIndicator bool   `json:"CurrentInternalEmployeeIndicator"`
	CurrentExternalEmployeeIndicator bool   `json:"CurrentExternalEmployeeIndicator"`
	FormattedName                    string `json:"FormattedName"`
	TitleCode                        string `json:"TitleCode"`
	AcademicTitleCode                string `json:"AcademicTitleCode"`
	FirstName                        string `json:"FirstName"`
	MiddleName                       string `json:"MiddleName"`
	LastName                         string `json:"LastName"`
	SecondLastName                   string `json:"SecondLastName"`
	NickName                         string `json:"NickName"`
	GenderCode                       string `json:"GenderCode"`
	LanguageCode                     string `json:"LanguageCode"`
	FormattedAddress                 string `json:"FormattedAddress"`
	CountryCode                      string `json:"CountryCode"`
	RegionCode                       string `json:"RegionCode"`
	AddressLine1                     string `json:"AddressLine1"`
	AddressLine2                     string `json:"AddressLine2"`
	HouseNumber                      string `json:"HouseNumber"`
	Street                           string `json:"Street"`
	AddressLine4                     string `json:"AddressLine4"`
	AddressLine5                     string `json:"AddressLine5"`
	City                             string `json:"City"`
	PostalCode                       string `json:"PostalCode"`
	Phone                            string `json:"Phone"`
	Mobile                           string `json:"Mobile"`
	Fax                              string `json:"Fax"`
	Email                            string `json:"Email"`
	UserValidityStartDate            string `json:"UserValidityStartDate"`
	UserValidityEndDate              string `json:"UserValidityEndDate"`
	UserPasswordPolicyCode           string `json:"UserPasswordPolicyCode"`
	UserLockedIndicator              bool   `json:"UserLockedIndicator"`
	TimeZoneCode                     string `json:"TimeZoneCode"`
	ManagerUUID                      string `json:"ManagerUUID"`
	ManagerFormattedName             string `json:"ManagerFormattedName"`
	JobName                          string `json:"JobName"`
	CreatedOn                        string `json:"CreatedOn"`
	CreatedBy                        string `json:"CreatedBy"`
	ChangedOn                        string `json:"ChangedOn"`
	ChangedBy                        string `json:"ChangedBy"`
	EntityLastChangedOn              string `json:"EntityLastChangedOn"`
	ToBusinessUser                   string `json:"BusinessUser"`
	ToManagerEmployeeBasicData       string `json:"ManagerEmployeeBasicData"`
}

type BusinessUserBusinessRoleAssignment struct {
	ObjectID            string `json:"ObjectID"`
	ParentObjectID      string `json:"ParentObjectID"`
	EmployeeID          string `json:"EmployeeID"`
	UserID              string `json:"UserID"`
	BusinessRoleID      string `json:"BusinessRoleID"`
	EntityLastChangedOn string `json:"EntityLastChangedOn"`
}

type EmployeeBasicData struct {
	ObjectID                         string `json:"ObjectID"`
	ETag                             string `json:"ETag"`
	EmployeeID                       string `json:"EmployeeID"`
	EmployeeUUID                     string `json:"EmployeeUUID"`
	UserID                           string `json:"UserID"`
	IdentityUUID                     string `json:"IdentityUUID"`
	BusinessPartnerID                string `json:"BusinessPartnerID"`
	CurrentInternalEmployeeIndicator bool   `json:"CurrentInternalEmployeeIndicator"`
	CurrentExternalEmployeeIndicator bool   `json:"CurrentExternalEmployeeIndicator"`
	FormattedName                    string `json:"FormattedName"`
	TitleCode                        string `json:"TitleCode"`
	AcademicTitleCode                string `json:"AcademicTitleCode"`
	FirstName                        string `json:"FirstName"`
	MiddleName                       string `json:"MiddleName"`
	LastName                         string `json:"LastName"`
	SecondLastName                   string `json:"SecondLastName"`
	NickName                         string `json:"NickName"`
	GenderCode                       string `json:"GenderCode"`
	LanguageCode                     string `json:"LanguageCode"`
	FormattedAddress                 string `json:"FormattedAddress"`
	CountryCode                      string `json:"CountryCode"`
	RegionCode                       string `json:"RegionCode"`
	AddressLine1                     string `json:"AddressLine1"`
	AddressLine2                     string `json:"AddressLine2"`
	HouseNumber                      string `json:"HouseNumber"`
	Street                           string `json:"Street"`
	AddressLine4                     string `json:"AddressLine4"`
	AddressLine5                     string `json:"AddressLine5"`
	City                             string `json:"City"`
	PostalCode                       string `json:"PostalCode"`
	Phone                            string `json:"Phone"`
	Mobile                           string `json:"Mobile"`
	Fax                              string `json:"Fax"`
	Email                            string `json:"Email"`
	UserValidityStartDate            string `json:"UserValidityStartDate"`
	UserValidityEndDate              string `json:"UserValidityEndDate"`
	UserPasswordPolicyCode           string `json:"UserPasswordPolicyCode"`
	UserLockedIndicator              bool   `json:"UserLockedIndicator"`
	TimeZoneCode                     string `json:"TimeZoneCode"`
	ManagerUUID                      string `json:"ManagerUUID"`
	ManagerFormattedName             string `json:"ManagerFormattedName"`
	JobName                          string `json:"JobName"`
	CreatedOn                        string `json:"CreatedOn"`
	CreatedBy                        string `json:"CreatedBy"`
	ChangedOn                        string `json:"ChangedOn"`
	ChangedBy                        string `json:"ChangedBy"`
	EntityLastChangedOn              string `json:"EntityLastChangedOn"`
	ToBusinessUser                   string `json:"BusinessUser"`
	ToManagerEmployeeBasicData       string `json:"ManagerEmployeeBasicData"`
}
