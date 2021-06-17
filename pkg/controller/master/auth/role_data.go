package auth

const (
	bootstrapAdminConfig   = "admincreated"
	defaultAdminLabelKey   = "auth.harvesterhci.io/bootstrapping"
	defaultAdminLabelValue = "admin-user"
	usernameLabelKey       = "harvesterhci.io/username"
	defaultAdminPassword   = "password"
)

var defaultAdminLabel = map[string]string{
	defaultAdminLabelKey: defaultAdminLabelValue,
}

// bootstrapAdmin checks if the bootstrapAdminConfig exists, if it does this indicates it has
// already created the admin user and should not attempt it again. Otherwise attempt to create the admin.
