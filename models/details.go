package models

import (

	u "twitter-hangouts/utils"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/ichtrojan/thoth"
	_"strings"
	"os"
	"fmt"
	"log"
	"errors"
	"time"
	"io/ioutil"
	"net/http"
	"bytes"
	"regexp"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/rs/xid"
	"encoding/json"
)
//Post Type details
type Details struct{
	gorm.Model
	Name   string `gorm:"type:varchar(100)" json:"name"`
	Email  string `gorm:"type:varchar(100)" json:"email"`
	Phone  string `gorm:"not null" json:"phone"`
	Reference  string `gorm:"not null"`
	Initiated  bool `gorm:"default:true"`
	Completed  bool `gorm:"default: false"`
	Verified  bool `gorm:"default: false"`
	Verified_by  string `gorm:"null;type:varchar(100)"`
}
type Payload struct {
	Txref  string `json:"txref"`
	SECKEY string `json:"SECKEY"`
}
type TxnVerificationResponse struct {
	Data    txnVerificationResponseData `json:"data"`
	Message string                      `json:"message"`
	Status  string                      `json:"status"`
}

type txnVerificationResponseData struct {
	Txid                            int         `json:"txid"`
	Txref                           string      `json:"txref"`
	Flwref                          string      `json:"flwref"`
	Devicefingerprint               string      `json:"devicefingerprint"`
	Cycle                           string      `json:"cycle"`
	Amount                          int         `json:"amount"`
	Currency                        string      `json:"currency"`
	Chargedamount                   int         `json:"chargedamount"`
	Appfee                          int         `json:"appfee"`
	Merchantfee                     int         `json:"merchantfee"`
	Merchantbearsfee                int         `json:"merchantbearsfee"`
	Chargecode                      string      `json:"chargecode"`
	Chargemessage                   string      `json:"chargemessage"`
	Authmodel                       string      `json:"authmodel"`
	IP                              string      `json:"ip"`
	Narration                       string      `json:"narration"`
	Status                          string      `json:"status"`
	Vbvcode                         string      `json:"vbvcode"`
	Vbvmessage                      string      `json:"vbvmessage"`
	Authurl                         string      `json:"authurl"`
	Acctcode                        interface{} `json:"acctcode"`
	Acctmessage                     interface{} `json:"acctmessage"`
	Paymenttype                     string      `json:"paymenttype"`
	Paymentid                       string      `json:"paymentid"`
	Fraudstatus                     string      `json:"fraudstatus"`
	Chargetype                      string      `json:"chargetype"`
	Createdday                      int         `json:"createdday"`
	Createddayname                  string      `json:"createddayname"`
	Createdweek                     int         `json:"createdweek"`
	Createdmonth                    int         `json:"createdmonth"`
	Createdmonthname                string      `json:"createdmonthname"`
	Createdquarter                  int         `json:"createdquarter"`
	Createdyear                     int         `json:"createdyear"`
	Createdyearisleap               bool        `json:"createdyearisleap"`
	Createddayispublicholiday       int         `json:"createddayispublicholiday"`
	Createdhour                     int         `json:"createdhour"`
	Createdminute                   int         `json:"createdminute"`
	Createdpmam                     string      `json:"createdpmam"`
	Created                         time.Time   `json:"created"`
	Customerid                      int         `json:"customerid"`
	Custphone                       interface{} `json:"custphone"`
	Custnetworkprovider             string      `json:"custnetworkprovider"`
	Custname                        string      `json:"custname"`
	Custemail                       string      `json:"custemail"`
	Custemailprovider               string      `json:"custemailprovider"`
	Custcreated                     time.Time   `json:"custcreated"`
	Accountid                       int         `json:"accountid"`
	Acctbusinessname                string      `json:"acctbusinessname"`
	Acctcontactperson               string      `json:"acctcontactperson"`
	Acctcountry                     string      `json:"acctcountry"`
	Acctbearsfeeattransactiontime   int         `json:"acctbearsfeeattransactiontime"`
	Acctparent                      int         `json:"acctparent"`
	Acctvpcmerchant                 string      `json:"acctvpcmerchant"`
	Acctalias                       interface{} `json:"acctalias"`
	Acctisliveapproved              int         `json:"acctisliveapproved"`
	Orderref                        string      `json:"orderref"`
	Paymentplan                     interface{} `json:"paymentplan"`
	Paymentpage                     interface{} `json:"paymentpage"`
	Raveref                         string      `json:"raveref"`
	Amountsettledforthistransaction int         `json:"amountsettledforthistransaction"`
	Card                            struct {
		Expirymonth    string `json:"expirymonth"`
		Expiryyear     string `json:"expiryyear"`
		CardBIN        string `json:"cardBIN"`
		Last4Digits    string `json:"last4digits"`
		Brand          string `json:"brand"`
		IssuingCountry string `json:"issuing_country"`
		CardTokens     []struct {
			Embedtoken string `json:"embedtoken"`
			Shortcode  string `json:"shortcode"`
			Expiry     string `json:"expiry"`
		} `json:"card_tokens"`
		Type          string `json:"type"`
		LifeTimeToken string `json:"life_time_token"`
	} `json:"card"`
	Meta []struct {
		ID                   int         `json:"id"`
		Metaname             string      `json:"metaname"`
		Metavalue            string      `json:"metavalue"`
		CreatedAt            time.Time   `json:"createdAt"`
		UpdatedAt            time.Time   `json:"updatedAt"`
		DeletedAt            interface{} `json:"deletedAt"`
		GetpaidTransactionID int         `json:"getpaidTransactionId"`
	} `json:"meta"`
} 

type Error struct {
    Mesage string `json:"message"`
}
var count int64

