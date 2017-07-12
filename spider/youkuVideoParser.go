package spider

import (
	"log"

	"strings"

	"github.com/benbjohnson/phantomjs"
)

//爬去优酷视频真实地址
func RunYoukuVideoParser(youkuVideoUrl string) (mp4url string) {
	//源地址 http://v.youku.com/v_show/id_XMjY5MjQ3NDQ2NA==.html?spm=a2hfu.20010077.m_210490.5~5!2~5~5!3~5~5!3~5~A&f=29102521&from=y1.3-fun-fun-904-10077.90023-210490.3-4
	//提换之后的地址 http://m.youku.com/video/id_XMjY5MjQ3NDQ2NA==.html?spm=a2hfu.20010077.m_210490.5~5!2~5~5!3~5~5!3~5~A&f=29102521&from=y1.3-fun-fun-904-10077.90023-210490.3-4
	youkuVideoUrl = strings.Replace(youkuVideoUrl, "http://v.youku.com/v_show/", "http://player.youku.com/embed/", -1)
	youkuVideoUrl = strings.Replace(youkuVideoUrl, "https://v.youku.com/v_show/", "http://m.youku.com/video/", -1)

	if err := phantomjs.DefaultProcess.Open(); err != nil {
		log.Fatal(err)
	}
	defer phantomjs.DefaultProcess.Close()
	//创建webpage
	page, err := phantomjs.CreateWebPage()
	if err != nil {
		log.Fatal(err)
	}
	defer page.Close()
	//设置webpage配置
	webPageSettings := phantomjs.WebPageSettings{
		JavascriptEnabled:             true,
		LoadImages:                    false,
		LocalToRemoteURLAccessEnabled: true, //local script can asscess the remote files
		UserAgent:                     "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Mobile Safari/537.36",
		Username:                      "dejavuzhou@qq.com",
		Password:                      "ZHou1987",
		XSSAuditingEnabled:            true,
		WebSecurityEnabled:            true,
		ResourceTimeout:               20,
	}

	if err := page.SetSettings(webPageSettings); err != nil {
		log.Fatal(err)
	}

	//http://www.dy2018.com/html/gndy/oumei/
	//http://www.dy2018.com/html/tv/oumeitv/index.html
	// if err := page.IncludeJS("http://cdn.bootcss.com/jquery/3.2.1/jquery.slim.min.js"); err != nil {
	// 	log.Fatal(err)
	// }
	if err := page.Open(youkuVideoUrl); err != nil {
		log.Fatal(err)
	}
	// 不支持 promise phantomjs
	// Read first link.
	json, err := page.Evaluate(`function() {
		document.getElementsByClassName('x-video-button')[0].click();
  		return document.getElementsByTagName('video')[0].src;
	}`)
	if err != nil {
		log.Fatal(err)
	}
	return json.(string)
	// var array = new Array();
	// var array = $('#header > div > div.bd2 > div.bd3 > div.bd3r > div.co_area2 > div.co_content8 > ul > table > tbody > tr:nth-child(2) > td:nth-child(2) > b > a:nth-child(2)').each(function(index,value){var href = $(value).attr('href');var name=$(value).attr('title');array.push({href:href,name:name})});
}