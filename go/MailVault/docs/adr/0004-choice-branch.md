ADR-0004

Separate Branch and Mail Account

Reason:

One Branch can have multiple mail accounts.

## Branch

ID
Code
Name
IsActive
CreatedAt
UpdatedAt
DeletedAt

        │
        │ 1
        │
        │
        ▼

## MailAccount

ID
BranchID
Email
Password
Provider
IsActive
LastSyncAt
CreatedAt
UpdatedAt
DeletedAt

        │
        │ 1
        │
        ▼

## Email

ID
MailAccountID
YahooMessageID
Subject
Sender
Recipient
CC
BCC
Body
ReceivedAt
HasAttachment
CreatedAt
UpdatedAt
DeletedAt

        │
        │ 1
        ▼

## Attachment

ID
EmailID
OriginalName
StoredName
MimeType
Size
StoragePath
