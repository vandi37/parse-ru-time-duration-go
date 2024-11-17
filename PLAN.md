# parse-ru-time-duration-go plan

## Main idea

It's a beginner project for parsing Russian time duration sentences to time.Duration format
It gets and sends data from http requests and responses

## Project structure

### Running the project

- func main
- application structure
- timeout exiting
- service mode

### Backend packages

- func for parsing a sting with Russian sentence to time.Duration format

    Example sentences:
        - "1 день 2 часа 10 минут"
        - "76 лет 23 месяца 89 секунд"
        - "12 секунд"
- Custom error handling with vanerrors
- Logging the result

### HTTP part

- Getting requests
- Parsing them
- Sending an error or a response

## Time plan

The project should take from one to four days

## Plan

- set application structure -> done
- Create parsing func -> done
- Create unit tests -> not done
- Create http server -> not done
- Create unit tests #2 -> not done
- v1.0.0 -> not done
