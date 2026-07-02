Title:
Separate Branch and MailAccount

Status:
Accepted

Context:
A branch may own multiple email accounts.

Decision:
Branch and MailAccount are separated into different entities.

Consequences:

- One Branch can have many MailAccounts.
- Easier to support Gmail, Yahoo, Outlook.
- Future-proof for multiple mailboxes.
