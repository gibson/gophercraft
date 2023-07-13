package realm

import (
	"time"

	"github.com/gibson/gophercraft/packet/synctime"
)

func (s *Session) HandleQueryTime(q *synctime.QueryTime) {
	s.Send(&synctime.QueryTimeResponse{
		Unix: int32(time.Now().Unix()),
	})
}
