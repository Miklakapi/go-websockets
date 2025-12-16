# Go-WebSockets

![license](https://img.shields.io/badge/license-MIT-blue)
![linux](https://img.shields.io/badge/os-Linux-green)
![language](https://img.shields.io/badge/language-Go_1.25.1-blue)
![version](https://img.shields.io/badge/version-1.0.0-success)
![status](https://img.shields.io/badge/status-production-green)

A lightweight WebSocket server written in Go using Gin and Gorilla WebSocket.
The project implements a real-time message broadcasting system based on the Hub-Client architecture.

## Table of Contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Features](#features)
* [Status](#status)

## General info
This project is a real-time WebSocket server written in Go.
It allows multiple clients to connect to a single WebSocket endpoint and exchange messages in real time.

The server is based on a central Hub, which:
* registers and unregisters clients,
* manages active connections,
* broadcasts messages to all connected clients except the sender.

Each client connection is handled concurrently using goroutines, ensuring high performance and scalability.

The application exposes:
* a simple HTTP health endpoint (/ping),
* a WebSocket endpoint (/ws) for real-time communication.

<p align="center" width="100%">
    <img src="https://github.com/Miklakapi/go-websockets/blob/main/README_IMAGES/Demo.gif"> 
</p>

## Technologies
Project is created with:

* Go 1.25.1
* Gin 1.11.0 - HTTP server and routing
* Gorilla WebSocket 1.5.3 - WebSocket implementation
* Node.js - CLI WebSocket client for testing

## Setup
### Go WebSocket Server
1. Install dependencies: ```go mod tidy```
2. Run the server: ```go run main.go```
3. The server will start on: http://localhost:8000

### WebSocket CLI Client
A simple Node.js client is included for testing purposes.
1. Install dependencies: ```npm install ```
2. Run the client: ```node main.mjs```
3. Type messages in the terminal to broadcast them to other connected clients.
4. Type exit to close the connection.

## Features
* Real-time WebSocket communication
* Hub-based architecture for client management
* Message broadcasting to all connected clients
* Automatic client registration and cleanup
* Ping/Pong mechanism to detect dead connections
* Simple CLI client for testing

## Status
The project's development has been completed.