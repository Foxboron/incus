package shared

import (
	"time"
)

type ImageProperties map[string]string

type ImageAlias struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Target      string `json:"target"`
}

type ImageAliases []ImageAlias

type ImageInfo struct {
	Aliases      ImageAliases      `json:"aliases"`
	Architecture int               `json:"architecture"`
	Fingerprint  string            `json:"fingerprint"`
	Filename     string            `json:"filename"`
	Properties   map[string]string `json:"properties"`
	Public       bool              `json:"public"`
	Size         int64             `json:"size"`
	CreationDate time.Time         `json:"created_at"`
	ExpiryDate   time.Time         `json:"expires_at"`
	UploadDate   time.Time         `json:"uploaded_at"`
}

/*
 * BriefImageInfo contains a subset of the fields in
 * ImageInfo, namely those which a user may update
 */
type BriefImageInfo struct {
	Properties map[string]string `json:"properties"`
	Public     bool              `json:"public"`
}

func (i *ImageInfo) Brief() BriefImageInfo {
	retstate := BriefImageInfo{
		Properties: i.Properties,
		Public:     i.Public}
	return retstate
}
