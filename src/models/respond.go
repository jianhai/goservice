package models

type Respond struct {
    ResultCode	int64  `json:"resultCode"`
    Message	string `json:"message"`
    Data	interface{}  `json:"data"`
}

func RespondError(error int64) Respond {
    resp :=  Respond {
        ResultCode  : error,
        Message  :  "Error Normal",
    }

    return resp
}
