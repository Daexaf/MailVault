Handler

Tugasnya:

menerima HTTP Request
membaca JSON
validasi format dasar (binding)
memanggil Service
mengembalikan Response

Dia tidak boleh query database.

===========================================

Service

Tugasnya:

Business Logic
Validasi bisnis
Orkestrasi
Memanggil beberapa Repository bila perlu

=====================================================

Repository

Tugasnya:

Query database
Insert
Update
Delete
Search

Tidak ada keputusan bisnis.

================================================

Database

Tugasnya:

Menjaga integritas data

Misalnya:

Foreign Key
Unique
Not Null

===============================================

User
│
▼
POST /mail-accounts
│
▼
Handler
│
▼
MailAccountService
│
├── BranchRepository.FindByID()
│
├── MailAccountRepository.ExistsByEmail()
│
└── MailAccountRepository.Create()
│
▼
SQL Server

Service adalah tempat:

Transformasi + Keputusan + Business Rule
