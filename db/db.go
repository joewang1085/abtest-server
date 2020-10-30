package db

import (
	"fmt"
	"time"

	"abtest-server/service"
	"github.com/go-abtest/sdk"
)

// DataInit 初始化数据
func DataInit() {

	/*
		Home
	*/
	homeLayer1ZoneA := &sdk.Zone{
		ProjectID: "Home",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A",

		Weight: sdk.Weight{
			Min: 1,
			Max: 40,
		},

		Value:      "A",
		Tag:        "原主页",
		UserGroups: []string{},
	}
	homeLayer1ZoneB := &sdk.Zone{
		ProjectID: "Home",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "B",

		Weight: sdk.Weight{
			Min: 41,
			Max: 70,
		},

		Value:      "B",
		Tag:        "新主页",
		UserGroups: []string{},
	}
	homeLayer1ZoneC := &sdk.Zone{
		ProjectID: "Home",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "C",

		Weight: sdk.Weight{
			Min: 71,
			Max: 85,
		},

		Value:      "C",
		Tag:        "原主页",
		UserGroups: []string{},
	}
	homeLayer1ZoneD := &sdk.Zone{
		ProjectID: "Home",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "D",

		Weight: sdk.Weight{
			Min: 86,
			Max: 100,
		},

		Value:      "D",
		Tag:        "原主页",
		UserGroups: []string{},
	}

	/*
		Color
	*/
	colorLayer1ZoneA := &sdk.Zone{
		ProjectID: "Color",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A",

		Weight: sdk.Weight{
			Min: 1,
			Max: 25,
		},

		Value:      "A",
		Tag:        "字体 黑色",
		UserGroups: []string{},
	}
	colorLayer1ZoneB := &sdk.Zone{
		ProjectID: "Color",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "B",

		Weight: sdk.Weight{
			Min: 26,
			Max: 75,
		},

		Value:      "B",
		Tag:        "字体 红色",
		UserGroups: []string{},
	}
	colorLayer1ZoneC := &sdk.Zone{
		ProjectID: "Color",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "C",

		Weight: sdk.Weight{
			Min: 76,
			Max: 100,
		},

		Value:      "C",
		Tag:        "字体 白色",
		UserGroups: []string{},
	}
	colorLayer2ZoneD := &sdk.Zone{
		ProjectID: "Color",
		Layer: sdk.Layer{
			LayerID: "Layer2",
			ParentZones: []*sdk.Zone{
				colorLayer1ZoneA,
				colorLayer1ZoneB,
				colorLayer1ZoneC,
			},
			TotalWeight: 100,
		},
		ZoneID: "D",

		Weight: sdk.Weight{
			Min: 1,
			Max: 50,
		},

		Value:      "D",
		Tag:        "背景 白色",
		UserGroups: []string{},
	}
	colorLayer2ZoneE := &sdk.Zone{
		ProjectID: "Color",
		Layer: sdk.Layer{
			LayerID: "Layer2",
			ParentZones: []*sdk.Zone{
				colorLayer1ZoneA,
				colorLayer1ZoneB,
				colorLayer1ZoneC,
			},
			TotalWeight: 100,
		},
		ZoneID: "E",

		Weight: sdk.Weight{
			Min: 51,
			Max: 100,
		},

		Value:      "E",
		Tag:        "背景 黑色",
		UserGroups: []string{},
	}

	/*
		ComplexColor
	*/
	complexColorLayer1ZoneA := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A",

		Weight: sdk.Weight{
			Min: 1,
			Max: 25,
		},

		Value:      "A",
		Tag:        "字体 黑色",
		UserGroups: []string{},
	}
	complexColorLayer1ZoneB := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "B",

		Weight: sdk.Weight{
			Min: 26,
			Max: 75,
		},

		Value:      "B",
		Tag:        "字体 红色",
		UserGroups: []string{},
	}
	complexColorLayer1ZoneC := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID:     "Layer1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "C",

		Weight: sdk.Weight{
			Min: 76,
			Max: 100,
		},

		Value:      "C",
		Tag:        "字体 白色",
		UserGroups: []string{},
	}
	complexColorLayer2_1ZoneD := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID: "Layer2-1",
			ParentZones: []*sdk.Zone{
				complexColorLayer1ZoneA,
			},
			TotalWeight: 100,
		},
		ZoneID: "D",

		Weight: sdk.Weight{
			Min: 1,
			Max: 100,
		},

		Value:      "D",
		Tag:        "背景 白色",
		UserGroups: []string{},
	}
	complexColorLayer2_2ZoneD := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID: "Layer2-2",
			ParentZones: []*sdk.Zone{
				complexColorLayer1ZoneB,
			},
			TotalWeight: 100,
		},
		ZoneID: "D",

		Weight: sdk.Weight{
			Min: 1,
			Max: 50,
		},

		Value:      "D",
		Tag:        "背景 白色",
		UserGroups: []string{},
	}
	complexColorLayer2_2ZoneE := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID: "Layer2-2",
			ParentZones: []*sdk.Zone{
				complexColorLayer1ZoneB,
			},
			TotalWeight: 100,
		},
		ZoneID: "E",

		Weight: sdk.Weight{
			Min: 51,
			Max: 100,
		},

		Value:      "E",
		Tag:        "背景 黑色",
		UserGroups: []string{},
	}
	complexColorLayer2_3ZoneE := &sdk.Zone{
		ProjectID: "ComplexColor",
		Layer: sdk.Layer{
			LayerID: "Layer2-3",
			ParentZones: []*sdk.Zone{
				complexColorLayer1ZoneC,
			},
			TotalWeight: 100,
		},
		ZoneID: "E",

		Weight: sdk.Weight{
			Min: 1,
			Max: 100,
		},

		Value:      "E",
		Tag:        "背景 黑色",
		UserGroups: []string{},
	}

	labHomeZones := []*sdk.Zone{
		// project Home
		homeLayer1ZoneA,
		homeLayer1ZoneB,
		homeLayer1ZoneC,
		homeLayer1ZoneD,
	}
	labColorZones := []*sdk.Zone{
		// project Color
		colorLayer1ZoneA,
		colorLayer1ZoneB,
		colorLayer1ZoneC,
		colorLayer2ZoneD,
		colorLayer2ZoneE,
	}
	labComplexColorZones := []*sdk.Zone{
		// project ComplexColor
		complexColorLayer1ZoneA,
		complexColorLayer1ZoneB,
		complexColorLayer1ZoneC,
		complexColorLayer2_1ZoneD,
		complexColorLayer2_2ZoneD,
		complexColorLayer2_2ZoneE,
		complexColorLayer2_3ZoneE,
	}

	fmt.Print("创建 Home 实验配置...")
	printPoint(3)
	service.CreateABTestConfig("Home", labHomeZones)
	fmt.Print("创建 Color 实验配置...")
	printPoint(3)
	service.CreateABTestConfig("Color", labColorZones)
	fmt.Print("创建 Complex Color 实验配置...")
	printPoint(3)
	service.CreateABTestConfig("ComplexColor", labComplexColorZones)
}

func printPoint(n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Print(".")
	}
	time.Sleep(time.Second)
	fmt.Println("Done!")
	time.Sleep(time.Second)
}
