// generated by protoc-gen-gcraft : DO NOT EDIT
package bnet

import (
	fmt "fmt"
	protocol "github.com/gibson/gophercraft/bnet/bgs/protocol"
	v2 "github.com/gibson/gophercraft/bnet/bgs/protocol/report/v2"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// shut the compiler up
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = protocol.E_MethodOptions

const ReportServiceHash = 0x3A4218FB

type ReportService interface {
	SubmitReport(*Conn, uint32, *v2.SubmitReportRequest)
}

func DispatchReportService(conn *Conn, svc ReportService, token uint32, method uint32, data []byte) error {
	switch method {
	case 1:
		var args v2.SubmitReportRequest
		err := proto.Unmarshal(data, &args)
		if err != nil {
			return err
		}
		svc.SubmitReport(conn, token, &args)
	}
	return nil
}

type EmptyReportService struct{}

func (e EmptyReportService) SubmitReport(conn *Conn, token uint32, args *v2.SubmitReportRequest) {
	conn.SendResponseCode(token, ERROR_RPC_NOT_IMPLEMENTED)
}

func (c *Conn) ReportService_Request_SubmitReport(args *v2.SubmitReportRequest) error {
	header, _, err := c.Request(ReportServiceHash, 1, args)
	if err != nil {
		return err
	}
	if Status(header.GetStatus()) != ERROR_OK {
		return fmt.Errorf("bnet: non-ok status 0x%08X", header.GetStatus())
	}
	return nil
}
