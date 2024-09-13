Blockchain-based Asset Management System: 
This project is a blockchain-inspired asset management system built using Go and mux as a web framework. It allows a financial institution to manage and track accounts as assets, supporting functionalities 
like creating, updating, querying, and retrieving the transaction history of accounts. This system provides transparency, security, and immutability for account-related transactions.

Features:
Create Accounts: Add new accounts with specific attributes such as DealerID, MSISDN, MPIN, Balance, Status, TransAmount, TransType, and Remarks.

Read Accounts: Retrieve account details based on DealerID.

Update Accounts: Modify existing account information.

Delete Accounts: Remove accounts from the system.

List All Accounts: Get a list of all accounts stored in the system.

Transaction History: Retrieve mock transaction history for any account.


API Endpoints
HTTP Method	Endpoint	Description
POST	/accounts-	Create a new account

GET	/accounts-	List all accounts

GET	/accounts/{dealerId}-	Get details of a specific account by DealerID

PUT	/accounts/{dealerId}-	Update details of a specific account by DealerID

DELETE	/accounts/{dealerId}-	Delete an account by DealerID

GET	/accounts/{dealerId}/history-	Get transaction history for a specific account by DealerID
