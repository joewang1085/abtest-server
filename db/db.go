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
		*
										Project : Order
											实验规则：
											all users
											[1,100]
												|
												/\
										/					\
									/					/		\
								/					/				\
								|					|				|
								A1				B1-1				B1-2
							layer "1":[1,50]	layer "1":[51,75] 	layer "1":[76,100]
								|						|		\/		|
							/		\					|		/\		|
						|				|				||				||
						A2-1			A2-2			B2-1			B2-2
			layer "A2":[1,50]	layer "A2":[51,100]	layer "B2":[1,50]	layer "B2":[51,100]
														\				/
															\		/
																\/
																B3
														layer "B3": [1:100]
		*
	*/
	orderA1 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A",

		Weight: sdk.Weight{
			Min: 1,
			Max: 50,
		},

		Value: "A1",
		Tag:   "A1",
	}
	orderA2_1 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "A2",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A2-1",

		Weight: sdk.Weight{
			Min: 1,
			Max: 50,
		},

		Value: "A2-1",
		Tag:   "A2-1",
	}
	orderA2_2 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "A2",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A2-2",

		Weight: sdk.Weight{
			Min: 51,
			Max: 100,
		},

		Value: "A2-2",
		Tag:   "A2-2",
	}
	orderB1_1 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "B1-1",

		Weight: sdk.Weight{
			Min: 51,
			Max: 75,
		},

		Value: "B1-1",
		Tag:   "B1-1",
	}
	orderB1_2 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "B1-2",

		Weight: sdk.Weight{
			Min: 76,
			Max: 100,
		},

		Value: "B1-2",
		Tag:   "B1-2",
	}
	orderB2_1 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "B2",
			ParentZones: []*sdk.Zone{orderB1_2, orderB1_1},
			TotalWeight: 100,
		},
		ZoneID: "B2-1",

		Weight: sdk.Weight{
			Min: 1,
			Max: 50,
		},

		Value: "B2-1",
		Tag:   "B2-1",
	}
	orderB2_2 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "B2",
			ParentZones: []*sdk.Zone{orderB1_2, orderB1_1},
			TotalWeight: 100,
		},
		ZoneID: "B2-2",

		Weight: sdk.Weight{
			Min: 51,
			Max: 100,
		},

		Value: "B2-2",
		Tag:   "B2-2",
	}
	orderB3 := &sdk.Zone{
		ProjectID: "Order",
		Layer: sdk.Layer{
			LayerID:     "B3",
			ParentZones: []*sdk.Zone{orderB2_2, orderB2_1},
			TotalWeight: 100,
		},
		ZoneID: "B3",

		Weight: sdk.Weight{
			Min: 1,
			Max: 100,
		},

		Value: "B3",
		Tag:   "B3",
	}

	/*
			*
	Project: Display
	实验规则：
	All users
	|||
	|||
	//\\
	//\\
	//\\
	//\\
	ABA1A2
	(A)(B)(A)(A)
	[1.20] [21,60][61,80][81,100]
			*
	*/
	displayA := &sdk.Zone{
		ProjectID: "Display",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A",

		Weight: sdk.Weight{
			Min: 1,
			Max: 20,
		},

		Value:      "A",
		Tag:        "A",
		UserGroups: []string{"All"},
	}
	displayB := &sdk.Zone{
		ProjectID: "Display",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "B",

		Weight: sdk.Weight{
			Min: 21,
			Max: 60,
		},

		Value:      "B",
		Tag:        "B",
		UserGroups: []string{"VIP", "Frequent"},
	}
	displayA1 := &sdk.Zone{
		ProjectID: "Display",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A1",

		Weight: sdk.Weight{
			Min: 61,
			Max: 80,
		},

		Value:      "A",
		Tag:        "A1",
		UserGroups: []string{"VIP", "Frequent"},
	}
	displayA2 := &sdk.Zone{
		ProjectID: "Display",
		Layer: sdk.Layer{
			LayerID:     "1",
			ParentZones: nil,
			TotalWeight: 100,
		},
		ZoneID: "A2",

		Weight: sdk.Weight{
			Min: 81,
			Max: 100,
		},

		Value:      "A",
		Tag:        "A2",
		UserGroups: []string{"VIP", "Frequent"},
	}

	labSearchZones := []*sdk.Zone{
		/*
						*
								project search
								实验规则：
									    all users
									    [1.100]
										   |
									   /       \
									  /         \
									 |           |
									 A           B
			            layer "1":[1,50]     layer "1":[51,100]


						*
		*/
		&sdk.Zone{
			ProjectID: "search",
			Layer: sdk.Layer{
				LayerID:     "1",
				ParentZones: nil,
				TotalWeight: 10,
			},
			ZoneID: "A",

			Weight: sdk.Weight{
				Min: 1,
				Max: 5,
			},

			Value: "A",
			Tag:   "A",
		},

		&sdk.Zone{
			ProjectID: "search",
			Layer: sdk.Layer{
				LayerID:     "1",
				ParentZones: nil,
				TotalWeight: 10,
			},
			ZoneID: "B",

			Weight: sdk.Weight{
				Min: 6,
				Max: 10,
			},

			Value: "B",
			Tag:   "B",
		},
	}
	labOrderZones := []*sdk.Zone{
		// project Order
		orderA1,
		orderA2_1,
		orderA2_2,
		orderB1_1,
		orderB1_2,
		orderB2_1,
		orderB2_2,
		orderB3,
	}
	labDisplayZones := []*sdk.Zone{
		// project Display
		displayA,
		displayB,
		displayA1,
		displayA2,
	}

	fmt.Print("创建 search 实验配置...")
	printPoint(3)
	service.CreateABTestConfig("search", labSearchZones)
	fmt.Print("创建 Order 实验配置...")
	printPoint(3)
	service.CreateABTestConfig("Order", labOrderZones)
	fmt.Print("创建 Display 实验配置...")
	printPoint(3)
	service.CreateABTestConfig("Display", labDisplayZones)
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
