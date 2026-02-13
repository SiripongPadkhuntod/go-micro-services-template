# 📘 Architecture.md

Dexter Transport Backend  
Clean Architecture + Hexagonal (Ports & Adapters)  
Tech Stack: Go + Gin + PostgreSQL

---

# 1. Overview

Dexter Transport เป็นระบบจัดการขนส่งที่ออกแบบด้วยแนวคิด Clean Architecture และ Hexagonal Architecture เพื่อให้:

- แยก Business Logic ออกจาก Infrastructure อย่างชัดเจน
- รองรับการขยายระบบในอนาคต
- เปลี่ยน Database / Framework ได้โดยไม่กระทบ Business
- รองรับการทำ Unit Test และ Mock ได้ง่าย
- สามารถแยกเป็น Microservice ได้ในอนาคต

---

# 2. Architectural Pattern

## 2.1 Clean Architecture

แยก Layer ออกเป็น:

- Handler (Delivery Mechanism)
- Service (Use Case / Business Logic)
- Repository (Interface)
- Repository Implementation (Postgres)
- Domain (Core Business Model)

Dependency จะต้องไหลเข้าหา Domain เสมอ

## 2.2 Hexagonal Architecture (Ports & Adapters)

- Port = Interface
- Adapter = Implementation

ทำให้สามารถเปลี่ยน implementation ได้โดยไม่กระทบ core business

---

# 3. High Level Dependency Flow

```
Handler → Service → Repository Interface
                             ↓
                    Postgres Repository
```

Rules สำคัญ:

- Domain ห้าม import layer อื่น
- Service ห้าม import Infrastructure โดยตรง
- Handler ห้ามคุย Database โดยตรง
- Repository Implementation ห้ามมี Business Logic

---

# 4. Project Structure Responsibility

## 4.1 Root Level

### cmd/

Entry point ของ application

- cmd/server/main.go → Start HTTP server
- cmd/migrate/migrate.go → Run database migration
- cmd/docs/main.go → Generate Swagger documentation

---

### config/

- config.go → Load environment configuration

---

### db/migration/

เก็บ SQL migration versioned

- 001_init_schema.up.sql
- 001_init_schema.down.sql

---

### docs/

Swagger generated files

---

## 4.2 internal/

Core application logic

---

# 5. internal/app Layer

Business Layer ของระบบ

---

## 5.1 adapter/

ใช้สำหรับเชื่อมต่อ External Service หรือ Cross Service Call

```
adapter/
 └── rest-api/
      ├── init.go
      ├── payload/
      │    └── profile-get-by-user-id.go
      └── profile-get-by-user-id.go
```

หน้าที่:
- เรียก API ภายนอก
- แปลง payload ↔ domain
- ไม่ควรมี business logic

---

## 5.2 domain/

Core Business Model

ตัวอย่างไฟล์:

- company.go
- user.go
- dealer.go
- delivery_batch.go
- delivery_item.go
- expense.go
- invoice.go
- errors.go

Rules:

- ห้าม import gin
- ห้าม import postgres
- ห้ามมี JSON tag / DB tag
- เป็น Pure Business Struct เท่านั้น

---

## 5.3 handler/

HTTP Layer

```
handler/
 ├── dto/
 │    ├── auth_dto.go
 │    ├── dealer_dto.go
 │    ├── delivery_dto.go
 │    ├── expense_dto.go
 │    └── invoice_dto.go
 ├── auth_handler.go
 ├── dealer_handler.go
 ├── delivery_handler.go
 ├── expense_handler.go
 ├── invoice_handler.go
 └── init.go
```

หน้าที่:

- Parse request
- Validate input
- Call service
- Return response

ข้อห้าม:

- ห้ามมี business logic
- ห้ามเรียก DB ตรง ๆ

---

## 5.4 port/

เก็บ Interface ทั้งหมด

- adapter.go
- service.go
- repository.go
- handler.go

Purpose:

- กำหนด Contract ของแต่ละ Layer
- รองรับ Mock สำหรับ Unit Test

---

## 5.5 repository/

Implementation ของ Data Layer

