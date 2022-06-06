package tfc

import "time"

type RegistryType string

func (rt RegistryType) String() string {
	return string(rt)
}

const (
	RegistryTypePrivate RegistryType = "private"
	RegistryTypePublic  RegistryType = "public"
)

type EntityType string

func (rt EntityType) String() string {
	return string(rt)
}

const (
	EntityTypeRegistryProviders                EntityType = "registry-providers"
	EntityTypeRegistryProviderVersions         EntityType = "registry-provider-versions"
	EntityTypeRegistryProviderVersionPlatforms EntityType = "registry-provider-version-platforms"
	EntityTypeRegistryModuleVersions           EntityType = "registry-module-versions"
)

type Links struct {
	First                *string `json:"first,omitempty"`
	Last                 *string `json:"last,omitempty"`
	Next                 *string `json:"next,omitempty"`
	Prev                 *string `json:"prev,omitempty"`
	Self                 *string `json:"self,omitempty"`
	Related              *string `json:"related,omitempty"`
	ProviderBinaryUpload *string `json:"provider-binary-upload,omitempty"`
	SHASUMsSigUpload     *string `json:"shasums-sig-upload,omitempty"`
	SHASUMsUpload        *string `json:"shasums-upload,omitempty"`
	Upload               *string `json:"upload,omitempty"`
}

type Permissions struct {
	CanDelete      *bool `json:"can-delete,omitempty"`
	CanUploadAsset *bool `json:"can-upload-asset,omitempty"`
}

type EntityDataAttributes struct {
	CreatedAt   time.Time   `json:"created-at"`
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	Permissions Permissions `json:"permissions"`
	UpdatedAt   time.Time   `json:"updated-at"`

	// optional fields

	KeyID              *string       `json:"key-id,omitempty"`
	Protocols          []string      `json:"protocols,omitempty"`
	RegistryName       *RegistryType `json:"registry-name,omitempty"`
	SHASUMsUploaded    *bool         `json:"shasums-uploaded,omitempty"`
	SHASUMsSigUploaded *bool         `json:"shasums-sig-uploaded,omitempty"`
	Version            *string       `json:"version,omitempty"`
}

type EntityData struct {
	ID   string     `json:"id"`
	Type EntityType `json:"type"`
}

type Entity struct {
	Data EntityData `json:"data"`
}

type Entities []Entity

func (ens Entities) EntitiesByType(entityType EntityType) (Entities, bool) {
	out := make(Entities, 0)
	for _, en := range ens {
		if en.Data.Type == entityType {
			out = append(out, en)
		}
	}
	return out, len(out) > 0
}

type Versions struct {
	Data  Entities `json:"data"`
	Links Links    `json:"links"`
}

type Relationships struct {
	Versions *Versions `json:"versions,omitempty"`

	Organization              *Entity `json:"organization,omitempty"`
	RegistryModule            *Entity `json:"registry-module,omitempty"`
	RegistryProvider          *Entity `json:"registry-provider,omitempty"`
	RegistryProviderPlatforms *Entity `json:"registry-provider-platforms,omitempty"`
	RegistryProviderVersion   *Entity `json:"registry-provider-version,omitempty"`
}

type ProviderDataAttributes struct {
}

type ProviderData struct {
	EntityData
	Attributes    ProviderDataAttributes `json:"attributes"`
	Links         Links                  `json:"links"`
	Relationships Relationships          `json:"relationships"`
}

type Provider struct {
	Data ProviderData `json:"data"`
}

type Providers []Provider

func (ps Providers) ByName(name string) (Provider, bool) {
	for _, p := range ps {
		if p.Data.Attributes.Name == name {
			return p, true
		}
	}
	return Provider{}, false
}

func (ps Providers) ByID(id string) (Provider, bool) {
	for _, p := range ps {
		if p.Data.ID == id {
			return p, true
		}
	}
	return Provider{}, false
}

type MetaPagination struct {
	CurrentPage *int `json:"current-page"`
	NextPage    *int `json:"next-page"`
	PageSize    *int `json:"page-size"`
	PrevPage    *int `json:"prev-page"`
	TotalCount  *int `json:"total-count"`
	TotalPages  *int `json:"total-pages"`
}

type Meta struct {
	Pagination MetaPagination `json:"pagination"`
}

