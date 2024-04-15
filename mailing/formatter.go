package mailing

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"transaction-processor/model"
)

func HTMLFormat(summary *model.TransactionSummary) (string, error) {
	// Read the template file
	summaryTemplateBytes, err := os.ReadFile("mailing/summary_template.html")
	if err != nil {
		return "", err
	}

	summaryTemplate := string(summaryTemplateBytes)

	// Define function to sort months
	funcMap := template.FuncMap{
		"printf": fmt.Sprintf,
		"MonthOrder": func() []string {
			return []string{
				"January", "February", "March", "April", "May", "June",
				"July", "August", "September", "October", "November", "December",
			}
		},
	}

	// Parse the template
	tmpl, err := template.New("summaryTemplate").Funcs(funcMap).Parse(summaryTemplate)
	if err != nil {
		return "", err
	}

	// Execute the Template for summary
	var result bytes.Buffer
	err = tmpl.Execute(&result, summary)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
