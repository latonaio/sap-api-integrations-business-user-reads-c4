# sap-api-integrations-business-user-reads-c4
sap-api-integrations-business-user-reads-c4 は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で ビジネスユーザ データを取得するマイクロサービスです。    
sap-api-integrations-business-user-reads-c4 には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-business-user-reads-c4 は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/businessuser/overview   

## 動作環境  
sap-api-integrations-business-user-reads-c4 は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-business-user-reads-c4 は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-business-user-reads-c4 が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/businessuser/overview    
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-business-user-reads-c4 には、次の API をコールするためのリソースが含まれています。  

* BusinessUserCollection（ビジネスユーザ - ビジネスユーザ）※ビジネスユーザの詳細データを取得するために、ToBusinessRoleAssignment、ToEmployeeBasicData、と合わせて取得されます。
* ToBusinessRoleAssignment（ビジネスユーザ - ビジネスロール割当）
* ToEmployeeBasicData（ビジネスユーザ - 従業員基本情報）
* BusinessUserBusinessRoleAssignmentCollection（ビジネスユーザ - ビジネスロール割当）
* EmployeeBasicDataCollection（ビジネスユーザ - 従業員基本情報）


## API への 値入力条件 の 初期値
sap-api-integrations-business-user-reads-c4 において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.BusinessUserCollection.UserID（ユーザID）
* inoutSDC.BusinessUserCollection.BusinessUserBusinessRoleAssignment.BusinessRoleID（ビジネスロールID）
* inoutSDC.BusinessUserCollection.BusinessUserBusinessRoleAssignment.EmployeeBasicData.EmployeeID（従業員ID）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Supplier" が指定されています。    
  
```
	"api_schema": "BusinessUserCollection",
	"accepter": ["BusinessUserCollection"],
	"business_user_code": "VINCENT",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "BusinessUserCollection",
	"accepter": ["All"],
	"business_user_code": "VINCENT",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP ビジネスユーザ  の ビジネスユーザデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "EmployeeBasicData" は、/SAP_API_Output_Formatter/type.go 内 の Type BusinessUserCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-business-user-reads-c4/SAP_API_Caller/caller.go#L63",
	"function": "sap-api-integrations-business-user-reads-c4/SAP_API_Caller.(*SAPAPICaller).BusinessUserCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E038C701ED298B0A06CCCA7F81B",
			"ETag": "2022-05-01T10:35:36+09:00",
			"EmployeeID": "7000002",
			"EmployeeUUID": "00163E03-8C70-1ED2-98B0-A06CCCA7F81B",
			"UserID": "VINCENT",
			"TechnicalUserID": "JYU71I0BE9C",
			"IdentityUUID": "00163E05-AE66-1ED3-8ED2-B36895238383",
			"BusinessPartnerID": "8000000211",
			"BusinessPartnerFormattedName": "Vincent Gregory",
			"DepartmentName": "",
			"CompanyName": "",
			"ManagerName": "",
			"EmailURI": "Vincent.Gregory@ondemand.com",
			"DecimalFormatCode": "X",
			"DecimalFormatCodeText": "1,234,567.89",
			"DateFormatCode": "2",
			"DateFormatCodeText": "MM/DD/YYYY",
			"TimeFormatCode": "1",
			"TimeFormatCodeText": "12-Hour Time",
			"TimeZoneCode": "PST",
			"TimeZoneCodeText": "(UTC-08:00) Pacific Time (Los Angeles, Tijuana, Yukon)",
			"LogonLanguageCode": "EN",
			"LogonLanguageCodeText": "English",
			"UserValidityStartDate": "2013-10-21T09:00:00+09:00",
			"UserValidityEndDate": "9999-12-31T09:00:00+09:00",
			"UserLockedIndicator": true,
			"UserCountedIndicator": false,
			"PasswordPolicyCode": "C_ALMIKA_BUSINESS_USER",
			"PasswordPolicyCodeText": "",
			"PasswordInactiveIndicator": false,
			"PasswordLockedIndicator": false,
			"UserAccountTypeCode": "A",
			"UserAccountTypeCodeText": "Dialog",
			"CreatedOn": "2013-10-21T09:00:00+09:00",
			"UserCreatedBy": "Michael Johnes",
			"EntityLastChangedOn": "2016-07-13T20:41:32+09:00",
			"UserChangedBy": "",
			"UserChangedOn": "",
			"BusinessUserBusinessRoleAssignment": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/BusinessUserCollection('00163E038C701ED298B0A06CCCA7F81B')/BusinessUserBusinessRoleAssignment",
			"BusinessUserSubscriptionAssignment": "",
			"EmployeeBasicData": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/BusinessUserCollection('00163E038C701ED298B0A06CCCA7F81B')/EmployeeBasicData"
		}
	],
	"time": "2022-05-01T21:19:29+09:00"
}
```