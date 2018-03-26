package main

// Version object
type Version struct {
	Number      string `json:"number"`
	ForceUpdate bool   `json:"forceUpdate"`
	URL         string `json:"url,omitempty"`
}
