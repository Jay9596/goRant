# goRant
An unofficial Go client for public devRant API

https://www.devrant.io/  

[![Go Report Card](https://goreportcard.com/badge/github.com/jay9596/goRant)](https://goreportcard.com/report/github.com/jay9596/goRant)

## Table of Content
1. [Installation](#Installation)
2. [Documentation](#Documentation)
3. [Getting Started](#Getting-Started)
4. [Todo](#Todo)


## Installation:
 ``` go get github.com/jay9596/goRant ```
 
## Documentation
  ``` go doc github.com/jay9596/goRant ```  
  Or visit:
  [GoDoc](https://godoc.org/github.com/Jay9596/goRant)

## Getting Started
* ### goRant Functions
    #### Assign a New client  
      devRant := goRant.New()
     #### Rants : Fetch rants
     ``` rants,err := devRant.Rants() ```
     #### Rant : Fetch a single rant
     ``` rant,comments,err := devRant.GetRant(rantID) ```
     #### Profile : Fetch a user profile
     ``` user,err := devRant.Profile("byte") ```
     #### Search : Search on devrant
     ``` rants,err := devRant.Search("vscode") ```
     #### Surprise : Fetch a random rant
     ``` rant,comments,err := devRant.Surprise() ```
     #### WeeklyRant : Fetch weekly rants
     ``` rants,err := devRant.WeeklyRant() ```
     #### Collabs : Fetch collabs 
     ``` rants,err := devRant.Collabs() ```
     #### Stories : Fetch stories
     ``` rants,err := devRant.Stories() ```
* ### goRant Types
     All _rant_ specified above represent a **Rant** struct  
     The _rants_ represent an array of Rant  **[ ]Rant**  
     The _user_ represent **User** struct  
     The _comments_ represent an array of Comment **[ ]Comment**  
 
 ## Todo
 - [x] Add supporter ++ support   
 - [ ] Add Tests  
