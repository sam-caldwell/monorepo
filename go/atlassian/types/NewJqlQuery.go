package AtlassianTypes

import "github.com/sam-caldwell/monorepo/go/list"

// NewJqlQuery - Create, configure and return a new JqlQuery structure
func NewJqlQuery(expand []string, fields []string, FieldByKeys bool,
	Jql string, MaxResults uint, startAt uint) (jql *JqlQuery, err error) {

	const (
		QueryValidation   = "strict"
		defaultMaxResults = 50
		MaxStartAt        = 1000
	)

	if len(expand) == 0 {
		expand = []string{
			"names",
			"schema",
			"operations",
		}
	}

	if len(fields) == 0 {
		fields = []string{
			"summary",
			"status",
		}
	}

	allowedExpandValues := []string{
		"renderedFields",
		"names",
		"schema",
		"transitions",
		"operations",
		"editmeta",
		"changelog",
	}

	if err = list.LeftSetElementsNotInRightSet(&expand, &allowedExpandValues); err != nil {
		return nil, err
	}

	//ToDo: Sanitize the fields list to ensure fields are valid

	//ToDo: Validate JQL Query (requires separate call to the Jira server

	if MaxResults == 0 {
		MaxResults = defaultMaxResults
	}

	if startAt > MaxStartAt {
		startAt = MaxStartAt //StartAt cannot be > MaxStartAt
	}

	jql = &JqlQuery{
		Expand:        expand,
		Fields:        fields,
		FieldsByKeys:  FieldByKeys,
		Jql:           Jql,
		MaxResults:    MaxResults,
		StartAt:       startAt,
		ValidateQuery: QueryValidation,
	}

	return jql, err

}
