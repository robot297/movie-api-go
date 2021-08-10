# Welcome to Summit API Day

Here you will find all information to get started with building and deploying your web service on heroku.

## Labs

What you need:

* Postman installed on your workstation

### Lab 1

Use the provided API endpoint to create a movie entry. ​Use your name as the 'Name' criteria.
Paste the response message & response code in the Teams chat.​

**Hint:** ***Utilize the accounts API documentation below.***

### Lab 2

* Using the data from Lab 1, update your name to <studentid>

**Hint:** *Utilize the accounts API documentation below.*

### Lab 3

* Using the data from Lab 2, use the provided API endpoint to add $10,000 to your new bank account! Paste the response message & response code in the Teams chat.

**Hint:** ***Utilize the accounts and transactions API documentation below.***

### Lab 4

* Similar to Lab 3, use the provided API endpoint to add $3,000,000 to your new bank account! You hit gold, baby!

**Hint:** ***Utilize the accounts and transactions API documentation below.***

### Lab 5

* Use the provided API endpoint to withdraw $1,000 from your bank account! This money won’t spend itself! Paste the response message & response code in the Teams chat.

**Hint:** ***Utilize the accounts and transactions API documentation below.***

### Lab 6

* The gig is up! Use the provided API endpoint to delete your account! Paste the response message & response code in the Teams chat.

**Hint:** ***Utilize the accounts API documentation below.***

## API Endpoints

### Base URL

[https://summit-movie-api.herokuapp.com/movie](https://summit-movie-api.herokuapp.com/movie)

### Movies

#### Get All Movies

| Web Method | End Point |
| ------- | ------ |
| **GET** | /movie |

Returns information for all movies.

==== Get one account ====
{| class="wikitable"
|-
|'''GET'''
| /simplebank/v1/accounts/<code>{id}</code>
|}

Returns information for one account.

==== Create an account ====
{| class="wikitable"
|-
|'''POST'''
|/simplebank/v1/accounts
|}

Creates an account.

Example:

<syntaxhighlight lang=json>
{   
    "owner": "Hulk",
    "type": "Savings" 
}
</syntaxhighlight>

==== Update an account using the account number ID====
{| class="wikitable"
|-
|'''PUT'''
|/simplebank/v1/accounts/<code>{id}</code>
|}

Updates the owner's name on an account.

Example:

<syntaxhighlight lang=json>
{
    "owner": "LP"
}
</syntaxhighlight>

==== Delete an account ====
{| class="wikitable"
|-
|'''DELETE'''
|/simplebank/v1/accounts/<code>{account_number}</code>
|}

Delete one account.

=== Transactions ===
==== Deposit Money to an account====
{| class="wikitable"
|-
|'''POST'''
|/simplebank/v1/transactions/deposit/<code>{account_number}</code>
|}

Deposits money into an account.

Example:

<syntaxhighlight lang=json>
{
    "amount": 100000000
}
</syntaxhighlight>

==== Withdraw Money from an account ====
{| class="wikitable"
|-
|'''POST'''
|/simplebank/v1/transactions/withdraw/<code>{account_number}</code>
|}

Withdraws money from an account.

### Reset
####  Reset data back to default
| Web Method | End Point |
| ------- | ------ |
| **DELETE** | /reset |

Reset data back to default.

```
{
	"info": {
		"_postman_id": "17ed9561-caf3-4e8c-9372-9636ad4d4a3a",
		"name": "SimpleBank",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "Get All Accounts",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {},
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/accounts/",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"accounts",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "Get All One Account",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {},
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/accounts/12345678",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"accounts",
						"12345678"
					]
				}
			},
			"response": []
		},
		{
			"name": "Create Account",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {}
			},
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"owner\": \"Hulk\",\n    \"type\": \"Savings\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/accounts",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"accounts"
					]
				}
			},
			"response": []
		},
		{
			"name": "Update Account Owners Name",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {}
			},
			"request": {
				"method": "PUT",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"owner\": \"LP\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/accounts/12345678",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"accounts",
						"12345678"
					]
				}
			},
			"response": []
		},
		{
			"name": "Delete an account",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {}
			},
			"request": {
				"method": "DELETE",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/accounts/11018614",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"accounts",
						"11018614"
					]
				}
			},
			"response": []
		},
		{
			"name": "Deposit Money",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {}
			},
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"amount\": 100000000\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/transactions/deposit/12345678",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"transactions",
						"deposit",
						"12345678"
					]
				}
			},
			"response": []
		},
		{
			"name": "Withdraw Money",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {}
			},
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"amount\": 1\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/simplebank/v1/transactions/withdraw/12345678",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"simplebank",
						"v1",
						"transactions",
						"withdraw",
						"12345678"
					]
				}
			},
			"response": []
		},
		{
			"name": "health",
			"protocolProfileBehavior": {
				"disabledSystemHeaders": {},
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "http://simplebank.wot-api-day.10.127.214.195.xip.io/health",
					"protocol": "http",
					"host": [
						"simplebank",
						"wot-api-day",
						"10",
						"127",
						"214",
						"195",
						"xip",
						"io"
					],
					"path": [
						"health"
					]
				}
			},
			"response": []
		}
	],
	"protocolProfileBehavior": {}
}
```