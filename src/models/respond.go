package models

type Respond struct {
    ResultCode	int64  `json:"resultCode"`
    Message	string `json:"message"`
    Data	interface{}  `json:"data"`
}

type TopicData struct {
    TopicId  int64 `json:"topicId"`
    Title string `json:"title"`
}

func RespondError(error int64) Respond {
    resp :=  Respond {
        ResultCode  : error,
        Message  :  "Error Normal",
    }

    return resp
}
