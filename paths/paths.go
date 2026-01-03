package paths

import (
	"fmt"
	"os"
	"strings"
)

var name = "Paths"

type FileSystemPath struct {
	path string
}

func (f FileSystemPath) String() string {
	return clean(f.path)
}

func HTML() FileSystemPath {
	return FileSystemPath{clean(Res().String() + "/html/templates/")}
}

func HTMLTemplates() FileSystemPath {
	return FileSystemPath{clean(Res().String() + "/html/")}
}

func HTMLPage(in string) string {
	return clean(HTML().String() + in + ".html")
}

func HTMLTemplate() string {
	return clean(HTMLTemplates().String() + "templates.html")
}

func Images() FileSystemPath {
	return FileSystemPath{clean(Res().String() + "/img")}
}

func Backups() FileSystemPath {
	return FileSystemPath{clean(Data().String() + "/backups")}
}

func Dumps() FileSystemPath {
	return FileSystemPath{clean(Data().String() + "/dumps")}
}

func Database() FileSystemPath {
	return FileSystemPath{clean(Data().String() + "/database")}
}

func Config() FileSystemPath {
	return FileSystemPath{clean(Data().String() + "/config")}
}

func Defaults() FileSystemPath {
	return FileSystemPath{clean(Data().String() + "/defaults")}
}

func Logs() FileSystemPath {
	return FileSystemPath{clean(Data().String() + "/logs")}
}

func Data() FileSystemPath {
	return FileSystemPath{clean("/data")}
}

func Res() FileSystemPath {
	return FileSystemPath{clean("./res")}
}

func Application() FileSystemPath {
	return FileSystemPath{clean(fullPath())}
}

func Seperator() string {
	return string(os.PathSeparator)
}

func (F *FileSystemPath) Is(in FileSystemPath) bool {
	return clean(F.path) == clean(in.path)
}

func fullPath() string {
	// Get the full path of the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("[%v] Error getting current directory [%v]", name, err.Error())
		panic(err)
	}
	return clean(dir)
}

func clean(path string) string {
	sep := string(os.PathSeparator) + string(os.PathSeparator)
	rtn := strings.ReplaceAll(path, sep, string(os.PathSeparator))
	return rtn
}
