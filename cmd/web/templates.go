package main

import (
	"html/template"
	"path/filepath"

	"github.com/rguichette/webclip/pkg/models"
)

//define a templateData type to act a the holding structures for any -->DYNAMIC<--  data the we want to pass to our HTML templates.
//At the moment, it only contains one field, but we'll add more to it as the build progresses.

type templateData struct {
	Clip  *models.Clip
	Clips []*models.Clip
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	//init new map to act as cache
	cache := map[string]*template.Template{}
	//use filepath.Glob() to get a slice of all filepaths with the .page.tmpl extension
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))

	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		//extract the file name from the full file path and assign it to the name variable
		name := filepath.Base(page)
		//Parse the page template file in to a template set.
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// use the parseGlob method to add any "layout" templates to the template se(just 'base' in this case)
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		//use the parseGlob to add any 'partial templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}
		//add the template set to the cache, using the name of the page as the key
		cache[name] = ts
		//return the map

	}
	return cache, nil
}
