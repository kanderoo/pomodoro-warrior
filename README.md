# Pomodoro-Warrior

A simple pomodoro timer that integrates with taskwarrior.

Written in Go using BubbleTea.

## Known Bugs
- The `timer` bubble in the [bubbles library](https://github.com/charmbracelet/bubbles/tree/master) has a bug where stopping and restarting the timer under a second causes tick duplication.