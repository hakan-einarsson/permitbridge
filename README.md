# PermitBridge Documentation

## Overview

PermitBridge is an authorization service built in Go. It uses a Role-Based Access Control (RBAC) model, making it easy to manage user permissions in a granular and organized manner. With PermitBridge, you can define roles with different sets of permissions and assign these roles to users.

## Getting Started

To get started with PermitBridge, you first need to define your roles and users. This can be done by sending HTTP requests to the appropriate endpoints.

Here are the main routes that you can use:

- `/`: The index route
- `/authorize`: The authorization route
- `/roles`: The roles route (GET request only)
- `/users`: The users route (GET request only)
- `/assets`: The assets route (GET request only)
- `/schema/roles`: The roles schema route (supports GET, POST, PATCH, PUT, DELETE requests)
- `/schema/users`: The users schema route (supports GET, POST, PATCH, PUT, DELETE requests)

## Role Definitions

In PermitBridge, a role is defined by a name and a set of assets. Each asset is associated with a set of permissions.

Example Role Definition:

```json
{
    "name": "superadmin",
    "assets": {
        "global": [
            "read",
            "write",
            "delete"
        ]
    }
}
```
## User Definitions

In PermitBridge, a user is defined by an ID and a set of roles.

Example User Definition:

```json
{
    "id": "123456789",
    "roles": [
        "superadmin"
    ]
}
```
## Further Development
PermitBridge is a powerful and flexible system, but there's always room for improvement. Here are a few potential enhancements:

- Integration with identity providers: While the current system allows administrators to manually assign roles to users, it could be enhanced to integrate with third-party identity providers. This would allow the system to automatically assign roles based on a user's group membership in the identity provider.

- Support for more complex access control models: The current system is based on an RBAC model, but it could be extended to support more complex models like ABAC or Context-Based Access Control (CBAC).

- Auditing and reporting capabilities: The system could be enhanced to provide more robust auditing and reporting capabilities, helping administrators track access control changes and understand how permissions are being used in the system.

- Performance enhancements: There are a variety of ways the system's performance could be enhanced, such as implementing caching, optimizing database queries, or moving to a faster protocol like gRPC.