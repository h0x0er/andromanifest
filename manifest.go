package andromanifest

import "encoding/xml"

type Manifest struct {
	XMLName xml.Name `xml:"manifest"`

	Package                  string `xml:"package,attr"`
	SdkVersion               string `xml:"compileSdkVersion,attr"`
	SharedUserID             string `xml:"sharedUserId,attr"`
	SharedUserLabel          string `xml:"sharedUserLabel"`
	PlatformBuildVersionName string `xml:"platformBuildVersionName,attr"`
	VersionCode              string `xml:"versionCode,attr"`

	Permissions     []Permission     `xml:"permission"`
	UsesPermissions []UsesPermission `xml:"uses-permission"`
	UsesFeatures    []UsesFeature    `xml:"uses-feature"`

	Application Application `xml:"application"`
}

func (m *Manifest) GetActivites() []Activity {
	return m.Application.Activities
}

func (m *Manifest) GetActivitesAliases() []ActivityAlias {
	return m.Application.ActivityAliases
}

func (m *Manifest) GetServices() []Service {
	return m.Application.Services
}

func (m *Manifest) GetReceivers() []Receiver {
	return m.Application.Receivers
}

func (m *Manifest) GetProviders() []Provider {
	return m.Application.Providers
}

type Permission struct {
	XMLName         xml.Name `xml:"permission"`
	Name            string   `xml:"name,attr"`
	ProtectionLevel string   `xml:"protectionLevel,attr,omitempty"`
}

type UsesPermission struct {
	XMLName xml.Name `xml:"uses-permission"`
	Name    string   `xml:"name,attr"`
}

type UsesFeature struct {
	XMLName  xml.Name `xml:"uses-feature"`
	Name     string   `xml:"name,attr"`
	Required string   `xml:"required,attr,omitempty"`
	Version  string   `xml:"version,attr,omitempty"`
}

type Application struct {
	XMLName xml.Name `xml:"application"`

	AllowBackup          string
	BackupAgent          string
	Icon                 string
	Label                string
	HardwareAccelerated  string
	Name                 string
	RestoreAnyVersion    string
	RequiredAccountType  string
	TaskAffinity         string
	UsesCleartextTraffic string
	SupportRtl           string

	UsesLibraries   []UsesLibrary   `xml:"uses-library,omitempty"`
	MetaDatas       []MetaData      `xml:"meta-data,omitempty"`
	Activities      []Activity      `xml:"activity"`
	ActivityAliases []ActivityAlias `xml:"activity-alias"`
	Services        []Service       `xml:"service"`
	Providers       []Provider      `xml:"provider"`
	Receivers       []Receiver      `xml:"receiver"`
}

type UsesLibrary struct {
	XMLName xml.Name `xml:"uses-library"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type MetaData struct {
	XMLName xml.Name `xml:"meta-data"`

	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Activity struct {
	XMLName            xml.Name `xml:"activity"`
	Exported           string   `xml:"exported,attr,omitempty"`
	Name               string   `xml:"name,attr,omitempty"`
	LaunchMode         string   `xml:"launchMode,attr,omitempty"`
	Theme              string   `xml:"theme,attr,omitempty"`
	TaskAffinity       string   `xml:"taskAffinity,attr,omitempty"`
	TargetActivity     string   `xml:"targetActivity,attr,omitempty"`
	ExcludeFromRecents string   `xml:"excludeFromRecents,attr,omitempty"`

	IntentFilters []IntentFilter `xml:"intent-filter,omitempty"`
}

type ActivityAlias struct {
	XMLName            xml.Name `xml:"activity-alias"`
	Exported           string   `xml:"exported,attr,omitempty"`
	Name               string   `xml:"name,attr,omitempty"`
	LaunchMode         string   `xml:"launchMode,attr,omitempty"`
	Theme              string   `xml:"theme,attr,omitempty"`
	TaskAffinity       string   `xml:"taskAffinity,attr,omitempty"`
	TargetActivity     string   `xml:"targetActivity,attr,omitempty"`
	ExcludeFromRecents string   `xml:"excludeFromRecents,attr,omitempty"`

	IntentFilters []IntentFilter `xml:"intent-filter,omitempty"`
}

type Service struct {
	XMLName    xml.Name `xml:"service"`
	Exported   string   `xml:"exported,attr,omitempty"`
	Name       string   `xml:"name,attr,omitempty"`
	Permission string   `xml:"permission,attr,omitempty"`

	IntentFilters []IntentFilter `xml:"intent-filter,omitempty"`
}

type Receiver struct {
	XMLName    xml.Name `xml:"receiver"`
	Exported   string   `xml:"exported,attr,omitempty"`
	Name       string   `xml:"name,attr,omitempty"`
	Permission string   `xml:"permission,attr,omitempty"`

	IntentFilters []IntentFilter `xml:"intent-filter,omitempty"`
}

type Provider struct {
	XMLName             xml.Name `xml:"provider"`
	Exported            string   `xml:"exported,attr,omitempty"`
	Name                string   `xml:"name,attr,omitempty"`
	Authorities         string   `xml:"authorities,omitempty"`
	GrantUriPermissions string   `xml:"grantUriPermissions,attr"`

	Permission string `xml:"permission,attr,omitempty"`

	IntentFilters []IntentFilter `xml:"intent-filter,omitempty"`
}

type IntentFilter struct {
	XMLName xml.Name `xml:"intent-filter"`

	Actions    []Action   `xml:"action,omitempty"`
	Categories []Category `xml:"category,omitempty"`
	Datas      []Data     `xml:"data,omitempty"`
}

type Action struct {
	XMLName xml.Name `xml:"action"`
	Name    string   `xml:"name,attr"`
}

type Category struct {
	XMLName xml.Name `xml:"category"`
	Name    string   `xml:"name,attr"`
}

type Data struct {
	XMLName xml.Name `xml:"data"`

	MimeType   string `xml:"mime-type,attr,omitempty"`
	Scheme     string `xml:"scheme,attr,omitempty"`
	Host       string `xml:"host,attr,omitempty"`
	PathPrefix string `xml:"pathPrefix,attr,omitempty"`
}
