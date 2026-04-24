# 🎸 Groupie Tracker

**Groupie Tracker** is a web-based platform built with **Go** that consumes a RESTful API to manipulate and visualize data about music bands, their concert locations, and scheduled dates. The project showcases the ability to handle complex data structures, JSON unmarshalling, and client-server architecture using only standard packages.

## 🚀 Key Features
* **Data Manipulation**: Successfully consumes four API endpoints: Artists, Locations, Dates, and Relation.
* **Data Visualization**: Displays band information through cards, lists, and detailed pages.
* **Client-Server Events**: Implements a feature that triggers communication between the client and server (request-response).
* **Error Handling**: Custom handlers for error pages to ensure the website and server never crash.
* **Standard Library**: Built strictly using standard Go packages as per the project constraints.

---
## Author
- aymane bouzerda (@bouzerda0)

---
## 🛠️ Tech Stack
* **Backend**: Go (Standard Library).
* **Frontend**: HTML5, CSS3.
* **Data Format**: JSON.
* **Architecture**: Layered Modular Structure.

---

## 📂 Project Structure
```text
groupie-tracker/
├── main.go            # Application entry point

├── handlers/          # Logic Layer: HTTP request handlers
├── models/            # Data Layer: Data structures (Structs)
├── utils/             # Data Layer: API fetch and JSON logic
├── templates/         # UI Layer: HTML template files
└── static/            # UI Layer: CSS and static assets

