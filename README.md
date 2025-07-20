
# GoSearch - English - [Türkçe](#türkçe-readme)


GoSearch is a **lightweight** and **blazingly fast** search engine implemented in Go, utilizing an inverted index stored in a SQLite database. It provides a RESTful API for indexing documents and performing search queries. The project uses the Gin web framework for handling HTTP requests and responses, and it supports basic text processing, such as tokenization, to facilitate efficient search operations. 

## Table of Contents
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)

## Features
- **Blazingly Fast**: Blazingly fast
- **Inverted Index**: Stores tokens and their associated document IDs in a SQLite database for efficient search.
- **Text Processing**: Supports tokenization, lowercasing, and removal of punctuation for clean search terms.
- **RESTful API**: Provides endpoints for adding documents to the index and searching for documents by query terms.
- **Lightweight**: Uses SQLite for storage, making it easy to set up and use without heavy dependencies.
- **Built with Gin**: Leverages the Gin framework for fast and scalable HTTP request handling.

## Installation

### Prerequisites
- Go (version 1.16 or higher)
- SQLite3
- Git (optional, for cloning the repository)

### Steps
1. **Clone the Repository** (if applicable):
   ```bash
   git clone https://github.com/yourusername/gosearch.git
   cd gosearch
   ```

2. **Install Dependencies**:
   Ensure you have the required Go modules by running:
   ```bash
   go mod tidy
   ```
   The project depends on:
   - `github.com/gin-gonic/gin` for the web framework
   - `github.com/mattn/go-sqlite3` for SQLite integration

3. **Set Up the Database**:
   The SQLite database (`search.db`) is automatically created when the application starts. No additional setup is required.

4. **Run the Application**:
   Start the server by running:
   ```bash
   go run .
   ```
   The server will listen on `http://localhost:8080`.

## Usage

### Adding a Document to the Index
Use the `/gosearch/add` endpoint to index a document. Provide a JSON payload with an `id` and `title`.

**Example Request**:
```bash
curl -X POST http://localhost:8080/gosearch/add \
-H "Content-Type: application/json" \
-d '{"id": 1, "title": "Sample document about Go programming"}'
```

**Example Response**:
```json
{
  "message": "Index added successfully",
  "data": ["sample", "document", "about", "go", "programming"]
}
```

### Searching for Documents
Use the `/gosearch/search` endpoint with a query parameter `q` to search for documents containing the specified term.

**Example Request**:
```bash
curl http://localhost:8080/gosearch/search?q=programming
```

**Example Response**:
```json
{
  "result": ["1"]
}
```

If no results are found:
```json
{
  "result": "Couldn't found any result."
}
```

## API Endpoints

| Method | Endpoint                | Description                              | Parameters                     |
|--------|-------------------------|------------------------------------------|--------------------------------|
| GET    | `/gosearch/search`      | Search for documents by query term       | `q` (query string, required)   |
| POST   | `/gosearch/add`         | Add a document to the inverted index     | JSON: `id` (int), `title` (string) |

## Project Structure
```
gosearch/
├── main.go               
├── core/                 
│   ├── core.go           
├── search.db             
└── go.mod                
```

### Key Components
- **main.go**: Initializes the Gin router, registers routes, and starts the HTTP server.
- **core/core.go**: Contains the core logic, including:
  - `InitDB()`: Sets up the SQLite database and creates the `inverted_index` table.
  - `Tokenize()`: Processes input text by lowercasing, removing punctuation, and splitting into tokens.
  - `InsertIndex()`: Adds tokens and document IDs to the inverted index.
  - `Search()`: Retrieves document IDs for a given search term.
- **inverted_index table**: Stores token-to-document mappings with columns `token` (TEXT) and `doc_id` (INTEGER).

---
---

# GoSearch <a name="türkçe-readme-section"></a>


GoSearch, Go dilinde geliştirilmiş **hafif** ve **çok hızlı** bir arama motorudur. SQLite veritabanında ters indeks (inverted index) kullanarak çalışır. RESTful API üzerinden belgeleri indeksleme ve arama sorguları gerçekleştirme imkanı sunar. Proje, HTTP isteklerini ve yanıtlarını yönetmek için Gin web çerçevesini kullanır ve temel metin işleme (tokenization) ile verimli arama işlemleri sağlar.

