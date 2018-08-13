package soteapi

import (
	"encoding/xml"
)

// GetCategory returns category
func (api *SoteAPI) GetCategory(sessionHash string, id int) (*CategoryWithTimestamp, error) {
	req := getCategoryRequest{SessionHash: sessionHash, ID: id}
	resp := getCategoryResponse{}
	err := serviceCategory.Call(stCategory, req, &resp, nil)
	return &resp.Data, err
}

type getCategoryRequest struct {
	XMLName     xml.Name `xml:"GetCategory"`
	SessionHash string   `xml:"GetCategory>_session_hash"`
	ID          int      `xml:"GetCategory>id"`
}

type getCategoryResponse struct {
	XMLName xml.Name              `xml:"GetCategoryResponse"`
	Data    CategoryWithTimestamp `xml:"GetCategoryResponse"`
}
