# Introduction
This repository contains my first GO project. Having a strong object-oriented background with .Net/C# I am curious to start with Golang.
While I am doing a lot of googling (with a good amount of trial & error) to learn the language, I am also reading these two books:

* [Learning GO](https://www.amazon.de/Learning-Go-Idiomatic-Real-World-Programming/dp/1098139291)
* [100 Go Mistakes and How to Avoid Them](https://www.amazon.de/100-Mistakes-How-Avoid-Them/dp/1617299596)

The learning objective is to implement a notification service (HTTP service) that acts as a message dispatcher to various messaging applications (e.g. MS Teams, Slack, ...).
Learning a new programming language is an evolutionary process. So I am going to start with a very simple solution that I will expand over time.

> [!CAUTION]
> The code may not reflect best practices and is not yet finalised in many places. 

# Quickstart
- [Build and run the notification service locally](/docs/howto_buildandrun.md)
- [Use the notification service API](/docs/howto_api.md)
- [Implement new receivers (e.g. to send messages to Twilio)](/docs/howto_receivers.md)
- [Functional and non-functional requirements](/docs/requirements.md)

# Notification Service

## System Design
A REST/JSON based service that dispatches notifications to 3rd party (chat or message) applications (receiver systems). 
The service should be extensible in the sense of integrating more receiver systems and custom routing rules of 
notifications.  

![Notification Service](/docs/assets/NotificationServiceComponents.png)

_The diagram uses [C4 notation](https://c4model.com/)_

**TODO describe each building block**

## Considerations and future improvements

### Security
- Implementing authentication and authorisation
- Implementing HTTPS
- Implementing HTTP CORS

## Extensibility
- Adding versioning of notification and concrete receivers
- Adding a property bag to notifications to send arbitrary data to receivers

## Reliability
- Adding notification persistence to core messaging middleware
- Adding retry logic for receivers in messaging middleware
- Adding circuit breaker utils for downstream receiver systems and concrete receivers
- Definition of notification delivering guarantees
- Implementing throttling and quotas (rate limiting)

## Architecture & Code
- Architecture tests
- Improved error handling
- Improved input validation
- Applying GO best-practices

### DevOps
- Monitoring of Service Level Indicators (SLI) e.g. Open Telemetry Metrics
- Alerting of SLIs
- Deployment and day two scenarios (breaking changes)
- Configuration files for receivers (instead of environment variables)
- Notification traceability (Correlation ID)



