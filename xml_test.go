/**
@author: Vikas K
**/
package main

import (
	"fmt"
	"testing"
)

var xmlRaw = `<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
<s:Header>
	<ActivityId CorrelationId="22540d5d-be76-4321-bb2d-8a5ddcb71dc2" xmlns="http://schemas.microsoft.com/2004/09/ServiceModel/Diagnostics">00000000-0000-0000-0000-000000000000</ActivityId>
</s:Header>
<s:Body>
	<objExecResponse xmlns="http://tempuri.org/">
		<objExecResult xmlns:a="http://schemas.datacontract.org/2004/07/BOService.Class" xmlns:i="http://www.w3.org/2001/XMLSchema-instance">
			<a:blnSuccess>true</a:blnSuccess>
			<a:intException>0</a:intException>
			<a:lstDataViews>
				<a:DataView i:type="b:Auditoriums_TableView" xmlns:b="http://schemas.datacontract.org/2004/07/BOService.Class.Views.Auditoriums">
					<b:Audi_chrStatus>Y </b:Audi_chrStatus>
					<b:Audi_chrStatusDesc>Yes</b:Audi_chrStatusDesc>
					<b:Audi_sntNum>1</b:Audi_sntNum>
					<b:Audi_strName>Yellow Bar </b:Audi_strName>
					<b:Audi_strShortName>Yellow Bar</b:Audi_strShortName>
					<b:Revamp_Master>N</b:Revamp_Master>
					<b:Venue_strCode>YBAD</b:Venue_strCode>
					<b:Venue_strName>Yellow Bar All Day: Mumbai</b:Venue_strName>
					<b:intTabStatus>0</b:intTabStatus>
				</a:DataView>
				<a:DataView i:type="b:Auditoriums_TableView" xmlns:b="http://schemas.datacontract.org/2004/07/BOService.Class.Views.Auditoriums">
					<b:Audi_chrStatus>Y </b:Audi_chrStatus>
					<b:Audi_chrStatusDesc>Yes</b:Audi_chrStatusDesc>
					<b:Audi_sntNum>2</b:Audi_sntNum>
					<b:Audi_strName>Yellow Comedy</b:Audi_strName>
					<b:Audi_strShortName>Yellow Com</b:Audi_strShortName>
					<b:Revamp_Master>N</b:Revamp_Master>
					<b:Venue_strCode>YBAD</b:Venue_strCode>
					<b:Venue_strName>Yellow Bar All Day: Mumbai</b:Venue_strName>
					<b:intTabStatus>0</b:intTabStatus>
				</a:DataView>
			</a:lstDataViews>
			<a:strData/>
			<a:strException/>
			<a:strVersion>20.11.27.01</a:strVersion>
		</objExecResult>
	</objExecResponse>
</s:Body>
</s:Envelope>`

func TestGetXmlMap(t *testing.T) {
	receivingXmlMap, err := GetXmlMap(xmlRaw, "->")
	if err != nil {
		fmt.Println("Err: ", err)
		return
	} else if receivingXmlMap == nil {
		fmt.Println("Received nil xml map")
		return
	}
	for key, value := range receivingXmlMap {
		fmt.Println(key, " > ", value)
	}
}
