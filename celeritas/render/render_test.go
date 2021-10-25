package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{
		"go_page",
		"go",
		"home",
		false,
		"Error Rendering Go Template",
	},
	{
		"go_page_no_template",
		"go",
		"no-file",
		true,
		"No Error Rendering Non-Existing Go Template",
	},
	{
		"jet_page",
		"jet",
		"home",
		false,
		"Error Rendering Jet Template",
	},
	{
		"jet_page_no_template",
		"jet",
		"no-file",
		true,
		"No Error Rendering Non-Existing Jet Template",
	},
	{
		"invalid_renderer_engine",
		"",
		"home",
		true,
		"No Error Rendering Invalid Template",
	},
}

func TestRender_Page(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	for _, e := range pageData {
		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"
		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s: ", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s: ", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}
}
