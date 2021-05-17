package main

import "fmt"

// TODO: ---- 28、多态 See: https://studygolang.com/articles/12598

/// 使用接口实现多态
type Income28 interface {
	calculate28() int
	source28()    string
}

type FixedBilling28 struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial28 struct {
	projectName string
	noOfHours  int
	hourlyRate int
}

func (fb FixedBilling28) calculate28() int {
	return fb.biddedAmount
}

func (fb FixedBilling28) source28() string {
	return fb.projectName
}

func (tm TimeAndMaterial28) calculate28() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial28) source28() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income28) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("\n Income28 From %s = $%d\n", income.source28(), income.calculate28())
		netincome += income.calculate28()
	}
	fmt.Printf("\n Net income28 of organisation = $%d", netincome)
}


/// 新增收益流
type Advertisement28 struct {
	adName     string
	CPC        int
	noOfClicks int
}

func (a Advertisement28) calculate28() int {
	return a.CPC * a.noOfClicks
}

func (a Advertisement28) source28() string {
	return a.adName
}

func day28()  {

	/// 使用接口实现多态
	project1 := FixedBilling28{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling28{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial28{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	/// incomeStreams := []Income28{project1, project2, project3}
	/// calculateNetIncome(incomeStreams)

    /// 新增收益流
	bannerAd := Advertisement28{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd  := Advertisement28{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income28{project1, project2, project3, bannerAd, popupAd}
	calculateNetIncome(incomeStreams)

}
