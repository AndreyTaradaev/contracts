package tools

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// имя файла конфига.
const (
	filename string = "config.json"
)

// Вспомогательная структура парсинг командной строки.
type Arg struct {
	Fileconfig string
	Help       bool
	Debug      bool
}

// helper for interfase stringer.
func (a Arg) String() string {
	var s string = ""
	if len(a.Fileconfig) != 0 {
		s += "-config " + a.Fileconfig
	}
	if a.Help {
		s += " -help"
	}
	if a.Debug {
		s += " -debug"
	}
	return s
}

// назавание файла конфигурации поумолчанию
func FileConfig() string {
	return filename
}

// =================работа с командной строкой=======================================
// парсим командную строку возвращаем путь до конфиг файла, debug, help.
func ParseFile() (Arg, error) {
	var a Arg
	exe, err := os.Executable()
	if err != nil {
		return a, err
	}
	// получаем каталог и добавляем имя конф-файла по умолчанию
	s := filepath.Dir(exe) + "/" + filename

	flag.StringVar(&a.Fileconfig, "config", s, "Path  config file")
	flag.BoolVar(&a.Help, "help", false, "man for app")
	flag.BoolVar(&a.Help, "h", false, "man for app")
	flag.BoolVar(&a.Debug, "debug", false, "enable debug")
	flag.BoolVar(&a.Debug, "d", false, "enable debug")
	flag.Parse()
	return a, nil
}

// выводим  последний элемент пути.
func GetBase(path string) string {
	return filepath.Base(path)
}

//
func GetAppExe() string{
	exe, err := os.Executable()
	if err!=nil {
		return ""
	}
	return exe
}
// Имя файла без расширения
func FileNamewithoutExt(path string) string {
	f := filepath.Base(path)
	n := strings.LastIndex(f, ".")
	if n == -1 {
		return f
	}
	b := []byte(f)
	f = string(b[0:n])
	return f
}

// проверка наличия файла
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// конвертирует строку в число если ошибка возвращает значение по умолчанию.
func GetIntDef(value string, def int64) int64 {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	return v
}