var reference string


func init(){
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}

// func (detail *Details) Validate() (map[string] interface{}, bool) {

// 	if !strings.Contains(detail.Email, "@") {
// 		return u.Message(false, "Email address is required"), false
// 	}

// 	if detail.Name == "" {
// 		return u.Message(false, "Name is required"), false
// 	}

// 	if len(detail.Name) < 3 {
// 		return u.Message(false, "Name is not valid"), false
// 	}

// 	if len(detail.Phone) < 11 || len(detail.Phone) > 11 {
// 		return u.Message(false, "Phone number is not valid"), false
// 	}

// 	//Email must be unique
// 	temp := &Details{}

// 	//check for errors and duplicate emails
// 	err := GetDB().Table("details").Where("email = ?", detail.Email).First(temp).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return u.Message(false, "Connection error. Please retry"), false
// 	}
// 	if temp.Email != "" {
// 		return u.Message(false, "Email address already in use by another user."), false
// 	}

// 	return u.Message(false, "Requirement passed"), true
// }
func (a Details) Validate() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Phone, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{11}$"))),
	)
}

func (detail *Details) Create() (map[string] interface{}) {

	err := detail.Validate()

	// if resp, ok := detail.Validate(); !ok {
	// 	return resp
	// }
	if err != nil {
		b, _ := json.Marshal(err)
		return u.Message(false, string(b))
	}

	//Email must be unique
	temp := &Details{}
	//check for errors and duplicate emails
	errors := GetDB().Table("details").Where("email = ?", detail.Email).First(temp).Error
	if errors != nil && errors != gorm.ErrRecordNotFound {
		return u.Errors(false, "Something went wrong. Please retry")
		
	}
	//For users that started the purchase process, but didnt complete and would love to come back to complete it.
	if temp.Email != "" {
		//This means someone has used this email to pay for stuffff, hehe.
		if temp.Completed == true {
			return u.Errors(false, "This email cannot be used due to some reasons or the other. Leemao")
		}else{
			//update reference and proceed to paying your debts mannnnnnn 
			detail.Reference = check_for_existing_ref()
			err := GetDB().Table("details").Where("email = ?", detail.Email).Update("reference", detail.Reference).Error

			if err != nil && err != gorm.ErrRecordNotFound {

				return u.Errors(false, string("Something went wrong. Please retry"))

			}
			mapD := map[string]string{"email": detail.Email, "reference": detail.Reference}
			response := u.Message(true, "Detail has been updated")
			
			response["details"] = mapD
			return response
		}
	}

	//For new users that wanna pay their debts.
	detail.Reference = check_for_existing_ref()

	GetDB().Create(detail)

	if detail.ID <= 0 {
		return u.Errors(false, string("Failed to create account, connection error"))
	}

	mapD := map[string]string{"email": detail.Email, "reference": detail.Reference}
    // mapB, _ := json.Marshal(mapD)
	response := u.Message(true, "Detail has been updated")
	
	response["details"] = mapD
	return response
}

func (detail *Payload) Confirm() (map[string] interface{}) {
	// temp := &Payload{}
	details := &Details{}
	logger, _ := thoth.Init("log")
	if detail.Txref == "" {
		return u.Message(false, "No reference code passed.")
	}else{

		err := GetDB().Table("details").Where("reference = ?", detail.Txref).First(details).Error

		if err != nil || err = gorm.ErrRecordNotFound {

			return u.Message(false, "Something went wrong with verifying reference code")

		}
		if details.Completed == true {

			return u.Message(false, "This payment been verified long ago.")
		}
		
		txref := detail.Txref

		results := confirm_reference_code(txref);

		data := TxnVerificationResponse{}

		_ = json.Unmarshal([]byte(results), &data)

		if data.Status == "success" && data.Data.Currency == "NGN" && data.Data.Chargedamount == 2000 {

			err := GetDB().Table("details").Where("reference = ? AND email = ?" , txref,data.Data.Custemail).Update("completed", true).Error

			if err != nil && err != gorm.ErrRecordNotFound {

				return u.Message(false, "Something went wrong with verifying reference code")
			}
			return u.Message(true, "This payment has been verified")
		}

		logger.Log(errors.New(results))
		return u.Message(false, "This payment could not be verified.")
	}

}

func confirm_reference_code(txref string) string {
	secKey, found := os.LookupEnv("RAVE_SECRET_KEY")
	if !found {
		log.Fatal("You need to set the \"RAVE_SECRET_KEY\" environment variable")
	}
	data := Payload{txref,secKey}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "Something went wrong here sha"
	}
	body := bytes.NewBuffer(payloadBytes)

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}	
	req, err := http.NewRequest("POST", return_api(), body)
	if err != nil {
		return "Something went wrong here sha 1"
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "Something went wrong here sha 3"
	}
	
	defer resp.Body.Close()

	result,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return "Something went wrong here sha 4"
		// log.Fatalln(err)
	}
	return string(result)


}
//Check for existing reference code 
func check_for_existing_ref() string{
	reference := generate_reference()
	GetDB().Table("details").Where("reference = ?", reference).Count(&count)
	// db.Model(&Details{}).Where("reference = ?", reference).Count(&count)
	if count >= 1 {
		return generate_reference()
	} else{
		return reference
	}
}

//Generate new reference code for payment bla bla
func generate_reference() string {
	guid := xid.New()
	reference := guid.String()
	return reference
}
//Return URL for live or test mode
func return_api() string {
	if os.Getenv("RAVE_MODE") != "live" {
		return os.Getenv("RAVE_API_TEST")
	}
	return os.Getenv("RAVE_API_LIVE")
}