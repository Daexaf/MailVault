# Learning #005 - SQL Server Connection

## Problem

GORM tidak dapat terhubung ke SQL Server.

Error:

```
no instance matching 'SQLEXPRESS01'
```

## Root Cause

TCP/IP pada SQL Server Configuration masih Disabled.

## Solution

1. SQL Server Configuration Manager
2. SQL Server Network Configuration
3. Protocols for SQLEXPRESS01
4. Enable TCP/IP
5. Restart SQL Server Service

## Lesson Learned

Jangan langsung menyalahkan kode.

Pastikan:

- SQL Service Running
- SQL Browser Running
- TCP/IP Enabled
- SQL Authentication benar
