package models

const SIGNATURE = "3DS¿SADA_SDA^4S4S44CSVSVS8FVS^SÇ"
const PKEY = "2s8d5_¡43ç/5&<>dç=¡werdevent"
const SIGN_TYPE = "SECURE_WERDEVENT_CREDENTIAL"

type Signature struct {
	Sign    string `json:"sign"`
	Type    string `json:"type"`
	Created string `json:"created"`
	UI      string `json:"ui"`
	BI      string `json:"bi"`
	CI      string `json:"CI"`
}
