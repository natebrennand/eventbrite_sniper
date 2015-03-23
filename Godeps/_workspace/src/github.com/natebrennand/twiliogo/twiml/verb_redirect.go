package twiml

type redirect struct {
	XMLName int `xml:"Redirect"`
	*RedirectOpts
	Target *string `xml:",chardata"`
}

// RedirectOpts configures a redirect block
type RedirectOpts struct {
	Method string `xml:"method,attr,omitempty"`
}

func addRedirect(t twimlResponse, opts *RedirectOpts, target *string) {
	t.appendContents(&redirect{RedirectOpts: opts, Target: target})
}
