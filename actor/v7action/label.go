package v7action

import (
	"code.cloudfoundry.org/cli/types"
)

func (actor *Actor) UpdateApplicationLabelsByApplicationName(appName string, spaceGUID string, labels map[string]types.NullString) (Warnings, error) {
	app, appWarnings, err := actor.GetApplicationByNameAndSpace(appName, spaceGUID)
	if err != nil {
		return appWarnings, err
	}
	app.Metadata = &Metadata{Labels: labels}
	_, updateWarnings, err := actor.UpdateApplication(app)
	return append(appWarnings, updateWarnings...), err
}

func (actor *Actor) UpdateOrganizationLabelsByOrganizationName(orgName string, labels map[string]types.NullString) (Warnings, error) {
	org, warnings, err := actor.GetOrganizationByName(orgName)
	if err != nil {
		return warnings, err
	}
	org.Metadata = &Metadata{Labels: labels}

	// org: v7action.Organization
	_, updateWarnings, err := actor.UpdateOrganization(org)
	return append(warnings, updateWarnings...), err
}
