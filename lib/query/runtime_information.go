package query

import (
	"os"
	"strings"

	"github.com/mithrandie/csvq/lib/parser"
	"github.com/mithrandie/csvq/lib/value"
)

const (
	UncommittedInformation  = "UNCOMMITTED"
	CreatedInformation      = "CREATED"
	UpdatedInformation      = "UPDATED"
	UpdatedViewsInformation = "UPDATED_VIEWS"
	LoadedTablesInformation = "LOADED_TABLES"
	WorkingDirectory        = "WORKING_DIRECTORY"
	VersionInformation      = "VERSION"
)

var RuntimeInformatinList = []string{
	UncommittedInformation,
	CreatedInformation,
	UpdatedInformation,
	UpdatedViewsInformation,
	LoadedTablesInformation,
	WorkingDirectory,
	VersionInformation,
}

func GetRuntimeInformation(expr parser.RuntimeInformation) (value.Primary, error) {
	var p value.Primary

	switch strings.ToUpper(expr.Name) {
	case UncommittedInformation:
		p = value.NewBoolean(!UncommittedViews.IsEmpty())
	case CreatedInformation:
		p = value.NewInteger(int64(UncommittedViews.CountCreatedTables()))
	case UpdatedInformation:
		p = value.NewInteger(int64(UncommittedViews.CountUpdatedTables()))
	case UpdatedViewsInformation:
		p = value.NewInteger(int64(UncommittedViews.CountUpdatedViews()))
	case LoadedTablesInformation:
		p = value.NewInteger(int64(len(ViewCache)))
	case WorkingDirectory:
		wd, err := os.Getwd()
		if err != nil {
			return p, err
		}
		p = value.NewString(wd)
	case VersionInformation:
		p = value.NewString(Version)
	default:
		return p, NewInvalidRuntimeInformationError(expr)
	}

	return p, nil
}
