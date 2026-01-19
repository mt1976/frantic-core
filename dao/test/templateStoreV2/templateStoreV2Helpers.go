package templateStoreV2

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/goforj/godump"
	ce "github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/dao/audit"
	"github.com/mt1976/frantic-core/dao/entities"
	"github.com/mt1976/frantic-core/logHandler"
)

func (record *TemplateStore) upgradeProcessing() error {
	return nil
}

func (record *TemplateStore) defaultProcessing() error {
	return nil
}

func (record *TemplateStore) validationProcessing() error {
	return nil
}

func (h *TemplateStore) postGetProcessing() error {
	return nil
}

func (record *TemplateStore) preDeleteProcessing() error {
	return nil
}

func templateClone(ctx context.Context, source TemplateStore) (TemplateStore, error) {
	_ = ctx
	_ = source
	panic("Not Implemented")
}

func assertTemplateStore(result any, field entities.Field, value any) (*TemplateStore, error) {
	x, ok := result.(*TemplateStore)
	if !ok {
		return nil, ce.ErrDAOAssertWrapper(tableName, field.String(), value,
			ce.ErrInvalidTypeWrapper(field.String(), fmt.Sprintf("%T", result), "*TemplateStore"))
	}
	return x, nil
}

func Login(ctx context.Context, sq string) (TemplateStore, error) {
	temp := buildUserStub(sq)
	var usr TemplateStore

	logHandler.TraceLogger.Printf("%v", godump.DumpStr(temp))

	usrList, err := GetAllWhere(Fields.UserCode, temp.UserCode)
	if err != nil || len(usrList) == 0 {
		if err != nil {
			logHandler.WarningLogger.Printf("Warning=[%v] User=[%v]", err.Error(), temp.UserName)
			return TemplateStore{}, err
		}
		usr, err = Add(ctx, sq)
		if err != nil {
			logHandler.ErrorLogger.Printf("Warning=[%v] User=[%v]", err.Error(), temp.UserName)
			return TemplateStore{}, err
		}
		return usr, nil
	}

	// Existing user(s): update login details, return first updated.
	u := usrList[0]
	u.LastHost, _ = os.Hostname()
	u.LastLogin = time.Now()
	if err = u.UpdateWithAction(ctx, audit.LOGIN, fmt.Sprintf("User %v logged in", u.UserName)); err != nil {
		logHandler.WarningLogger.Printf("Warning=[%v] User=[%v]", err.Error(), u.UserName)
		return TemplateStore{}, err
	}
	usr = u
	return usr, nil
}

func (u *TemplateStore) SetName(name string) error {
	if name == "" {
		return ce.ErrEmptyName
	}
	if len(name) > 50 {
		return ce.ErrNameTooLong
	}
	u.RealName = name
	return nil
}

func (u *TemplateStore) buildUserCode() string {
	return fmt.Sprintf("%v%v%v", u.UID, cfg.Display.Delimiter, u.UserName)
}

func Add(ctx context.Context, sq string) (TemplateStore, error) {
	testu := buildUserStub(sq)

	u, err := Create(ctx, testu.UserName, testu.UID, testu.RealName, testu.Email, testu.GID)
	if err != nil {
		logHandler.ErrorLogger.Printf("Error: '%v'", err.Error())
		return TemplateStore{}, err
	}

	return u, nil
}

func buildUserStub(sq string) TemplateStore {
	currentUser, _ := user.Current()
	hostname, _ := os.Hostname()

	stub := TemplateStore{}
	stub.ID = 0
	stub.UID = fmt.Sprintf("%v%04v", currentUser.Uid, sq)
	stub.UserName = currentUser.Username
	stub.RealName = currentUser.Name
	stub.GID = currentUser.Gid
	stub.Email = strings.ToLower(fmt.Sprintf("%v@%v.com", currentUser.Username, hostname))
	stub.UserCode = strings.ToLower(fmt.Sprintf("%v%v%s", stub.UID, cfg.Display.Delimiter, currentUser.Username))
	return stub
}
