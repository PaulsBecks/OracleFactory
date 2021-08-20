package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func TransformParameterType(parameter interface{}, parameterType string) (interface{}, error) {
	switch parameterType {
	case "uint256":
		f, ok := parameter.(float64)
		if ok {
			return big.NewInt(int64(f)), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "uint64":
		f, ok := parameter.(float64)
		if ok {
			return uint64(f), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "uint8":
		f, ok := parameter.(float64)
		if ok {
			return uint8(f), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "string":
		s, ok := parameter.(string)
		if ok {
			return s, nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "bytes":
		s, ok := parameter.(string)
		if ok {
			return []byte(s), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "bytes32":
		s, ok := parameter.(string)
		if ok {
			var arr [32]byte
			copy(arr[:], []byte(s))
			return arr, nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "bool":
		b, ok := parameter.(bool)
		if ok {
			return b, nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	case "address":
		s, ok := parameter.(string)
		if ok {
			return common.HexToAddress(s), nil
		}
		return nil, fmt.Errorf("Unable to parse parameter. %v", parameter)
	}
	return nil, fmt.Errorf("Unable to transform paramter of type: %T", parameter)
}
