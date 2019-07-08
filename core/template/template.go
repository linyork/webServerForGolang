package template

import (
    "fmt"
    "github.com/gin-contrib/multitemplate"
    "github.com/gin-gonic/gin"
    "gin/core/config"
    "os"
    "path/filepath"
    "runtime"
    "strings"
)

func changeSlash(path string) string {
    if runtime.GOOS == "windows" {
        return strings.Replace(path, "/", "\\", -1)
    }
    return path
}

func LoadTemplate(templatesDir string) multitemplate.Renderer {
    // 常數
    templatesDir = changeSlash(templatesDir)
    layoutDir := changeSlash("/pc/layouts/")
    commonDir := changeSlash("/pc/blocks/commons/")
    blockDir := changeSlash("/pc/blocks/")
    slash := changeSlash("/")
    fileLastName := ".html"

    render := multitemplate.NewRenderer()

    allBlocks, allLayouts, allCommons := make([]string, 0), make([]string, 0), make([]string, 0)

    err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
            return err
        }

        if !info.IsDir() && strings.HasSuffix(path, fileLastName) {
            if strings.HasPrefix(path, templatesDir + layoutDir) {
                allLayouts = append(allLayouts, path)
            } else if strings.HasPrefix(path, templatesDir + commonDir) {
                allCommons = append(allCommons, path)
            } else if strings.HasPrefix(path, templatesDir + blockDir) {
                allBlocks = append(allBlocks, path)
            }
        }

        return nil
    })
    if err != nil {
        panic(err.Error())
    }

    for _, block := range allBlocks {
        viewCopy := make([]string, 0)
        
        for _, layout := range allLayouts {
            viewCopy = append(viewCopy, layout)
        }

        viewCopy = append(viewCopy, block)

        for _, common := range allCommons {
            viewCopy = append(viewCopy, common)
        }

        // 比對路徑
        blockPath := block[0:strings.LastIndex(block, slash)]
        for _, blockCheck := range allBlocks {
            blockCheckPath := blockCheck[0:strings.LastIndex(blockCheck, slash)]
            if strings.HasPrefix(blockCheckPath, blockPath) && blockPath != blockCheckPath {

                viewCopy = append(viewCopy, blockCheck)
            }
        }

        identify := strings.Replace(block, templatesDir + blockDir, "", -1)
        identify = strings.Replace(identify, fileLastName, "", -1)
        identify = strings.Replace(identify, "\\", "/", -1)

        render.AddFromFiles(identify, viewCopy...)
    }
    return render
}

func SetHTMLRender(ginEngine *gin.Engine) {
    render := multitemplate.NewRenderer()
    templatesDir := config.Template.Path
    allBlocks, allLayouts := make([]string, 0), make([]string, 0)

    err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
            return err
        }

        if !info.IsDir() && strings.HasSuffix(path, ".html") {
            if strings.HasPrefix(path, templatesDir+"/layouts/") {
                allLayouts = append(allLayouts, path)
            } else if strings.HasPrefix(path, templatesDir+"/blocks/") {
                allBlocks = append(allBlocks, path)
            }
        }

        return nil
    })

    if err != nil {
        panic(err.Error())
    }

    for _, block := range allBlocks {
        viewCopy := make([]string, 0)
        viewCopy = append(viewCopy, block)

        for _, layout := range allLayouts {
            viewCopy = append(viewCopy, layout)
        }

        identify := strings.Replace(block, config.Template.Path+"/blocks/", "", -1)
        identify = strings.Replace(identify, ".html", "", -1)
        identify = strings.Replace(identify, "/", ".", -1)

        render.AddFromFiles(identify, viewCopy...)
    }

    ginEngine.HTMLRender = render
}
