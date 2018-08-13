package soteapi

import (
	"encoding/xml"
)

// CountCategory returns categories count
func (api *SoteAPI) CountCategory(sessionHash string) (int, error) {
	req := countCategoryRequest{SessionHash: sessionHash}
	resp := countCategoryResponse{}
	err := serviceCategory.Call(stCategory, req, &resp, nil)
	return resp.Count, err
}

type countCategoryRequest struct {
	XMLName     xml.Name `xml:"CountCategory"`
	SessionHash string   `xml:"CountCategory>_session_hash"`
}

type countCategoryResponse struct {
	XMLName xml.Name `xml:"CountCategoryResponse"`
	Count   int      `xml:"CountCategoryResponse>_count"`
}
