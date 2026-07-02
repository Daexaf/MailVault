# ADR 0001 - Project Structure

## Status

Accepted

## Context

MailVault akan dikembangkan dalam jangka panjang dan diperkirakan memiliki banyak modul seperti:

- Email Worker
- Scheduler
- Storage
- Database
- Web Dashboard

Agar project tetap mudah dipelihara, diperlukan struktur folder yang konsisten.

## Decision

Project menggunakan struktur:

- cmd
- internal
- docs
- web
- storage

Dengan pola:

Handler

↓

Service

↓

Repository

↓

Database

## Consequences

- Mudah dikembangkan
- Mudah dipahami
- Tidak mencampur business logic dengan infrastructure
