package mock

import (
	"github.com/harrisonzhao/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
