package environment

import (
    "crypto/sha256"
    // "encoding/base64"
    "encoding/hex"
    // "database/sql"
    // "fmt"
    // "strconv"
)

/**
	Utility methods
**/

var LocalMemoryMode bool
var LocalMemoryMap = make(map[string]string)

func HashAndStoreMessage(message string, hashString *string) bool {
	bv := []byte(message) 
    hash := sha256.Sum256(bv)

    *hashString = hex.EncodeToString(hash[:])
    if(StoreKV(*hashString, message)) {
    	return true
    }

    return false
}

// queries database using hash, populate message, return success
// if unsuccessful lookup return false
// if hash not found return false
func LookupHash(hash string, message *string) bool {
    if LocalMemoryMode {
    	if val, ok := LocalMemoryMap[hash]; ok {
    		*message = val
    		return true
		} else {
			return false
		}
	} else if DatabaseGetValue(hash, message) {
		return true
	}
    return false
}

func StoreKV(hashKey string, messageValue string) bool {
    if LocalMemoryMode {
    	LocalMemoryMap[hashKey] = messageValue
	} else if DatabaseInsertKV(hashKey, messageValue) {
		return true
	}
	return false
}
