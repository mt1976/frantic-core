package ioHelpers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/mt1976/frantic-core/commonConfig"
	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/idHelpers"
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-core/paths"
	"github.com/mt1976/frantic-core/timing"
)

var name = "IO"

func GetDBFileName(name string) string {
	if name == "" {
		panic(fmt.Errorf("db name is required"))
	}

	cfg := commonConfig.Get()
	sep := "-"

	name = cfg.GetApplication_Name() + sep + name

	name = strings.ToLower(name)

	path := "%s" + paths.Seperator() + paths.Database().String() + paths.Seperator() + "%s.db"

	xx := fmt.Sprintf(path, paths.Application().String(), name)
	//logger.InfoLogger.Println("DBN: DB File Name: ", xx)

	return clean(xx)
}

func Dump(tableName string, where paths.FileSystemPath, action string, recordID string, yy any) {
	cfg := commonConfig.Get()
	sep := "-"

	logHandler.DatabaseLogger.Printf("[SUPPORT] [%v] Dump to '%v'", strings.ToUpper(tableName), where.String())
	id := idHelpers.GetUUID()
	if action != "" {
		id = id + sep + cfg.GetApplication_Name() + sep + tableName + sep + strings.ToTitle(action) + sep + recordID
	}
	id = id + ".json"

	if where.Is(paths.Backups()) {
		id = tableName + sep + action + ".bk"
	}

	path := where.String()

	b, err := json.Marshal(yy)
	if err != nil {
		logHandler.WarningLogger.Printf("[SUPPORT] [%v] [Marshalling] Error=[%v]", strings.ToUpper(action), err.Error())
		return
	}
	output := string(b)

	_, err = Write(id, path, output)
	if err != nil {
		logHandler.ErrorLogger.Printf("[SUPPORT] [%v] [Write] Error=[%v]", strings.ToUpper(action), err.Error())
	}
}

func Backup(table, location string) {
	//path := BACKUPS.path
	sep := "-"
	cfg := commonConfig.Get()
	table = strings.ToLower(cfg.GetApplication_Name() + sep + table)
	logHandler.EventLogger.Printf("Backup=[%v] [%v.db] to [%v]", strings.ToLower(table), table, location)

	// sleep for 1 second
	time.Sleep(1 * time.Second)
	timing := timing.Start(table, "Backup", "")
	//dateTime := time.Now().Format("20060102150405")
	toPath := clean(location) // location has all the path info
	toFile := clean(toPath + paths.Seperator() + table + ".db")

	fromPath := clean(paths.Database().String())
	fromFile := clean(paths.Application().String() + paths.Database().String() + paths.Seperator() + table + ".db")
	logHandler.EventLogger.Printf("Backup=[%v] Path=[%v]", strings.ToLower(table), paths.Application().String())
	logHandler.EventLogger.Printf("Backup=[%v] Database=[%v.db]", strings.ToLower(table), table)
	logHandler.EventLogger.Printf("Backup=[%v] From=[%v] %v", strings.ToLower(table), fromPath, toFile)
	logHandler.EventLogger.Printf("Backup=[%v] To=[%v]", strings.ToLower(table), toPath)

	// remove last char from path
	//toPath = toPath[:len(toPath)-1]
	//fromPath = fromPath[:len(fromPath)-1]

	logHandler.InfoLogger.Printf("BackupPaths=[%v] [%v] to [%v]", strings.ToLower(table), fromPath, toPath)
	logHandler.InfoLogger.Printf("BackupFiles=[%v] [%v] to [%v]", strings.ToLower(table), fromFile, toFile)

	err := CopyFile(fromFile, toFile)
	if err != nil {
		logHandler.ErrorLogger.Printf("Backup=[%v] [%v] to [%v] Error=[%v]", strings.ToLower(name), fromPath, toPath, err.Error())
		panic(err)
	}
	timing.Stop(1)
	logHandler.EventLogger.Printf("Backup=[%v] COMPLETE", strings.ToLower(table))
}

// File copies a single file from src to dst
func CopyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	src = clean(src)
	dst = clean(dst)

	if srcfd, err = os.Open(src); err != nil {
		return commonErrors.WrapOSError(err)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return commonErrors.WrapOSError(err)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return commonErrors.WrapOSError(err)
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return commonErrors.WrapOSError(err)
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func MkDir(path string) error {
	// Create the folder
	path = clean(path)
	logHandler.InfoLogger.Printf("[%v] Creating folder Path=[%v]", strings.ToUpper(name), path)
	return os.MkdirAll(path, os.ModeSticky|os.ModePerm)
}

func Dir(path string) ([]string, error) {
	path = clean(path)
	// Get all folders in the backup directory
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, commonErrors.WrapOSError(err)
	}
	var folders []string
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders, nil
}

func DeleteFolder(path string) error {
	path = clean(path)
	// Delete the folder
	logHandler.InfoLogger.Printf("[DELETE] [%v] Deleting folder Path=[%v]", strings.ToUpper(name), path)
	return os.RemoveAll(path)
	//return nil
}

func clean(path string) string {
	sep := string(os.PathSeparator) + string(os.PathSeparator)
	rtn := strings.ReplaceAll(path, sep, string(os.PathSeparator))
	return rtn
}
