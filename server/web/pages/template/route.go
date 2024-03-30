package template

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteWebPages(route gin.IRouter) {
	route.GET("/", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Indexhtml))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/html; charset=utf-8", Indexhtml)
	})

	route.GET("/apple-splash-1125-2436.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash11252436jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash11252436jpg)
	})

	route.GET("/apple-splash-1136-640.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash1136640jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash1136640jpg)
	})

	route.GET("/apple-splash-1170-2532.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash11702532jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash11702532jpg)
	})

	route.GET("/apple-splash-1242-2208.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash12422208jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash12422208jpg)
	})

	route.GET("/apple-splash-1242-2688.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash12422688jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash12422688jpg)
	})

	route.GET("/apple-splash-1284-2778.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash12842778jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash12842778jpg)
	})

	route.GET("/apple-splash-1334-750.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash1334750jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash1334750jpg)
	})

	route.GET("/apple-splash-1536-2048.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash15362048jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash15362048jpg)
	})

	route.GET("/apple-splash-1620-2160.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash16202160jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash16202160jpg)
	})

	route.GET("/apple-splash-1668-2224.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash16682224jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash16682224jpg)
	})

	route.GET("/apple-splash-1668-2388.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash16682388jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash16682388jpg)
	})

	route.GET("/apple-splash-1792-828.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash1792828jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash1792828jpg)
	})

	route.GET("/apple-splash-2048-1536.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash20481536jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash20481536jpg)
	})

	route.GET("/apple-splash-2048-2732.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash20482732jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash20482732jpg)
	})

	route.GET("/apple-splash-2160-1620.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash21601620jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash21601620jpg)
	})

	route.GET("/apple-splash-2208-1242.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash22081242jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash22081242jpg)
	})

	route.GET("/apple-splash-2224-1668.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash22241668jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash22241668jpg)
	})

	route.GET("/apple-splash-2388-1668.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash23881668jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash23881668jpg)
	})

	route.GET("/apple-splash-2436-1125.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash24361125jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash24361125jpg)
	})

	route.GET("/apple-splash-2532-1170.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash25321170jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash25321170jpg)
	})

	route.GET("/apple-splash-2688-1242.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash26881242jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash26881242jpg)
	})

	route.GET("/apple-splash-2732-2048.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash27322048jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash27322048jpg)
	})

	route.GET("/apple-splash-2778-1284.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash27781284jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash27781284jpg)
	})

	route.GET("/apple-splash-640-1136.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash6401136jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash6401136jpg)
	})

	route.GET("/apple-splash-750-1334.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash7501334jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash7501334jpg)
	})

	route.GET("/apple-splash-828-1792.jpg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesapplesplash8281792jpg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/jpeg", Serverwebpagestemplatepagesapplesplash8281792jpg)
	})

	route.GET("/asset-manifest.json", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesassetmanifestjson))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/json", Serverwebpagestemplatepagesassetmanifestjson)
	})

	route.GET("/browserconfig.xml", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesbrowserconfigxml))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/xml; charset=utf-8", Serverwebpagestemplatepagesbrowserconfigxml)
	})

	route.GET("/dlnaicon-120.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesdlnaicon120png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepagesdlnaicon120png)
	})

	route.GET("/dlnaicon-48.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesdlnaicon48png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepagesdlnaicon48png)
	})

	route.GET("/favicon-16x16.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesfavicon16x16png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepagesfavicon16x16png)
	})

	route.GET("/favicon-32x32.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesfavicon32x32png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepagesfavicon32x32png)
	})

	route.GET("/favicon.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesfaviconico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Serverwebpagestemplatepagesfaviconico)
	})

	route.GET("/icon.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesiconpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepagesiconpng)
	})

	route.GET("/index.html", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesindexhtml))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/html; charset=utf-8", Serverwebpagestemplatepagesindexhtml)
	})

	route.GET("/logo.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepageslogopng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepageslogopng)
	})

	route.GET("/mstile-150x150.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesmstile150x150png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Serverwebpagestemplatepagesmstile150x150png)
	})

	route.GET("/site.webmanifest", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagessitewebmanifest))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/manifest+json", Serverwebpagestemplatepagessitewebmanifest)
	})

	route.GET("/static/js/2.64714af8.chunk.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjs264714af8chunkjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Serverwebpagestemplatepagesstaticjs264714af8chunkjs)
	})

	route.GET("/static/js/2.64714af8.chunk.js.LICENSE.txt", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjs264714af8chunkjsLICENSEtxt))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/plain; charset=utf-8", Serverwebpagestemplatepagesstaticjs264714af8chunkjsLICENSEtxt)
	})

	route.GET("/static/js/2.64714af8.chunk.js.map", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjs264714af8chunkjsmap))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/json", Serverwebpagestemplatepagesstaticjs264714af8chunkjsmap)
	})

	route.GET("/static/js/main.bb243e45.chunk.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjsmainbb243e45chunkjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Serverwebpagestemplatepagesstaticjsmainbb243e45chunkjs)
	})

	route.GET("/static/js/main.bb243e45.chunk.js.map", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjsmainbb243e45chunkjsmap))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/json", Serverwebpagestemplatepagesstaticjsmainbb243e45chunkjsmap)
	})

	route.GET("/static/js/runtime-main.f542387e.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjsruntimemainf542387ejs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Serverwebpagestemplatepagesstaticjsruntimemainf542387ejs)
	})

	route.GET("/static/js/runtime-main.f542387e.js.map", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Serverwebpagestemplatepagesstaticjsruntimemainf542387ejsmap))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/json", Serverwebpagestemplatepagesstaticjsruntimemainf542387ejsmap)
	})
}
