package api

import "time"

type workItemResponse struct {
	Links struct {
		HTML struct {
			HREF string `json:"href"`
		} `json:"html"`
		// there are more here, fields, self, workItemComments, workItemRevisions, workItemType, and workItemUpdates
	} `json:"_links"`
	Fields struct {
		AssignedTo     usersResponse `json:"System.AssignedTo"`
		ChangedDate    time.Time     `json:"System.ChangedDate"`
		CreatedDate    time.Time     `json:"System.CreatedDate"`
		CreatedBy      usersResponse `json:"System.CreatedBy"`
		Description    string        `json:"System.Description"`
		DueDate        time.Time     `json:"Microsoft.VSTS.Scheduling.DueDate"` // ??
		IterationPath  string        `json:"System.IterationPath"`
		TeamProject    string        `json:"System.TeamProject"`
		Priority       int           `json:"Microsoft.VSTS.Common.Priority"`
		Reason         string        `json:"System.Reason"`
		ResolvedReason string        `json:"Microsoft.VSTS.Common.ResolvedReason"`
		ResolvedDate   time.Time     `json:"Microsoft.VSTS.Common.ResolvedDate"`
		StoryPoints    float64       `json:"Microsoft.VSTS.Scheduling.StoryPoints"`
		State          string        `json:"System.State"`
		Tags           string        `json:"System.Tags"`
		Title          string        `json:"System.Title"`
		WorkItemType   string        `json:"System.WorkItemType"`
	} `json:"fields"`
	Relations []struct {
		Rel string `json:"rel"`
		URL string `json:"url"`
	} `json:"relations"`
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type usersResponse struct {
	Descriptor  string `json:"descriptor"`
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
	ImageURL    string `json:"imageUrl"`
	UniqueName  string `json:"uniqueName"`
	URL         string `json:"url"`
}
type resolutionResponse struct {
	AllowedValues []string `json:"allowedValues"`
	Name          string   `json:"name"`
	ReferenceName string   `json:"referenceName"`
}
type workConfigResponse struct {
	Name          string `json:"name"`
	ReferenceName string `json:"referenceName"`
	States        []struct {
		Category string `json:"category"`
		Name     string `json:"name"`
	} `json:"states"`
	FieldInstances []struct {
		ReferenceName string `json:"referenceName"`
	} `json:"fieldInstances"`
	Fields []struct {
		ReferenceName string `json:"referenceName"`
	} `json:"fields"`
}

type workConfigStatus string

// These seem to be the default statuses
const workConfigCompletedStatus = workConfigStatus("Completed")
const workConfigInProgressStatus = workConfigStatus("InProgress")
const workConfigProposedStatus = workConfigStatus("Proposed")
const workConfigRemovedStatus = workConfigStatus("Removed")
const workConfigResolvedStatus = workConfigStatus("Resolved")

type changelogField struct {
	NewValue   interface{} `json:"newValue"`
	OldValue   interface{} `json:"oldvalue"`
	customerID string
	refType    string
}

type changelogResponse struct {
	Fields      map[string]changelogField `json:"fields"`
	ID          int64                     `json:"id"`
	RevisedDate time.Time                 `json:"revisedDate"`
	URL         string                    `json:"url"`
	Relations   struct {
		Added []struct {
			Attributes struct {
				Name string `json:"name"`
			} `json:"attributes"`
			URL string `json:"url"`
		} `json:"added"`
		Removed []struct {
			Attributes struct {
				Name string `json:"name"`
			} `json:"attributes"`
			URL string `json:"url"`
		} `json:"removed"`
	} `json:"relations"`
	RevisedBy usersResponse `json:"revisedBy"`
}

type sprintsResponse struct {
	Attributes struct {
		FinishDate time.Time `json:"finishDate"`
		StartDate  time.Time `json:"startDate"`
		TimeFrame  string    `json:"timeFrame"` // past, current, future
	} `json:"attributes"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"url"`
}

type teamsResponse struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	IdentityURL string `json:"identityUrl"`
	Name        string `json:"name"`
	ProjectID   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	URL         string `json:"url"`
}

type workItemsResponse struct {
	AsOf    time.Time `json:"asOf"`
	Columns []struct {
		Name          string `json:"name"`
		ReferenceName string `json:"referenceName"`
		URL           string `json:"url"`
	} `json:"columns"`
	QueryResultType string `json:"queryResultType"`
	QueryType       string `json:"queryType"`
	SortColumns     []struct {
		Descending bool `json:"descending"`
		Field      struct {
			Name          string `json:"name"`
			ReferenceName string `json:"referenceName"`
			URL           string `json:"url"`
		} `json:"field"`
	} `json:"sortColumns"`
	WorkItems []struct {
		ID  int64  `json:"id"`
		URL string `json:"url"`
	} `json:"workItems"`
}

type projectResponseLight struct {
	ID             string `json:"id"`
	LastUpdateTime string `json:"lastUpdateTime"` // not in TFS
	Name           string `json:"name"`
	State          string `json:"state"`
}

type projectResponse struct {
	projectResponseLight
	Revision    int64  `json:"revision"`
	State       string `json:"state"`
	URL         string `json:"url"`
	Visibility  string `json:"visibility"`
	Description string `json:"description"`
}

type usersResponseAzure struct {
	Identity usersResponse `json:"identity"`
}