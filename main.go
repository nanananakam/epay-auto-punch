package main

import (
	"fmt"
	"github.com/sclevine/agouti"
	"os"
	"time"
)

func main() {
	driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless"}), agouti.Debug)

	err := driver.Start()
	if err != nil {
		driver.Stop()
		panic(err)
	}

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Navigate("https://prb01.payroll.co.jp/epayc/")
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Find("input[name=\"copCd\"]").Fill(os.Getenv("ePayWorkCopCd"))
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Find("input[name=\"empCd\"]").Fill(os.Getenv("ePayWorkEmpCd"))
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Find("input[name=\"password\"]").Fill(os.Getenv("ePayWorkPassword"))
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Find("button[type=\"submit\"]").Click()
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Navigate("https://prb01.payroll.co.jp/epayc/mainPersonal.do?op=doSso&fwdSyscd=work&concd=calendar")
	if err != nil {
		driver.Stop()
		panic(err)
	}

	now := time.Now()

	fromValue, err := page.All(".work-control.work-calendar").At(now.Day() - 1).All("input[type=\"text\"]").At(0).Attribute("value")
	if err != nil {
		driver.Stop()
		panic(err)
	}

	if fromValue == "" {
		err = page.All(".work-control.work-calendar").At(now.Day() - 1).All("input[type=\"text\"]").At(0).Fill(fmt.Sprintf("%02d%02d", now.Hour(), now.Minute()))
		if err != nil {
			driver.Stop()
			panic(err)
		}
		fmt.Println("出勤打刻")
	} else {
		err = page.All(".work-control.work-calendar").At(now.Day() - 1).All("input[type=\"text\"]").At(1).Fill(fmt.Sprintf("%02d%02d", now.Hour(), now.Minute()))
		if err != nil {
			driver.Stop()
			panic(err)
		}
		fmt.Println("退勤打刻")
	}

	err = page.Find("#contentsRight > div.wrapperCenter.mt10 > a.buttonLBright.lastchild").Click()
	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = page.Find("#navigation > ul > li.lastchild").Click()

	if err != nil {
		driver.Stop()
		panic(err)
	}

	err = driver.Stop()
	if err != nil {
		panic(err)
	}
	fmt.Println("正常完了")

}
