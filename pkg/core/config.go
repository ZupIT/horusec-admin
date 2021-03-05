package core

type Configuration struct {
	*General `json:",inline,omitempty"`
	*Auth    `json:",inline,omitempty"`
	*Manager `json:",inline,omitempty"`
}
