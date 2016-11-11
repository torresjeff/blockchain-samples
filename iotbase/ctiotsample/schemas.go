package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
	)

var schemas = `

{
    "API": {
        "createAssetContainer": {
            "description": "Creates a new container",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this container",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all containers",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "barcode"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsContainer": {
            "description": "Delete all containers from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "description": "Match mode plus array of property : value pairs",
                        "properties": {
                            "match": {
                                "description": "Select match mode, missing property counts as not matched",
                                "enum": [
                                    "ALL",
                                    "ANY",
                                    "NONE"
                                ],
                                "type": "string"
                            },
                            "select": {
                                "description": "Array of property : value pairs to match",
                                "items": {
                                    "properties": {
                                        "qprop": {
                                            "description": "Qualified property name, e.g. container.barcode",
                                            "type": "string"
                                        },
                                        "value": {
                                            "description": "Property value to be matched",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssets"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetContainer": {
            "description": "Delete a container from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "maxItems": 1,
                        "minItems": 1,
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetContainer": {
            "description": "Delete one or more properties from a container's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "a container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteWorldState": {
            "description": "**** WARNING *** Clears the entire contents of world state, redeploy the contract after using this, in debugging mode, will require a restart",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteWorldState"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "initContract": {
            "description": "Sets contract version and nickname",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "nickname": {
                                "default": "CTIORSAMPLE",
                                "description": "The nickname of the current contract instance",
                                "type": "string"
                            },
                            "version": {
                                "description": "The version number of the current contract instance",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssetsContainer": {
            "description": "Returns the state of all containers, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "description": "Match mode plus array of property : value pairs",
                        "properties": {
                            "match": {
                                "description": "Select match mode, missing property counts as not matched",
                                "enum": [
                                    "ALL",
                                    "ANY",
                                    "NONE"
                                ],
                                "type": "string"
                            },
                            "select": {
                                "description": "Array of property : value pairs to match",
                                "items": {
                                    "properties": {
                                        "qprop": {
                                            "description": "Qualified property name, e.g. container.barcode",
                                            "type": "string"
                                        },
                                        "value": {
                                            "description": "Property value to be matched",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssets"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of container states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "A container's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This container's world state container ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "A list of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the container's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The container's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This container has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "container": {
                                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "A container's ID",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "The carrier in possession of this container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "Common properties for all containers",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "container": {
                                                "description": "An chaincode event emitted by a contract invoke",
                                                "properties": {
                                                    "name": {
                                                        "description": "The chaincode event's name",
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "The chaincode event's properties",
                                                        "properties": {},
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The container's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this container",
                                        "properties": {
                                            "container": {
                                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "A container's ID",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "The carrier in possession of this container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "Common properties for all containers",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetContainer": {
            "description": "Returns the state a container",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "a container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A container's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This container's world state container ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "A list of alert names",
                            "items": {
                                "description": "An alert name",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the container's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The container's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This container has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "container": {
                                    "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                    "properties": {
                                        "barcode": {
                                            "description": "A container's ID",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all containers",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of a container's contents in degrees Celsuis",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "container": {
                                    "description": "An chaincode event emitted by a contract invoke",
                                    "properties": {
                                        "name": {
                                            "description": "The chaincode event's name",
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "The chaincode event's properties",
                                            "properties": {},
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The container's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this container",
                            "properties": {
                                "container": {
                                    "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                    "properties": {
                                        "barcode": {
                                            "description": "A container's ID",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all containers",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of a container's contents in degrees Celsuis",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetHistoryContainer": {
            "description": "Returns the history of a container",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "barcode": {
                                "description": "A container's ID",
                                "type": "string"
                            },
                            "end": {
                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                "format": "date-time",
                                "sample": "yyyy-mm-dd hh:mm:ss",
                                "type": "string"
                            },
                            "start": {
                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                "format": "date-time",
                                "sample": "yyyy-mm-dd hh:mm:ss",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of container states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "A container's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This container's world state container ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "A list of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the container's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The container's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This container has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "container": {
                                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "A container's ID",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "The carrier in possession of this container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "Common properties for all containers",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "container": {
                                                "description": "An chaincode event emitted by a contract invoke",
                                                "properties": {
                                                    "name": {
                                                        "description": "The chaincode event's name",
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "The chaincode event's properties",
                                                        "properties": {},
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The container's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this container",
                                        "properties": {
                                            "container": {
                                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "A container's ID",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "The carrier in possession of this container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "Common properties for all containers",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readRecentStates": {
            "description": "Returns the state of recently updated assets",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readRecentStates"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of container states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "A container's complete state",
                                "properties": {
                                    "AssetKey": {
                                        "description": "This container's world state container ID",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "A list of alert names",
                                        "items": {
                                            "description": "An alert name",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the container's ID, declared in the contract code",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "The container's asset class",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "This container has no active alerts",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The contract event that created this state, for example updateAssetContainer",
                                        "properties": {
                                            "container": {
                                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "A container's ID",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "The carrier in possession of this container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "Common properties for all containers",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "The chaincode event emitted on invoke exit, if any",
                                        "properties": {
                                            "container": {
                                                "description": "An chaincode event emitted by a contract invoke",
                                                "properties": {
                                                    "name": {
                                                        "description": "The chaincode event's name",
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "The chaincode event's properties",
                                                        "properties": {},
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "The container's asset class prefix in world state",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "Properties that have been received or calculated for this container",
                                        "properties": {
                                            "container": {
                                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "A container's ID",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "The carrier in possession of this container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "Common properties for all containers",
                                                        "properties": {
                                                            "appdata": {
                                                                "description": "Application managed information as an array of key:value pairs",
                                                                "items": {
                                                                    "properties": {
                                                                        "K": {
                                                                            "type": "string"
                                                                        },
                                                                        "V": {
                                                                            "type": "string"
                                                                        }
                                                                    },
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "devicetimestamp": {
                                                                "description": "A timestamp recoded by the device that sent the current event",
                                                                "type": "string"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching the blockchain",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching the blockchain",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readWorldState": {
            "description": "Returns the entire contents of world state",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readWorldState"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "setCreateOnUpdate": {
            "description": "Allow updateAsset to create a container upon receipt of its first event",
            "properties": {
                "args": {
                    "items": {
                        "setCreateOnUpdate": {
                            "type": "boolean"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setCreateOnUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level for the contract",
            "properties": {
                "args": {
                    "items": {
                        "logLevel": {
                            "enum": [
                                "CRITICAL",
                                "ERROR",
                                "WARNING",
                                "NOTICE",
                                "INFO",
                                "DEBUG"
                            ],
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setLoggingLevel"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetContainer": {
            "description": "Update a contaner's state with one or more property changes",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this container",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all containers",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
                                            },
                                            "location": {
                                                "description": "A geographical coordinate",
                                                "properties": {
                                                    "latitude": {
                                                        "type": "number"
                                                    },
                                                    "longitude": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of a container's contents in degrees Celsuis",
                                        "type": "number"
                                    }
                                },
                                "required": [
                                    "barcode"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "container": {
            "description": "The changeable properties for a container, also considered its 'event' as a partial state",
            "properties": {
                "barcode": {
                    "description": "A container's ID",
                    "type": "string"
                },
                "carrier": {
                    "description": "The carrier in possession of this container",
                    "type": "string"
                },
                "common": {
                    "description": "Common properties for all containers",
                    "properties": {
                        "appdata": {
                            "description": "Application managed information as an array of key:value pairs",
                            "items": {
                                "properties": {
                                    "K": {
                                        "type": "string"
                                    },
                                    "V": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "devicetimestamp": {
                            "description": "A timestamp recoded by the device that sent the current event",
                            "type": "string"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                    "description": "Temperature of a container's contents in degrees Celsuis",
                    "type": "number"
                }
            },
            "required": [
                "barcode"
            ],
            "type": "object"
        },
        "containerstate": {
            "description": "A container's complete state",
            "properties": {
                "AssetKey": {
                    "description": "This container's world state container ID",
                    "type": "string"
                },
                "alerts": {
                    "description": "A list of alert names",
                    "items": {
                        "description": "An alert name",
                        "enum": [
                            "OVERTTEMP"
                        ],
                        "type": "string"
                    },
                    "type": "array"
                },
                "assetIDpath": {
                    "description": "Qualified property path to the container's ID, declared in the contract code",
                    "type": "string"
                },
                "class": {
                    "description": "The container's asset class",
                    "type": "string"
                },
                "compliant": {
                    "description": "This container has no active alerts",
                    "type": "boolean"
                },
                "eventin": {
                    "description": "The contract event that created this state, for example updateAssetContainer",
                    "properties": {
                        "container": {
                            "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                            "properties": {
                                "barcode": {
                                    "description": "A container's ID",
                                    "type": "string"
                                },
                                "carrier": {
                                    "description": "The carrier in possession of this container",
                                    "type": "string"
                                },
                                "common": {
                                    "description": "Common properties for all containers",
                                    "properties": {
                                        "appdata": {
                                            "description": "Application managed information as an array of key:value pairs",
                                            "items": {
                                                "properties": {
                                                    "K": {
                                                        "type": "string"
                                                    },
                                                    "V": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "devicetimestamp": {
                                            "description": "A timestamp recoded by the device that sent the current event",
                                            "type": "string"
                                        },
                                        "location": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "temperature": {
                                    "description": "Temperature of a container's contents in degrees Celsuis",
                                    "type": "number"
                                }
                            },
                            "required": [
                                "barcode"
                            ],
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "eventout": {
                    "description": "The chaincode event emitted on invoke exit, if any",
                    "properties": {
                        "container": {
                            "description": "An chaincode event emitted by a contract invoke",
                            "properties": {
                                "name": {
                                    "description": "The chaincode event's name",
                                    "type": "string"
                                },
                                "payload": {
                                    "description": "The chaincode event's properties",
                                    "properties": {},
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "prefix": {
                    "description": "The container's asset class prefix in world state",
                    "type": "string"
                },
                "state": {
                    "description": "Properties that have been received or calculated for this container",
                    "properties": {
                        "container": {
                            "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                            "properties": {
                                "barcode": {
                                    "description": "A container's ID",
                                    "type": "string"
                                },
                                "carrier": {
                                    "description": "The carrier in possession of this container",
                                    "type": "string"
                                },
                                "common": {
                                    "description": "Common properties for all containers",
                                    "properties": {
                                        "appdata": {
                                            "description": "Application managed information as an array of key:value pairs",
                                            "items": {
                                                "properties": {
                                                    "K": {
                                                        "type": "string"
                                                    },
                                                    "V": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "devicetimestamp": {
                                            "description": "A timestamp recoded by the device that sent the current event",
                                            "type": "string"
                                        },
                                        "location": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "temperature": {
                                    "description": "Temperature of a container's contents in degrees Celsuis",
                                    "type": "number"
                                }
                            },
                            "required": [
                                "barcode"
                            ],
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "txnid": {
                    "description": "Transaction UUID matching the blockchain",
                    "type": "string"
                },
                "txnts": {
                    "description": "Transaction timestamp matching the blockchain",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "containerstatearray": {
            "description": "Array of container states, can mix asset classes",
            "items": {
                "patternProperties": {
                    "^CON": {
                        "description": "A container's complete state",
                        "properties": {
                            "AssetKey": {
                                "description": "This container's world state container ID",
                                "type": "string"
                            },
                            "alerts": {
                                "description": "A list of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "enum": [
                                        "OVERTTEMP"
                                    ],
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetIDpath": {
                                "description": "Qualified property path to the container's ID, declared in the contract code",
                                "type": "string"
                            },
                            "class": {
                                "description": "The container's asset class",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "This container has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAssetContainer",
                                "properties": {
                                    "container": {
                                        "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                        "properties": {
                                            "barcode": {
                                                "description": "A container's ID",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this container",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all containers",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of a container's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "barcode"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "container": {
                                        "description": "An chaincode event emitted by a contract invoke",
                                        "properties": {
                                            "name": {
                                                "description": "The chaincode event's name",
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "The chaincode event's properties",
                                                "properties": {},
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "prefix": {
                                "description": "The container's asset class prefix in world state",
                                "type": "string"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this container",
                                "properties": {
                                    "container": {
                                        "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                        "properties": {
                                            "barcode": {
                                                "description": "A container's ID",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this container",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all containers",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of a container's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "barcode"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    }
                },
                "type": "object"
            },
            "minItems": 0,
            "type": "array"
        },
        "containerstateexternal": {
            "patternProperties": {
                "^CON": {
                    "description": "A container's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This container's world state container ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "A list of alert names",
                            "items": {
                                "description": "An alert name",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the container's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The container's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This container has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "container": {
                                    "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                    "properties": {
                                        "barcode": {
                                            "description": "A container's ID",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all containers",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of a container's contents in degrees Celsuis",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "container": {
                                    "description": "An chaincode event emitted by a contract invoke",
                                    "properties": {
                                        "name": {
                                            "description": "The chaincode event's name",
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "The chaincode event's properties",
                                            "properties": {},
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "The container's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this container",
                            "properties": {
                                "container": {
                                    "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                    "properties": {
                                        "barcode": {
                                            "description": "A container's ID",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all containers",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of a container's contents in degrees Celsuis",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "invokeevent": {
            "description": "An chaincode event emitted by a contract invoke",
            "properties": {
                "name": {
                    "description": "The chaincode event's name",
                    "type": "string"
                },
                "payload": {
                    "description": "The chaincode event's properties",
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "ioteventcommon": {
            "description": "Common properties for all containers",
            "properties": {
                "appdata": {
                    "description": "Application managed information as an array of key:value pairs",
                    "items": {
                        "properties": {
                            "K": {
                                "type": "string"
                            },
                            "V": {
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "devicetimestamp": {
                    "description": "A timestamp recoded by the device that sent the current event",
                    "type": "string"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "stateFilter": {
            "description": "Match mode plus array of property : value pairs",
            "properties": {
                "match": {
                    "description": "Select match mode, missing property counts as not matched",
                    "enum": [
                        "ALL",
                        "ANY",
                        "NONE"
                    ],
                    "type": "string"
                },
                "select": {
                    "description": "Array of property : value pairs to match",
                    "items": {
                        "properties": {
                            "qprop": {
                                "description": "Qualified property name, e.g. container.barcode",
                                "type": "string"
                            },
                            "value": {
                                "description": "Property value to be matched",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        }
    }
}`


	var readAssetSchemas as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(schemas), nil
	}
	func init() {
		as.AddRoute("readAssetSchemas", "query", as.SystemClass, readAssetSchemas)
	}
	