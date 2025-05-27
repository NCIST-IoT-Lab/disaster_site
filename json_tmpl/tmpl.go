package tmpl

import (
	"os"
	"path/filepath"
	"runtime"
)

const (
    eventTemplate = "event.jtgo"
)

// GetEventTemplate 获取事件模板，实时读取文件内容
func GetEventTemplate() string {
    // 获取当前文件所在目录
    _, filename, _, _ := runtime.Caller(0)
    dir := filepath.Dir(filename)
    
    // 从文件系统读取文件
    templatePath := filepath.Join(dir, eventTemplate)
    template, err := os.ReadFile(templatePath)
    if err != nil {
        return ""
    }
    return string(template)
}