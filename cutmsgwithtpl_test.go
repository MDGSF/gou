// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"encoding/json"
	"testing"
)

func genJSONMap(strJSON string) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	if err := json.Unmarshal([]byte(strJSON), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func testCutMsg(t *testing.T, tplJSON, msgJSON, expectResultJSON string) {
	tpl, err := genJSONMap(tplJSON)
	if err != nil {
		t.Logf("tpl, err = %v", err)
	}

	msg, err := genJSONMap(msgJSON)
	if err != nil {
		t.Logf("msg, err = %v", err)
	}

	expectResult, err := genJSONMap(expectResultJSON)
	if err != nil {
		t.Logf("expectResult, err = %v", err)
	}

	result := CutMsgWithTemplate(tpl, msg)

	resultJSONOut, _ := json.Marshal(result)
	expectResultJSONOut, _ := json.Marshal(expectResult)
	if string(resultJSONOut) != string(expectResultJSONOut) {
		t.Fatalf("not equal, result = %v, expectResult = %v\n", string(resultJSONOut), string(expectResultJSONOut))
	}
}

func TestCutMsgWithTemplate1(t *testing.T) {

	tplJSON := `
	{
		"frame_id": 0
	}
	`

	msgJSON := `
	{
		"frame_id": 293
	}
	`

	testCutMsg(t, tplJSON, msgJSON, msgJSON)
}

func TestCutMsgWithTemplate2(t *testing.T) {

	tplJSON := `
	{
		"frame_id": 0
	}
	`

	msgJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vb_warning": 0,
			"headway": 0.6000000238418579,
			"vehicle_id": 1,
			"frame_id": 293,
			"fcw": 0,
			"headway_warning": 0
		}
	}
	`

	expectResultJSON := `
	{
		"frame_id": 293
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate3(t *testing.T) {

	tplJSON := `
	{
		"frame_id": 0,
		"vehicle_warning": {
			"warning_level": 0,
			"vehicle_id": 0,
			"headway_warning": 0
		}
	}
	`

	msgJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vb_warning": 0,
			"headway": 0.6000000238418579,
			"vehicle_id": 1,
			"frame_id": 293,
			"fcw": 0,
			"headway_warning": 0
		}
	}
	`

	expectResultJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vehicle_id": 1,
			"headway_warning": 0
		}
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate4(t *testing.T) {

	tplJSON := `
	{
		"lane": [{
			"type": 0,
			"perspective_view_pts": 0,
			"width": 0,
			"start": 0,
			"end": 0,
			"perspective_view_poly_coeff": 0,
			"label": 0,
			"warning": 0,
			"color": 0
		}]
	}
	`

	msgJSON := `
	{
		"lane": [{
			"type": 1,
			"perspective_view_pts": [
				[6.745694637298584, 520.8117065429688],
				[46.79523468017578, 511.2055969238281],
				[81.6808853149414, 502.873779296875]
			],
			"width": 0.11999999731779099,
			"start": [-6.0, 520.8117065429688],
			"end": [47.7599983215332, 418.2877197265625],
			"confidence": 0.9847350716590881,
			"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
			"bird_view_poly_coeff": [-4.901395797729492, -0.07936269044876099, 0.0007704439340159297, -6.120264828268773e-08],
			"label": 0,
			"warning": false,
			"bird_view_pts": [
				[14.319999694824219, -5.900000095367432],
				[14.799999237060547, -5.980000019073486],
				[15.28000259399414, -5.980000019073486]
			],
			"color": 4
		}]
	}
	`

	expectResultJSON := `
	{
		"lane": [{
			"type": 1,
			"perspective_view_pts": [
				[6.745694637298584, 520.8117065429688],
				[46.79523468017578, 511.2055969238281],
				[81.6808853149414, 502.873779296875]
			],
			"width": 0.11999999731779099,
			"start": [-6.0, 520.8117065429688],
			"end": [47.7599983215332, 418.2877197265625],
			"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
			"label": 0,
			"warning": false,
			"color": 4
		}]
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate5(t *testing.T) {

	tplJSON := `
	{
		"lane": [{
			"type": 0,
			"perspective_view_pts": 0,
			"width": 0,
			"start": 0,
			"end": 0,
			"perspective_view_poly_coeff": 0,
			"label": 0,
			"warning": 0,
			"color": 0
		}]
	}
	`

	msgJSON := `
	{
		"lane": [
			{
				"type": 1,
				"perspective_view_pts": [
					[6.745694637298584, 520.8117065429688],
					[46.79523468017578, 511.2055969238281],
					[81.6808853149414, 502.873779296875]
				],
				"width": 0.11999999731779099,
				"start": [-6.0, 520.8117065429688],
				"end": [47.7599983215332, 418.2877197265625],
				"confidence": 0.9847350716590881,
				"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
				"bird_view_poly_coeff": [-4.901395797729492, -0.07936269044876099, 0.0007704439340159297, -6.120264828268773e-08],
				"label": 0,
				"warning": false,
				"bird_view_pts": [
					[14.319999694824219, -5.900000095367432],
					[14.799999237060547, -5.980000019073486],
					[15.28000259399414, -5.980000019073486]
				],
				"color": 4
			},
			{
				"type": 2,
				"perspective_view_pts": [
					[6.745694637298584, 520.8117065429688],
					[46.79523468017578, 511.2055969238281],
					[81.6808853149414, 502.873779296875]
				],
				"width": 0.11999999731779099,
				"start": [-6.0, 520.8117065429688],
				"end": [47.7599983215332, 418.2877197265625],
				"confidence": 0.9847350716590881,
				"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
				"bird_view_poly_coeff": [-4.901395797729492, -0.07936269044876099, 0.0007704439340159297, -6.120264828268773e-08],
				"label": 0,
				"warning": false,
				"bird_view_pts": [
					[14.319999694824219, -5.900000095367432],
					[14.799999237060547, -5.980000019073486],
					[15.28000259399414, -5.980000019073486]
				],
				"color": 5
			}
		]
	}
	`

	expectResultJSON := `
	{
		"lane": [{
			"type": 1,
			"perspective_view_pts": [
				[6.745694637298584, 520.8117065429688],
				[46.79523468017578, 511.2055969238281],
				[81.6808853149414, 502.873779296875]
			],
			"width": 0.11999999731779099,
			"start": [-6.0, 520.8117065429688],
			"end": [47.7599983215332, 418.2877197265625],
			"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
			"label": 0,
			"warning": false,
			"color": 4
		},
		{
			"type": 2,
			"perspective_view_pts": [
				[6.745694637298584, 520.8117065429688],
				[46.79523468017578, 511.2055969238281],
				[81.6808853149414, 502.873779296875]
			],
			"width": 0.11999999731779099,
			"start": [-6.0, 520.8117065429688],
			"end": [47.7599983215332, 418.2877197265625],
			"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
			"label": 0,
			"warning": false,
			"color": 5
		}]
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate6(t *testing.T) {

	tplJSON := `
	{
		"frame_id": 0,
		"vehicle_warning": {
			"warning_level": 0,
			"vehicle_id": 0,
			"headway_warning": 0
		},
		"lane": [{
			"type": 0,
			"perspective_view_pts": 0,
			"width": 0,
			"start": 0,
			"end": 0,
			"perspective_view_poly_coeff": 0,
			"label": 0,
			"warning": 0,
			"color": 0
		}]
	}
	`

	msgJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vb_warning": 0,
			"headway": 0.6000000238418579,
			"vehicle_id": 1,
			"frame_id": 293,
			"fcw": 0,
			"headway_warning": 0
		},
		"lane": [
			{
				"type": 1,
				"perspective_view_pts": [
					[6.745694637298584, 520.8117065429688],
					[46.79523468017578, 511.2055969238281],
					[81.6808853149414, 502.873779296875]
				],
				"width": 0.11999999731779099,
				"start": [-6.0, 520.8117065429688],
				"end": [47.7599983215332, 418.2877197265625],
				"confidence": 0.9847350716590881,
				"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
				"bird_view_poly_coeff": [-4.901395797729492, -0.07936269044876099, 0.0007704439340159297, -6.120264828268773e-08],
				"label": 0,
				"warning": false,
				"bird_view_pts": [
					[14.319999694824219, -5.900000095367432],
					[14.799999237060547, -5.980000019073486],
					[15.28000259399414, -5.980000019073486]
				],
				"color": 4
			},
			{
				"type": 2,
				"perspective_view_pts": [
					[6.745694637298584, 520.8117065429688],
					[46.79523468017578, 511.2055969238281],
					[81.6808853149414, 502.873779296875]
				],
				"width": 0.11999999731779099,
				"start": [-6.0, 520.8117065429688],
				"end": [47.7599983215332, 418.2877197265625],
				"confidence": 0.9847350716590881,
				"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
				"bird_view_poly_coeff": [-4.901395797729492, -0.07936269044876099, 0.0007704439340159297, -6.120264828268773e-08],
				"label": 0,
				"warning": false,
				"bird_view_pts": [
					[14.319999694824219, -5.900000095367432],
					[14.799999237060547, -5.980000019073486],
					[15.28000259399414, -5.980000019073486]
				],
				"color": 5
			}
		]
	}
	`

	expectResultJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vehicle_id": 1,
			"headway_warning": 0
		},
		"lane": [{
			"type": 1,
			"perspective_view_pts": [
				[6.745694637298584, 520.8117065429688],
				[46.79523468017578, 511.2055969238281],
				[81.6808853149414, 502.873779296875]
			],
			"width": 0.11999999731779099,
			"start": [-6.0, 520.8117065429688],
			"end": [47.7599983215332, 418.2877197265625],
			"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
			"label": 0,
			"warning": false,
			"color": 4
		},
		{
			"type": 2,
			"perspective_view_pts": [
				[6.745694637298584, 520.8117065429688],
				[46.79523468017578, 511.2055969238281],
				[81.6808853149414, 502.873779296875]
			],
			"width": 0.11999999731779099,
			"start": [-6.0, 520.8117065429688],
			"end": [47.7599983215332, 418.2877197265625],
			"perspective_view_poly_coeff": [2325.809326171875, -4.467968463897705, 0.0, 0.0],
			"label": 0,
			"warning": false,
			"color": 5
		}]
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate7(t *testing.T) {

	tplJSON := `
	{
		"frame_id": 0
	}
	`

	msgJSON := `
	{
		"test": 293
	}
	`

	expectResultJSON := `
	{}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate8(t *testing.T) {

	tplJSON := `
	{
		"arrtest": [
			{
				"item": 0,
				"name": 0,
				"key": 0
			}
		]
	}
	`

	msgJSON := `
	{
		"arrtest": [
			{
				"item": 123,
				"name": "item1"
			}
		]
	}
	`

	expectResultJSON := `
	{
		"arrtest": [
			{
				"item": 123,
				"name": "item1"
			}
		]
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate9(t *testing.T) {

	tplJSON := `
	{
		"arrtest": [
			{
				"item": 0,
				"name": 0,
				"key": 0
			}
		]
	}
	`

	msgJSON := `
	{
		"arrtestaaaaa": [
			{
				"item": 123,
				"name": "item1"
			}
		]
	}
	`

	expectResultJSON := `
	{
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate10(t *testing.T) {

	tplJSON := `
	{
		"frame_id": 0,
		"vehicle_warning": {
			"warning_level": 0,
			"vehicle_id": 0,
			"headway_warning": 0
		},
		"arrtest": [
			{
				"item": 0,
				"name": 0,
				"key": 0
			}
		]
	}
	`

	msgJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vb_warning": 0,
			"headway": 0.6000000238418579,
			"vehicle_id": 1,
			"frame_id": 293,
			"fcw": 0,
			"headway_warning": 0
		},
		"arrtestaaaaa": [
			{
				"item": 123,
				"name": "item1"
			}
		]
	}
	`

	expectResultJSON := `
	{
		"frame_id": 293,
		"vehicle_warning": {
			"warning_level": 2,
			"vehicle_id": 1,
			"headway_warning": 0
		}
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate11(t *testing.T) {

	tplJSON := `
	{
		"arrtest": [
			{
				"item": 0,
				"name": 0,
				"key": 0
			},
			{
				"keyboard": 0,
				"unknownvalue": "test"
			}
		]
	}
	`

	msgJSON := `
	{
		"arrtestaaaaa": [
			{
				"item": 123,
				"name": "item1"
			}
		]
	}
	`

	expectResultJSON := `
	{
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate12(t *testing.T) {

	tplJSON := `
	{
		"unknownvalue": "test"
	}
	`

	msgJSON := `
	{
		"arrtestaaaaa": [
			{
				"item": 123,
				"name": "item1"
			}
		]
	}
	`

	expectResultJSON := `
	{
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate13(t *testing.T) {

	tplJSON := `
	{
		"arrtest": [
			{
				"item": 0,
				"name": 0,
				"obj": {
					"objname": 0
				}
			}
		]
	}
	`

	msgJSON := `
	{
		"arrtest": [
			{
				"item": 123,
				"name": "item1",
				"obj": {
					"objname": "I'm object 1",
					"objkey": "kkkkkk 11111"
				}
			},
			{
				"item": 123,
				"name": "item1",
				"obj": {
					"objname": "I'm object 2",
					"objkey": "kkkkkk 22222"
				}
			}
		]
	}
	`

	expectResultJSON := `
	{
		"arrtest": [
			{
				"item": 123,
				"name": "item1",
				"obj": {
					"objname": "I'm object 1"
				}
			},
			{
				"item": 123,
				"name": "item1",
				"obj": {
					"objname": "I'm object 2"
				}
			}
		]
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}

func TestCutMsgWithTemplate14(t *testing.T) {

	tplJSON := `
	{
		"lane": [{
			"type": 0,
			"perspective_view_pts": 0,
			"width": 0,
			"start": 0,
			"end": 0,
			"perspective_view_poly_coeff": 0,
			"label": 0,
			"warning": 0,
			"color": 0
		}]
	}
	`

	msgJSON := `
	{
		"lane": []
	}
	`

	expectResultJSON := `
	{
		"lane": []
	}
	`

	testCutMsg(t, tplJSON, msgJSON, expectResultJSON)
}
