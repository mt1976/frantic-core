package templateStore

// Data Access Object template
// Version: 0.3.0
// Updated on: 2025-12-31

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/goforj/godump"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/mathHelpers"
)

//TODO: RENAME "template" TO THE NAME OF THE DOMAIN ENTITY
//TODO: Implement the validate function to process the domain entity
//TODO: Implement the calculate function to process the domain entity
//TODO: Implement the isDuplicateOf function to process the domain entity
//TODO: Implement the postGetProcessing function to process the domain entity

// upgradeProcessing performs any necessary upgrade processing for the TemplateStore record.
// This processing is triggered directly after the record has been retrieved from the database.
func (record *TemplateStore) upgradeProcessing() error {
	//TODO: Add any upgrade processing here
	//This processing is triggered directly after the record has been retrieved from the database
	return nil
}

// defaultProcessing performs default calculations for the TemplateStore record.
// This processing is triggered directly before the record is saved to the database.
func (record *TemplateStore) defaultProcessing() error {
	//TODO: Add any default calculations here
	//This processing is triggered directly before the record is saved to the database
	return nil
}

// validationProcessing performs validation checks for the TemplateStore record.
// This processing is triggered directly before the record is saved to the database and after the default calculations.
func (record *TemplateStore) validationProcessing() error {
	//TODO: Add any record validation here
	//This processing is triggered directly before the record is saved to the database and after the default calculations
	return nil
}

// postGetProcessing performs any necessary processing for the TemplateStore record after it has been retrieved from the database.
// This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing.
func (h *TemplateStore) postGetProcessing() error {
	//TODO: Add any post get processing here
	//This processing is triggered directly after the record has been retrieved from the database and after the upgrade processing
	return nil
}

// preDeleteProcessing performs any necessary processing for the TemplateStore record before it is deleted from the database.
// This processing is triggered directly before the record is deleted from the database.
func (record *TemplateStore) preDeleteProcessing() error {
	//TODO: Add any pre delete processing here
	//This processing is triggered directly before the record is deleted from the database
	return nil
}

// templateClone creates a clone of the given TemplateStore record.
// This function is used to duplicate a TemplateStore record.
func templateClone(ctx context.Context, source TemplateStore) (TemplateStore, error) {
	//TODO: Add any clone processing here
	panic("Not Implemented")
}

// // prepare performs preparation steps for the TemplateStore record before it is saved to the database.
// func (u *TemplateStore) prepare() (TemplateStore, error) {
// 	//os.Exit(0)
// 	// logHandler.ErrorLogger.Printf("ACT: VAL Validate")
// 	user, err := u.dup()
// 	if err == commonErrors.ErrorDuplicate {
// 		return *u, nil
// 	}
// 	if err != nil {
// 		return user, err
// 	}
// 	return *u, nil
// }

// // calculate performs calculations for the TemplateStore record before it is saved to the database.
// func (u *TemplateStore) calculate() error {
// 	// Calculate the duration in days between the start and end dates
// 	return nil
// }

// // dup checks for duplicate TemplateStore records based on UserCode and UserName.
// func (u *TemplateStore) dup() (TemplateStore, error) {

// 	// logHandler.InfoLogger.Printf("CHK: CheckUniqueCode %v", name)

// 	// Get all status
// 	userList, err := GetAll()
// 	if err != nil {
// 		logHandler.ErrorLogger.Printf("Error Getting all Users: %v", err.Error())
// 		return TemplateStore{}, err
// 	}

// 	// range through status list, if status code is found and deletedby is empty then return error
// 	for _, s := range userList {
// 		//s.Dump(strings.ToUpper(code) + "-uchk-" + s.Code)
// 		testValue := strings.ToUpper(u.UserCode)
// 		checkValue := strings.ToUpper(s.UserCode)
// 		// logHandler.InfoLogger.Printf("CHK: TestValue:[%v] CheckValue:[%v]", testValue, checkValue)
// 		// logHandler.InfoLogger.Printf("CHK: Code:[%v] s.Code:[%v] s.Audit.DeletedBy:[%v]", testCode, s.Code, s.Audit.DeletedBy)
// 		if checkValue == testValue && s.Audit.DeletedBy == "" {
// 			logHandler.WarningLogger.Printf("[%v] DUPLICATE UID [%v] already in use for [%v]", TableName, testValue, s.UserName)
// 			return s, commonErrors.ErrorDuplicate
// 		}
// 		testValue = strings.ToUpper(u.UserName)
// 		checkValue = strings.ToUpper(s.UserName)
// 		// logHandler.InfoLogger.Printf("CHK: TestValue:[%v] CheckValue:[%v]", testValue, checkValue)
// 		// logHandler.InfoLogger.Printf("CHK: Code:[%v] s.Code:[%v] s.Audit.DeletedBy:[%v]", testCode, s.Code, s.Audit.DeletedBy)
// 		if checkValue == testValue && s.Audit.DeletedBy == "" {
// 			logHandler.WarningLogger.Printf("[%v] DUPLICATE User Name [%v] already in use for [%v]", TableName, testValue, s.UserName)
// 			return s, commonErrors.ErrorDuplicate
// 		}
// 	}

