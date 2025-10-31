# Tetrix — A Terminal Falling Blocks Game

**Tetrix** is a simple falling-blocks game built for educational purposes.  
It runs entirely in the terminal and demonstrates basic concepts of game loops, rendering, and input handling using Go.

---

## 🎯 Purpose

This project was created as a learning exercise to explore:
- Terminal-based rendering techniques in Go.
- Structuring a small game loop with clean separation of logic and UI.
- Handling real-time input in a console environment.
- Applying Go idioms and clean architecture even in a simple project.

It is **not** a commercial product, and it is **not affiliated with or endorsed by The Tetris Company**.

---

## 🧩 Legal Notice

> “Tetris®” is a registered trademark of Tetris Holding LLC.  
> This project, **Tetrix**, is an independent educational clone inspired by the classic falling blocks concept.  
> It does **not** reproduce the original Tetris game’s artwork, sounds, or exact visual design.  
> The purpose of this code is purely educational — to study programming patterns and terminal game mechanics in Go.

If you wish to distribute, modify, or extend this project, please ensure you:
1. Do not use the name “Tetris” in your binary, repository, or distribution title.
2. Avoid replicating any proprietary visual designs or assets from the original Tetris.
3. Keep this disclaimer in your documentation.

---

## 🧠 Features

- Runs entirely in your terminal (no graphics libraries required)
- Smooth block falling and rotation
- Scoring and level system
- Basic sound effects (optional)
- Configurable board size

---

## 🚀 Run Locally

### Requirements
- Go 1.24 or later
- A terminal that supports ANSI escape codes

### Installation

```bash
git clone https://github.com/fsjorgeluis/tetrix
cd tetrix
go run ./cmd/game/main.go