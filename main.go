package main

import (
	"fmt"
	"github.com/google/uuid"
	"nuonuosdk/api"
	"strings"
)

func main() {

	appKey := "90***18"  // appKey
	appSecret := "E4***40B7" // appSecret
	token := "77*****vxi"// TOKEN
	taxnum := "33990***199" // 税号
	url := "https://sdk.nuonuo.com/open/v1/services" // 请求地址
	method := "nuonuo.ElectronInvoice.queryInvoiceResult" // API方法名
	senid := strings.Replace(uuid.New().String(), "-", "", -1) //唯一标识，无需修改，保持默认即可
	// 请求参数
	body := "{\"isOfferInvoiceDetail\":\"1\",\"orderNos\":[],\"serialNos\":[\"23051609540002437585\"]}"

	result := api.SendPostSyncRequest(url, senid, appKey, appSecret, token, taxnum, method, body)

	fmt.Println(result)
}
