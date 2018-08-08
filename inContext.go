// Command example runs a sample webserver that uses go-i18n/v2/i18n.
package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// get from config
var isPhraseAppEnabled bool

func init()  {
	flag.BoolVar(&isPhraseAppEnabled,"phraseApp", false, "Enable PhraseApp mode")
	flag.Parse()
}

var apiToken = os.Getenv("PHRASE_APP_TOKEN")

func translate(s string) string  {
	if isPhraseAppEnabled {
		return "{{__phrase_" + s + "__}}"
	} else {
		return s
	}
}

var funcs = template.FuncMap{
"translate": translate,
}

var page = template.Must(template.New("").Funcs(funcs).Parse(`
<!DOCTYPE html>
<html lang= {{ .CurrentLocale }}>
<body>
<h1>{{ translate .Title }}</h1>
{{range .Paragraphs}}<p>{{ translate . }}</p>{{end}}
<script>
window.PHRASEAPP_CONFIG = {
   projectId: {{ .apiToken }}
};
(function() {
   var phraseapp = document.createElement('script'); phraseapp.type = 'text/javascript'; phraseapp.async = true;
   phraseapp.src = ['https://', 'phraseapp.com/assets/in-context-editor/2.0/app.js?', new Date().getTime()].join('');
   var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(phraseapp, s);
})();
</script>
</body>
</html>
`))

func main() {

	bundle := &i18n.Bundle{DefaultLanguage: language.Greek}
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _,lang := range []string{"en" ,"el"} {
		bundle.MustLoadMessageFile(fmt.Sprintf("active.%v.json", lang))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		name := r.FormValue("name")
		if name == "" {
			name = "Alex"
		}

		myInvoicesCount := 10

		helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "HelloPerson",
			},
			TemplateData: map[string]interface{}{
				"Name": name,
			},
		})

		myInvoices := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "invoices",
			},
			TemplateData: map[string]interface{}{
				"Count": myInvoicesCount,
			},
			PluralCount: myInvoicesCount,
		})

		err := page.Execute(w, map[string]interface{}{
			"apiToken": apiToken,
			"Title": helloPerson,
			"CurrentLocale": language.Greek.String(),
			"Paragraphs": []string{
				myInvoices,
			},
		})
		if err != nil {
			panic(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}