type ListProvidersResponse struct {
	Data  Providers `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

type (
	CreateProviderRequestDataAttributes struct {
		Name         string `json:"name"`
		Namespace    string `json:"namespace"`
		RegistryName string `json:"registry-name"`
	}

	CreateProviderRequestData struct {
		Attributes CreateProviderRequestDataAttributes `json:"attributes"`
		Type       string                              `json:"type"`
	}

	CreateProviderRequest struct {
		Data CreateProviderRequestData `json:"data"`
	}
)

func NewCreateProviderRequest(providerName, orgNamespace string, registryType RegistryType) CreateProviderRequest {
	m := CreateProviderRequest{
		Data: CreateProviderRequestData{
			Type: string(EntityTypeRegistryProviders),
			Attributes: CreateProviderRequestDataAttributes{
				Name:         providerName,
				Namespace:    orgNamespace,
				RegistryName: string(registryType),
			},
		},
	}

	return m
}

type (
	CreateProviderVersionRequestDataAttributes struct {
		Version   string   `json:"version"`
		KeyID     string   `json:"key-id"`
		Protocols []string `json:"protocols"`
	}
	CreateProviderVersionRequestData struct {
		Type       string                                     `json:"type"`
		Attributes CreateProviderVersionRequestDataAttributes `json:"attributes"`
	}

	CreateProviderVersionRequest struct {
		Data CreateProviderVersionRequestData `json:"data"`
	}
)

func NewCreateProviderVersionRequest(vers, keyID string, protocols []string) CreateProviderVersionRequest {
	m := CreateProviderVersionRequest{
		Data: CreateProviderVersionRequestData{
			Type: string(EntityTypeRegistryProviderVersions),
			Attributes: CreateProviderVersionRequestDataAttributes{
				Protocols: protocols,
				Version:   vers,
				KeyID:     keyID,
			},
		},
	}
	return m
}

type (
	CreateProviderVersionResponseDataAttributes struct {
		CreatedAt          string      `json:"created-at"`
		KeyID              string      `json:"key-id"`
		Permissions        Permissions `json:"permissions"`
		Protocols          []string    `json:"protocols"`
		ShasumsSigUploaded bool        `json:"shasums-sig-uploaded"`
		ShasumsUploaded    bool        `json:"shasums-uploaded"`
		UpdatedAt          string      `json:"updated-at"`
		Version            string      `json:"version"`
	}

	CreateProviderVersionResponseData struct {
		EntityData
		Attributes    CreateProviderVersionResponseDataAttributes `json:"attributes"`
		Links         Links                                       `json:"links"`
		Relationships Relationships                               `json:"relationships"`
	}

	CreateProviderVersionResponse struct {
		Data CreateProviderVersionResponseData `json:"data"`
	}
)

type (
	CreateProviderVersionPlatformRequestDataAttributes struct {
		Arch     string `json:"arch"`
		Filename string `json:"filename"`
		OS       string `json:"os"`
		Shasum   string `json:"shasum"`
	}

	CreateProviderVersionPlatformRequestData struct {
		Attributes CreateProviderVersionPlatformRequestDataAttributes `json:"attributes"`
		Type       string                                             `json:"type"`
	}

	CreateProviderVersionPlatformRequest struct {
		Data CreateProviderVersionPlatformRequestData `json:"data"`
	}
)

func NewCreateProviderVersionPlatformRequest(os, arch, shasum, filename string) CreateProviderVersionPlatformRequest {
	m := CreateProviderVersionPlatformRequest{
		Data: CreateProviderVersionPlatformRequestData{
			Type: string(EntityTypeRegistryProviderVersionPlatforms),
			Attributes: CreateProviderVersionPlatformRequestDataAttributes{
				Arch:     arch,
				Filename: filename,
				OS:       os,
				Shasum:   shasum,
			},
		},
	}

	return m
}

type (
	CreateProviderVersionPlatformResponseDataAttributes struct {
		Arch                   string      `json:"arch"`
		Filename               string      `json:"filename"`
		Os                     string      `json:"os"`
		Permissions            Permissions `json:"permissions"`
		ProviderBinaryUploaded bool        `json:"provider-binary-uploaded"`
		Shasum                 string      `json:"shasum"`
	}

	CreateProviderVersionPlatformResponseData struct {
		EntityData
		Attributes    CreateProviderVersionPlatformResponseDataAttributes `json:"attributes"`
		Links         Links                                               `json:"links"`
		Relationships Relationships                                       `json:"relationships"`
	}

	CreateProviderVersionPlatformResponse struct {
		Data CreateProviderVersionPlatformResponseData `json:"data"`
	}
)

type (
	CreateModuleVersionRequestDataAttributes struct {
		Version string `json:"version"`
	}

	CreateModuleVersionRequestData struct {
		Attributes CreateModuleVersionRequestDataAttributes `json:"attributes"`
		Type       string                                   `json:"type"`
	}

	CreateModuleVersionRequest struct {
		Data CreateModuleVersionRequestData `json:"data"`
	}
)

func NewCreateModuleVersionRequest(version string) CreateModuleVersionRequest {
	m := CreateModuleVersionRequest{
		Data: CreateModuleVersionRequestData{
			Type: string(EntityTypeRegistryModuleVersions),
			Attributes: CreateModuleVersionRequestDataAttributes{
				Version: version,
			},
		},
	}

	return m
}

type (
	CreateModuleVersionResponseDataAttributes struct {
		CreatedAt time.Time `json:"created-at"`
		Source    string    `json:"source"`
		Status    string    `json:"status"`
		UpdatedAt time.Time `json:"updated-at"`
		Version   string    `json:"version"`
	}

	CreateModuleVersionResponseData struct {
		EntityData
		Attributes    CreateModuleVersionResponseDataAttributes `json:"attributes"`
		Links         Links                                     `json:"links"`
		Relationships Relationships                             `json:"relationships"`
	}

	CreateModuleVersionResponse struct {
		Data CreateModuleVersionResponseData `json:"data"`
	}
)
