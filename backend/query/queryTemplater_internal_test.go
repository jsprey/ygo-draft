package query

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"text/template"
)

func TestNewSqlQueryTemplater(t *testing.T) {
	t.Run("create new templater", func(t *testing.T) {
		// when
		templater, err := NewSqlQueryTemplater()

		// then
		require.NoError(t, err)
		require.NotNil(t, templater)
	})

	t.Run("creating new templater fails of invalid template", func(t *testing.T) {
		// given
		originalTemplate := TemplateContentSelectCardByID
		defer func() { TemplateContentSelectCardByID = originalTemplate }()

		TemplateContentSelectCardByID = "SELECT * FROM public.cards {{}{}}"

		// when
		templater, err := NewSqlQueryTemplater()

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "")
		assert.Nil(t, templater)
	})
}

func TestSqlQueryTemplater_Template(t *testing.T) {
	t.Run("template successfully", func(t *testing.T) {
		// given
		myTemplate, err := template.New("myTemplate").Parse("This {{.Message}}")
		templater := sqlQueryTemplater{Templates: map[string]*template.Template{
			myTemplate.Name(): myTemplate,
		}}

		// when
		messageObject := struct {
			Message string
		}{Message: "is a template that I generated myself!"}
		result, err := templater.Template("myTemplate", messageObject)

		// then
		require.NoError(t, err)
		require.Equal(t, "This is a template that I generated myself!", result)
	})

	t.Run("error on executing template", func(t *testing.T) {
		// given
		emptyTemplate := template.New("emptyTemplate")
		templater := sqlQueryTemplater{Templates: map[string]*template.Template{
			emptyTemplate.Name(): emptyTemplate,
		}}

		// when
		messageObject := struct {
			Message string
		}{Message: "is a template that I generated myself!"}
		_, err := templater.Template("emptyTemplate", messageObject)

		// then
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to execute sqlTemplate")
	})

	t.Run("templating non existing template", func(t *testing.T) {
		// given
		templater := sqlQueryTemplater{Templates: map[string]*template.Template{}}

		// when
		_, err := templater.Template("myNonexistingTemplate", "myMessageObject")

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, ErrorTemplateDoesNotExist)
	})
}
