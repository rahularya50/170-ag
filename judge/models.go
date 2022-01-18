package judge

import (
	"170-ag/proto/schemas"
	"os"

	"google.golang.org/protobuf/proto"
)

func decodeJudgingProto(data []byte) (*schemas.JudgingRequest, error) {
	state := &schemas.JudgingRequest{}
	if err := proto.Unmarshal(data, state); err != nil {
		return nil, err
	}
	return state, nil
}

func storeJudgingProto(judgingRequest *schemas.JudgingRequest) error {
	data, err := proto.Marshal(judgingRequest)
	if err != nil {
		return err
	}
	os.WriteFile("/workdir/judging_proto", data, 0644)
	return nil
}

func loadJudgingProto() (*schemas.JudgingRequest, error) {
	data, err := os.ReadFile("/workdir/judging_proto")
	if err != nil {
		return nil, err
	}
	judgingProto, err := decodeJudgingProto(data)
	if err != nil {
		return nil, err
	}
	return judgingProto, nil
}

func decodeGradingProto(data []byte) (*schemas.GradingResponse, error) {
	state := &schemas.GradingResponse{}
	if err := proto.Unmarshal(data, state); err != nil {
		return nil, err
	}
	return state, nil
}

func storeGradingProto(gradingResponse *schemas.GradingResponse) error {
	data, err := proto.Marshal(gradingResponse)
	if err != nil {
		return err
	}
	os.WriteFile("/workdir/grading_proto", data, 0644)
	return nil
}

func loadGradingProto() (*schemas.GradingResponse, error) {
	data, err := os.ReadFile("/workdir/grading_proto")
	if err != nil {
		return nil, err
	}
	gradingProto, err := decodeGradingProto(data)
	if err != nil {
		return nil, err
	}
	return gradingProto, nil
}
