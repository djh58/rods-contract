package directrequestocr

import (
	"fmt"

	"github.com/ava-labs/coreth/accounts/abi"
	"github.com/pkg/errors"
)

type reportCodec struct {
	reportTypes abi.Arguments
}

func getReportTypes() (abi.Arguments, error) {
	bytes32ArrType, err := abi.NewType("bytes32[]", "", []abi.ArgumentMarshaling{})
	if err != nil {
		return nil, fmt.Errorf("unable to create an ABI type object for bytes32[]")
	}
	bytesArrType, err := abi.NewType("bytes[]", "", []abi.ArgumentMarshaling{})
	if err != nil {
		return nil, fmt.Errorf("unable to create an ABI type object for bytes[]")
	}
	return abi.Arguments([]abi.Argument{
		{Name: "ids", Type: bytes32ArrType},
		{Name: "results", Type: bytesArrType},
		{Name: "errors", Type: bytesArrType},
	}), nil
}

func NewReportCodec() (*reportCodec, error) {
	reportTypes, err := getReportTypes()
	if err != nil {
		return nil, err
	}
	return &reportCodec{
		reportTypes: reportTypes,
	}, nil
}

func sliceToByte32(slice []byte) [32]byte {
	var res [32]byte
	copy(res[:], slice[:32])
	return res
}

func (c *reportCodec) EncodeReport(requests []*ProcessedRequest) ([]byte, error) {
	size := len(requests)
	if size == 0 {
		return []byte{}, nil
	}
	ids := make([][32]byte, size)
	results := make([][]byte, size)
	errors := make([][]byte, size)
	for i := 0; i < size; i++ {
		ids[i] = sliceToByte32(requests[i].RequestID)
		results[i] = requests[i].Result
		errors[i] = requests[i].Error
	}
	return c.reportTypes.Pack(ids, results, errors)
}

func (c *reportCodec) DecodeReport(raw []byte) ([]*ProcessedRequest, error) {
	reportElems := map[string]interface{}{}
	if err := c.reportTypes.UnpackIntoMap(reportElems, raw); err != nil {
		return nil, errors.WithMessage(err, "unable to unpack elements from raw report")
	}

	idsIface, idsOK := reportElems["ids"]
	resultsIface, resultsOK := reportElems["results"]
	errorsIface, errorsOK := reportElems["errors"]
	if !idsOK || !resultsOK || !errorsOK {
		return nil, fmt.Errorf("missing arrays in raw report, ids: %v, results: %v, errors: %v", idsOK, resultsOK, errorsOK)
	}

	ids, idsOK := idsIface.([][32]byte)
	results, resultsOK := resultsIface.([][]byte)
	errors, errorsOK := errorsIface.([][]byte)
	if !idsOK || !resultsOK || !errorsOK {
		return nil, fmt.Errorf("unable to cast part of raw report into array type, ids: %v, results: %v, errors: %v", idsOK, resultsOK, errorsOK)
	}

	size := len(ids)
	if len(results) != size || len(errors) != size {
		return nil, fmt.Errorf("unequal sizes of arrays parsed from raw report, ids: %v, results: %v, errors: %v", len(ids), len(results), len(errors))
	}
	if size == 0 {
		return []*ProcessedRequest{}, nil
	}

	decoded := make([]*ProcessedRequest, size)
	for i := 0; i < size; i++ {
		decoded[i] = &ProcessedRequest{
			RequestID: ids[i][:],
			Result:    results[i],
			Error:     errors[i],
		}
	}
	return decoded, nil
}
