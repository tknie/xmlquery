package xmlquery

import (
	"io"
	"os"
	"testing"
)

const resultOk = `<?xml version="1.0" encoding="UTF-8"?>
<test>
  <property name="bookmarks"/>
  <style/>
  <group>
    <groupHeader>
      <band height="40"/>
      <property name="cheight" value="pixel"/>
      <textField>
        <reportElement mode="Opaque" x="0" y="0" width="515" height="40" forecolor="#FFFFFF" backcolor="#0000FF" uuid="5e6f64bc-a822-4ba3-b80c-edd7d8fdd382">
          <property name="cx" value="pixel"/>
          <property name="height" value="pixel"/>
        </reportElement>
        <textFieldExpression><![CDATA["First Group Header"]]></textFieldExpression>
      </textField>
    </groupHeader>
  </group>
</test>`

func TestModifications(t *testing.T) {
	f, err := os.Open("tests/1.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := io.Reader(f)
	doc, err := Parse(r)
	FindEach(doc, "//groupHeader", func(i int, node *Node) {
		bandNode := node.SelectElement("band")
		MoveChildNodes(bandNode, node)
		RemoveWithCriterium(node, "//property", func(node *Node) bool {
			for _, a := range node.Attr {
				if a.Name.Local == "name" {
					if a.Value == "y" {
						return true
					}
				}
			}
			return false
		})
		// FindEach(node, "//property", func(i int, n *Node) {
		// 	if n.PrevSibling == nil {
		// 		n.FirstChild = n.NextSibling
		// 	} else {
		// 		n.PrevSibling.NextSibling = n.NextSibling
		// 		if n.NextSibling != nil {
		// 			n.NextSibling.PrevSibling = n.PrevSibling
		// 		}
		// 	}
		// })
	})
	if resultOk != doc.OutputXMLWithOptions(WithOutputSelf(), WithIndentation("  "), WithEmptyTagSupport(), WithoutPreserveSpace()) {
		t.Fatal("Result not ok")
	}
}
