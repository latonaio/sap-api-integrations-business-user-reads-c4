package responses

type ToEmployeeBasicData struct {
	D struct {
		Results struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
