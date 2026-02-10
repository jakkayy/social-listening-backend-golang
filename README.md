# Social Listening Backend (Golang)

> **Production-grade Social Listening & Market Intelligence Backend**
> วิเคราะห์เสียงลูกค้าจาก Social Comment → แปลงเป็น Insight + Alert เชิงกลยุทธ์

---

##  ภาพรวมโปรเจกต์

โปรเจกต์นี้คือ **Backend System สำหรับ Social Listening** ที่ออกแบบในระดับ production จริง ไม่ใช่ demo หรือ tutorial

ระบบสามารถ:

* เก็บ Social Comments
* วิเคราะห์ Sentiment / Intent
* ตรวจจับ Trend ตามช่วงเวลา (Window-based)
* ตรวจจับ Keyword Spike เชิงธุรกิจ
* สร้าง Alert พร้อม Deduplication
* สรุปผลรายวัน (Daily Insight Snapshot)

เหมาะสำหรับ:

* Social Listening Platform
* Market / Customer Intelligence
* Marketing Analytics
* Consulting Dashboard

---

##  Core Concept

```
Social Comments
      ↓
Ingestion Worker
      ↓
Sentiment / Intent Analysis
      ↓
Trend & Keyword Detection
      ↓
Alert Engine (Deduplicated)
      ↓
Daily Insight Snapshot
      ↓
API for Dashboard / Report
```

> แยก **Worker** ออกจาก **API** อย่างชัดเจน เพื่อความเสถียรและ scalability

---

##  Architecture Overview

### 1. Worker (Background Process)

หน้าที่หลัก:

* Collect comments (Mock / Future: Facebook API)
* Analyze sentiment & intent
* Save raw data + analysis
* Detect:

  * Negative sentiment spike (Window-based)
  * Keyword spike (Business keywords)
* Generate alerts (with deduplication)
* Create daily insight snapshot

### 2. API Server (Gin)

หน้าที่หลัก:

* Serve data from database
* Endpoints:

  * `/api/overview`
  * `/api/alerts`
  * `/api/daily-insights`

> API **ไม่ทำงานหนัก** → อ่านข้อมูลจาก DB เท่านั้น

### 3. Database (PostgreSQL)

ใช้เป็น **Single Source of Truth**

ตารางหลัก:

* `comments`
* `comment_analysis`
* `alerts`
* `daily_insights`

---

##  Project Structure

```
cmd/
 ├─ api/                # HTTP API (Gin)
 │   └─ main.go
 └─ worker/             # Background Worker
     └─ main.go

internal/
 ├─ api/handler/        # HTTP Handlers
 ├─ config/             # DB / App config
 ├─ domain/             # Business entities
 ├─ ingestion/          # Data collectors
 ├─ processing/         # Sentiment / Intent logic
 ├─ insight/            # Trend & keyword logic
 └─ storage/            # Repository layer

migrations/             # SQL schema migrations
```

---

##  Key Features (เชิงเทคนิค)

###  Window-based Trend Detection

* เปรียบเทียบข้อมูล **10 นาทีล่าสุด** กับ **10 นาทีก่อนหน้า**
* ไม่ใช้ in-memory state
* Worker restart แล้วไม่พัง

###  Alert Deduplication

* Alert ประเภทเดียวกัน
* จะไม่ถูกสร้างซ้ำภายในช่วงเวลาที่กำหนด (เช่น 30 นาที)

###  Keyword Spike Detection

* ตรวจจับคำเชิงธุรกิจ เช่น:

  * แพง
  * ช้า
  * โกง
  * บริการแย่
* แจ้งเตือนพร้อม context ว่า "ปัญหาคืออะไร"

###  Daily Insight Snapshot

* สรุปข้อมูลรายวัน (1 row / 1 day)
* ใช้กับ dashboard / report / slide ได้ทันที

---

## 📊 Daily Insight Snapshot Example

| Date       | Total | Positive | Neutral | Negative | Top Keywords | Alerts |
| ---------- | ----- | -------- | ------- | -------- | ------------ | ------ |
| 2026-02-10 | 120   | 65       | 30      | 25       | แพง, ช้า     | 3      |

---

##  How to Run

### 1. Start Database (Docker)

```bash
docker run -d \
  -p 5433:5432 \
  -e POSTGRES_DB=social_listening \
  -e POSTGRES_USER=sl_user \
  -e POSTGRES_PASSWORD=sl_pass \
  postgres:15
```

### 2. Run Migrations

```bash
psql -h localhost -p 5433 -U sl_user -d social_listening -f migrations/init.sql
psql -h localhost -p 5433 -U sl_user -d social_listening -f migrations/add_annalyzed_at.sql
psql -h localhost -p 5433 -U sl_user -d social_listening -f migrations/alert.sql
psql -h localhost -p 5433 -U sl_user -d social_listening -f migrations/daily_insights.sql
```

### 3. Run Worker

```bash
go run cmd/worker/main.go
```

### 4. Run API

```bash
go run cmd/api/main.go
```

---

##  Example APIs

### Get Alerts

```
GET /api/alerts
```

### Get Daily Insights

```
GET /api/daily-insights
```

---

##  Design Decisions

* **Worker-first architecture** → รองรับงานหนัก
* **Database as Source of Truth** → consistency
* **Repository Pattern** → testable & maintainable
* **Idempotent operations** → production-ready
* **Time-window analytics** → monitoring-grade logic

---

##  Use Cases

* Social Listening Platform
* Brand Monitoring System
* Marketing Intelligence
* Customer Experience Analytics
* Consulting / Strategy Dashboard

---

##  Future Improvements

* Real Facebook / Social API integration
* NLP-based keyword extraction
* Frontend dashboard (Next.js)
* Alert severity levels
* Multi-brand / multi-client support

---

##  Author Notes

โปรเจกต์นี้ออกแบบเพื่อแสดง:

* System design thinking
* Backend engineering skill (Mid–Senior)
* Data-driven architecture

> เน้น "คิดเป็นระบบ" มากกว่าแค่เขียน API

---

**Status:**  Production-grade backend foundation complete
