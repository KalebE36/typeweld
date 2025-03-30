# 🔥 TypeWeld

> **TypeWeld:** Define once, use everywhere — Seamless type-safe communication across languages.

TypeWeld is an **open-source type-safe communication framework** that allows you to define data structures and services using a simple, intuitive DSL with a `.twld` extension. From this definition, TypeWeld generates type-safe, language-specific code that seamlessly integrates with your backend and frontend systems.

Inspired by **Protocol Buffers (gRPC)** and **Thrift**, TypeWeld goes beyond by focusing on:
- ✅ **Type Safety** — Ensuring consistent types across multiple languages.
- ⚡ **Custom Wire Protocols** — Optimized for performance and security.
- 🔗 **Modular Code Generation** — Easily extend to new languages or platforms.

---

## 🚀 Why TypeWeld?

- **Define Once, Use Everywhere**: Define your types and services in `.twld` files and generate language-specific code.
- **Type-Safe Communication**: Prevent bugs by maintaining strict type safety across different environments.
- **Multi-Language Support**: Generate code for multiple target languages like Go, Python, and JavaScript.
- **High Performance**: Optimized serialization/deserialization with minimal overhead.
- **Extensible DSL**: A custom and intuitive DSL that is easy to learn and use.

---

## 📚 How It Works

### 1. **Define Types and Services**
Create a `.twld` file with your type definitions and services.

```plaintext
// example.twld
package mypackage;

message Person {
  string name = 1;
  int age = 2;
}

message Greeting {
  string text = 1;
  Person sender = 2;
}

service Greeter {
  rpc SayHello(Greeting) returns (Greeting);
}