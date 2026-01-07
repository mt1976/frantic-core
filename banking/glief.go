package banking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
	validations "github.com/robfordww/finident"
)

type GLIEF struct {
	Meta struct {
		GoldenCopy struct {
			PublishDate time.Time `json:"publishDate"`
		} `json:"goldenCopy"`
		Pagination struct {
			CurrentPage int `json:"currentPage"`
			PerPage     int `json:"perPage"`
			From        int `json:"from"`
			To          int `json:"to"`
			Total       int `json:"total"`
			LastPage    int `json:"lastPage"`
		} `json:"pagination"`
	} `json:"meta"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Lei    string `json:"lei"`
			Entity struct {
				LegalName struct {
					Name     string `json:"name"`
					Language string `json:"language"`
				} `json:"legalName"`
				OtherNames               []interface{} `json:"otherNames"`
				TransliteratedOtherNames []interface{} `json:"transliteratedOtherNames"`
				LegalAddress             struct {
					Language                    string      `json:"language"`
					AddressLines                []string    `json:"addressLines"`
					AddressNumber               interface{} `json:"addressNumber"`
					AddressNumberWithinBuilding interface{} `json:"addressNumberWithinBuilding"`
					MailRouting                 interface{} `json:"mailRouting"`
					City                        string      `json:"city"`
					Region                      string      `json:"region"`
					Country                     string      `json:"country"`
					PostalCode                  string      `json:"postalCode"`
				} `json:"legalAddress"`
				HeadquartersAddress struct {
					Language                    string      `json:"language"`
					AddressLines                []string    `json:"addressLines"`
					AddressNumber               interface{} `json:"addressNumber"`
					AddressNumberWithinBuilding interface{} `json:"addressNumberWithinBuilding"`
					MailRouting                 interface{} `json:"mailRouting"`
					City                        string      `json:"city"`
					Region                      string      `json:"region"`
					Country                     string      `json:"country"`
					PostalCode                  string      `json:"postalCode"`
				} `json:"headquartersAddress"`
				RegisteredAt struct {
					ID    string      `json:"id"`
					Other interface{} `json:"other"`
				} `json:"registeredAt"`
				RegisteredAs interface{} `json:"registeredAs"`
				Jurisdiction string      `json:"jurisdiction"`
				Category     interface{} `json:"category"`
				LegalForm    struct {
					ID    string `json:"id"`
					Other string `json:"other"`
				} `json:"legalForm"`
				AssociatedEntity struct {
					Lei  interface{} `json:"lei"`
					Name interface{} `json:"name"`
				} `json:"associatedEntity"`
				Status     string `json:"status"`
				Expiration struct {
					Date   interface{} `json:"date"`
					Reason interface{} `json:"reason"`
				} `json:"expiration"`
				SuccessorEntity struct {
					Lei  interface{} `json:"lei"`
					Name interface{} `json:"name"`
				} `json:"successorEntity"`
				OtherAddresses []interface{} `json:"otherAddresses"`
			} `json:"entity"`
			Registration struct {
				InitialRegistrationDate time.Time `json:"initialRegistrationDate"`
				LastUpdateDate          time.Time `json:"lastUpdateDate"`
				Status                  string    `json:"status"`
				NextRenewalDate         time.Time `json:"nextRenewalDate"`
				ManagingLou             string    `json:"managingLou"`
				CorroborationLevel      string    `json:"corroborationLevel"`
				ValidatedAt             struct {
					ID    string      `json:"id"`
					Other interface{} `json:"other"`
				} `json:"validatedAt"`
				ValidatedAs                interface{}   `json:"validatedAs"`
				OtherValidationAuthorities []interface{} `json:"otherValidationAuthorities"`
			} `json:"registration"`
			Bic interface{} `json:"bic"`
		} `json:"attributes"`
		Relationships struct {
			ManagingLou struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"managing-lou"`
			LeiIssuer struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"lei-issuer"`
			FieldModifications struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"field-modifications"`
			DirectParent struct {
				Links struct {
					ReportingException string `json:"reporting-exception"`
				} `json:"links"`
			} `json:"direct-parent"`
			UltimateParent struct {
				Links struct {
					ReportingException string `json:"reporting-exception"`
				} `json:"links"`
			} `json:"ultimate-parent"`
			DirectChildren struct {
				Links struct {
					RelationshipRecords string `json:"relationship-records"`
					Related             string `json:"related"`
				} `json:"links"`
			} `json:"direct-children"`
			UltimateChildren struct {
				Links struct {
					RelationshipRecords string `json:"relationship-records"`
					Related             string `json:"related"`
				} `json:"links"`
			} `json:"ultimate-children"`
			Isins struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"isins"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}

var gliefAPIURI = "https://api.gleif.org/api/v1/lei-records?filter[isin]=%v"

func Lookup_LEI(inISIN string) (string, error) {

	isinOK, err := validations.ValidateISIN(inISIN)
	if err == nil {
		//log.Printf("%q is Valid %v", inISIN, isinOK)
		if isinOK {
			uri := fmt.Sprintf(gliefAPIURI, inISIN)
			//log.Printf("URI=[%q]", uri)

			resp, err := http.Get(uri)
			if err != nil {
				//fmt.Println("No response from request")
				logHandler.ErrorLogger.Printf("No response from request: %v", err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body) // response body is []byte
			//fmt.Println(string(body))
			var result GLIEF
			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
				//				fmt.Println("Can not unmarshal JSON")
				logHandler.ErrorLogger.Printf("Can not unmarshal JSON: %v", err)
			}
			//fmt.Printf("result: %v\n", result)
			//spew.Dump(result)
			return result.Data[0].Attributes.Lei, nil
		}
	} else {
		return "", commonErrors.ErrValidationWrapper(err)
	}

	return "", nil
}
