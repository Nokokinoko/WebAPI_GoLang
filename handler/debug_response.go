package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func debugResponse(context echo.Context, requestUri string, fields map[string]interface{}, response interface{}) error {
	message := map[string]interface{}{
		"disconn":     true,
		"server_uri":  context.Path(),
		"request_uri": requestUri,
		"response":    response,
	}
	if fields != nil && 0 < len(fields) {
		message["fields"] = fields
	}
	return context.JSON(http.StatusOK, message)
}

// --- Balance ---
func debugResponseBalance() interface{} {
	return "100"
}

// --- List Collections ---
func debugResponseListCollections() interface{} {
	return []map[string]interface{}{
		{
			"displayName": "Name_1",
			"fqCn":        "com.qsn.qgn.qcn1",
			"address":     "Address_1",
			"createdAt":   1577804400,
		},
		{
			"displayName": "Name_2",
			"fqCn":        "com.qsn.qgn.qcn2",
			"address":     "Address_2",
			"createdAt":   1609426800,
		},
	}
}

// --- Collection Detail ---
func debugResponseCollectionDetail() interface{} {
	return map[string]interface{}{
		"address":     "Address_1",
		"fqCn":        "com.qsn.qgn.qcn1",
		"displayName": "Name_1",
		"createdAt":   1577804400,
		"totalSupply": map[string]interface{}{
			"Supply1": 10,
			"Supply2": 20,
		},
		"description": "Description_1",
		"imageUrl":    "https://example",
		"status":      "Status_OK",
	}
}

// --- Clone ---
func debugResponseClone() interface{} {
	return map[string]interface{}{
		"txHash":      "Hash_1",
		"description": "Description_1",
		"details": map[string]interface{}{
			"KeyDetail": "ValDetail",
		},
	}
}

// --- Mint ---
func debugResponseMint() interface{} {
	return map[string]interface{}{
		"txHash":      "Hash_1",
		"description": "Description_1",
		"details": map[string]interface{}{
			"KeyDetail": "ValDetail",
		},
	}
}

// --- Burn ---
func debugResponseBurn() interface{} {
	return map[string]interface{}{
		"txHash":      "Hash_1",
		"description": "Description_1",
		"details": map[string]interface{}{
			"KeyDetail": "ValDetail",
		},
	}
}

// --- Burn to Mint ---
func debugResponseBurnToMint() interface{} {
	return map[string]interface{}{
		"txHash":      "Hash_1",
		"description": "Description_1",
		"details": map[string]interface{}{
			"KeyDetail": "ValDetail",
		},
	}
}

// --- Transfer ---
func debugResponseTransfer() interface{} {
	return map[string]interface{}{
		"txHash":      "Hash_1",
		"description": "Description_1",
		"details": map[string]interface{}{
			"KeyDetail": "ValDetail",
		},
	}
}

// --- List Tokens ---
func debugResponseListTokens() interface{} {
	return []map[string]interface{}{
		{
			"displayName": "Name_1",
			"fqTn":        "com.qsn.qgn.qcn1.qtn1",
			"totalSupply": 10,
		},
		{
			"displayName": "Name_2",
			"fqTn":        "com.qsn.qgn.qcn1.qtn2",
			"totalSupply": 20,
		},
	}
}

// --- List Tokens by Collection ---
func debugResponseListTokensByCollection() interface{} {
	return []map[string]interface{}{
		{
			"displayName": "Name_1",
			"fqTn":        "com.qsn.qgn.qcn1.qtn1",
			"totalSupply": 10,
		},
		{
			"displayName": "Name_2",
			"fqTn":        "com.qsn.qgn.qcn1.qtn2",
			"totalSupply": 20,
		},
	}
}

// --- Token Detail ---
func debugResponseTokenDetail() interface{} {
	return map[string]interface{}{
		"displayName":       "Name_1",
		"description":       "Description_1",
		"imageUrl":          "https://example",
		"fqTn":              "com.qsn.qgn.qcn1.qtn1",
		"collectionAddress": "Address_1",
		"totalSupply":       10,
		"properties": map[string]interface{}{
			"Properties_1": map[string]interface{}{
				"displayName": "NameProperties_1",
				"value":       "ValueProperties_1",
			},
		},
		"mutableProperties": map[string]interface{}{
			"MutableProperties_1": map[string]interface{}{
				"displayName": "NameMutableProperties_1",
				"value":       "ValueMutableProperties_1",
			},
		},
		"createdAt": 1577804400,
	}
}

// --- List User Tokens ---
func debugResponseListUserTokens() interface{} {
	return []map[string]interface{}{
		{
			"fqTn":        "com.qsn.qgn.qcn1.qtn1",
			"totalSupply": 10,
			"amount":      1,
		},
		{
			"fqTn":        "com.qsn.qgn.qcn1.qtn2",
			"totalSupply": 20,
			"amount":      2,
		},
	}
}

// --- List User Tokens by Collection ---
func debugResponseListUserTokensByCollection() interface{} {
	return []map[string]interface{}{
		{
			"fqTn":        "com.qsn.qgn.qcn1.qtn1",
			"totalSupply": 10,
			"amount":      1,
		},
		{
			"fqTn":        "com.qsn.qgn.qcn1.qtn2",
			"totalSupply": 20,
			"amount":      2,
		},
	}
}

// --- User Token Detail ---
func debugResponseUserTokenDetail() interface{} {
	return map[string]interface{}{
		"fqTn":        "com.qsn.qgn.qcn1.qtn1",
		"totalSupply": 10,
		"amount":      1,
	}
}

// --- List Users by Token ---
func debugResponseListUsersByToken() interface{} {
	return []map[string]interface{}{
		{
			"address": "Addres_1",
			"amount":  1,
		},
		{
			"address": "Addres_2",
			"amount":  2,
		},
	}
}
