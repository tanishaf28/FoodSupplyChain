# 🥗 Food Supply Chain Management using Blockchain

A decentralized, transparent, and secure food supply chain management system built with **Hyperledger Fabric**, **React.js**, **Node.js**, and **GoLang**. This project ensures traceability, authenticity, and trust at every step  from manufacturer to consumer.

---

## 📦 Project Architecture

- **Frontend**: React.js
- **Middleware**: Node.js with Fabric SDK
- **Blockchain Network**: Hyperledger Fabric (v2.x)
- **Smart Contracts**: Chaincode written in Go
- **Certificate Authority**: Fabric CA (for each organization)
- **Database**: LevelDB (via Fabric)

---

## 🏗️ Network Setup

- **Organizations**: 
  - `ManufacturerOrg`
  - `MiddleMenOrg` (Wholesaler, Distributor, Retailer)
  - `ConsumerOrg`
- **Peers**: 5 peers across 3 orgs
- **Orderer**: RAFT-based
- **Channel**: `SupplychainChannel`

---

## 🔗 Chaincode Functionality

- `InitLedger`
- `CreateUser`
- `SignIn`
- `CreateProduct`
- `UpdateProduct` (for each supply chain stage)
- `QueryProduct`
- `QueryAll`

---

## 🚀 Features

- 👤 Role-based user system (Manufacturer, Wholesaler, Distributor, Retailer, Consumer)
- 🔍 Product traceability across the entire chain
- 🔐 Tamper-proof, immutable ledger
- 📬 Email notifications on product events (planned)
- 🧪 REST APIs for all chaincode operations

---

## 💻 Running the Project

### 1. Start the Network
```bash
./network.sh up createChannel -ca -s couchdb
./network.sh deployCC -ccn supplychain -ccp ../chaincode/ -ccl go
2. Start Backend (Node.js)
cd server
npm install
node app.js
3. Start Frontend (React.js)
cd client
npm install
npm start
📊 Future Enhancements
✅ Real-time product tracking dashboard

📱Increased Scalability

🧠 Machine learning for anomaly detection

📁 Folder Structure
Scm/
   scm
      ├── artifacts/            
      ├── bin/            
      ├── chaincode/        
      ├── network-scripts/           
      ├── web-app/            
      └── README.md
   high-throughput
   test-network

🙋‍♀️ Author
Tanisha Fonseca.
Master’s Student, Concordia University
🔗 GitHub: tanishaf28

📜 License
This project is open-source and available under the MIT License.
