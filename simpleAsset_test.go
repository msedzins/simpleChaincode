package main

import (     
 	"testing"     
 	"github.com/hyperledger/fabric/core/chaincode/shim" 
 ) 
 
 
 func TestSet(t *testing.T) {  
 	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleAsset))  
 	if stub == nil {   
 		t.Fatalf("MockStub creation failed")  
 	}  
 	
 	args := [][]byte{   
 		[]byte("set"), []byte("key1"),    
 		[]byte("value1")}   
 	
 	invokeResult := stub.MockInvoke("12345", args)  
 	if invokeResult.Status != 200 {   
 		t.Errorf("Set returned non-OK status, got: %d, want: %d. Error: %s", 
 			invokeResult.Status, 200, invokeResult.Message)  
 		return
 	}  
 	
 	value, _ := stub.GetState("key1")  
 	if string(value) != "value1"  {   t.Errorf("Bad value for the key! %s", value)  } 
 } 
 
 func TestGet(t *testing.T) {  
 	stub := shim.NewMockStub("mockChaincodeStub", new(SimpleAsset))  
 	if stub == nil {   
 		t.Fatalf("MockStub creation failed")  
 	}  

	//Insert dummy data
 	args := [][]byte{   
		[]byte("set"), []byte("key2"),    
		[]byte("value2")}   
 	
 	invokeResult := stub.MockInvoke("12345", args)  
 	if invokeResult.Status != 200 {   
 		t.Errorf("Set returned non-OK status, got: %d, want: %d. Error: %s", 
 			invokeResult.Status, 200, invokeResult.Message)  
 		return
 	}  
	//dummy data adding finished
 	
 	args2 := [][]byte{   
 		[]byte("get"), []byte("key2")}   
 	
 	invokeResult2 := stub.MockInvoke("12345", args2)  
 	if invokeResult2.Status != 200 {   
 		t.Errorf("Get returned non-OK status, got: %d, want: %d. Error: %s", 
 			invokeResult2.Status, 200, invokeResult2.Message)  
 		return
 	}  
 	
 	value := string(invokeResult2.Payload)  
 	if string(value) != "value2"  {   t.Errorf("Bad value for the key! %s", value)  } 
 } 
 