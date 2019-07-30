# fake-useragent

## Install
```
go get github.com/JaeGerW2016/fake-useragent
```

## Usage
```
package main

import (
	"github.com/JaeGerW2016/fake-useragent"
	"log"
)

func main() {
	// recommend to use
	random := fake_useragent.Random()
	log.Printf("Random: %s", random)

	chrome := fake_useragent.Chrome()
	log.Printf("Chrome: %s", chrome)

	internetExplorer := fake_useragent.InternetExplorer()
	log.Printf("IE: %s", internetExplorer)

	firefox := fake_useragent.Firefox()
	log.Printf("Firefox: %s", firefox)

	safari := fake_useragent.Safari()
	log.Printf("Safari: %s", safari)

	android := fake_useragent.Android()
	log.Printf("Android: %s", android)

	macOSX := fake_useragent.MacOSX()
	log.Printf("MacOSX: %s", macOSX)

	ios := fake_useragent.IOS()
	log.Printf("IOS: %s", ios)

	linux := fake_useragent.Linux()
	log.Printf("Linux: %s", linux)

	iphone := fake_useragent.IPhone()
	log.Printf("IPhone: %s", iphone)

	ipad := fake_useragent.IPad()
	log.Printf("IPad: %s", ipad)

	computer := fake_useragent.Computer()
	log.Printf("Computer: %s", computer)

	mobile := fake_useragent.Mobile()
	log.Printf("Mobile: %s", mobile)
}
```
