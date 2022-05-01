package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-business-user-reads-c4/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBusinessUser(userID, businessRoleID, employeeID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "BusinessUserCollection":
			func() {
				c.BusinessUserCollection(userID)
				wg.Done()
			}()
		case "BusinessRoleAssignment":
			func() {
				c.BusinessRoleAssignment(businessRoleID)
				wg.Done()
			}()
		case "EmployeeBasicData":
			func() {
				c.EmployeeBasicData(employeeID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) BusinessUserCollection(userID string) {
	businessUserCollectionData, err := c.callBusinessUserSrvAPIRequirementBusinessUserCollection("BusinessUserCollection", userID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessUserCollectionData)

	businessRoleAssignmentData, err := c.callToBusinessRoleAssignment(businessUserCollectionData[0].ToBusinessUserBusinessRoleAssignment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessRoleAssignmentData)

	employeeBasicDataData, err := c.callToEmployeeBasicData(businessUserCollectionData[0].ToEmployeeBasicData)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(employeeBasicDataData)

}

func (c *SAPAPICaller) callBusinessUserSrvAPIRequirementBusinessUserCollection(api, userID string) ([]sap_api_output_formatter.BusinessUserCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBusinessUserCollection(req, userID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessUserCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBusinessRoleAssignment(url string) ([]sap_api_output_formatter.ToBusinessUserBusinessRoleAssignment, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBusinessUserBusinessRoleAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToEmployeeBasicData(url string) (*sap_api_output_formatter.ToEmployeeBasicData, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToEmployeeBasicData(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BusinessRoleAssignment(businessRoleID string) {
	data, err := c.callBusinessUserSrvAPIRequirementBusinessRoleAssignment("BusinessUserBusinessRoleAssignmentCollection", businessRoleID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBusinessUserSrvAPIRequirementBusinessRoleAssignment(api, businessRoleID string) ([]sap_api_output_formatter.BusinessUserBusinessRoleAssignment, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBusinessRoleAssignment(req, businessRoleID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessUserBusinessRoleAssignment(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) EmployeeBasicData(employeeID string) {
	data, err := c.callBusinessUserSrvAPIRequirementEmployeeBasicData("EmployeeBasicDataCollection", employeeID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBusinessUserSrvAPIRequirementEmployeeBasicData(api, employeeID string) ([]sap_api_output_formatter.EmployeeBasicData, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithEmployeeBasicData(req, employeeID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEmployeeBasicData(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithBusinessUserCollection(req *http.Request, userID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("UserID eq '%s'", userID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithBusinessRoleAssignment(req *http.Request, businessRoleID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BusinessRoleID eq '%s'", businessRoleID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithEmployeeBasicData(req *http.Request, employeeID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("EmployeeID eq '%s'", employeeID))
	req.URL.RawQuery = params.Encode()
}
