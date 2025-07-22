# 📦 URL Shortener API - Project Overview

## 🔍 Project Description

The **URL Shortener API** is a modern, secure, and user-friendly backend service that allows users to convert long URLs into short, shareable links. Built with clean RESTful principles and JWT-based authentication, the platform enables users to create, manage, and retrieve shortened URLs, all while ensuring performance and scalability with Redis-based rate limiting.

This project is perfect for developers and recruiters looking to see expertise in:

* Secure authentication (JWT + Cookies)
* REST API design
* Redis rate limiting
* Clean MVC architecture
* URL shortening logic with collision handling

## ⚙️ How It Works

1. **Signup/Login**: Users register and authenticate using email and password. JWTs are issued and stored securely in HTTP-only cookies.
2. **Shorten URL**: Authenticated users can send a POST request with a long URL. The system generates a unique short code and stores the mapping in the database.
3. **Redirection**: When someone accesses the short URL, they’re instantly redirected to the original long URL.
4. **Rate Limiting**: Redis is used to track user request counts and limit the number of URL creation requests per time window, preventing abuse.
5. **Get All URLs**: Logged-in users can retrieve all URLs they've shortened.
6. **Delete URL**: Users can delete any of their own shortened URLs.

## 🧠 Key Features

* 🔒 **Secure Auth**: Login/signup using JWT stored in secure cookies.
* 📉 **Rate Limiter**: Limits API calls per user using Redis to avoid spamming.
* 🔗 **Dynamic Redirection**: Instantly redirects users from a short code to the original URL.
* 🧹 **Clean API**: Built using clear, predictable RESTful principles.
* 🧑‍💼 **Multi-user Support**: Each user gets a personalized URL shortening experience.

## 🛠 Tech Stack

* **Node.js + Express** - REST API
* **Redis** - Caching + Rate Limiting
* **PostgreSQL** - Relational data storage
* **JWT + Cookies** - Secure Authentication
* **Render** - Cloud Deployment

## 🚀 Why It Matters

This project highlights backend engineering skills including:

* Authentication flow
* Performance optimization
* Middleware usage
* Redis integration
* Real-world architecture practices

Whether you're a recruiter evaluating backend capabilities or a developer looking for inspiration, this URL shortener showcases production-level backend design with clarity, security, and scalability at its core.
🔗 [View Full API Documentation](https://.postman.co/workspace/deevictor~332aa9cb-c83d-44c7-9d2e-32a5782b1309/collection/29434244-fd235b1a-6d41-4d1f-ad42-d66b0e9d34d0?action=share&creator=29434244&active-environment=29434244-05fa3ce9-3898-426a-a099-e415c16c1e4a
)
