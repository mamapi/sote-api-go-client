package soteapi

import (
	"encoding/xml"
)

// AddProducer adds a new producer
func (api *SoteAPI) AddProducer(sessionHash string, producer Producer) (int, error) {
	req := addProducerRequest{
		Inner: addProducerRequestInner{
			SessionHash: sessionHash,
			Producer:    producer,
		},
	}
	resp := addProducerResponse{}
	err := serviceProducer.Call(stProducer, req, &resp, nil)
	return resp.ID, err
}

type addProducerRequest struct {
	XMLName xml.Name                `xml:"AddProducer"`
	Inner   addProducerRequestInner `xml:"AddProducer"`
}

type addProducerRequestInner struct {
	SessionHash string `xml:"_session_hash"`
	Producer
}

type addProducerResponse struct {
	XMLName xml.Name `xml:"AddProducerResponse"`
	ID      int      `xml:"AddProducerResponse>id"`
}
