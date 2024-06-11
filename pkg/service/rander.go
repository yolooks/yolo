package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/yolooks/yolo/pkg/template/cmd"
	"github.com/yolooks/yolo/pkg/template/config"
	"github.com/yolooks/yolo/pkg/template/controller"
	"github.com/yolooks/yolo/pkg/template/model"
	"github.com/yolooks/yolo/pkg/template/router"
	"github.com/yolooks/yolo/pkg/template/service"
	"github.com/yolooks/yolo/pkg/template/util/cm"
)

type Rander struct {
	ProjectName string
	ProjectPort int
}

func NewRander(projectName string, projectPort int) *Rander {
	return &Rander{
		ProjectName: projectName,
		ProjectPort: projectPort,
	}
}

func (r *Rander) InitDir() error {
	dirs := []string{"cmd", "pkg", "etc", "script"}
	for _, dir := range dirs {
		curPath := filepath.Join(r.ProjectName, dir)
		if err := os.MkdirAll(curPath, os.ModePerm); err != nil {
			return err
		}
		fmt.Printf("create dir: %s success\n", curPath)
	}
	return nil
}

func (r *Rander) InitPkg() error {
	command := fmt.Sprintf("cd %s/pkg && mkdir -p {config,controller,model,router,service,util/cm}", r.ProjectName)
	if err := exec.Command("/bin/bash", "-c", command).Run(); err != nil {
		return err
	}
	fmt.Println("pkg init success")

	// rander cmd
	cmdTpl, err := r.GenerateFile(cmd.TPL)
	if err != nil {
		return err
	}
	cmdFile := filepath.Join(r.ProjectName, "cmd", "server.go")
	if err := os.WriteFile(cmdFile, cmdTpl, 0644); err != nil {
		return err
	}
	fmt.Println("cmd rander success")

	// rander config
	etcTpl, err := r.GenerateFile(config.ETC_TPL)
	if err != nil {
		return err
	}
	etcFile := filepath.Join(r.ProjectName, "etc", "dev.yaml")
	if err := os.WriteFile(etcFile, etcTpl, 0644); err != nil {
		return err
	}

	configTpl, err := r.GenerateFile(config.CONFIG_TPL)
	if err != nil {
		return err
	}
	configFile := filepath.Join(r.ProjectName, "pkg", "config", "config.go")
	if err := os.WriteFile(configFile, configTpl, 0644); err != nil {
		return err
	}
	fmt.Println("config rander success")

	configLogTpl, err := r.GenerateFile(config.LOG_TPL)
	if err != nil {
		return err
	}
	configLogFile := filepath.Join(r.ProjectName, "pkg", "config", "log.go")
	if err := os.WriteFile(configLogFile, configLogTpl, 0644); err != nil {
		return err
	}
	fmt.Println("config log rander success")

	configDefineTpl, err := r.GenerateFile(config.CONFIG_DEFINE_TPL)
	if err != nil {
		return err
	}
	configDefineFile := filepath.Join(r.ProjectName, "pkg", "config", "define.go")
	if err := os.WriteFile(configDefineFile, configDefineTpl, 0644); err != nil {
		return err
	}
	fmt.Println("config define rander success")

	// rander service
	serviceTpl, err := r.GenerateFile(service.SERVICE_TPL)
	if err != nil {
		return err
	}
	serviceFile := filepath.Join(r.ProjectName, "pkg", "service", "service.go")
	if err := os.WriteFile(serviceFile, serviceTpl, 0644); err != nil {
		return err
	}
	fmt.Println("service rander success")

	// rander controller
	typeDefTpl, err := r.GenerateFile(controller.TYPEDEF_TPL)
	if err != nil {
		return err
	}
	typeDefFile := filepath.Join(r.ProjectName, "pkg", "controller", "typedef.go")
	if err := os.WriteFile(typeDefFile, typeDefTpl, 0644); err != nil {
		return err
	}
	fmt.Println("controller typedef rander success")

	controllerTpl, err := r.GenerateFile(controller.CONTROLLER_TPL)
	if err != nil {
		return err
	}
	controllerFile := filepath.Join(r.ProjectName, "pkg", "controller", "controller.go")
	if err := os.WriteFile(controllerFile, controllerTpl, 0644); err != nil {
		return err
	}
	fmt.Println("controller rander success")

	// rander router
	urlTpl, err := r.GenerateFile(router.URL_TPL)
	if err != nil {
		return err
	}
	urlFile := filepath.Join(r.ProjectName, "pkg", "router", "url.go")
	if err := os.WriteFile(urlFile, urlTpl, 0644); err != nil {
		return err
	}
	middlewareTpl, err := r.GenerateFile(router.MID_TPL)
	if err != nil {
		return err
	}
	middlewareFile := filepath.Join(r.ProjectName, "pkg", "router", "middleware.go")
	if err := os.WriteFile(middlewareFile, middlewareTpl, 0644); err != nil {
		return err
	}
	fmt.Println("router url rander success")

	// rander util
	cmdTpl, err = r.GenerateFile(cm.TPL)
	if err != nil {
		return err
	}
	cmdFile = filepath.Join(r.ProjectName, "pkg", "util", "cm", "cm.go")
	if err := os.WriteFile(cmdFile, cmdTpl, 0644); err != nil {
		return err
	}
	fmt.Println("util cm rander success")

	// rander model
	modelTpl, err := r.GenerateFile(model.TPL)
	if err != nil {
		return err
	}
	modelFile := filepath.Join(r.ProjectName, "pkg", "model", "base.go")
	if err := os.WriteFile(modelFile, modelTpl, 0644); err != nil {
		return err
	}
	fmt.Println("model rander success")
	return nil
}

func (r *Rander) GenerateFile(tpl string) ([]byte, error) {
	tmpl, err := template.New("yolo").Parse(tpl)
	if err != nil {
		fmt.Println("template file parse error: ", err)
		return nil, err
	}

	var tplOutput bytes.Buffer
	if err := tmpl.Execute(&tplOutput, r); err != nil {
		fmt.Println("exec template file get output error: ", err)
		return nil, err
	}
	return tplOutput.Bytes(), nil
}

func (r *Rander) RunGoMod() error {
	command := fmt.Sprintf("cd %s && go mod init %s", r.ProjectName, r.ProjectName)
	if err := exec.Command("/bin/bash", "-c", command).Run(); err != nil {
		return err
	}
	fmt.Printf("go mod init %s success\n", r.ProjectName)

	command = fmt.Sprintf("cd %s && go mod tidy", r.ProjectName)
	if err := exec.Command("/bin/bash", "-c", command).Run(); err != nil {
		return err
	}
	fmt.Println("go mod tidy success")
	return nil
}
