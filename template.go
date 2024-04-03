package smtppool

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"path/filepath"
)

type EmailTpl struct {
	subject *template.Template
	body    *template.Template
}

type HTML struct {
	Subject string
	Body    string
}

// InitEmailTpl loads the template.
func InitEmailTpl(subj, tplFile string) (*EmailTpl, error) {
	out := &EmailTpl{}

	if tplFile != "" {
		tpl, err := template.New(filepath.Base(tplFile)).Funcs(sprig.FuncMap()).ParseFiles(tplFile)

		if err != nil {
			return nil, errors.New(fmt.Sprintf("error parsing template file: %s: %v", tplFile, err))
		}
		out.body = tpl
	}

	// Subject template string.
	if subj != "" {
		tpl, err := template.New("subject").Parse(subj)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error parsing template subject: %s: %v", tplFile, err))
		}

		out.subject = tpl
	}
	return out, nil
}

// GetHTML stuffs EmailTpl with data. Returns body and subject
// as string
func GetHTML(tpl *EmailTpl, data interface{}) (*HTML, error) {
	var (
		subj = &bytes.Buffer{}
		out  = &bytes.Buffer{}
	)

	if tpl != nil {
		if tpl.subject != nil {
			if err := tpl.subject.Execute(subj, data); err != nil {
				return nil, err
			}
		}

		if tpl.body != nil {
			if err := tpl.body.Execute(out, data); err != nil {
				return nil, err
			}
		}
	}

	return &HTML{
		Subject: subj.String(),
		Body:    out.String(),
	}, nil
}
