# Tubes Stima 2 by GasTubes
> WikiRace Game Solver written in Go. Based on the concept of Breadth First Search (BFS) and Iterative Deepening Search (IDS) algorithm.


## Table of Contents
* [General Information](#general-information)
* [IDS and BFS Implementation](#ids-and-bfs-implementation)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Overview](#overview)
* [Setup](#setup)
* [Project Status](#project-status)
* [Room for Improvement](#room-for-improvement)
* [Library](#library)
* [Acknowledgements](#acknowledgements)
* [Authors](#authors)


## General Information
![GrafWikiRace](https://miro.medium.com/v2/resize:fit:1400/1*jxmEbVn2FFWybZsIicJCWQ.png)

WikiRace or Wiki Game is a game involving Wikipedia, a free online encyclopedia maintained by various volunteers around the world, where players start at a Wikipedia article and have to browse other articles on Wikipedia (by clicking on the link within each article) to go to another article that has been predetermined in the shortest time or fewest clicks (articles).

The program is made to solve the WikiRace Game problem and to win the game by implementing the IDS and BFS algorithm. It is a web-based application which receives input in the form of algorithm type, initial article title, and title purpose article. The program provides output in the form of the number of articles examined, the number of articles traversed, the article browsing route (from the initial article to the destination article), and search time (in ms).

## IDS and BFS Implementation 
mmm


## Technologies Used
The whole program was written in Go.


## Features
- [x] Receive input in the form of algorithm type, initial article title, and title purpose article
- [x] Choose the algorithm (IDS or BFS) through input from the user
- [x] Provides output in the form of the number of articles examined, the number of articles traversed, the article browsing route (from the initial article to the destination article), and search time (in ms)
- [x] Can at least issue one of the shortest routes that is less than 5 minutes in each game


## Overview
![Overview](src/Image/Overview.jpg)


## Setup
### Installation
- Download and install [Go](https://go.dev/doc/install) 
- Install the whole modules and libraries used in the source code
- Download the whole folders and files in this repository or do clone the repository

### Compilation 
To run in local:

Frontend
1. Clone this repository in your own local directory

    `git clone https://github.com/ZachS17/Tubes2_GasTubes.git`

2. Open the command line and change the directory to 'Frontend' folder

    `cd Tubes2_GasTubes/src/Frontend`
    
3. Run `npm install` on the command line
4. Run `npm run start` on the command line

Backend 
1. Navigate to `/src/Backend` folder
2. Run `go get .` to make sure all dependencies installed
3. Run `go run api.go`
4. **Recommended**: test the API using [Postman](https://www.postman.com/downloads/) or thunder client on vscode extensions


## Project Status
Project is: _complete_

All the specifications were implemented.


## Room for Improvement
- A faster or more efficient algorithm to make the program run quicker
- A better UI/UX to satisfy the users of this application


## Library
* [Node JS](https://nodejs.org/en/)
* [React](https://reactjs.org/)
* [Bootstrap](https://getbootstrap.com/)
* [Axios](https://axios-http.com/docs/intro)
* [Golang](https://go.dev/)
* [Echo](https://echo.labstack.com/)

for testing purposes:
* [Postman](https://www.postman.com/downloads/)

  
## Acknowledgements
- This project was based on [Spesifikasi Tugas Besar 2 Stima](https://informatika.stei.itb.ac.id/~rinaldi.munir/Stmik/2023-2024/Tubes2-Stima-2024.pdf)
- Thanks to God
- Thanks to Mrs. Masayu Leylia Khodra, Mrs. Nur Ulfa Maulidevi, and Mr. Rinaldi as our lecturers
- Thanks to academic assistants
- This project was created to fulfill our Big Project for IF2211 Algorithm Strategies


## Authors
| No. | Name | Student ID |
| :---: | :---: | :---: |
| 1. | Habibi Galang Trianda | 10023457 |
| 2. | Nelsen Putra | 13520130 |
| 3. | Zachary Samuel Tobing | 13522016 |
