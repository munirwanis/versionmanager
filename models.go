package main

// Version object
type Version struct {
	Number      string `json:"number"`
	ForceUpdate bool   `json:"forceUpdate"`
	URL         string `json:"url,omitempty"`
}

// Error object
type Error struct {
	Message string `json:"message"`
}
