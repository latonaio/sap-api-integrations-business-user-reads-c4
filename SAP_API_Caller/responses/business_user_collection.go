package responses

type BusinessUserCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ObjectID                           string `json:"ObjectID"`
			ETag                               string `json:"ETag"`
			EmployeeID                         string `json:"EmployeeID"`
			EmployeeUUID                       string `json:"EmployeeUUID"`
			UserID                             string `json:"UserID"`
			TechnicalUserID                    string `json:"TechnicalUserID"`
			IdentityUUID                       string `json:"IdentityUUID"`
			BusinessPartnerID                  string `json:"BusinessPartnerID"`
			BusinessPartnerFormattedName       string `json:"BusinessPartnerFormattedName"`
			DepartmentName                     string `json:"DepartmentName"`
			CompanyName                        string `json:"CompanyName"`
			ManagerName                        string `json:"ManagerName"`
			EmailURI                           string `json:"EmailURI"`
			DecimalFormatCode                  string `json:"DecimalFormatCode"`
			DecimalFormatCodeText              string `json:"DecimalFormatCodeText"`
			DateFormatCode                     string `json:"DateFormatCode"`
			DateFormatCodeText                 string `json:"DateFormatCodeText"`
			TimeFormatCode                     string `json:"TimeFormatCode"`
			TimeFormatCodeText                 string `json:"TimeFormatCodeText"`
			TimeZoneCode                       string `json:"TimeZoneCode"`
			TimeZoneCodeText                   string `json:"TimeZoneCodeText"`
			LogonLanguageCode                  string `json:"LogonLanguageCode"`
			LogonLanguageCodeText              string `json:"LogonLanguageCodeText"`
			UserValidityStartDate              string `json:"UserValidityStartDate"`
			UserValidityEndDate                string `json:"UserValidityEndDate"`
			UserLockedIndicator                bool   `json:"UserLockedIndicator"`
			UserCountedIndicator               bool   `json:"UserCountedIndicator"`
			PasswordPolicyCode                 string `json:"PasswordPolicyCode"`
			PasswordPolicyCodeText             string `json:"PasswordPolicyCodeText"`
			PasswordInactiveIndicator          bool   `json:"PasswordInactiveIndicator"`
			PasswordLockedIndicator            bool   `json:"PasswordLockedIndicator"`
			UserAccountTypeCode                string `json:"UserAccountTypeCode"`
			UserAccountTypeCodeText            string `json:"UserAccountTypeCodeText"`
			CreatedOn                          string `json:"CreatedOn"`
			UserCreatedBy                      string `json:"UserCreatedBy"`
			EntityLastChangedOn                string `json:"EntityLastChangedOn"`
			UserChangedBy                      string `json:"UserChangedBy"`
			UserChangedOn                      string `json:"UserChangedOn"`
			BusinessUserBusinessRoleAssignment struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"BusinessUserBusinessRoleAssignment"`
			EmployeeBasicData struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"EmployeeBasicData"`
		} `json:"results"`
	} `json:"d"`
}