## İçindekiler
- [Özellikler](#özellikler)
- [Kurulum](#kurulum)
- [Kullanım](#kullanım)
- [API Endpoint'leri](#api-endpointleri)
- [Proje Yapısı](#proje-yapısı)

## Özellikler
- **Aşırı hızlı**: Aşırı hızlı
- **Ters İndeks**: Token'ları ve ilgili belge ID'lerini SQLite veritabanında saklayarak verimli arama sağlar.
- **Metin İşleme**: Küçük harfe çevirme, noktalama işaretlerini kaldırma ve token'lara ayırma işlemlerini destekler.
- **RESTful API**: Belgeleri indeksleme ve sorgu terimleriyle arama için endpointler sunar.
- **Hafif Yapı**: SQLite kullanıldığı için kurulumu kolaydır ve ağır bağımlılıklar gerektirmez.
- **Gin ile Geliştirildi**: Hızlı ve ölçeklenebilir HTTP istek yönetimi için Gin framework kullanır.

## Kurulum

### Gereksinimler
- Go (sürüm 1.16 veya üstü)
- SQLite3
- Git (opsiyonel, depoyu klonlamak için)

### Adımlar
1. **Depoyu Klonlayın** (isteğe bağlı):
   ```bash
   git clone https://github.com/kullaniciadiniz/gosearch.git
   cd gosearch
   ```

2. **Bağımlılıkları Yükleyin**:
   Gerekli Go modüllerini yüklemek için:
   ```bash
   go mod tidy
   ```
   Proje şu bağımlılıklara sahiptir:
   - `github.com/gin-gonic/gin` (web çerçevesi)
   - `github.com/mattn/go-sqlite3` (SQLite entegrasyonu)

3. **Veritabanını Kurun**:
   SQLite veritabanı (`search.db`), uygulama çalıştığında otomatik olarak oluşturulur. Ekstra bir kurulum gerekmez.

4. **Uygulamayı Çalıştırın**:
   Sunucuyu başlatmak için:
 handwriting
   ```bash
   go run .
   ```
   Sunucu `http://localhost:8080` adresinde çalışacaktır.

## Kullanım

### Belgeyi İndeksleme
`/gosearch/add` uç noktası ile bir belgeyi indeksleyin. JSON formatında `id` ve `title` içeren bir veri gönderin.

**Örnek İstek**:
```bash
curl -X POST http://localhost:8080/gosearch/add \
-H "Content-Type: application/json" \
-d '{"id": 1, "title": "Go programlama hakkında örnek belge"}'
```

**Örnek Yanıt**:
```json
{
  "message": "Index added successfully",
  "data": ["go", "programlama", "hakkinda", "ornek", "belge"]
}
```

### Belgelerde Arama
`/gosearch/search` uç noktası ile `q` parametresi kullanarak belgelerde arama yapın.

**Örnek İstek**:
```bash
curl http://localhost:8080/gosearch/search?q=programlama
```

**Örnek Yanıt**:
```json
{
  "result": ["1"]
}
```

Sonuç bulunamazsa:
```json
{
  "result": "Couldn't found any result."
}
```

## API Endpoint'leri

| Metod | Uç Nokta                | Açıklama                                 | Parametreler                   |
|-------|-------------------------|------------------------------------------|--------------------------------|
| GET   | `/gosearch/search`      | Sorgu terimine göre belgeleri ara        | `q` (sorgu string’i, zorunlu)  |
| POST  | `/gosearch/add`         | Ters indekse belge ekle                  | JSON: `id` (int), `title` (string) |

## Proje Yapısı
```
gosearch/
├── main.go               
├── core/                  
│   ├── core.go           
├── search.db           
└── go.mod                
```

### Ana Bileşenler
- **main.go**: Gin router’ını başlatır, rotaları kaydeder ve HTTP sunucusunu çalıştırır.
- **core/core.go**: Temel mantığı içerir:
  - `InitDB()`: SQLite veritabanını kurar ve `inverted_index` tablosunu oluşturur.
  - `Tokenize()`: Giriş metnini küçük harfe çevirir, noktalama işaretlerini kaldırır ve token’lara ayırır.
  - `InsertIndex()`: Token’ları ve belge ID’lerini ters indekse ekler.
  - `Search()`: Verilen arama terimi için belge ID’lerini getirir.
- **inverted_index tablosu**: Token-belge eşlemelerini `token` (TEXT) ve `doc_id` (INTEGER) sütunlarıyla saklar.