```
repository/
 ├── init.go
 └── postgres-repository/
      ├── init.go
      ├── user_repository.go
      ├── dealer_repository.go
      ├── delivery_repository.go
      ├── expense_repository.go
      ├── invoice_repository.go
      └── entity/
           ├── user_entity.go
           ├── dealer_entity.go
           ├── delivery_entity.go
           ├── expense_entity.go
           └── invoice_entity.go
```

หน้าที่:

- เขียน SQL
- Map entity ↔ domain
- Handle transaction ระดับ data

---

## 5.6 service/

Business Logic Layer

```
service/
 ├── init.go
 ├── auth_service.go
 ├── dealer_service.go
 ├── delivery_service.go
 ├── expense_service.go
 └── invoice_service.go
```

หน้าที่:

- Validate business rule
- คำนวณยอดเงิน
- ควบคุม workflow
- เรียก repository ผ่าน interface

ข้อห้าม:

- ห้ามเขียน SQL
- ห้าม import postgres โดยตรง

---

## 5.7 utils/

Utility function ที่ใช้ใน layer app เท่านั้น

---

# 6. internal Layer อื่น ๆ

## 6.1 constant/

เก็บ constant กลางของระบบ

- role.go
- status.go
- error_code.go

---

## 6.2 infrastructure/

Implementation ของ External Systems

```
infrastructure/
 ├── db-client/
 ├── gin-client/
 ├── https-client/
 ├── middleware-client/
 ├── migrate/
 └── object-storage-client/
```

หน้าที่:

- Database connection
- HTTP server initialization
- Middleware
- External HTTP client
- Object storage client

Infrastructure สามารถเปลี่ยนได้โดยไม่กระทบ Business

---

## 6.3 property/

- property.go → Application property configuration

---

## 6.4 router/

- binding.go
- router.go

หน้าที่:

- จัดการ route grouping
- Bind handler เข้ากับ gin

---

## 6.5 server/

- context.go
- init.go
- server.go

หน้าที่:

- Initialize application
- Inject dependency
- Start server

---

# 7. pkg/

Reusable shared component

```
pkg/
 ├── model/v1/
 │    ├── base_response.go
 │    └── pagination.go
 └── v1/dto/
      ├── default.go
      └── health.go
```

ใช้สำหรับ:

- Standard API response
- Pagination model
- Shared DTO

---

# 8. Transaction Strategy

Transaction ควรถูกควบคุมที่ Service Layer

Pattern:

```
Service
 └── WithTransaction(fn)
       ├── begin
       ├── execute fn(repoWithTx)
       ├── commit
       └── rollback
```

เหตุผล:

- Service เข้าใจ business boundary
- Repository ไม่ควรรู้ flow ใหญ่ของ use case

---

# 9. Authentication & Authorization

Authentication:

- JWT Token

Authorization:

- Role-based access control

Role ตัวอย่าง:

- admin
- staff
- driver

Middleware อยู่ใน:

```
infrastructure/middleware-client/
```

---

# 10. Database Strategy

- PostgreSQL
- Versioned migration
- UUID primary key
- snake_case naming
- created_at / updated_at ทุก table
- รองรับ soft delete (deleted_at)

---

# 11. Scalability Strategy

รองรับในอนาคต:

1. Multi Company Support
   - company_id ทุก table

2. Microservice Migration
   - สามารถแยก delivery-service, invoice-service, auth-service ได้

3. Replace Database
   - เปลี่ยน Postgres → MySQL ได้โดยแก้เฉพาะ repository adapter

---

# 12. Testing Strategy

Unit Test:
- Mock repository ผ่าน port interface
- Test service layer เป็นหลัก

Integration Test:
- ทดสอบกับ Postgres จริง

HTTP Test:
- ทดสอบ handler ด้วย gin test context

---

# 13. Production Readiness

- Structured logging
- Graceful shutdown
- Health check endpoint
- Dockerized deployment
- Environment separation (dev / staging / prod)
- CI/CD pipeline

---

# 14. Summary

Dexter Transport Backend ถูกออกแบบให้:

- แยก Business ออกจาก Infrastructure
- รองรับการเติบโตของบริษัทในอนาคต
- รองรับการเปลี่ยน Technology
- Test ได้ง่าย
- พร้อมสำหรับ Production Deployment

