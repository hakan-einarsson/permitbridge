# PermitBridge
The Authorization Service is a role-based access control (RBAC) system, designed to handle authorization requests in a microservice architecture.

## Key Features
Role-Based Access Control: The service employs an RBAC model where permissions are associated with roles, and users are assigned to these roles. This setup provides a flexible way to manage access control in a system with potentially many users and diverse types of permissions.

- JSON Schema: The roles and permissions are defined in a JSON schema, which offers a clear and easy-to-modify structure for access control rules. The JSON-based configuration makes the service highly adaptable to various access control requirements.

- HTTP API: The service exposes an HTTP API for making authorization requests. Clients can send a request to this API, including the user's roles and the requested action. The service returns a decision indicating whether the action is permitted.

- Scalability and Future Enhancements: The service is designed with scalability and extensibility in mind. It could be expanded in the future to support more complex access control models, such as Attribute-Based Access Control (ABAC). Additionally, the current HTTP API could be supplemented or replaced with a faster protocol like gRPC for improved performance.

## Example JSON Schema
Below is an example of what the JSON schema might look like:

```json
{
    "roles": [
        {
            "name": "superadmin",
            "assets": {
                "global": [
                    "read",
                    "write",
                    "delete"
                ]
            }
        },
        {
            "name": "guest",
            "assets": {
                "posts": [
                    "read",
                ]
            }
        }
    ]
}
```
Global is a special asset that applies to all assets in the system. In this example, the superadmin role has full access to all assets, while the guest role only has read access to the posts asset.