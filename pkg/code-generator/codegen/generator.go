package codegen

import (
	"bytes"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/solo-io/solo-kit/pkg/code-generator/templates"
)

const fileHeader = `// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

`

type File struct {
	Filename string
	Content  string
}

type Files []File

func GenerateFiles(project *Project) (Files, error) {
	files, err := generateFilesForProject(project)
	if err != nil {
		return nil, err
	}
	for _, res := range project.Resources {
		// only generate files for the resources in our group, otherwise we import
		if res.GroupName != project.GroupName {
			continue
		}
		fs, err := generateFilesForResource(res)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}
	for _, grp := range project.ResourceGroups {
		fs, err := generateFilesForResourceGroup(grp)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}

	for _, res := range project.XDSResources {
		fs, err := generateFilesForXdsResource(res)
		if err != nil {
			return nil, err
		}
		files = append(files, fs...)
	}
	for i := range files {
		files[i].Content = fileHeader + files[i].Content
	}
	return files, nil
}

func generateFilesForXdsResource(resource *XDSResource) (Files, error) {
	var v Files
	for suffix, tmpl := range map[string]*template.Template{
		"_xds.sk.sk.go": templates.XdsTemplate,
	} {
		content, err := generateXdsResourceFile(resource, tmpl)
		if err != nil {
			return nil, err
		}
		v = append(v, File{
			Filename: strcase.ToSnake(resource.Name) + suffix,
			Content:  content,
		})
	}
	return v, nil
}
func generateFilesForResource(resource *Resource) (Files, error) {
	var v Files
	for suffix, tmpl := range map[string]*template.Template{
		".sk.go":             templates.ResourceTemplate,
		"_client.sk.go":      templates.ResourceClientTemplate,
		"_client_test.sk.go": templates.ResourceClientTestTemplate,
		"_reconciler.sk.go":  templates.ResourceReconcilerTemplate,
	} {
		content, err := generateResourceFile(resource, tmpl)
		if err != nil {
			return nil, err
		}
		v = append(v, File{
			Filename: strcase.ToSnake(resource.Name) + suffix,
			Content:  content,
		})
	}
	return v, nil
}

func generateFilesForResourceGroup(rg *ResourceGroup) (Files, error) {
	var v Files
	for suffix, tmpl := range map[string]*template.Template{
		"_snapshot.sk.go":              templates.ResourceGroupSnapshotTemplate,
		"_snapshot_emitter.sk.go":      templates.ResourceGroupEmitterTemplate,
		"_snapshot_emitter_test.sk.go": templates.ResourceGroupEmitterTestTemplate,
		"_event_loop.sk.go":            templates.ResourceGroupEventLoopTemplate,
		"_event_loop_test.sk.go":       templates.ResourceGroupEventLoopTestTemplate,
	} {
		content, err := generateResourceGroupFile(rg, tmpl)
		if err != nil {
			return nil, err
		}
		v = append(v, File{
			Filename: strcase.ToSnake(rg.GoName) + suffix,
			Content:  content,
		})
	}
	return v, nil
}

func generateFilesForProject(project *Project) (Files, error) {
	var v Files
	for suffix, tmpl := range map[string]*template.Template{
		"_suite_test.sk.go": templates.ProjectTestSuiteTemplate,
	} {
		content, err := generateProjectFile(project, tmpl)
		if err != nil {
			return nil, err
		}
		v = append(v, File{
			Filename: strcase.ToSnake(project.Name) + suffix,
			Content:  content,
		})
	}
	return v, nil
}

func generateXdsResourceFile(resource *XDSResource, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, resource); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func generateResourceFile(resource *Resource, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, resource); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func generateResourceGroupFile(rg *ResourceGroup, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, rg); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func generateProjectFile(project *Project, tmpl *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, project); err != nil {
		return "", err
	}
	return buf.String(), nil
}
