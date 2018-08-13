package soteapi

import (
	"encoding/xml"
)

// GetCategoryList returns categories list
func (api *SoteAPI) GetCategoryList(sessionHash string, offset, limit int) ([]CategoryWithTimestamp, error) {
	req := getCategoryListRequest{SessionHash: sessionHash, Offset: offset, Limit: limit}
	resp := getCategoryListResponse{}
	err := serviceCategory.Call(stCategory, req, &resp, nil)
	return resp.Items, err
}

type getCategoryListRequest struct {
	XMLName     xml.Name `xml:"GetCategoryList"`
	SessionHash string   `xml:"GetCategoryList>_session_hash"`
	Offset      int      `xml:"GetCategoryList>_offset"`
	Limit       int      `xml:"GetCategoryList>_limit"`
}

type getCategoryListResponse struct {
	XMLName xml.Name                `xml:"GetCategoryListResponse"`
	Items   []CategoryWithTimestamp `xml:"GetCategoryListResponse>item"`
}
