# Soal Interview: API & SQL

## Jawaban
2.  *select distinct c.id, c.name,o.id,o.order_date,o.total,oi.qty FROM customers as c full outer JOIN orders as o ON o.customer_id = c.id left join order_items as oi on o.id=oi.order_id order by c.id;*

## 1. Membuat API Add Data & Get Data List (In-Memory, Tanpa Database)

**Deskripsi:**

Buatlah sebuah API sederhana menggunakan bahasa pemrograman apapun dengan fungsionalitas:

### a. Add Customer

- **Endpoint:** POST `/api/customer/add`
- **Request body (JSON):**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789"
  }
  ```
- **Response (JSON):**
  ```json
  {
    "message": "Customer added successfully"
  }
  ```

### b. Get Customer List

- **Endpoint:** GET `/api/customer/get-all`
- **Response (JSON):**
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "08123456789"
    },
    ...
  ]
  ```

**Catatan:**

- Data customer **tidak perlu disimpan di database**, cukup menggunakan variable/list/array in-memory di aplikasi.
- Tidak perlu fitur edit/delete/autentikasi.
- Kirimkan kode lengkap & instruksi menjalankan (jika perlu).

---

## 2. Query SQL 

### Struktur Tabel

#### Inisialisasi Table (DDL)

```sql
CREATE TABLE customer (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);

CREATE TABLE `order` (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    total INT,
    FOREIGN KEY (customer_id) REFERENCES customer(id)
);

CREATE TABLE order_item (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    item_name VARCHAR(100),
    qty INT,
    price_per_item INT,
    FOREIGN KEY (order_id) REFERENCES `order`(id)
);
```

#### Insert Sample Data

```sql
INSERT INTO customer (name, email) VALUES
('John Doe', 'john@abc.com'),
('Jane Smith', 'jane@abc.com'),
('Bob', 'bob@abc.com');

INSERT INTO `order` (customer_id, order_date, total) VALUES
(1, '2024-06-01', 50000),
(1, '2024-06-03', 75000),
(2, '2024-06-04', 30000);

INSERT INTO order_item (order_id, item_name, qty, price_per_item) VALUES
(1, 'Keyboard', 1, 50000),
(2, 'Mouse', 1, 25000),
(2, 'Headset', 1, 50000),
(3, 'Mouse', 1, 30000);
```

---

### Soal Query:

#### **Tampilkan seluruh customer beserta (jika ada) data order-nya, yaitu:**

- `customer_id`
- `customer_name`
- `order_id`
- `order_date`
- `total_nominal`
- `jumlah_item_dibeli` (jumlah total item pada order tsb)

Jika customer belum pernah order, tetap ditampilkan (kolom order dan jumlah item = NULL/0).

**Contoh Output:**

| customer\_id | customer\_name | order\_id | order\_date | total\_nominal | jumlah\_item\_dibeli |
| ------------ | -------------- | --------- | ----------- | -------------- | -------------------- |
| 1            | John Doe       | 1         | 2024-06-01  | 50000          | 1                    |
| 1            | John Doe       | 2         | 2024-06-03  | 75000          | 2                    |
| 2            | Jane Smith     | 3         | 2024-06-04  | 30000          | 1                    |
| 3            | Bob            | NULL      | NULL        | NULL           | 0                    |


---

### **Instruksi:**

- Tulis query seperti di atas.
- Hasilkan output seperti contoh tabel.
- Query dapat digunakan baik di MySQL maupun PostgreSQL dengan sedikit penyesuaian pada penamaan tabel.

---
