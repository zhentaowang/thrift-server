// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package server

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type RESCODE int64

const (
	RESCODE__200 RESCODE = 200
	RESCODE__400 RESCODE = 400
	RESCODE__403 RESCODE = 403
	RESCODE__404 RESCODE = 404
	RESCODE__500 RESCODE = 500
	RESCODE__503 RESCODE = 503
)

func (p RESCODE) Int() int {
	switch p {
	case RESCODE__200:
		return 200
	case RESCODE__400:
		return 400
	case RESCODE__403:
		return 403
	case RESCODE__404:
		return 404
	case RESCODE__500:
		return 500
	case RESCODE__503:
		return 503
	default:
		return 200
	}
}

func (p RESCODE) String() string {
	switch p {
	case RESCODE__200:
		return "_200"
	case RESCODE__400:
		return "_400"
	case RESCODE__403:
		return "_403"
	case RESCODE__404:
		return "_404"
	case RESCODE__500:
		return "_500"
	case RESCODE__503:
		return "_503"
	}
	return "<UNSET>"
}

func RESCODEFromString(s string) (RESCODE, error) {
	switch s {
	case "_200":
		return RESCODE__200, nil
	case "400":
		return RESCODE__400, nil
	case "_404":
		return RESCODE__404, nil
	case "_500":
		return RESCODE__500, nil
	case "_503":
		return RESCODE__503, nil
	}
	return RESCODE(0), fmt.Errorf("not a valid RESCODE string")
}

func RESCODEPtr(v RESCODE) *RESCODE { return &v }

func (p RESCODE) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *RESCODE) UnmarshalText(text []byte) error {
	q, err := RESCODEFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type EXCCODE int64

const (
	EXCCODE_PARAMNOTFOUND   EXCCODE = 2001
	EXCCODE_SERVICENOTFOUND EXCCODE = 2002
)

func (p EXCCODE) String() string {
	switch p {
	case EXCCODE_PARAMNOTFOUND:
		return "PARAMNOTFOUND"
	case EXCCODE_SERVICENOTFOUND:
		return "SERVICENOTFOUND"
	}
	return "<UNSET>"
}

func EXCCODEFromString(s string) (EXCCODE, error) {
	switch s {
	case "PARAMNOTFOUND":
		return EXCCODE_PARAMNOTFOUND, nil
	case "SERVICENOTFOUND":
		return EXCCODE_SERVICENOTFOUND, nil
	}
	return EXCCODE(0), fmt.Errorf("not a valid EXCCODE string")
}

func EXCCODEPtr(v EXCCODE) *EXCCODE { return &v }

func (p EXCCODE) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *EXCCODE) UnmarshalText(text []byte) error {
	q, err := EXCCODEFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - ParamJSON
//  - ServiceName
type Request struct {
	ParamJSON   []byte `interface:"paramJSON,1,required" json:"paramJSON"`
	ServiceName string `interface:"serviceName,2,required" json:"serviceName"`
}

func NewRequest() *Request {
	return &Request{}
}

func (p *Request) GetParamJSON() []byte {
	return p.ParamJSON
}

func (p *Request) GetServiceName() string {
	return p.ServiceName
}
func (p *Request) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetParamJSON bool = false
	var issetServiceName bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetParamJSON = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetServiceName = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetParamJSON {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ParamJSON is not set"))
	}
	if !issetServiceName {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ServiceName is not set"))
	}
	return nil
}

func (p *Request) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ParamJSON = v
	}
	return nil
}

func (p *Request) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ServiceName = v
	}
	return nil
}

func (p *Request) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Request"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("paramJSON", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:paramJSON: ", p), err)
	}
	if err := oprot.WriteBinary(p.ParamJSON); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.paramJSON (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:paramJSON: ", p), err)
	}
	return err
}

func (p *Request) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("serviceName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:serviceName: ", p), err)
	}
	if err := oprot.WriteString(string(p.ServiceName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.serviceName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:serviceName: ", p), err)
	}
	return err
}

func (p *Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Request(%+v)", *p)
}

// Attributes:
//  - ResponeCode
//  - ResponseJSON
type Response struct {
	ResponeCode  RESCODE `interface:"responeCode,1,required" json:"responeCode"`
	ResponseJSON []byte  `interface:"responseJSON,2,required" json:"responseJSON"`
}

func NewResponse() *Response {
	return &Response{}
}

func (p *Response) GetResponeCode() RESCODE {
	return p.ResponeCode
}

func (p *Response) GetResponseJSON() []byte {
	return p.ResponseJSON
}
func (p *Response) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetResponeCode bool = false
	var issetResponseJSON bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetResponeCode = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetResponseJSON = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetResponeCode {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ResponeCode is not set"))
	}
	if !issetResponseJSON {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ResponseJSON is not set"))
	}
	return nil
}

func (p *Response) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := RESCODE(v)
		p.ResponeCode = temp
	}
	return nil
}

func (p *Response) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ResponseJSON = v
	}
	return nil
}

func (p *Response) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Response"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Response) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("responeCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:responeCode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ResponeCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.responeCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:responeCode: ", p), err)
	}
	return err
}

func (p *Response) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("responseJSON", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:responseJSON: ", p), err)
	}
	if err := oprot.WriteBinary(p.ResponseJSON); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.responseJSON (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:responseJSON: ", p), err)
	}
	return err
}

func (p *Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Response(%+v)", *p)
}

// Attributes:
//  - ExceptionCode
//  - ExceptionMess
type ServiceException struct {
	ExceptionCode EXCCODE `interface:"exceptionCode,1,required" json:"exceptionCode"`
	ExceptionMess string  `interface:"exceptionMess,2,required" json:"exceptionMess"`
}

func NewServiceException() *ServiceException {
	return &ServiceException{}
}

func (p *ServiceException) GetExceptionCode() EXCCODE {
	return p.ExceptionCode
}

func (p *ServiceException) GetExceptionMess() string {
	return p.ExceptionMess
}
func (p *ServiceException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetExceptionCode bool = false
	var issetExceptionMess bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetExceptionCode = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetExceptionMess = true
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetExceptionCode {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ExceptionCode is not set"))
	}
	if !issetExceptionMess {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ExceptionMess is not set"))
	}
	return nil
}

func (p *ServiceException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EXCCODE(v)
		p.ExceptionCode = temp
	}
	return nil
}

func (p *ServiceException) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ExceptionMess = v
	}
	return nil
}

func (p *ServiceException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ServiceException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ServiceException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("exceptionCode", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:exceptionCode: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.ExceptionCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.exceptionCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:exceptionCode: ", p), err)
	}
	return err
}

func (p *ServiceException) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("exceptionMess", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:exceptionMess: ", p), err)
	}
	if err := oprot.WriteString(string(p.ExceptionMess)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.exceptionMess (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:exceptionMess: ", p), err)
	}
	return err
}

func (p *ServiceException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ServiceException(%+v)", *p)
}

func (p *ServiceException) Error() string {
	return p.String()
}