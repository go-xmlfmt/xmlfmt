package xmlfmt_test

import (
	"fmt"
	"testing"

	"github.com/go-xmlfmt/xmlfmt"
)

const xml1 = `<root><this><is>a</is><test /><message><!-- with comment --><org><cn>Some org-or-other</cn><ph>Wouldnt you like to know</ph></org><contact><fn>Pat</fn><ln>Califia</ln></contact></message></this></root>`

const xml2 = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns"><soapenv:Header/><soapenv:Body><ns:request><ns:customer><ns:id>123</ns:id><ns:name type="NCHZ">John Brown</ns:name></ns:customer></ns:request></soapenv:Body></soapenv:Envelope>`

const xml3 = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:_xmlns="xmlns" _xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" _xmlns:ns="http://example.com/ns"><Header xmlns="http://schemas.xmlsoap.org/soap/envelope/"></Header><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><request xmlns="http://example.com/ns"><customer xmlns="http://example.com/ns"><id xmlns="http://example.com/ns">123</id><name xmlns="http://example.com/ns" type="NCHZ">John Brown</name></customer></request></Body></Envelope>`

func Example_output() {
	x3 := xmlfmt.FormatXML(xml3, "\t", " ")
	x2 := xmlfmt.FormatXML(xml2, "x ", " ")
	_ = x2
	_ = x3
	x1 := xmlfmt.FormatXML(xml1, "", "  ")
	fmt.Println(x1)
	// Output: <root>
	//   <this>
	//     <is>
	//       a</is>
	//     <test />
	//     <message>
	//       <!-- with comment -->
	//       <org>
	//         <cn>
	//           Some org-or-other</cn>
	//         <ph>
	//           Wouldnt you like to know</ph>
	//         </org>
	//       <contact>
	//         <fn>
	//           Pat</fn>
	//         <ln>
	//           Califia</ln>
	//         </contact>
	//       </message>
	//     </this>
	//   </root>
	//
}

const w1 = `..<root>
..  <this>
..    <is>
..      a</is>
..    <test />
..    <message>
..      <!-- with comment -->
..      <org>
..        <cn>
..          Some org-or-other</cn>
..        <ph>
..          Wouldnt you like to know</ph>
..        </org>
..      <contact>
..        <fn>
..          Pat</fn>
..        <ln>
..          Califia</ln>
..        </contact>
..      </message>
..    </this>
..  </root>
..`

const w2 = `x <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns">
x  <soapenv:Header/>
x  <soapenv:Body>
x   <ns:request>
x    <ns:customer>
x     <ns:id>
x      123</ns:id>
x     <ns:name type="NCHZ">
x      John Brown</ns:name>
x     </ns:customer>
x    </ns:request>
x   </soapenv:Body>
x  </soapenv:Envelope>
x `

const w3 = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:_xmlns="xmlns" _xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" _xmlns:ns="http://example.com/ns">
 <Header xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  </Header>
 <Body xmlns="http://schemas.xmlsoap.org/soap/envelope/">
  <request xmlns="http://example.com/ns">
   <customer xmlns="http://example.com/ns">
    <id xmlns="http://example.com/ns">
     123</id>
    <name xmlns="http://example.com/ns" type="NCHZ">
     John Brown</name>
    </customer>
   </request>
  </Body>
 </Envelope>
`

func TestFormatXML_t0(t *testing.T) {
	xmlfmt.NL = "\n"
}

func TestFormatXML_t1(t *testing.T) {
	x1 := xmlfmt.FormatXML(xml1, "..", "  ")
	if x1 != w1 {
		t.Errorf("got:\n%s, want:\n%s.", x1, w1)
	}
}

func TestFormatXML_t2(t *testing.T) {
	x2 := xmlfmt.FormatXML(xml2, "x ", " ")
	if x2 != w2 {
		t.Errorf("got:\n%s, want:\n%s.", x2, w2)
	}
}

func TestFormatXML_t3(t *testing.T) {
	x3 := xmlfmt.FormatXML(xml3, "", " ")
	if x3 != w3 {
		t.Errorf("got:\n%s, want:\n%s.", x3, w3)
	}
}
