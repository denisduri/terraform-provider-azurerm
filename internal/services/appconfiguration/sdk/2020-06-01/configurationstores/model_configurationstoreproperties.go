package configurationstores

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

type ConfigurationStoreProperties struct {
	CreationDate               *string                               `json:"creationDate,omitempty"`
	Encryption                 *EncryptionProperties                 `json:"encryption,omitempty"`
	Endpoint                   *string                               `json:"endpoint,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnectionReference `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState                    `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess                  `json:"publicNetworkAccess,omitempty"`
}

func (o ConfigurationStoreProperties) GetCreationDateAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o ConfigurationStoreProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}
