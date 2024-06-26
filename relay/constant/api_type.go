package constant

import (
	"one-api/common"
)

const (
	APITypeOpenAI = iota
	APITypeAnthropic
	APITypePaLM
	APITypeBaidu
	APITypeZhipu
	APITypeAli
	APITypeXunfei
	APITypeAIProxyLibrary
	APITypeTencent
	APITypeGemini
	APITypeZhipu_v4
	APITypeOllama
	APITypePerplexity
	APITypeAws
	APITypeCohere

	APITypeDummy // this one is only for count, do not add any channel after this
)

func ChannelType2APIType(channelType int) int {
	apiType := -1
	switch channelType {
	case common.ChannelTypeOpenAI:
		apiType = APITypeOpenAI
	case common.ChannelTypeAzure:
		apiType = APITypeOpenAI
	case common.ChannelTypeMoonshot:
		apiType = APITypeOpenAI
	case common.ChannelTypeLingYiWanWu:
		apiType = APITypeOpenAI
	case common.ChannelType360:
		apiType = APITypeOpenAI
	case common.ChannelTypeAnthropic:
		apiType = APITypeAnthropic
	case common.ChannelTypeBaidu:
		apiType = APITypeBaidu
	case common.ChannelTypePaLM:
		apiType = APITypePaLM
	case common.ChannelTypeZhipu:
		apiType = APITypeZhipu
	case common.ChannelTypeAli:
		apiType = APITypeAli
	case common.ChannelTypeXunfei:
		apiType = APITypeXunfei
	case common.ChannelTypeAIProxyLibrary:
		apiType = APITypeAIProxyLibrary
	case common.ChannelTypeTencent:
		apiType = APITypeTencent
	case common.ChannelTypeGemini:
		apiType = APITypeGemini
	case common.ChannelTypeZhipu_v4:
		apiType = APITypeZhipu_v4
	case common.ChannelTypeOllama:
		apiType = APITypeOllama
	case common.ChannelTypePerplexity:
		apiType = APITypePerplexity
	case common.ChannelTypeAws:
		apiType = APITypeAws
	case common.ChannelTypeCohere:
		apiType = APITypeCohere
	}
	return apiType
}