// 	// Return nil if the code is unique

// 	return *u, nil
// }

// End of Model Helpers

// Insert additional functions below this line

func Login(ctx context.Context) {

	temp := buildUserStub()
	var usr TemplateStore

	//godump.DumpJSON(temp)
	logHandler.TraceLogger.Printf("%v", godump.DumpStr(temp))

	//os.Exit(0)

	usrList, err := GetAllWhere(Fields.UserCode, temp.UserCode)
	logHandler.InfoLogger.Printf("Login UserCode=[%v] Count=[%v]", temp.UserCode, len(usrList))
	if err != nil || len(usrList) == 0 {
		if err != nil {
			logHandler.WarningLogger.Printf("Warning=[%v] User=[%v]", err.Error(), temp.UserName)
		}
		logHandler.InfoLogger.Printf("User=[%v] does not exist, creating", temp.UserName)
		usr, err = Add(ctx)
		if err != nil {
			logHandler.ErrorLogger.Printf("Warning=[%v] User=[%v]", err.Error(), temp.UserName)
			panic(err)
		}
	}
	logHandler.InfoLogger.Printf("Login UserCode=[%v] Count=[%v]", temp.UserCode, len(usrList))
	if len(usrList) >= 1 {
		usr = usrList[0]
	}

	usr.LastLogin = time.Now()
	usr.LastHost, _ = os.Hostname()
	//u.Dump("login")
	//xx := audit.LOGIN.WithMessage(fmt.Sprintf("User %v logged in", usr.UserName))
	err = usr.UpdateWithAction(ctx, audit.LOGIN, fmt.Sprintf("User %v logged in", usr.UserName))
	if err != nil {
		logHandler.WarningLogger.Printf("Warning=[%v] User=[%v]", err.Error(), usr.UserName)
	}
}

func (u *TemplateStore) SetName(name string) error {
	if name == "" {
		return commonErrors.ErrorEmptyName
	}
	if len(name) > 50 {
		return commonErrors.ErrorNameTooLong
	}
	u.RealName = name
	return nil
}

func (u *TemplateStore) buildUserCode() string {
	return fmt.Sprintf("%v%v%v", u.UID, cfg.Display.Delimiter, u.UserName)
}

func Add(ctx context.Context) (TemplateStore, error) {
	// Create a new User object

	//godump.Dump(currentUser)
	//username := fmt.Sprintf("%v_%s", currentUser.Uid, currentUser.Username)
	// Check if the user already exists, if not create
	testu := buildUserStub()

	// oldu, dupErr := testu.dup()
	// if dupErr == commonErrors.ErrorDuplicate {
	// 	logHandler.WarningLogger.Printf("A user already exists for [%v]", testu.UserName)
	// 	return oldu, nil
	// }

	u, err := Create(ctx, testu.UserName, testu.UID, testu.RealName, testu.Email, testu.GID)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error: '%v'", err.Error())
		return TemplateStore{}, err
	}

	//u.Dump("ADD")

	// Return the new User object
	return u, nil
}

func buildUserStub() TemplateStore {
	currentUser, _ := user.Current()
	hostname, _ := os.Hostname()

	stub := TemplateStore{}
	stub.ID = 0
	stub.UID = fmt.Sprintf("%v%v", currentUser.Uid, mathHelpers.RandomInt(10))
	stub.UserName = currentUser.Username
	stub.RealName = currentUser.Name
	stub.GID = currentUser.Gid
	stub.Email = strings.ToLower(fmt.Sprintf("%v@%v.com", currentUser.Username, hostname))
	stub.UserCode = strings.ToLower(fmt.Sprintf("%v%v%s", stub.UID, cfg.Display.Delimiter, currentUser.Username))
	//godump.DumpJSON(stub)
	return stub
}
