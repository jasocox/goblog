package view

import "net/http"

var (
	execHeader = templateExecer{header, nil}
	execFooter = templateExecer{footer, nil}
)

type templateExecer struct {
	name string
	data interface{}
}

func (e templateExecer) executeTemplate(w http.ResponseWriter) error {
	return templates.ExecuteTemplate(w, e.name, e.data)
}

func executeTemplates(execers []templateExecer, w http.ResponseWriter) (err error) {
	for _, e := range execers {
		err = e.executeTemplate(w)

		if err != nil {
			break
		}
	}

	return
}

func execLayout(w http.ResponseWriter, exec func() error) (err error) {
	err = execHeader.executeTemplate(w)

	if err == nil {
		err = exec()
	}

	if err == nil {
		err = execFooter.executeTemplate(w)
	}

	return
}
