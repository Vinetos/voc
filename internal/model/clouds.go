package model

import (
	"strings"
)

type CloudConfig struct {
	Clouds map[string]Cloud `yaml:"clouds"`
}

type Cloud struct {
	Auth               AuthInfo       `yaml:"auth"`
	RegionName         string         `yaml:"region_name,omitempty"`
	Interface          string         `yaml:"interface,omitempty"`
	IdentityAPIVersion string         `yaml:"identity_api_version,omitempty"`
	AuthType           string         `yaml:"auth_type,omitempty"`
	Verify             *bool          `yaml:"verify,omitempty"`
	CACert             string         `yaml:"cacert,omitempty"`
	ClientCert         string         `yaml:"cert,omitempty"`
	ClientKey          string         `yaml:"key,omitempty"`
	EndpointOverride   string         `yaml:"endpoint_override,omitempty"`
	AdditionalFields   map[string]any `yaml:",inline"` // pour tolérer des champs non mappés
}

type AuthInfo struct {
	AuthURL                     string `yaml:"auth_url"`
	Username                    string `yaml:"username,omitempty"`
	Password                    string `yaml:"password,omitempty"`
	UserID                      string `yaml:"user_id,omitempty"`
	ProjectID                   string `yaml:"project_id,omitempty"`
	ProjectName                 string `yaml:"project_name,omitempty"`
	UserDomainID                string `yaml:"user_domain_id,omitempty"`
	UserDomainName              string `yaml:"user_domain_name,omitempty"`
	ProjectDomainID             string `yaml:"project_domain_id,omitempty"`
	ProjectDomainName           string `yaml:"project_domain_name,omitempty"`
	DomainID                    string `yaml:"domain_id,omitempty"`
	DomainName                  string `yaml:"domain_name,omitempty"`
	Token                       string `yaml:"token,omitempty"`
	ApplicationCredentialID     string `yaml:"application_credential_id,omitempty"`
	ApplicationCredentialName   string `yaml:"application_credential_name,omitempty"`
	ApplicationCredentialSecret string `yaml:"application_credential_secret,omitempty"`
}

func (config CloudConfig) ColumnHeaders() []string {
	return strings.Split("CLOUD,REGION,PROJECT,USER", ",")
}

func (config CloudConfig) RowData() [][]string {
	// Map the data to a printable format
	var data [][]string
	for CloudName, Cloud := range config.Clouds {
		row := []string{
			CloudName,
			Cloud.RegionName,
			Cloud.Auth.ProjectName,
			Cloud.Auth.Username,
		}
		data = append(data, row)
	}
	return data
}

func (config CloudConfig) ShortName() string {
	return "Clouds"
}

func (config CloudConfig) ListClouds() ([]string, error) {
	var cloudNames []string
	for name := range config.Clouds {
		cloudNames = append(cloudNames, name)
	}

	return cloudNames, nil
}
