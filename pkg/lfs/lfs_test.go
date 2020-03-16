package lfs

import (
	"testing"
)

type testData struct {
	input          string
	padding        int
	expectedResult string
}

func Test(t *testing.T) {
	for _, test := range testCases {
		result := Obfp(test.input, test.padding)

		if result != test.expectedResult {
			t.Errorf("input: %s, padding: %d, result: %s, expected: %s\n", test.input, test.padding, result, test.expectedResult)
		}
	}
}

var testCases = [...]testData{
	// Empty input
	{
		input:          "",
		padding:        0,
		expectedResult: "VeryEvilPineapple",
	},
	// Test padding, positive and negative cases. Also, same input -> same output regardless of padding size.
	{
		input:          "asdf",
		padding:        -42,
		expectedResult: "RatherDeceitfulSwing",
	},
	{
		input:          "asdf",
		padding:        0,
		expectedResult: "RatherDeceitfulSwing",
	},
	{
		input:          "asdf",
		padding:        2,
		expectedResult: "RatherDeceitfulSwing5012",
	},
	{
		input:          "asdf",
		padding:        4,
		expectedResult: "RatherDeceitfulSwing5012F6C6",
	},
	{
		input:          "asdf",
		padding:        8,
		expectedResult: "RatherDeceitfulSwing5012F6C60B27661C",
	},
	{
		input:          "asdf",
		padding:        999,
		expectedResult: "RatherDeceitfulSwing5012F6C60B27661C",
	},
	// Test a few unique UUID:s
	{
		input:          "ac968750-7ca2-4dde-908b-aacbbed2f470",
		padding:        1,
		expectedResult: "StupidlyCatastrophicBobF4",
	},
	{
		input:          "3e3278cd-6030-400d-9c0d-ef9be0119853",
		padding:        5,
		expectedResult: "NextFuriousBobA2DEC84AEE",
	},
	{
		input:          "6745dc33-2fbd-4311-8884-10aab93199a9",
		padding:        7,
		expectedResult: "SoCatastrophicElephant7F2343BF6927EA",
	},
	// Big data blob
	{
		input: `mc093284750932nv2ono2hvfnoh3fo9ch3fxh23omhf293u4hfcqoiuwnhfc093u4hfc2938hnfc209u3hfc092hu3fc092nu3hfc92u3h4fc92nu3h4nfc923h40fc92h340fu2h34fc9u2nh3409uh2304hufc2093u4hfc0
		
		
		fcn9n2j43fc 9hu23cfj32fc2
		fc234ufh2o3ihfoh4f92c3hnfc928h43c92mj3fc23
		cfhfcliuw hfroiwuehgoiwuregoiwuecpowi hcpoqiwjecpoiqwhecp9824r+9u3h4f9283 h4f8w73hfwo83fou3wh4fcpoqihfp2u3h4fc983h4fcpu3nh4fcpoh3pf2h34pfc8h3p48hcqp348hfcqp384hfcpq834nfcpq9834hfcpq3h4fc`,
		padding:        0,
		expectedResult: "LoggedFastTree",
	},
}
