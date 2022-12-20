package main

import (
	"encoding/json"
	"io"
	"sync"
)

type request struct {
	Method string           `json:"method"`
	Body *json.RawMessage 	`json:"body"`
}

type TobCodecCorrect struct {
	closer io.Closer
	closed bool
	writer io.Writer
	codec  sync.Map
	req	   request
}

type TobCodecIncorrect struct {
	closer io.Closer
	closed bool
	writer io.Writer
	codec  sync.Map
	req	   request
}

// ok: servercodec-readrequestbody-unhandled-nil
func (cc *TobCodecCorrect) ReadRequestBody(body interface{}) error {
	if body == nil {
		return nil
	}

	var result [1]interface{}
	result[0] = body
	return json.Unmarshal(*cc.req.Body, &result)
}

// ok: servercodec-readrequestbody-unhandled-nil
func (cc *TobCodecCorrect) ReadRequestBody(body interface{}) error {
	var result [1]interface{}
	if body != nil {
		result[0] = body
		return json.Unmarshal(*cc.req.Body, &result)
	}
}

// ruleid: servercodec-readrequestbody-unhandled-nil
func (ci *TobCodecIncorrect) ReadRequestBody(body interface{}) error {
	var result [1]interface{}
	result[0] = body
	return json.Unmarshal(*ci.req.Body, &result)
}



