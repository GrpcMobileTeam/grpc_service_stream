package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/protobuf/proto"
	"bytes"
	"encoding/json"
	"gin-proto-test/example"
	. "github.com/smartystreets/goconvey/convey"
)

// protoc build code =>
// protoc HouseContractListInput.proto HouseContractListOutput.proto --go_out=../

func TestProtoMarshal(t *testing.T) {
	Convey("Proto请求测试", t, func() {
		router := setupRouter()
		w := httptest.NewRecorder()
		body := &example.HouseContractListInput{KeyWordName: "搜索"}
		inData, _ := proto.Marshal(body)

		req, _ := http.NewRequest("POST", "/test", bytes.NewReader(inData))
		req.Header.Set("Content-Type", "application/x-protobuf")
		router.ServeHTTP(w, req)

		var output example.HouseContractListOutput
		_ = proto.Unmarshal(w.Body.Bytes(), &output)
		So(int(output.TotalRow), ShouldEqual, 2)
	})
}

func TestJsonMarshal(t *testing.T) {
	Convey("Json请求测试", t, func() {
		router := setupRouter()
		w := httptest.NewRecorder()
		body := &example.HouseContractListInput{KeyWordName: "搜索"}
		inData, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/test", bytes.NewReader(inData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		var output example.HouseContractListOutput
		_ = json.Unmarshal(w.Body.Bytes(), &output)

		So(2, ShouldEqual, int(output.TotalRow))
	})
}
