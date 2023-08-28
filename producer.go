package main

import (
	"creative-e2e-test/proto/ads"
	"github.freewheel.tv/bricks/message.v2"
)

const (
	kafkaBrokerAddress = "center.global.kafka.us-east-1.stg.data.aws.fwmrm.net:9092"
)

func ProduceMarketAdMsg(num int) error {
	mamsg := &ads.MarketAdIngestion{
		ExternalId: "5:20414868447:508828045-j:508828045",
		Creative: &ads.MarketAdIngestion_Creative{
			Name:     "vpaid.2023.05.19-19.29-6622210",
			Duration: 15,
			UniversalAdId: &ads.MarketAdIngestion_Creative_UniversalAdId{
				IdRegistry: "GDCM",
				IdValue:    "196151874-1",
			},
			AdParameter: "{\"mediaFiles\":[{\"id\":\"740961540-160\",\"delivery\":\"progressive\",\"type\":\"application/x-mpegURL\",\"bitrate\":99,\"width\":256,\"height{\"id\":\"skip\",\"config\":{\"acceptMedia\":\"^video/\"}}]}",
			Renditions: []*ads.MarketAdIngestion_Creative_Rendition{
				{
					Source:        "https://static.adsafeprotected.com/ias/v1/vpaid.2023.05.19-19.29-6622210.js",
					ContentTypeId: 23,
					CreativeApiId: 2,
					Bitrate:       0,
					Width:         1920,
					Height:        1080,
				},
				{
					Source:        "https://gcdn.2mdn.net/api/manifest/index.m3u8",
					ContentTypeId: 68,
					Bitrate:       99,
					Width:         256,
					Height:        144,
					NeedDownload:  true,
				},
				{
					Source:        "https://gcdn.2mdn.net/videoplayback/id/f5e6f71598ba269b/itag/18/source/web/file.mp4",
					ContentTypeId: 45,
					Bitrate:       301,
					Width:         640,
					Height:        360,
					NeedDownload:  true,
				},
			},
		},
		Clickthrough: "https://adclick.g.doubleclick.net",
		Domain:       "reachmena.com",
		CchKey:       "adsafeprotected.com/196151874",
		VastWrapper:  "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>",
		VastContent:  "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<VAST version=\"4.0\"></VAST>",
	}
	mamp, err := message.NewProducer[*ads.MarketAdIngestion]([]string{kafkaBrokerAddress})
	if err != nil {
		return err
	}
	for i := 0; i < num; i++ {
		err = mamp.SendSyncMessage("fw_ads_market_ad_ingestion", nil, mamsg)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProduceCCHMappingMsg(num int) error {
	ccmmsg := &ads.CCHCreativeMapping{
		AdId:          "123",
		MrmCreativeId: 124,
	}
	ccmp, err := message.NewProducer[*ads.CCHCreativeMapping]([]string{kafkaBrokerAddress})
	if err != nil {
		return err
	}
	for i := 0; i < num; i++ {
		err = ccmp.SendSyncMessage("fw_ads_cch_creative_mapping", nil, ccmmsg)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProduceJittMsf(num int) error {
	jittmsg := &ads.JustInTimeTranscoding{
		AdId: "cb8379946a5b98e434b55282cb3bacf718250564521985408848",
		MrmCreative: &ads.JustInTimeTranscoding_MRMCreative{
			Id:                 10310405,
			IsInternal:         false,
			CreateMapping:      false,
			CreativeCategory:   "MARKETS_MODULE",
			InboundMkplOrderId: 65595,
			SourceNetworkId:    169843,
		},
		Creative: &ads.JustInTimeTranscoding_Creative{
			Name:               "MichiganCreative_PioneerGoodnightsleepV4-5607138.mp4",
			Duration:           29,
			UseClientMezzanine: false,
			Mezzanine: &ads.JustInTimeTranscoding_Creative_Rendition{
				Source:        "https://assets.springserve.com/video_creatives/000/694/507/MichiganCreative_PioneerGoodnightsleepV4-5607138.mp4",
				ContentTypeId: 45,
				Bitrate:       10000,
				Width:         1920,
				Height:        1080,
			},
		},
		TranscodePackageId:     291,
		PlayerProfileNetworkId: 169843,
		PlayerProfileId:        12665,
	}

	jittmp, err := message.NewProducer[*ads.JustInTimeTranscoding]([]string{kafkaBrokerAddress})
	if err != nil {
		return err
	}
	for i := 0; i < num; i++ {
		err = jittmp.SendSyncMessage("fw_ads_just_in_time_transcoding", nil, jittmsg)
		if err != nil {
			return err
		}
	}
	return nil
}
