package encoder

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	httpstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
	httpStd "net/http"
	"strings"
)

type httpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	baseContentType = "application"
)

func contentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

// ErrorResponse 提供统一格式的错误信息响应
func ErrorResponse(w httpStd.ResponseWriter, r *httpStd.Request, err error) {
	se := new(httpResponse)
	gs, ok := status.FromError(err)
	if !ok {
		se.Code = httpStd.StatusInternalServerError
		se.Message = err.Error()
	} else {
		se.Code = httpstatus.FromGRPCCode(gs.Code())
		se.Message = gs.Message()
	}
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(httpStd.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", codec.Name())
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}

// OKResponse 提供统一格式 200 HTTP OK 格式的响应信息
func OKResponse(w httpStd.ResponseWriter, r *httpStd.Request, reply interface{}) error {
	if reply == nil {
		return nil
	}
	if rd, ok := reply.(http.Redirector); ok {
		url, code := rd.Redirect()
		httpStd.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(&httpResponse{
		Code:    200,
		Data:    reply,
		Message: "OK",
	})
	if err != nil {
		return err
	}
	// 在Response Header中写入编码器的scheme
	w.Header().Set("Content-Type", contentType(codec.Name()))
	w.Write(data)
	return nil
}
