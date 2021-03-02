package domain

import (
	"bitbucket.org/latonaio/aion-core/pkg/go-client/msclient"
	"bitbucket.org/latonaio/aion-core/pkg/log"
)

type KanbanWriter struct {
	ConnectionKey string      `json:"connection_key"`
	Metadata      []*Metadata `json:"metadata"`
	DeviceName    string
}

type Metadata struct {
	PropertyName string      `json:"property_name"`
	Body         interface{} `json:"body"`
}

func (kw *KanbanWriter) BuildOutputData() []msclient.Option{
	var options []msclient.Option
	if kw.ConnectionKey != "" {
		options = append(options,msclient.SetConnectionKey(kw.ConnectionKey))
	}

	if kw.Metadata == nil || len(kw.Metadata) == 0 {
		log.Print("metadata is empty")
		return nil
	}
	data := map[string]interface{}{}
	for _, v := range kw.Metadata {
		data[v.PropertyName] = v.Body
	}
	options = append(options,msclient.SetMetadata(data))

	if kw.DeviceName != "" {
		options = append(options,msclient.SetDeviceName(kw.DeviceName))
	}

	return options
}
